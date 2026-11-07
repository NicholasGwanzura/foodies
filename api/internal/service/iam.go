package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"foodies/api/internal/db"
)

// CreateUser adds a sign-in account. It starts with no credentials, so the
// first-use step sets them.
func (s *Service) CreateUser(ctx context.Context, p db.CreateUserParams) (db.User, error) {
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateUser(ctx, p)
	if err != nil {
		return row, wrap(err, "create user")
	}
	if err := s.q(ctx).EnsureCredentials(ctx, row.ID); err != nil {
		return row, wrap(err, "create credentials")
	}
	return row, s.audit(ctx, "iam.user.write", "users", row.ID, nil, nil, row)
}

func (s *Service) syncMfa(ctx context.Context, user uuid.UUID) error {
	needs, err := s.q(ctx).UserRequiresMfa(ctx, user)
	if err != nil {
		return fmt.Errorf("check mfa: %w", err)
	}
	return wrap(s.q(ctx).SetUserMfaRequired(ctx, db.SetUserMfaRequiredParams{ID: user, MfaRequired: needs}), "store mfa flag")
}

// CreateRoleAssignment grants a role and refreshes whether the user needs MFA.
func (s *Service) CreateRoleAssignment(ctx context.Context, p db.CreateRoleAssignmentParams) (db.RoleAssignment, error) {
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRoleAssignment(ctx, p)
	if err != nil {
		return row, wrap(err, "create role assignment")
	}
	if err := s.syncMfa(ctx, row.UserID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "iam.role_assignment.write", "role_assignments", row.ID, nil, nil, row)
}

// DeleteRole removes a role, unless the seed owns it.
func (s *Service) DeleteRole(ctx context.Context, id uuid.UUID) error {
	before, err := s.GetRole(ctx, id)
	if err != nil {
		return err
	}
	if before.System {
		return errState("Seeded roles can be edited but not deleted.")
	}
	if err := s.q(ctx).DeleteRole(ctx, id); err != nil {
		return wrap(err, "delete role")
	}
	return s.audit(ctx, "iam.role.write", "roles", id, nil, before, nil)
}

// LockUser stops sign-in and ends the user's sessions.
func (s *Service) LockUser(ctx context.Context, p db.ActLockUserParams) (db.User, error) {
	return s.endAccess(ctx, p.ID, "iam.user.lock", func(actor *uuid.UUID) (db.User, error) {
		return s.q(ctx).ActLockUser(ctx, db.ActLockUserParams{ID: p.ID, UpdatedBy: actor})
	})
}

// DisableUser ends the account for good.
func (s *Service) DisableUser(ctx context.Context, p db.ActDisableUserParams) (db.User, error) {
	return s.endAccess(ctx, p.ID, "iam.user.disable", func(actor *uuid.UUID) (db.User, error) {
		return s.q(ctx).ActDisableUser(ctx, db.ActDisableUserParams{ID: p.ID, UpdatedBy: actor})
	})
}

func (s *Service) endAccess(ctx context.Context, id uuid.UUID, action string, apply func(*uuid.UUID) (db.User, error)) (db.User, error) {
	before, err := s.GetUser(ctx, id)
	if err != nil {
		return before, err
	}
	row, err := apply(actorID(ctx))
	if err != nil {
		return before, stateError(err, "do this to the account")
	}
	if err := s.q(ctx).RevokeUserSessions(ctx, id); err != nil {
		return row, wrap(err, "revoke sessions")
	}
	return row, s.audit(ctx, action, "users", id, nil, before, row)
}

// UnlockUser lets a locked user sign in again.
func (s *Service) UnlockUser(ctx context.Context, p db.ActUnlockUserParams) (db.User, error) {
	before, err := s.GetUser(ctx, p.ID)
	if err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActUnlockUser(ctx, p)
	if err != nil {
		return before, stateError(err, "unlock this account")
	}
	if err := s.q(ctx).ResetFailedLogins(ctx, p.ID); err != nil {
		return row, wrap(err, "reset attempts")
	}
	return row, s.audit(ctx, "iam.user.unlock", "users", p.ID, nil, before, row)
}

// ResetCredentialsUser wipes password, PIN and authenticator and ends every
// session. The user then needs the first-use step to sign in again.
func (s *Service) ResetCredentialsUser(ctx context.Context, p db.ActResetCredentialsUserParams) (db.User, error) {
	q := s.q(ctx)
	before, err := s.GetUser(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := q.SetPassword(ctx, db.SetPasswordParams{UserID: p.ID}); err != nil {
		return before, wrap(err, "clear password")
	}
	if err := q.ClearPin(ctx, p.ID); err != nil {
		return before, wrap(err, "clear pin")
	}
	if err := q.ClearMfa(ctx, p.ID); err != nil {
		return before, wrap(err, "clear mfa")
	}
	if err := q.RevokeUserSessions(ctx, p.ID); err != nil {
		return before, wrap(err, "revoke sessions")
	}
	if err := q.SetCredentialStatus(ctx, db.SetCredentialStatusParams{ID: p.ID, CredentialStatus: ptr("unset")}); err != nil {
		return before, wrap(err, "mark credentials unset")
	}
	row, err := q.GetUser(ctx, p.ID)
	if err != nil {
		return before, wrap(err, "read user")
	}
	return row, s.audit(ctx, "iam.user.reset_credentials", "users", p.ID, nil, before, row)
}

// CreateEmployeeStatusChange records a status. A separated status ends access at once.
func (s *Service) CreateEmployeeStatusChange(ctx context.Context, p db.CreateEmployeeStatusChangeParams) (db.EmployeeStatusChange, error) {
	q := s.q(ctx)
	p.CreatedBy = actorID(ctx)
	row, err := q.CreateEmployeeStatusChange(ctx, p)
	if err != nil {
		return row, wrap(err, "create status change")
	}
	if err := q.SetEmployeeStatus(ctx, db.SetEmployeeStatusParams{ID: p.EmployeeID, EmploymentStatus: &p.Status}); err != nil {
		return row, wrap(err, "set employee status")
	}
	switch p.Status {
	case "resigned", "terminated", "retired":
		user, err := q.GetUserByEmployee(ctx, p.EmployeeID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return row, wrap(err, "read user")
		}
		if err == nil {
			if _, err := s.DisableUser(ctx, db.ActDisableUserParams{ID: user.ID}); err != nil {
				return row, err
			}
		}
		if err := q.RevokeEmployeeIdentities(ctx, p.EmployeeID); err != nil {
			return row, wrap(err, "revoke whatsapp")
		}
	}
	return row, s.audit(ctx, "iam.employee_status_change.write", "employees", p.EmployeeID, nil, nil, row)
}

// applyTransfer moves an employee once the transfer is approved.
func (s *Service) applyTransfer(ctx context.Context, id uuid.UUID) error {
	q := s.q(ctx)
	t, err := q.GetEmployeeTransfer(ctx, id)
	if err != nil {
		return wrap(err, "read transfer")
	}
	err = q.ApplyEmployeeTransfer(ctx, db.ApplyEmployeeTransferParams{
		ID: t.EmployeeID, SiteID: t.ToSiteID, DepartmentID: t.ToDepartmentID, JobTitle: t.NewJobTitle, ManagerEmployeeID: t.NewManagerEmployeeID,
	})
	if err != nil {
		return wrap(err, "apply transfer")
	}
	_, err = q.SetEmployeeTransferStatus(ctx, db.SetEmployeeTransferStatusParams{ID: id, Status: "effective", UpdatedBy: actorID(ctx)})
	return wrap(err, "mark transfer effective")
}
