package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

// contentHash freezes what is being signed. Status and version fields stay in,
// because a signature covers one exact state of the record.
func (s *Service) contentHash(ctx context.Context, table string, id uuid.UUID) (*string, error) {
	sub, ok := subjects[table]
	if !ok {
		return nil, errInvalid("/subject_type", "Unknown record type.")
	}
	row, err := sub.get(ctx, s.q(ctx), id)
	if err != nil {
		return nil, notFound(err, "record")
	}
	raw, err := json.Marshal(row)
	if err != nil {
		return nil, fmt.Errorf("encode subject: %w", err)
	}
	sum := sha256.Sum256(raw)
	return ptr(hex.EncodeToString(sum[:])), nil
}

// CreateSignatureRequest asks someone to confirm a record.
func (s *Service) CreateSignatureRequest(ctx context.Context, p db.CreateSignatureRequestRowParams) (db.SignatureRequest, error) {
	var zero db.SignatureRequest
	if p.SignerKind == "external" && p.ExternalName == nil {
		return zero, errInvalid("/external_name", "Name the external party.")
	}
	if p.SignerKind == "employee" && p.SignerEmployeeID == nil && p.SignerRoleID == nil {
		return zero, errInvalid("/signer_employee_id", "Choose who signs.")
	}
	hash, err := s.contentHash(ctx, p.SubjectType, p.SubjectID)
	if err != nil {
		return zero, err
	}
	p.ContentHash = hash
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateSignatureRequestRow(ctx, p)
	if err != nil {
		return zero, wrap(err, "create signature request")
	}
	return row, s.audit(ctx, "signatures.signature_request.write", "signature_requests", row.ID, nil, nil, row)
}

// SignSignatureRequest records a signature with the signer's own PIN. A witness
// on a shared device passes their own session here, so the PIN is always theirs.
func (s *Service) SignSignatureRequest(ctx context.Context, id uuid.UUID, pin string) (db.SignatureRequest, error) {
	q := s.q(ctx)
	p := auth.From(ctx)
	row, err := q.GetSignatureRequest(ctx, id)
	if err != nil {
		return row, notFound(err, "signature request")
	}
	if row.Status != "pending" {
		return row, errState("This signature is no longer open.")
	}
	if row.SignerKind == "employee" {
		allowed := row.SignerEmployeeID != nil && *row.SignerEmployeeID == p.EmployeeID
		if !allowed && row.SignerRoleID != nil && row.SignerEmployeeID == nil {
			allowed, err = q.UserHoldsRole(ctx, db.UserHoldsRoleParams{UserID: p.UserID, RoleID: *row.SignerRoleID})
			if err != nil {
				return row, fmt.Errorf("check role: %w", err)
			}
		}
		if !allowed {
			return row, Fail(http.StatusForbidden, "forbidden", "Not your signature", "Only the named signer can confirm this.")
		}
	}
	if err := s.CheckPin(ctx, p.UserID, pin); err != nil {
		return row, err
	}
	hash, err := s.contentHash(ctx, row.SubjectType, row.SubjectID)
	if err != nil {
		return row, err
	}
	if row.ContentHash != nil && hash != nil && *row.ContentHash != *hash {
		return row, Fail(http.StatusConflict, "conflict", "Record changed", "The record changed after this signature was requested. Ask for a new one.")
	}
	signed, err := q.RecordSignature(ctx, db.RecordSignatureParams{ID: id, SignedChannel: &p.Channel, DeviceID: p.DeviceID, ContentHash: hash, AttestedBy: &p.UserID})
	if err != nil {
		return row, stateError(err, "sign this")
	}
	return signed, s.audit(ctx, "signatures.signature_request.sign", row.SubjectType, row.SubjectID, nil, nil, signed)
}

// DeclineSignatureRequest refuses to sign, with a reason kept in the audit trail.
func (s *Service) DeclineSignatureRequest(ctx context.Context, id uuid.UUID, reason string) (db.SignatureRequest, error) {
	row, err := s.q(ctx).DeclineSignature(ctx, db.DeclineSignatureParams{ID: id, UpdatedBy: actorID(ctx)})
	if err != nil {
		return row, stateError(err, "decline this")
	}
	return row, s.audit(ctx, "signatures.signature_request.decline", row.SubjectType, row.SubjectID, nil, nil, map[string]string{"reason": reason})
}

// requireSignatures blocks closing a record with signatures outstanding.
func (s *Service) requireSignatures(ctx context.Context, table string, id uuid.UUID) error {
	n, err := s.q(ctx).CountPendingSignatures(ctx, db.CountPendingSignaturesParams{SubjectType: table, SubjectID: id})
	if err != nil {
		return fmt.Errorf("count signatures: %w", err)
	}
	if n > 0 {
		return &Error{Status: http.StatusUnprocessableEntity, Code: "signature_required", Title: "Signatures outstanding",
			Detail:   fmt.Sprintf("%d signature(s) are still needed.", n),
			Blockers: []Blocker{{Kind: "signature_missing", Message: "Collect the outstanding signatures first"}}}
	}
	return nil
}

// voidSignatures cancels signatures after an edit, which changes what was signed.
func (s *Service) voidSignatures(ctx context.Context, table string, id uuid.UUID) error {
	_, err := s.q(ctx).VoidSignatures(ctx, db.VoidSignaturesParams{SubjectType: table, SubjectID: id})
	return wrap(err, "void signatures")
}
