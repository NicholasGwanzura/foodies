package httpapi

import (
	"context"
	"sort"
	"strings"

	"github.com/google/uuid"

	"foodies/api/internal/auth"
	"foodies/api/internal/http/apigen"
	"foodies/api/internal/service"
)

func tokenPair(p service.TokenPair) apigen.TokenPair {
	out := apigen.TokenPair{Status: apigen.TokenPairStatus(p.Status)}
	if p.AccessToken != "" {
		out.AccessToken = &p.AccessToken
		out.RefreshToken = &p.RefreshToken
		out.ExpiresIn = &p.ExpiresIn
	}
	if p.MfaChallengeID != "" {
		out.MfaChallengeId = &p.MfaChallengeID
	}
	return out
}

// Login signs in with employee number and password.
func (s *Server) Login(ctx context.Context, req apigen.LoginRequestObject) (apigen.LoginResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	pair, err := s.service.Login(ctx, req.Body.EmployeeNumber, str(req.Body.Password), requestIP(ctx), req.Body.DeviceId)
	if err != nil {
		return nil, err
	}
	return apigen.Login200JSONResponse(tokenPair(pair)), nil
}

// VerifyMfa finishes a two-factor sign-in.
func (s *Server) VerifyMfa(ctx context.Context, req apigen.VerifyMfaRequestObject) (apigen.VerifyMfaResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	challenge := ""
	if req.Body.MfaChallengeId != nil {
		challenge = *req.Body.MfaChallengeId
	}
	pair, err := s.service.VerifyMfa(ctx, challenge, req.Body.Code, requestIP(ctx))
	if err != nil {
		return nil, err
	}
	return apigen.VerifyMfa200JSONResponse(tokenPair(pair)), nil
}

// RefreshToken rotates the refresh token.
func (s *Server) RefreshToken(ctx context.Context, req apigen.RefreshTokenRequestObject) (apigen.RefreshTokenResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	pair, err := s.service.Refresh(ctx, str(req.Body.RefreshToken))
	if err != nil {
		return nil, err
	}
	return apigen.RefreshToken200JSONResponse(tokenPair(pair)), nil
}

// Logout ends the session.
func (s *Server) Logout(ctx context.Context, _ apigen.LogoutRequestObject) (apigen.LogoutResponseObject, error) {
	if err := s.service.Logout(ctx); err != nil {
		return nil, err
	}
	return apigen.Logout204Response{}, nil
}

// ChangePassword sets a new password.
func (s *Server) ChangePassword(ctx context.Context, req apigen.ChangePasswordRequestObject) (apigen.ChangePasswordResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	if err := s.service.ChangePassword(ctx, str(req.Body.CurrentPassword), str(req.Body.NewPassword)); err != nil {
		return nil, err
	}
	return apigen.ChangePassword204Response{}, nil
}

// SetPin sets the signing PIN.
func (s *Server) SetPin(ctx context.Context, req apigen.SetPinRequestObject) (apigen.SetPinResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	if err := s.service.SetPin(ctx, str(req.Body.Pin), str(req.Body.CurrentSecret)); err != nil {
		return nil, err
	}
	return apigen.SetPin204Response{}, nil
}

// UnlockWithPin switches user on an enrolled device.
func (s *Server) UnlockWithPin(ctx context.Context, req apigen.UnlockWithPinRequestObject) (apigen.UnlockWithPinResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	pair, err := s.service.UnlockWithPin(ctx, req.Body.DeviceId, req.Body.EmployeeNumber, str(req.Body.Pin), requestIP(ctx))
	if err != nil {
		return nil, err
	}
	return apigen.UnlockWithPin200JSONResponse(tokenPair(pair)), nil
}

// EnrolMfa starts two-factor enrolment.
func (s *Server) EnrolMfa(ctx context.Context, _ apigen.EnrolMfaRequestObject) (apigen.EnrolMfaResponseObject, error) {
	secret, uri, err := s.service.EnrolMfa(ctx)
	if err != nil {
		return nil, err
	}
	return apigen.EnrolMfa200JSONResponse{Secret: secret, OtpauthUri: uri}, nil
}

// ConfirmMfa finishes two-factor enrolment.
func (s *Server) ConfirmMfa(ctx context.Context, req apigen.ConfirmMfaRequestObject) (apigen.ConfirmMfaResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	if err := s.service.ConfirmMfa(ctx, req.Body.Code); err != nil {
		return nil, err
	}
	return apigen.ConfirmMfa204Response{}, nil
}

func scopeFields(scopeType string, id *uuid.UUID) map[string]any {
	m := map[string]any{"scope_type": scopeType}
	if id != nil {
		m["scope_id"] = id
	}
	return m
}

// GetMe returns the signed-in user.
func (s *Server) GetMe(ctx context.Context, _ apigen.GetMeRequestObject) (apigen.GetMeResponseObject, error) {
	me, err := s.service.GetMe(ctx)
	if err != nil {
		return nil, err
	}
	user, err := output[apigen.User](me.User)
	if err != nil {
		return nil, err
	}
	employee, err := output[apigen.Employee](me.Employee)
	if err != nil {
		return nil, err
	}
	roles := make([]map[string]any, 0, len(me.Grants))
	for _, g := range me.Grants {
		roles = append(roles, map[string]any{"key": g.RoleKey, "name": g.RoleName, "scope": scopeFields(g.ScopeType, g.ScopeID)})
	}
	return output[apigen.GetMe200JSONResponse](map[string]any{"user": user, "employee": employee, "roles": roles})
}

// GetMyPermissions returns what the signed-in user holds, and where.
func (s *Server) GetMyPermissions(ctx context.Context, _ apigen.GetMyPermissionsRequestObject) (apigen.GetMyPermissionsResponseObject, error) {
	p := auth.From(ctx)
	byKey := map[string][]map[string]any{}
	for _, g := range p.Grants {
		for _, key := range g.Permissions {
			byKey[key] = append(byKey[key], scopeFields(g.ScopeType, g.ScopeID))
		}
	}
	keys := make([]string, 0, len(byKey))
	for k := range byKey {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	perms := make([]map[string]any, 0, len(keys))
	for _, k := range keys {
		perms = append(perms, map[string]any{"key": k, "scopes": byKey[k]})
	}
	return output[apigen.GetMyPermissions200JSONResponse](map[string]any{"permissions": perms})
}

// ListMySessions lists the signed-in user's sessions.
func (s *Server) ListMySessions(ctx context.Context, _ apigen.ListMySessionsRequestObject) (apigen.ListMySessionsResponseObject, error) {
	rows, err := s.service.ListMySessions(ctx)
	if err != nil {
		return nil, err
	}
	data, err := outputs[apigen.Session](rows)
	if err != nil {
		return nil, err
	}
	return apigen.ListMySessions200JSONResponse{Data: data}, nil
}

// RevokeMySession ends one of the signed-in user's sessions.
func (s *Server) RevokeMySession(ctx context.Context, req apigen.RevokeMySessionRequestObject) (apigen.RevokeMySessionResponseObject, error) {
	id, err := uuid.Parse(req.SessionId)
	if err != nil {
		return nil, service.Fail(404, "not_found", "Not found", "No such session.")
	}
	if err := s.service.RevokeMySession(ctx, id); err != nil {
		return nil, err
	}
	return apigen.RevokeMySession204Response{}, nil
}

// ListPermissions lists every permission the API checks.
func (s *Server) ListPermissions(_ context.Context, _ apigen.ListPermissionsRequestObject) (apigen.ListPermissionsResponseObject, error) {
	seen := map[string]bool{}
	data := []apigen.Permission{}
	for _, op := range apigen.Operations {
		if op.Permission == "" || seen[op.Permission] {
			continue
		}
		seen[op.Permission] = true
		module, _, _ := strings.Cut(op.Permission, ".")
		data = append(data, apigen.Permission{Key: op.Permission, Module: module, Description: op.Summary})
	}
	sort.Slice(data, func(i, j int) bool { return data[i].Key < data[j].Key })
	return apigen.ListPermissions200JSONResponse{Data: data}, nil
}
