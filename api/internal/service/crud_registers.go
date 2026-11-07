// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListAssetClass returns a page of asset classes.
func (s *Service) ListAssetClass(ctx context.Context, p db.ListAssetClassParams) ([]db.AssetClass, error) {
	return wrap2(s.q(ctx).ListAssetClass(ctx, p))("list asset classes")
}

// GetAssetClass returns one row of asset classes.
func (s *Service) GetAssetClass(ctx context.Context, id uuid.UUID) (db.AssetClass, error) {
	row, err := s.q(ctx).GetAssetClass(ctx, id)
	if err != nil {
		return row, notFound(err, "asset classes")
	}
	return row, nil
}

// CreateAssetClass adds one row to asset classes.
func (s *Service) CreateAssetClass(ctx context.Context, p db.CreateAssetClassParams) (db.AssetClass, error) {
	var zero db.AssetClass
	if err := s.guard(ctx, "CreateAssetClass", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateAssetClass(ctx, p)
	if err != nil {
		return zero, wrap(err, "create asset classes")
	}
	if err := s.audit(ctx, "assets.asset_class.write", "asset_classes", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateAssetClass", row)
}

// UpdateAssetClass changes one row of asset classes.
func (s *Service) UpdateAssetClass(ctx context.Context, p db.UpdateAssetClassParams) (db.AssetClass, error) {
	before, err := s.GetAssetClass(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateAssetClass(ctx, p)
	if err != nil {
		return before, wrap(err, "update asset classes")
	}
	if err := s.voidSignatures(ctx, "asset_classes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "assets.asset_class.write", "asset_classes", row.ID, nil, before, row)
}

// ListAsset returns a page of assets.
func (s *Service) ListAsset(ctx context.Context, p db.ListAssetParams) ([]db.Asset, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListAsset(ctx, p))("list assets")
}

// GetAsset returns one row of assets.
func (s *Service) GetAsset(ctx context.Context, id uuid.UUID) (db.Asset, error) {
	row, err := s.q(ctx).GetAsset(ctx, id)
	if err != nil {
		return row, notFound(err, "assets")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateAsset adds one row to assets.
func (s *Service) CreateAsset(ctx context.Context, p db.CreateAssetParams) (db.Asset, error) {
	var zero db.Asset
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateAsset", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateAsset(ctx, p)
	if err != nil {
		return zero, wrap(err, "create assets")
	}
	if err := s.audit(ctx, "assets.asset.write", "assets", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateAsset", row)
}

// UpdateAsset changes one row of assets.
func (s *Service) UpdateAsset(ctx context.Context, p db.UpdateAssetParams) (db.Asset, error) {
	before, err := s.GetAsset(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateAsset(ctx, p)
	if err != nil {
		return before, wrap(err, "update assets")
	}
	if err := s.voidSignatures(ctx, "assets", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "assets.asset.write", "assets", row.ID, &row.SiteID, before, row)
}

// ListCapexRequest returns a page of capex requests.
func (s *Service) ListCapexRequest(ctx context.Context, p db.ListCapexRequestParams) ([]db.CapexRequest, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListCapexRequest(ctx, p))("list capex requests")
}

// GetCapexRequest returns one row of capex requests.
func (s *Service) GetCapexRequest(ctx context.Context, id uuid.UUID) (db.CapexRequest, error) {
	row, err := s.q(ctx).GetCapexRequest(ctx, id)
	if err != nil {
		return row, notFound(err, "capex requests")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateCapexRequest adds one row to capex requests.
func (s *Service) CreateCapexRequest(ctx context.Context, p db.CreateCapexRequestParams) (db.CapexRequest, error) {
	var zero db.CapexRequest
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateCapexRequest", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "CPX", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCapexRequest(ctx, p)
	if err != nil {
		return zero, wrap(err, "create capex requests")
	}
	if err := s.audit(ctx, "assets.capex_request.write", "capex_requests", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateCapexRequest", row)
}

// UpdateCapexRequest changes one row of capex requests.
func (s *Service) UpdateCapexRequest(ctx context.Context, p db.UpdateCapexRequestParams) (db.CapexRequest, error) {
	before, err := s.GetCapexRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCapexRequest(ctx, p)
	if err != nil {
		return before, wrap(err, "update capex requests")
	}
	if err := s.voidSignatures(ctx, "capex_requests", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "assets.capex_request.write", "capex_requests", row.ID, &row.SiteID, before, row)
}

// SubmitCapexRequest is the submit action on capex requests.
func (s *Service) SubmitCapexRequest(ctx context.Context, p db.ActSubmitCapexRequestParams) (db.CapexRequest, error) {
	before, err := s.GetCapexRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "assets.capex_request.submit", "capex_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitCapexRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitCapexRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "submit capex requests")
	}
	if err := s.audit(ctx, "assets.capex_request.submit", "capex_requests", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "capex_requests", row.ID, &row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitCapexRequest", row)
}

// ListAssetTransfer returns a page of asset transfers.
func (s *Service) ListAssetTransfer(ctx context.Context, p db.ListAssetTransferParams) ([]db.AssetTransfer, error) {
	return wrap2(s.q(ctx).ListAssetTransfer(ctx, p))("list asset transfers")
}

// GetAssetTransfer returns one row of asset transfers.
func (s *Service) GetAssetTransfer(ctx context.Context, id uuid.UUID) (db.AssetTransfer, error) {
	row, err := s.q(ctx).GetAssetTransfer(ctx, id)
	if err != nil {
		return row, notFound(err, "asset transfers")
	}
	return row, nil
}

// CreateAssetTransfer adds one row to asset transfers.
func (s *Service) CreateAssetTransfer(ctx context.Context, p db.CreateAssetTransferParams) (db.AssetTransfer, error) {
	var zero db.AssetTransfer
	if err := s.guard(ctx, "CreateAssetTransfer", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "ATR", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateAssetTransfer(ctx, p)
	if err != nil {
		return zero, wrap(err, "create asset transfers")
	}
	if err := s.audit(ctx, "assets.asset_transfer.write", "asset_transfers", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateAssetTransfer", row)
}

// SubmitAssetTransfer is the submit action on asset transfers.
func (s *Service) SubmitAssetTransfer(ctx context.Context, p db.ActSubmitAssetTransferParams) (db.AssetTransfer, error) {
	before, err := s.GetAssetTransfer(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "assets.asset_transfer.submit", "asset_transfers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitAssetTransfer", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitAssetTransfer(ctx, p)
	if err != nil {
		return before, stateError(err, "submit asset transfers")
	}
	if err := s.audit(ctx, "assets.asset_transfer.submit", "asset_transfers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "asset_transfers", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitAssetTransfer", row)
}

// ReceiveAssetTransfer is the acknowledge receipt action on asset transfers.
func (s *Service) ReceiveAssetTransfer(ctx context.Context, p db.ActReceiveAssetTransferParams) (db.AssetTransfer, error) {
	before, err := s.GetAssetTransfer(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "assets.asset_transfer.receive", "asset_transfers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReceiveAssetTransfer", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReceiveAssetTransfer(ctx, p)
	if err != nil {
		return before, stateError(err, "receive asset transfers")
	}
	if err := s.audit(ctx, "assets.asset_transfer.receive", "asset_transfers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReceiveAssetTransfer", row)
}

// ListAssetVerification returns a page of asset verifications.
func (s *Service) ListAssetVerification(ctx context.Context, p db.ListAssetVerificationParams) ([]db.AssetVerification, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListAssetVerification(ctx, p))("list asset verifications")
}

// GetAssetVerification returns one row of asset verifications.
func (s *Service) GetAssetVerification(ctx context.Context, id uuid.UUID) (db.AssetVerification, error) {
	row, err := s.q(ctx).GetAssetVerification(ctx, id)
	if err != nil {
		return row, notFound(err, "asset verifications")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateAssetVerification adds one row to asset verifications.
func (s *Service) CreateAssetVerification(ctx context.Context, p db.CreateAssetVerificationParams) (db.AssetVerification, error) {
	var zero db.AssetVerification
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateAssetVerification", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "AVR", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateAssetVerification(ctx, p)
	if err != nil {
		return zero, wrap(err, "create asset verifications")
	}
	if err := s.audit(ctx, "assets.asset_verification.write", "asset_verifications", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateAssetVerification", row)
}

// UpdateAssetVerification changes one row of asset verifications.
func (s *Service) UpdateAssetVerification(ctx context.Context, p db.UpdateAssetVerificationParams) (db.AssetVerification, error) {
	before, err := s.GetAssetVerification(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateAssetVerification(ctx, p)
	if err != nil {
		return before, wrap(err, "update asset verifications")
	}
	if err := s.voidSignatures(ctx, "asset_verifications", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "assets.asset_verification.write", "asset_verifications", row.ID, &row.SiteID, before, row)
}

// SubmitAssetVerification is the submit action on asset verifications.
func (s *Service) SubmitAssetVerification(ctx context.Context, p db.ActSubmitAssetVerificationParams) (db.AssetVerification, error) {
	before, err := s.GetAssetVerification(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "assets.asset_verification.submit", "asset_verifications", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitAssetVerification", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitAssetVerification(ctx, p)
	if err != nil {
		return before, stateError(err, "submit asset verifications")
	}
	if err := s.audit(ctx, "assets.asset_verification.submit", "asset_verifications", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitAssetVerification", row)
}

// ReviewAssetVerification is the review action on asset verifications.
func (s *Service) ReviewAssetVerification(ctx context.Context, p db.ActReviewAssetVerificationParams) (db.AssetVerification, error) {
	before, err := s.GetAssetVerification(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "assets.asset_verification.review", "asset_verifications", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReviewAssetVerification", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "asset_verifications", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActReviewAssetVerification(ctx, p)
	if err != nil {
		return before, stateError(err, "review asset verifications")
	}
	if err := s.audit(ctx, "assets.asset_verification.review", "asset_verifications", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReviewAssetVerification", row)
}

// AssetVerificationChildren returns the rows of a record's child tables, by API field.
func (s *Service) AssetVerificationChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListAssetVerificationLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list asset_verification_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplaceAssetVerificationLines swaps a record's asset verification lines for the given rows.
func (s *Service) ReplaceAssetVerificationLines(ctx context.Context, id uuid.UUID, rows []db.InsertAssetVerificationLinesParams) error {
	if err := s.q(ctx).DeleteAssetVerificationLines(ctx, id); err != nil {
		return wrap(err, "clear asset_verification_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertAssetVerificationLines(ctx, row); err != nil {
			return wrap(err, "add asset_verification_lines")
		}
	}
	return nil
}

// ListAssetDisposal returns a page of asset disposals.
func (s *Service) ListAssetDisposal(ctx context.Context, p db.ListAssetDisposalParams) ([]db.AssetDisposal, error) {
	return wrap2(s.q(ctx).ListAssetDisposal(ctx, p))("list asset disposals")
}

// GetAssetDisposal returns one row of asset disposals.
func (s *Service) GetAssetDisposal(ctx context.Context, id uuid.UUID) (db.AssetDisposal, error) {
	row, err := s.q(ctx).GetAssetDisposal(ctx, id)
	if err != nil {
		return row, notFound(err, "asset disposals")
	}
	return row, nil
}

// CreateAssetDisposal adds one row to asset disposals.
func (s *Service) CreateAssetDisposal(ctx context.Context, p db.CreateAssetDisposalParams) (db.AssetDisposal, error) {
	var zero db.AssetDisposal
	if err := s.guard(ctx, "CreateAssetDisposal", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "ADS", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateAssetDisposal(ctx, p)
	if err != nil {
		return zero, wrap(err, "create asset disposals")
	}
	if err := s.audit(ctx, "assets.asset_disposal.write", "asset_disposals", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateAssetDisposal", row)
}

// UpdateAssetDisposal changes one row of asset disposals.
func (s *Service) UpdateAssetDisposal(ctx context.Context, p db.UpdateAssetDisposalParams) (db.AssetDisposal, error) {
	before, err := s.GetAssetDisposal(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateAssetDisposal(ctx, p)
	if err != nil {
		return before, wrap(err, "update asset disposals")
	}
	if err := s.voidSignatures(ctx, "asset_disposals", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "assets.asset_disposal.write", "asset_disposals", row.ID, nil, before, row)
}

// SubmitAssetDisposal is the submit action on asset disposals.
func (s *Service) SubmitAssetDisposal(ctx context.Context, p db.ActSubmitAssetDisposalParams) (db.AssetDisposal, error) {
	before, err := s.GetAssetDisposal(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "assets.asset_disposal.submit", "asset_disposals", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitAssetDisposal", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitAssetDisposal(ctx, p)
	if err != nil {
		return before, stateError(err, "submit asset disposals")
	}
	if err := s.audit(ctx, "assets.asset_disposal.submit", "asset_disposals", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "asset_disposals", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitAssetDisposal", row)
}

// ConfirmRemovalAssetDisposal is the confirm physical removal action on asset disposals.
func (s *Service) ConfirmRemovalAssetDisposal(ctx context.Context, p db.ActConfirmRemovalAssetDisposalParams) (db.AssetDisposal, error) {
	before, err := s.GetAssetDisposal(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "assets.asset_disposal.confirm_removal", "asset_disposals", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ConfirmRemovalAssetDisposal", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActConfirmRemovalAssetDisposal(ctx, p)
	if err != nil {
		return before, stateError(err, "confirm removal asset disposals")
	}
	if err := s.audit(ctx, "assets.asset_disposal.confirm_removal", "asset_disposals", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ConfirmRemovalAssetDisposal", row)
}

// ListKeyItem returns a page of key items.
func (s *Service) ListKeyItem(ctx context.Context, p db.ListKeyItemParams) ([]db.KeyItem, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListKeyItem(ctx, p))("list key items")
}

// GetKeyItem returns one row of key items.
func (s *Service) GetKeyItem(ctx context.Context, id uuid.UUID) (db.KeyItem, error) {
	row, err := s.q(ctx).GetKeyItem(ctx, id)
	if err != nil {
		return row, notFound(err, "key items")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateKeyItem adds one row to key items.
func (s *Service) CreateKeyItem(ctx context.Context, p db.CreateKeyItemParams) (db.KeyItem, error) {
	var zero db.KeyItem
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateKeyItem", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateKeyItem(ctx, p)
	if err != nil {
		return zero, wrap(err, "create key items")
	}
	if err := s.audit(ctx, "keys.key_item.write", "key_items", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateKeyItem", row)
}

// UpdateKeyItem changes one row of key items.
func (s *Service) UpdateKeyItem(ctx context.Context, p db.UpdateKeyItemParams) (db.KeyItem, error) {
	before, err := s.GetKeyItem(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateKeyItem(ctx, p)
	if err != nil {
		return before, wrap(err, "update key items")
	}
	if err := s.voidSignatures(ctx, "key_items", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "keys.key_item.write", "key_items", row.ID, &row.SiteID, before, row)
}

// ListKeyMovement returns a page of key movements.
func (s *Service) ListKeyMovement(ctx context.Context, p db.ListKeyMovementParams) ([]db.KeyMovement, error) {
	return wrap2(s.q(ctx).ListKeyMovement(ctx, p))("list key movements")
}

// GetKeyMovement returns one row of key movements.
func (s *Service) GetKeyMovement(ctx context.Context, id uuid.UUID) (db.KeyMovement, error) {
	row, err := s.q(ctx).GetKeyMovement(ctx, id)
	if err != nil {
		return row, notFound(err, "key movements")
	}
	return row, nil
}

// CreateKeyMovement adds one row to key movements.
func (s *Service) CreateKeyMovement(ctx context.Context, p db.CreateKeyMovementParams) (db.KeyMovement, error) {
	var zero db.KeyMovement
	if err := s.guard(ctx, "CreateKeyMovement", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateKeyMovement(ctx, p)
	if err != nil {
		return zero, wrap(err, "create key movements")
	}
	if err := s.audit(ctx, "keys.key_movement.write", "key_movements", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateKeyMovement", row)
}

// ReturnKeyMovement is the sign back in action on key movements.
func (s *Service) ReturnKeyMovement(ctx context.Context, p db.ActReturnKeyMovementParams) (db.KeyMovement, error) {
	before, err := s.GetKeyMovement(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "keys.key_movement.return", "key_movements", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReturnKeyMovement", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReturnKeyMovement(ctx, p)
	if err != nil {
		return before, stateError(err, "return key movements")
	}
	if err := s.audit(ctx, "keys.key_movement.return", "key_movements", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReturnKeyMovement", row)
}

// ListKeyDuplicationRequest returns a page of key duplication requests.
func (s *Service) ListKeyDuplicationRequest(ctx context.Context, p db.ListKeyDuplicationRequestParams) ([]db.KeyDuplicationRequest, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListKeyDuplicationRequest(ctx, p))("list key duplication requests")
}

// GetKeyDuplicationRequest returns one row of key duplication requests.
func (s *Service) GetKeyDuplicationRequest(ctx context.Context, id uuid.UUID) (db.KeyDuplicationRequest, error) {
	row, err := s.q(ctx).GetKeyDuplicationRequest(ctx, id)
	if err != nil {
		return row, notFound(err, "key duplication requests")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateKeyDuplicationRequest adds one row to key duplication requests.
func (s *Service) CreateKeyDuplicationRequest(ctx context.Context, p db.CreateKeyDuplicationRequestParams) (db.KeyDuplicationRequest, error) {
	var zero db.KeyDuplicationRequest
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateKeyDuplicationRequest", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "KDR", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateKeyDuplicationRequest(ctx, p)
	if err != nil {
		return zero, wrap(err, "create key duplication requests")
	}
	if err := s.audit(ctx, "keys.key_duplication_request.write", "key_duplication_requests", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateKeyDuplicationRequest", row)
}

// UpdateKeyDuplicationRequest changes one row of key duplication requests.
func (s *Service) UpdateKeyDuplicationRequest(ctx context.Context, p db.UpdateKeyDuplicationRequestParams) (db.KeyDuplicationRequest, error) {
	before, err := s.GetKeyDuplicationRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateKeyDuplicationRequest(ctx, p)
	if err != nil {
		return before, wrap(err, "update key duplication requests")
	}
	if err := s.voidSignatures(ctx, "key_duplication_requests", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "keys.key_duplication_request.write", "key_duplication_requests", row.ID, &row.SiteID, before, row)
}

// SubmitKeyDuplicationRequest is the submit action on key duplication requests.
func (s *Service) SubmitKeyDuplicationRequest(ctx context.Context, p db.ActSubmitKeyDuplicationRequestParams) (db.KeyDuplicationRequest, error) {
	before, err := s.GetKeyDuplicationRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "keys.key_duplication_request.submit", "key_duplication_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitKeyDuplicationRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitKeyDuplicationRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "submit key duplication requests")
	}
	if err := s.audit(ctx, "keys.key_duplication_request.submit", "key_duplication_requests", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "key_duplication_requests", row.ID, &row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitKeyDuplicationRequest", row)
}

// ListKeyAudit returns a page of key audits.
func (s *Service) ListKeyAudit(ctx context.Context, p db.ListKeyAuditParams) ([]db.KeyAudit, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListKeyAudit(ctx, p))("list key audits")
}

// GetKeyAudit returns one row of key audits.
func (s *Service) GetKeyAudit(ctx context.Context, id uuid.UUID) (db.KeyAudit, error) {
	row, err := s.q(ctx).GetKeyAudit(ctx, id)
	if err != nil {
		return row, notFound(err, "key audits")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateKeyAudit adds one row to key audits.
func (s *Service) CreateKeyAudit(ctx context.Context, p db.CreateKeyAuditParams) (db.KeyAudit, error) {
	var zero db.KeyAudit
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateKeyAudit", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "KAU", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateKeyAudit(ctx, p)
	if err != nil {
		return zero, wrap(err, "create key audits")
	}
	if err := s.audit(ctx, "keys.key_audit.write", "key_audits", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateKeyAudit", row)
}

// UpdateKeyAudit changes one row of key audits.
func (s *Service) UpdateKeyAudit(ctx context.Context, p db.UpdateKeyAuditParams) (db.KeyAudit, error) {
	before, err := s.GetKeyAudit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateKeyAudit(ctx, p)
	if err != nil {
		return before, wrap(err, "update key audits")
	}
	if err := s.voidSignatures(ctx, "key_audits", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "keys.key_audit.write", "key_audits", row.ID, &row.SiteID, before, row)
}

// CompleteKeyAudit is the sign off action on key audits.
func (s *Service) CompleteKeyAudit(ctx context.Context, p db.ActCompleteKeyAuditParams) (db.KeyAudit, error) {
	before, err := s.GetKeyAudit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "keys.key_audit.complete", "key_audits", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CompleteKeyAudit", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "key_audits", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompleteKeyAudit(ctx, p)
	if err != nil {
		return before, stateError(err, "complete key audits")
	}
	if err := s.audit(ctx, "keys.key_audit.complete", "key_audits", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CompleteKeyAudit", row)
}

// KeyAuditChildren returns the rows of a record's child tables, by API field.
func (s *Service) KeyAuditChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListKeyAuditLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list key_audit_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplaceKeyAuditLines swaps a record's key audit lines for the given rows.
func (s *Service) ReplaceKeyAuditLines(ctx context.Context, id uuid.UUID, rows []db.InsertKeyAuditLinesParams) error {
	if err := s.q(ctx).DeleteKeyAuditLines(ctx, id); err != nil {
		return wrap(err, "clear key_audit_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertKeyAuditLines(ctx, row); err != nil {
			return wrap(err, "add key_audit_lines")
		}
	}
	return nil
}

// ListVisitorEntry returns a page of visitor entries.
func (s *Service) ListVisitorEntry(ctx context.Context, p db.ListVisitorEntryParams) ([]db.VisitorEntry, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListVisitorEntry(ctx, p))("list visitor entries")
}

// GetVisitorEntry returns one row of visitor entries.
func (s *Service) GetVisitorEntry(ctx context.Context, id uuid.UUID) (db.VisitorEntry, error) {
	row, err := s.q(ctx).GetVisitorEntry(ctx, id)
	if err != nil {
		return row, notFound(err, "visitor entries")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateVisitorEntry adds one row to visitor entries.
func (s *Service) CreateVisitorEntry(ctx context.Context, p db.CreateVisitorEntryParams) (db.VisitorEntry, error) {
	var zero db.VisitorEntry
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateVisitorEntry", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateVisitorEntry(ctx, p)
	if err != nil {
		return zero, wrap(err, "create visitor entries")
	}
	if err := s.audit(ctx, "visitors.visitor_entry.write", "visitor_entries", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateVisitorEntry", row)
}

// UpdateVisitorEntry changes one row of visitor entries.
func (s *Service) UpdateVisitorEntry(ctx context.Context, p db.UpdateVisitorEntryParams) (db.VisitorEntry, error) {
	before, err := s.GetVisitorEntry(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateVisitorEntry(ctx, p)
	if err != nil {
		return before, wrap(err, "update visitor entries")
	}
	if err := s.voidSignatures(ctx, "visitor_entries", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "visitors.visitor_entry.write", "visitor_entries", row.ID, &row.SiteID, before, row)
}

// SignOutVisitorEntry is the record departure action on visitor entries.
func (s *Service) SignOutVisitorEntry(ctx context.Context, p db.ActSignOutVisitorEntryParams) (db.VisitorEntry, error) {
	before, err := s.GetVisitorEntry(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "visitors.visitor_entry.sign_out", "visitor_entries", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SignOutVisitorEntry", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSignOutVisitorEntry(ctx, p)
	if err != nil {
		return before, stateError(err, "sign out visitor entries")
	}
	if err := s.audit(ctx, "visitors.visitor_entry.sign_out", "visitor_entries", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SignOutVisitorEntry", row)
}

// ListLostPropertyItem returns a page of lost property items.
func (s *Service) ListLostPropertyItem(ctx context.Context, p db.ListLostPropertyItemParams) ([]db.LostPropertyItem, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListLostPropertyItem(ctx, p))("list lost property items")
}

// GetLostPropertyItem returns one row of lost property items.
func (s *Service) GetLostPropertyItem(ctx context.Context, id uuid.UUID) (db.LostPropertyItem, error) {
	row, err := s.q(ctx).GetLostPropertyItem(ctx, id)
	if err != nil {
		return row, notFound(err, "lost property items")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateLostPropertyItem adds one row to lost property items.
func (s *Service) CreateLostPropertyItem(ctx context.Context, p db.CreateLostPropertyItemParams) (db.LostPropertyItem, error) {
	var zero db.LostPropertyItem
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateLostPropertyItem", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "LPR", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateLostPropertyItem(ctx, p)
	if err != nil {
		return zero, wrap(err, "create lost property items")
	}
	if err := s.audit(ctx, "lost_property.lost_property_item.write", "lost_property_items", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateLostPropertyItem", row)
}

// UpdateLostPropertyItem changes one row of lost property items.
func (s *Service) UpdateLostPropertyItem(ctx context.Context, p db.UpdateLostPropertyItemParams) (db.LostPropertyItem, error) {
	before, err := s.GetLostPropertyItem(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateLostPropertyItem(ctx, p)
	if err != nil {
		return before, wrap(err, "update lost property items")
	}
	if err := s.voidSignatures(ctx, "lost_property_items", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "lost_property.lost_property_item.write", "lost_property_items", row.ID, &row.SiteID, before, row)
}

// ReturnLostPropertyItem is the release to a verified claimant action on lost property items.
func (s *Service) ReturnLostPropertyItem(ctx context.Context, p db.ActReturnLostPropertyItemParams) (db.LostPropertyItem, error) {
	before, err := s.GetLostPropertyItem(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "lost_property.lost_property_item.return", "lost_property_items", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReturnLostPropertyItem", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReturnLostPropertyItem(ctx, p)
	if err != nil {
		return before, stateError(err, "return lost property items")
	}
	if err := s.audit(ctx, "lost_property.lost_property_item.return", "lost_property_items", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReturnLostPropertyItem", row)
}

// DisposeLostPropertyItem is the dispose at the due date action on lost property items.
func (s *Service) DisposeLostPropertyItem(ctx context.Context, p db.ActDisposeLostPropertyItemParams) (db.LostPropertyItem, error) {
	before, err := s.GetLostPropertyItem(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "lost_property.lost_property_item.dispose", "lost_property_items", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "DisposeLostPropertyItem", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActDisposeLostPropertyItem(ctx, p)
	if err != nil {
		return before, stateError(err, "dispose lost property items")
	}
	if err := s.audit(ctx, "lost_property.lost_property_item.dispose", "lost_property_items", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "DisposeLostPropertyItem", row)
}

// ListDocument returns a page of documents.
func (s *Service) ListDocument(ctx context.Context, p db.ListDocumentParams) ([]db.Document, error) {
	return wrap2(s.q(ctx).ListDocument(ctx, p))("list documents")
}

// GetDocument returns one row of documents.
func (s *Service) GetDocument(ctx context.Context, id uuid.UUID) (db.Document, error) {
	row, err := s.q(ctx).GetDocument(ctx, id)
	if err != nil {
		return row, notFound(err, "documents")
	}
	return row, nil
}

// CreateDocument adds one row to documents.
func (s *Service) CreateDocument(ctx context.Context, p db.CreateDocumentParams) (db.Document, error) {
	var zero db.Document
	if err := s.guard(ctx, "CreateDocument", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDocument(ctx, p)
	if err != nil {
		return zero, wrap(err, "create documents")
	}
	if err := s.audit(ctx, "documents.document.write", "documents", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDocument", row)
}

// UpdateDocument changes one row of documents.
func (s *Service) UpdateDocument(ctx context.Context, p db.UpdateDocumentParams) (db.Document, error) {
	before, err := s.GetDocument(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDocument(ctx, p)
	if err != nil {
		return before, wrap(err, "update documents")
	}
	if err := s.voidSignatures(ctx, "documents", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "documents.document.write", "documents", row.ID, nil, before, row)
}

// ListDocumentRevision returns a page of document revisions.
func (s *Service) ListDocumentRevision(ctx context.Context, p db.ListDocumentRevisionParams) ([]db.DocumentRevision, error) {
	return wrap2(s.q(ctx).ListDocumentRevision(ctx, p))("list document revisions")
}

// GetDocumentRevision returns one row of document revisions.
func (s *Service) GetDocumentRevision(ctx context.Context, id uuid.UUID) (db.DocumentRevision, error) {
	row, err := s.q(ctx).GetDocumentRevision(ctx, id)
	if err != nil {
		return row, notFound(err, "document revisions")
	}
	return row, nil
}

// CreateDocumentRevision adds one row to document revisions.
func (s *Service) CreateDocumentRevision(ctx context.Context, p db.CreateDocumentRevisionParams) (db.DocumentRevision, error) {
	var zero db.DocumentRevision
	if err := s.guard(ctx, "CreateDocumentRevision", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDocumentRevision(ctx, p)
	if err != nil {
		return zero, wrap(err, "create document revisions")
	}
	if err := s.audit(ctx, "documents.document_revision.write", "document_revisions", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDocumentRevision", row)
}

// UpdateDocumentRevision changes one row of document revisions.
func (s *Service) UpdateDocumentRevision(ctx context.Context, p db.UpdateDocumentRevisionParams) (db.DocumentRevision, error) {
	before, err := s.GetDocumentRevision(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDocumentRevision(ctx, p)
	if err != nil {
		return before, wrap(err, "update document revisions")
	}
	if err := s.voidSignatures(ctx, "document_revisions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "documents.document_revision.write", "document_revisions", row.ID, nil, before, row)
}

// SubmitDocumentRevision is the submit action on document revisions.
func (s *Service) SubmitDocumentRevision(ctx context.Context, p db.ActSubmitDocumentRevisionParams) (db.DocumentRevision, error) {
	before, err := s.GetDocumentRevision(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "documents.document_revision.submit", "document_revisions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitDocumentRevision", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitDocumentRevision(ctx, p)
	if err != nil {
		return before, stateError(err, "submit document revisions")
	}
	if err := s.audit(ctx, "documents.document_revision.submit", "document_revisions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "document_revisions", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitDocumentRevision", row)
}

// IssueDocumentRevision is the issue and notify affected people action on document revisions.
func (s *Service) IssueDocumentRevision(ctx context.Context, p db.ActIssueDocumentRevisionParams) (db.DocumentRevision, error) {
	before, err := s.GetDocumentRevision(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "documents.document_revision.issue", "document_revisions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "IssueDocumentRevision", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "document_revisions", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActIssueDocumentRevision(ctx, p)
	if err != nil {
		return before, stateError(err, "issue document revisions")
	}
	if err := s.audit(ctx, "documents.document_revision.issue", "document_revisions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "IssueDocumentRevision", row)
}

// WithdrawDocumentRevision is the withdraw action on document revisions.
func (s *Service) WithdrawDocumentRevision(ctx context.Context, p db.ActWithdrawDocumentRevisionParams) (db.DocumentRevision, error) {
	before, err := s.GetDocumentRevision(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "documents.document_revision.withdraw", "document_revisions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "WithdrawDocumentRevision", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActWithdrawDocumentRevision(ctx, p)
	if err != nil {
		return before, stateError(err, "withdraw document revisions")
	}
	if err := s.audit(ctx, "documents.document_revision.withdraw", "document_revisions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "WithdrawDocumentRevision", row)
}

// ListDocumentRequest returns a page of document requests.
func (s *Service) ListDocumentRequest(ctx context.Context, p db.ListDocumentRequestParams) ([]db.DocumentRequest, error) {
	return wrap2(s.q(ctx).ListDocumentRequest(ctx, p))("list document requests")
}

// GetDocumentRequest returns one row of document requests.
func (s *Service) GetDocumentRequest(ctx context.Context, id uuid.UUID) (db.DocumentRequest, error) {
	row, err := s.q(ctx).GetDocumentRequest(ctx, id)
	if err != nil {
		return row, notFound(err, "document requests")
	}
	return row, nil
}

// CreateDocumentRequest adds one row to document requests.
func (s *Service) CreateDocumentRequest(ctx context.Context, p db.CreateDocumentRequestParams) (db.DocumentRequest, error) {
	var zero db.DocumentRequest
	if err := s.guard(ctx, "CreateDocumentRequest", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "DRQ", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDocumentRequest(ctx, p)
	if err != nil {
		return zero, wrap(err, "create document requests")
	}
	if err := s.audit(ctx, "documents.document_request.write", "document_requests", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDocumentRequest", row)
}

// UpdateDocumentRequest changes one row of document requests.
func (s *Service) UpdateDocumentRequest(ctx context.Context, p db.UpdateDocumentRequestParams) (db.DocumentRequest, error) {
	before, err := s.GetDocumentRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDocumentRequest(ctx, p)
	if err != nil {
		return before, wrap(err, "update document requests")
	}
	if err := s.voidSignatures(ctx, "document_requests", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "documents.document_request.write", "document_requests", row.ID, nil, before, row)
}

// CloseDocumentRequest is the close action on document requests.
func (s *Service) CloseDocumentRequest(ctx context.Context, p db.ActCloseDocumentRequestParams) (db.DocumentRequest, error) {
	before, err := s.GetDocumentRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "documents.document_request.close", "document_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseDocumentRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "document_requests", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseDocumentRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "close document requests")
	}
	if err := s.audit(ctx, "documents.document_request.close", "document_requests", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseDocumentRequest", row)
}

// ListDocumentAcknowledgement returns a page of document acknowledgements.
func (s *Service) ListDocumentAcknowledgement(ctx context.Context, p db.ListDocumentAcknowledgementParams) ([]db.DocumentAcknowledgement, error) {
	return wrap2(s.q(ctx).ListDocumentAcknowledgement(ctx, p))("list document acknowledgements")
}

// GetDocumentAcknowledgement returns one row of document acknowledgements.
func (s *Service) GetDocumentAcknowledgement(ctx context.Context, id uuid.UUID) (db.DocumentAcknowledgement, error) {
	row, err := s.q(ctx).GetDocumentAcknowledgement(ctx, id)
	if err != nil {
		return row, notFound(err, "document acknowledgements")
	}
	return row, nil
}

// AcknowledgeDocumentAcknowledgement is the acknowledge action on document acknowledgements.
func (s *Service) AcknowledgeDocumentAcknowledgement(ctx context.Context, p db.ActAcknowledgeDocumentAcknowledgementParams) (db.DocumentAcknowledgement, error) {
	before, err := s.GetDocumentAcknowledgement(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "documents.document_acknowledgement.acknowledge", "document_acknowledgements", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "AcknowledgeDocumentAcknowledgement", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAcknowledgeDocumentAcknowledgement(ctx, p)
	if err != nil {
		return before, stateError(err, "acknowledge document acknowledgements")
	}
	if err := s.audit(ctx, "documents.document_acknowledgement.acknowledge", "document_acknowledgements", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "AcknowledgeDocumentAcknowledgement", row)
}

// ListCourse returns a page of courses.
func (s *Service) ListCourse(ctx context.Context, p db.ListCourseParams) ([]db.Course, error) {
	return wrap2(s.q(ctx).ListCourse(ctx, p))("list courses")
}

// GetCourse returns one row of courses.
func (s *Service) GetCourse(ctx context.Context, id uuid.UUID) (db.Course, error) {
	row, err := s.q(ctx).GetCourse(ctx, id)
	if err != nil {
		return row, notFound(err, "courses")
	}
	return row, nil
}

// CreateCourse adds one row to courses.
func (s *Service) CreateCourse(ctx context.Context, p db.CreateCourseParams) (db.Course, error) {
	var zero db.Course
	if err := s.guard(ctx, "CreateCourse", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCourse(ctx, p)
	if err != nil {
		return zero, wrap(err, "create courses")
	}
	if err := s.audit(ctx, "training.course.write", "courses", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateCourse", row)
}

// UpdateCourse changes one row of courses.
func (s *Service) UpdateCourse(ctx context.Context, p db.UpdateCourseParams) (db.Course, error) {
	before, err := s.GetCourse(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCourse(ctx, p)
	if err != nil {
		return before, wrap(err, "update courses")
	}
	if err := s.voidSignatures(ctx, "courses", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "training.course.write", "courses", row.ID, nil, before, row)
}

// ListTrainingRequirement returns a page of training requirements.
func (s *Service) ListTrainingRequirement(ctx context.Context, p db.ListTrainingRequirementParams) ([]db.TrainingRequirement, error) {
	return wrap2(s.q(ctx).ListTrainingRequirement(ctx, p))("list training requirements")
}

// GetTrainingRequirement returns one row of training requirements.
func (s *Service) GetTrainingRequirement(ctx context.Context, id uuid.UUID) (db.TrainingRequirement, error) {
	row, err := s.q(ctx).GetTrainingRequirement(ctx, id)
	if err != nil {
		return row, notFound(err, "training requirements")
	}
	return row, nil
}

// CreateTrainingRequirement adds one row to training requirements.
func (s *Service) CreateTrainingRequirement(ctx context.Context, p db.CreateTrainingRequirementParams) (db.TrainingRequirement, error) {
	var zero db.TrainingRequirement
	if err := s.guard(ctx, "CreateTrainingRequirement", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateTrainingRequirement(ctx, p)
	if err != nil {
		return zero, wrap(err, "create training requirements")
	}
	if err := s.audit(ctx, "training.training_requirement.write", "training_requirements", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateTrainingRequirement", row)
}

// UpdateTrainingRequirement changes one row of training requirements.
func (s *Service) UpdateTrainingRequirement(ctx context.Context, p db.UpdateTrainingRequirementParams) (db.TrainingRequirement, error) {
	before, err := s.GetTrainingRequirement(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateTrainingRequirement(ctx, p)
	if err != nil {
		return before, wrap(err, "update training requirements")
	}
	if err := s.voidSignatures(ctx, "training_requirements", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "training.training_requirement.write", "training_requirements", row.ID, nil, before, row)
}

// DeleteTrainingRequirement removes one row of training requirements.
func (s *Service) DeleteTrainingRequirement(ctx context.Context, id uuid.UUID) error {
	before, err := s.GetTrainingRequirement(ctx, id)
	if err != nil {
		return err
	}
	if err := s.q(ctx).DeleteTrainingRequirement(ctx, id); err != nil {
		return wrap(err, "delete training requirements")
	}
	return s.audit(ctx, "training_requirements.delete", "training_requirements", id, nil, before, nil)
}

// ListTrainingSession returns a page of training sessions.
func (s *Service) ListTrainingSession(ctx context.Context, p db.ListTrainingSessionParams) ([]db.TrainingSession, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListTrainingSession(ctx, p))("list training sessions")
}

// GetTrainingSession returns one row of training sessions.
func (s *Service) GetTrainingSession(ctx context.Context, id uuid.UUID) (db.TrainingSession, error) {
	row, err := s.q(ctx).GetTrainingSession(ctx, id)
	if err != nil {
		return row, notFound(err, "training sessions")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateTrainingSession adds one row to training sessions.
func (s *Service) CreateTrainingSession(ctx context.Context, p db.CreateTrainingSessionParams) (db.TrainingSession, error) {
	var zero db.TrainingSession
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateTrainingSession", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "TRN", p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateTrainingSession(ctx, p)
	if err != nil {
		return zero, wrap(err, "create training sessions")
	}
	if err := s.audit(ctx, "training.training_session.write", "training_sessions", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateTrainingSession", row)
}

// UpdateTrainingSession changes one row of training sessions.
func (s *Service) UpdateTrainingSession(ctx context.Context, p db.UpdateTrainingSessionParams) (db.TrainingSession, error) {
	before, err := s.GetTrainingSession(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateTrainingSession(ctx, p)
	if err != nil {
		return before, wrap(err, "update training sessions")
	}
	if err := s.voidSignatures(ctx, "training_sessions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "training.training_session.write", "training_sessions", row.ID, row.SiteID, before, row)
}

// CloseTrainingSession is the close attendance action on training sessions.
func (s *Service) CloseTrainingSession(ctx context.Context, p db.ActCloseTrainingSessionParams) (db.TrainingSession, error) {
	before, err := s.GetTrainingSession(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "training.training_session.close", "training_sessions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseTrainingSession", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "training_sessions", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseTrainingSession(ctx, p)
	if err != nil {
		return before, stateError(err, "close training sessions")
	}
	if err := s.audit(ctx, "training.training_session.close", "training_sessions", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseTrainingSession", row)
}

// TrainingSessionChildren returns the rows of a record's child tables, by API field.
func (s *Service) TrainingSessionChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListTrainingParticipants(ctx, id)
		if err != nil {
			return nil, wrap(err, "list training_participants")
		}
		out["participants"] = childRows(rows)
	}
	return out, nil
}

// ReplaceTrainingParticipants swaps a record's training participants for the given rows.
func (s *Service) ReplaceTrainingParticipants(ctx context.Context, id uuid.UUID, rows []db.InsertTrainingParticipantsParams) error {
	if err := s.q(ctx).DeleteTrainingParticipants(ctx, id); err != nil {
		return wrap(err, "clear training_participants")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertTrainingParticipants(ctx, row); err != nil {
			return wrap(err, "add training_participants")
		}
	}
	return nil
}

// ListTrainingAssessment returns a page of training assessments.
func (s *Service) ListTrainingAssessment(ctx context.Context, p db.ListTrainingAssessmentParams) ([]db.TrainingAssessment, error) {
	return wrap2(s.q(ctx).ListTrainingAssessment(ctx, p))("list training assessments")
}

// GetTrainingAssessment returns one row of training assessments.
func (s *Service) GetTrainingAssessment(ctx context.Context, id uuid.UUID) (db.TrainingAssessment, error) {
	row, err := s.q(ctx).GetTrainingAssessment(ctx, id)
	if err != nil {
		return row, notFound(err, "training assessments")
	}
	return row, nil
}

// CreateTrainingAssessment adds one row to training assessments.
func (s *Service) CreateTrainingAssessment(ctx context.Context, p db.CreateTrainingAssessmentParams) (db.TrainingAssessment, error) {
	var zero db.TrainingAssessment
	if err := s.guard(ctx, "CreateTrainingAssessment", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateTrainingAssessment(ctx, p)
	if err != nil {
		return zero, wrap(err, "create training assessments")
	}
	if err := s.audit(ctx, "training.training_assessment.write", "training_assessments", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateTrainingAssessment", row)
}

// ListCompetency returns a page of competencies.
func (s *Service) ListCompetency(ctx context.Context, p db.ListCompetencyParams) ([]db.Competency, error) {
	return wrap2(s.q(ctx).ListCompetency(ctx, p))("list competencies")
}

// GetCompetency returns one row of competencies.
func (s *Service) GetCompetency(ctx context.Context, id uuid.UUID) (db.Competency, error) {
	row, err := s.q(ctx).GetCompetency(ctx, id)
	if err != nil {
		return row, notFound(err, "competencies")
	}
	return row, nil
}

// CreateCompetency adds one row to competencies.
func (s *Service) CreateCompetency(ctx context.Context, p db.CreateCompetencyParams) (db.Competency, error) {
	var zero db.Competency
	if err := s.guard(ctx, "CreateCompetency", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCompetency(ctx, p)
	if err != nil {
		return zero, wrap(err, "create competencies")
	}
	if err := s.audit(ctx, "training.competency.write", "competencies", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateCompetency", row)
}

// UpdateCompetency changes one row of competencies.
func (s *Service) UpdateCompetency(ctx context.Context, p db.UpdateCompetencyParams) (db.Competency, error) {
	before, err := s.GetCompetency(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCompetency(ctx, p)
	if err != nil {
		return before, wrap(err, "update competencies")
	}
	if err := s.voidSignatures(ctx, "competencies", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "training.competency.write", "competencies", row.ID, nil, before, row)
}

// ListCertificate returns a page of certificates.
func (s *Service) ListCertificate(ctx context.Context, p db.ListCertificateParams) ([]db.Certificate, error) {
	return wrap2(s.q(ctx).ListCertificate(ctx, p))("list certificates")
}

// GetCertificate returns one row of certificates.
func (s *Service) GetCertificate(ctx context.Context, id uuid.UUID) (db.Certificate, error) {
	row, err := s.q(ctx).GetCertificate(ctx, id)
	if err != nil {
		return row, notFound(err, "certificates")
	}
	return row, nil
}

// CreateCertificate adds one row to certificates.
func (s *Service) CreateCertificate(ctx context.Context, p db.CreateCertificateParams) (db.Certificate, error) {
	var zero db.Certificate
	if err := s.guard(ctx, "CreateCertificate", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCertificate(ctx, p)
	if err != nil {
		return zero, wrap(err, "create certificates")
	}
	if err := s.audit(ctx, "training.certificate.write", "certificates", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateCertificate", row)
}

// UpdateCertificate changes one row of certificates.
func (s *Service) UpdateCertificate(ctx context.Context, p db.UpdateCertificateParams) (db.Certificate, error) {
	before, err := s.GetCertificate(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCertificate(ctx, p)
	if err != nil {
		return before, wrap(err, "update certificates")
	}
	if err := s.voidSignatures(ctx, "certificates", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "training.certificate.write", "certificates", row.ID, nil, before, row)
}

// ListTrainingRequest returns a page of training requests.
func (s *Service) ListTrainingRequest(ctx context.Context, p db.ListTrainingRequestParams) ([]db.TrainingRequest, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListTrainingRequest(ctx, p))("list training requests")
}

// GetTrainingRequest returns one row of training requests.
func (s *Service) GetTrainingRequest(ctx context.Context, id uuid.UUID) (db.TrainingRequest, error) {
	row, err := s.q(ctx).GetTrainingRequest(ctx, id)
	if err != nil {
		return row, notFound(err, "training requests")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateTrainingRequest adds one row to training requests.
func (s *Service) CreateTrainingRequest(ctx context.Context, p db.CreateTrainingRequestParams) (db.TrainingRequest, error) {
	var zero db.TrainingRequest
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateTrainingRequest", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "TRQ", p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateTrainingRequest(ctx, p)
	if err != nil {
		return zero, wrap(err, "create training requests")
	}
	if err := s.audit(ctx, "training.training_request.write", "training_requests", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateTrainingRequest", row)
}

// UpdateTrainingRequest changes one row of training requests.
func (s *Service) UpdateTrainingRequest(ctx context.Context, p db.UpdateTrainingRequestParams) (db.TrainingRequest, error) {
	before, err := s.GetTrainingRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateTrainingRequest(ctx, p)
	if err != nil {
		return before, wrap(err, "update training requests")
	}
	if err := s.voidSignatures(ctx, "training_requests", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "training.training_request.write", "training_requests", row.ID, row.SiteID, before, row)
}

// SubmitTrainingRequest is the submit action on training requests.
func (s *Service) SubmitTrainingRequest(ctx context.Context, p db.ActSubmitTrainingRequestParams) (db.TrainingRequest, error) {
	before, err := s.GetTrainingRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "training.training_request.submit", "training_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitTrainingRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitTrainingRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "submit training requests")
	}
	if err := s.audit(ctx, "training.training_request.submit", "training_requests", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "training_requests", row.ID, row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitTrainingRequest", row)
}

// ListLeaveType returns a page of leave types.
func (s *Service) ListLeaveType(ctx context.Context, p db.ListLeaveTypeParams) ([]db.LeaveType, error) {
	return wrap2(s.q(ctx).ListLeaveType(ctx, p))("list leave types")
}

// GetLeaveType returns one row of leave types.
func (s *Service) GetLeaveType(ctx context.Context, id uuid.UUID) (db.LeaveType, error) {
	row, err := s.q(ctx).GetLeaveType(ctx, id)
	if err != nil {
		return row, notFound(err, "leave types")
	}
	return row, nil
}

// CreateLeaveType adds one row to leave types.
func (s *Service) CreateLeaveType(ctx context.Context, p db.CreateLeaveTypeParams) (db.LeaveType, error) {
	var zero db.LeaveType
	if err := s.guard(ctx, "CreateLeaveType", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateLeaveType(ctx, p)
	if err != nil {
		return zero, wrap(err, "create leave types")
	}
	if err := s.audit(ctx, "hr.leave_type.write", "leave_types", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateLeaveType", row)
}

// UpdateLeaveType changes one row of leave types.
func (s *Service) UpdateLeaveType(ctx context.Context, p db.UpdateLeaveTypeParams) (db.LeaveType, error) {
	before, err := s.GetLeaveType(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateLeaveType(ctx, p)
	if err != nil {
		return before, wrap(err, "update leave types")
	}
	if err := s.voidSignatures(ctx, "leave_types", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.leave_type.write", "leave_types", row.ID, nil, before, row)
}

// ListLeaveBalance returns a page of leave balances.
func (s *Service) ListLeaveBalance(ctx context.Context, p db.ListLeaveBalanceParams) ([]db.LeaveBalance, error) {
	return wrap2(s.q(ctx).ListLeaveBalance(ctx, p))("list leave balances")
}

// GetLeaveBalance returns one row of leave balances.
func (s *Service) GetLeaveBalance(ctx context.Context, id uuid.UUID) (db.LeaveBalance, error) {
	row, err := s.q(ctx).GetLeaveBalance(ctx, id)
	if err != nil {
		return row, notFound(err, "leave balances")
	}
	return row, nil
}

// ListLeaveRequest returns a page of leave requests.
func (s *Service) ListLeaveRequest(ctx context.Context, p db.ListLeaveRequestParams) ([]db.LeaveRequest, error) {
	return wrap2(s.q(ctx).ListLeaveRequest(ctx, p))("list leave requests")
}

// GetLeaveRequest returns one row of leave requests.
func (s *Service) GetLeaveRequest(ctx context.Context, id uuid.UUID) (db.LeaveRequest, error) {
	row, err := s.q(ctx).GetLeaveRequest(ctx, id)
	if err != nil {
		return row, notFound(err, "leave requests")
	}
	return row, nil
}

// CreateLeaveRequest adds one row to leave requests.
func (s *Service) CreateLeaveRequest(ctx context.Context, p db.CreateLeaveRequestParams) (db.LeaveRequest, error) {
	var zero db.LeaveRequest
	if err := s.guard(ctx, "CreateLeaveRequest", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "LVE", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateLeaveRequest(ctx, p)
	if err != nil {
		return zero, wrap(err, "create leave requests")
	}
	if err := s.audit(ctx, "hr.leave_request.write", "leave_requests", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateLeaveRequest", row)
}

// UpdateLeaveRequest changes one row of leave requests.
func (s *Service) UpdateLeaveRequest(ctx context.Context, p db.UpdateLeaveRequestParams) (db.LeaveRequest, error) {
	before, err := s.GetLeaveRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateLeaveRequest(ctx, p)
	if err != nil {
		return before, wrap(err, "update leave requests")
	}
	if err := s.voidSignatures(ctx, "leave_requests", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.leave_request.write", "leave_requests", row.ID, nil, before, row)
}

// SubmitLeaveRequest is the submit action on leave requests.
func (s *Service) SubmitLeaveRequest(ctx context.Context, p db.ActSubmitLeaveRequestParams) (db.LeaveRequest, error) {
	before, err := s.GetLeaveRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "hr.leave_request.submit", "leave_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitLeaveRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitLeaveRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "submit leave requests")
	}
	if err := s.audit(ctx, "hr.leave_request.submit", "leave_requests", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "leave_requests", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitLeaveRequest", row)
}

// CancelLeaveRequest is the cancel action on leave requests.
func (s *Service) CancelLeaveRequest(ctx context.Context, p db.ActCancelLeaveRequestParams) (db.LeaveRequest, error) {
	before, err := s.GetLeaveRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "hr.leave_request.cancel", "leave_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CancelLeaveRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActCancelLeaveRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "cancel leave requests")
	}
	if err := s.audit(ctx, "hr.leave_request.cancel", "leave_requests", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CancelLeaveRequest", row)
}

// ListExitClearance returns a page of exit clearances.
func (s *Service) ListExitClearance(ctx context.Context, p db.ListExitClearanceParams) ([]db.ExitClearance, error) {
	return wrap2(s.q(ctx).ListExitClearance(ctx, p))("list exit clearances")
}

// GetExitClearance returns one row of exit clearances.
func (s *Service) GetExitClearance(ctx context.Context, id uuid.UUID) (db.ExitClearance, error) {
	row, err := s.q(ctx).GetExitClearance(ctx, id)
	if err != nil {
		return row, notFound(err, "exit clearances")
	}
	return row, nil
}

// CreateExitClearance adds one row to exit clearances.
func (s *Service) CreateExitClearance(ctx context.Context, p db.CreateExitClearanceParams) (db.ExitClearance, error) {
	var zero db.ExitClearance
	if err := s.guard(ctx, "CreateExitClearance", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "EXT", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateExitClearance(ctx, p)
	if err != nil {
		return zero, wrap(err, "create exit clearances")
	}
	if err := s.audit(ctx, "hr.exit_clearance.write", "exit_clearances", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateExitClearance", row)
}

// UpdateExitClearance changes one row of exit clearances.
func (s *Service) UpdateExitClearance(ctx context.Context, p db.UpdateExitClearanceParams) (db.ExitClearance, error) {
	before, err := s.GetExitClearance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateExitClearance(ctx, p)
	if err != nil {
		return before, wrap(err, "update exit clearances")
	}
	if err := s.voidSignatures(ctx, "exit_clearances", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.exit_clearance.write", "exit_clearances", row.ID, nil, before, row)
}

// SignSectionExitClearance is the sign off a department section action on exit clearances.
func (s *Service) SignSectionExitClearance(ctx context.Context, p db.ActSignSectionExitClearanceParams) (db.ExitClearance, error) {
	before, err := s.GetExitClearance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "hr.exit_clearance.sign_section", "exit_clearances", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SignSectionExitClearance", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSignSectionExitClearance(ctx, p)
	if err != nil {
		return before, stateError(err, "sign section exit clearances")
	}
	if err := s.audit(ctx, "hr.exit_clearance.sign_section", "exit_clearances", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SignSectionExitClearance", row)
}

// CompleteExitClearance is the hand over to payroll action on exit clearances.
func (s *Service) CompleteExitClearance(ctx context.Context, p db.ActCompleteExitClearanceParams) (db.ExitClearance, error) {
	before, err := s.GetExitClearance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "hr.exit_clearance.complete", "exit_clearances", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CompleteExitClearance", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "exit_clearances", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompleteExitClearance(ctx, p)
	if err != nil {
		return before, stateError(err, "complete exit clearances")
	}
	if err := s.audit(ctx, "hr.exit_clearance.complete", "exit_clearances", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CompleteExitClearance", row)
}

// ExitClearanceChildren returns the rows of a record's child tables, by API field.
func (s *Service) ExitClearanceChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListExitClearanceSections(ctx, id)
		if err != nil {
			return nil, wrap(err, "list exit_clearance_sections")
		}
		out["sections"] = childRows(rows)
	}
	return out, nil
}

// ReplaceExitClearanceSections swaps a record's exit clearance sections for the given rows.
func (s *Service) ReplaceExitClearanceSections(ctx context.Context, id uuid.UUID, rows []db.InsertExitClearanceSectionsParams) error {
	if err := s.q(ctx).DeleteExitClearanceSections(ctx, id); err != nil {
		return wrap(err, "clear exit_clearance_sections")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertExitClearanceSections(ctx, row); err != nil {
			return wrap(err, "add exit_clearance_sections")
		}
	}
	return nil
}

// ListRoster returns a page of rosters.
func (s *Service) ListRoster(ctx context.Context, p db.ListRosterParams) ([]db.Roster, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListRoster(ctx, p))("list rosters")
}

// GetRoster returns one row of rosters.
func (s *Service) GetRoster(ctx context.Context, id uuid.UUID) (db.Roster, error) {
	row, err := s.q(ctx).GetRoster(ctx, id)
	if err != nil {
		return row, notFound(err, "rosters")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateRoster adds one row to rosters.
func (s *Service) CreateRoster(ctx context.Context, p db.CreateRosterParams) (db.Roster, error) {
	var zero db.Roster
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateRoster", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "ROS", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRoster(ctx, p)
	if err != nil {
		return zero, wrap(err, "create rosters")
	}
	if err := s.audit(ctx, "hr.roster.write", "rosters", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateRoster", row)
}

// UpdateRoster changes one row of rosters.
func (s *Service) UpdateRoster(ctx context.Context, p db.UpdateRosterParams) (db.Roster, error) {
	before, err := s.GetRoster(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRoster(ctx, p)
	if err != nil {
		return before, wrap(err, "update rosters")
	}
	if err := s.voidSignatures(ctx, "rosters", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.roster.write", "rosters", row.ID, &row.SiteID, before, row)
}

// SubmitRoster is the submit action on rosters.
func (s *Service) SubmitRoster(ctx context.Context, p db.ActSubmitRosterParams) (db.Roster, error) {
	before, err := s.GetRoster(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "hr.roster.submit", "rosters", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitRoster", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitRoster(ctx, p)
	if err != nil {
		return before, stateError(err, "submit rosters")
	}
	if err := s.audit(ctx, "hr.roster.submit", "rosters", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "rosters", row.ID, &row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitRoster", row)
}

// RosterChildren returns the rows of a record's child tables, by API field.
func (s *Service) RosterChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListRosterShifts(ctx, id)
		if err != nil {
			return nil, wrap(err, "list roster_shifts")
		}
		out["shifts"] = childRows(rows)
	}
	return out, nil
}

// ReplaceRosterShifts swaps a record's roster shifts for the given rows.
func (s *Service) ReplaceRosterShifts(ctx context.Context, id uuid.UUID, rows []db.InsertRosterShiftsParams) error {
	if err := s.q(ctx).DeleteRosterShifts(ctx, id); err != nil {
		return wrap(err, "clear roster_shifts")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertRosterShifts(ctx, row); err != nil {
			return wrap(err, "add roster_shifts")
		}
	}
	return nil
}

// ListAttendanceRecord returns a page of attendance records.
func (s *Service) ListAttendanceRecord(ctx context.Context, p db.ListAttendanceRecordParams) ([]db.AttendanceRecord, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListAttendanceRecord(ctx, p))("list attendance records")
}

// GetAttendanceRecord returns one row of attendance records.
func (s *Service) GetAttendanceRecord(ctx context.Context, id uuid.UUID) (db.AttendanceRecord, error) {
	row, err := s.q(ctx).GetAttendanceRecord(ctx, id)
	if err != nil {
		return row, notFound(err, "attendance records")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateAttendanceRecord adds one row to attendance records.
func (s *Service) CreateAttendanceRecord(ctx context.Context, p db.CreateAttendanceRecordParams) (db.AttendanceRecord, error) {
	var zero db.AttendanceRecord
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateAttendanceRecord", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateAttendanceRecord(ctx, p)
	if err != nil {
		return zero, wrap(err, "create attendance records")
	}
	if err := s.audit(ctx, "hr.attendance_record.write", "attendance_records", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateAttendanceRecord", row)
}

// UpdateAttendanceRecord changes one row of attendance records.
func (s *Service) UpdateAttendanceRecord(ctx context.Context, p db.UpdateAttendanceRecordParams) (db.AttendanceRecord, error) {
	before, err := s.GetAttendanceRecord(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateAttendanceRecord(ctx, p)
	if err != nil {
		return before, wrap(err, "update attendance records")
	}
	if err := s.voidSignatures(ctx, "attendance_records", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.attendance_record.write", "attendance_records", row.ID, &row.SiteID, before, row)
}

// ListUniformIssue returns a page of uniform issues.
func (s *Service) ListUniformIssue(ctx context.Context, p db.ListUniformIssueParams) ([]db.UniformIssue, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListUniformIssue(ctx, p))("list uniform issues")
}

// GetUniformIssue returns one row of uniform issues.
func (s *Service) GetUniformIssue(ctx context.Context, id uuid.UUID) (db.UniformIssue, error) {
	row, err := s.q(ctx).GetUniformIssue(ctx, id)
	if err != nil {
		return row, notFound(err, "uniform issues")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateUniformIssue adds one row to uniform issues.
func (s *Service) CreateUniformIssue(ctx context.Context, p db.CreateUniformIssueParams) (db.UniformIssue, error) {
	var zero db.UniformIssue
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateUniformIssue", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "UNI", p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateUniformIssue(ctx, p)
	if err != nil {
		return zero, wrap(err, "create uniform issues")
	}
	if err := s.audit(ctx, "hr.uniform_issue.write", "uniform_issues", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateUniformIssue", row)
}

// UpdateUniformIssue changes one row of uniform issues.
func (s *Service) UpdateUniformIssue(ctx context.Context, p db.UpdateUniformIssueParams) (db.UniformIssue, error) {
	before, err := s.GetUniformIssue(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateUniformIssue(ctx, p)
	if err != nil {
		return before, wrap(err, "update uniform issues")
	}
	if err := s.voidSignatures(ctx, "uniform_issues", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.uniform_issue.write", "uniform_issues", row.ID, row.SiteID, before, row)
}

// ListRecruitmentRequisition returns a page of recruitment requisitions.
func (s *Service) ListRecruitmentRequisition(ctx context.Context, p db.ListRecruitmentRequisitionParams) ([]db.RecruitmentRequisition, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListRecruitmentRequisition(ctx, p))("list recruitment requisitions")
}

// GetRecruitmentRequisition returns one row of recruitment requisitions.
func (s *Service) GetRecruitmentRequisition(ctx context.Context, id uuid.UUID) (db.RecruitmentRequisition, error) {
	row, err := s.q(ctx).GetRecruitmentRequisition(ctx, id)
	if err != nil {
		return row, notFound(err, "recruitment requisitions")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateRecruitmentRequisition adds one row to recruitment requisitions.
func (s *Service) CreateRecruitmentRequisition(ctx context.Context, p db.CreateRecruitmentRequisitionParams) (db.RecruitmentRequisition, error) {
	var zero db.RecruitmentRequisition
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateRecruitmentRequisition", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "REQ", p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRecruitmentRequisition(ctx, p)
	if err != nil {
		return zero, wrap(err, "create recruitment requisitions")
	}
	if err := s.audit(ctx, "hr.recruitment_requisition.write", "recruitment_requisitions", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateRecruitmentRequisition", row)
}

// UpdateRecruitmentRequisition changes one row of recruitment requisitions.
func (s *Service) UpdateRecruitmentRequisition(ctx context.Context, p db.UpdateRecruitmentRequisitionParams) (db.RecruitmentRequisition, error) {
	before, err := s.GetRecruitmentRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRecruitmentRequisition(ctx, p)
	if err != nil {
		return before, wrap(err, "update recruitment requisitions")
	}
	if err := s.voidSignatures(ctx, "recruitment_requisitions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.recruitment_requisition.write", "recruitment_requisitions", row.ID, row.SiteID, before, row)
}

// SubmitRecruitmentRequisition is the submit action on recruitment requisitions.
func (s *Service) SubmitRecruitmentRequisition(ctx context.Context, p db.ActSubmitRecruitmentRequisitionParams) (db.RecruitmentRequisition, error) {
	before, err := s.GetRecruitmentRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "hr.recruitment_requisition.submit", "recruitment_requisitions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitRecruitmentRequisition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitRecruitmentRequisition(ctx, p)
	if err != nil {
		return before, stateError(err, "submit recruitment requisitions")
	}
	if err := s.audit(ctx, "hr.recruitment_requisition.submit", "recruitment_requisitions", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "recruitment_requisitions", row.ID, row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitRecruitmentRequisition", row)
}

// ListCandidate returns a page of candidates.
func (s *Service) ListCandidate(ctx context.Context, p db.ListCandidateParams) ([]db.Candidate, error) {
	return wrap2(s.q(ctx).ListCandidate(ctx, p))("list candidates")
}

// GetCandidate returns one row of candidates.
func (s *Service) GetCandidate(ctx context.Context, id uuid.UUID) (db.Candidate, error) {
	row, err := s.q(ctx).GetCandidate(ctx, id)
	if err != nil {
		return row, notFound(err, "candidates")
	}
	return row, nil
}

// CreateCandidate adds one row to candidates.
func (s *Service) CreateCandidate(ctx context.Context, p db.CreateCandidateParams) (db.Candidate, error) {
	var zero db.Candidate
	if err := s.guard(ctx, "CreateCandidate", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCandidate(ctx, p)
	if err != nil {
		return zero, wrap(err, "create candidates")
	}
	if err := s.audit(ctx, "hr.candidate.write", "candidates", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateCandidate", row)
}

// UpdateCandidate changes one row of candidates.
func (s *Service) UpdateCandidate(ctx context.Context, p db.UpdateCandidateParams) (db.Candidate, error) {
	before, err := s.GetCandidate(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCandidate(ctx, p)
	if err != nil {
		return before, wrap(err, "update candidates")
	}
	if err := s.voidSignatures(ctx, "candidates", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.candidate.write", "candidates", row.ID, nil, before, row)
}

// ListPerformanceReview returns a page of performance reviews.
func (s *Service) ListPerformanceReview(ctx context.Context, p db.ListPerformanceReviewParams) ([]db.PerformanceReview, error) {
	return wrap2(s.q(ctx).ListPerformanceReview(ctx, p))("list performance reviews")
}

// GetPerformanceReview returns one row of performance reviews.
func (s *Service) GetPerformanceReview(ctx context.Context, id uuid.UUID) (db.PerformanceReview, error) {
	row, err := s.q(ctx).GetPerformanceReview(ctx, id)
	if err != nil {
		return row, notFound(err, "performance reviews")
	}
	return row, nil
}

// CreatePerformanceReview adds one row to performance reviews.
func (s *Service) CreatePerformanceReview(ctx context.Context, p db.CreatePerformanceReviewParams) (db.PerformanceReview, error) {
	var zero db.PerformanceReview
	if err := s.guard(ctx, "CreatePerformanceReview", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePerformanceReview(ctx, p)
	if err != nil {
		return zero, wrap(err, "create performance reviews")
	}
	if err := s.audit(ctx, "hr.performance_review.write", "performance_reviews", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreatePerformanceReview", row)
}

// UpdatePerformanceReview changes one row of performance reviews.
func (s *Service) UpdatePerformanceReview(ctx context.Context, p db.UpdatePerformanceReviewParams) (db.PerformanceReview, error) {
	before, err := s.GetPerformanceReview(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePerformanceReview(ctx, p)
	if err != nil {
		return before, wrap(err, "update performance reviews")
	}
	if err := s.voidSignatures(ctx, "performance_reviews", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.performance_review.write", "performance_reviews", row.ID, nil, before, row)
}

// CompletePerformanceReview is the complete action on performance reviews.
func (s *Service) CompletePerformanceReview(ctx context.Context, p db.ActCompletePerformanceReviewParams) (db.PerformanceReview, error) {
	before, err := s.GetPerformanceReview(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "hr.performance_review.complete", "performance_reviews", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CompletePerformanceReview", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "performance_reviews", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompletePerformanceReview(ctx, p)
	if err != nil {
		return before, stateError(err, "complete performance reviews")
	}
	if err := s.audit(ctx, "hr.performance_review.complete", "performance_reviews", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CompletePerformanceReview", row)
}

// ListDisciplinaryCase returns a page of disciplinary cases.
func (s *Service) ListDisciplinaryCase(ctx context.Context, p db.ListDisciplinaryCaseParams) ([]db.DisciplinaryCase, error) {
	return wrap2(s.q(ctx).ListDisciplinaryCase(ctx, p))("list disciplinary cases")
}

// GetDisciplinaryCase returns one row of disciplinary cases.
func (s *Service) GetDisciplinaryCase(ctx context.Context, id uuid.UUID) (db.DisciplinaryCase, error) {
	row, err := s.q(ctx).GetDisciplinaryCase(ctx, id)
	if err != nil {
		return row, notFound(err, "disciplinary cases")
	}
	return row, nil
}

// CreateDisciplinaryCase adds one row to disciplinary cases.
func (s *Service) CreateDisciplinaryCase(ctx context.Context, p db.CreateDisciplinaryCaseParams) (db.DisciplinaryCase, error) {
	var zero db.DisciplinaryCase
	if err := s.guard(ctx, "CreateDisciplinaryCase", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "DSC", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDisciplinaryCase(ctx, p)
	if err != nil {
		return zero, wrap(err, "create disciplinary cases")
	}
	if err := s.audit(ctx, "hr.disciplinary_case.write", "disciplinary_cases", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDisciplinaryCase", row)
}

// UpdateDisciplinaryCase changes one row of disciplinary cases.
func (s *Service) UpdateDisciplinaryCase(ctx context.Context, p db.UpdateDisciplinaryCaseParams) (db.DisciplinaryCase, error) {
	before, err := s.GetDisciplinaryCase(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDisciplinaryCase(ctx, p)
	if err != nil {
		return before, wrap(err, "update disciplinary cases")
	}
	if err := s.voidSignatures(ctx, "disciplinary_cases", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.disciplinary_case.write", "disciplinary_cases", row.ID, nil, before, row)
}

// ListGrievance returns a page of grievances.
func (s *Service) ListGrievance(ctx context.Context, p db.ListGrievanceParams) ([]db.Grievance, error) {
	return wrap2(s.q(ctx).ListGrievance(ctx, p))("list grievances")
}

// GetGrievance returns one row of grievances.
func (s *Service) GetGrievance(ctx context.Context, id uuid.UUID) (db.Grievance, error) {
	row, err := s.q(ctx).GetGrievance(ctx, id)
	if err != nil {
		return row, notFound(err, "grievances")
	}
	return row, nil
}

// CreateGrievance adds one row to grievances.
func (s *Service) CreateGrievance(ctx context.Context, p db.CreateGrievanceParams) (db.Grievance, error) {
	var zero db.Grievance
	if err := s.guard(ctx, "CreateGrievance", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "GRE", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateGrievance(ctx, p)
	if err != nil {
		return zero, wrap(err, "create grievances")
	}
	if err := s.audit(ctx, "hr.grievance.write", "grievances", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateGrievance", row)
}

// UpdateGrievance changes one row of grievances.
func (s *Service) UpdateGrievance(ctx context.Context, p db.UpdateGrievanceParams) (db.Grievance, error) {
	before, err := s.GetGrievance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateGrievance(ctx, p)
	if err != nil {
		return before, wrap(err, "update grievances")
	}
	if err := s.voidSignatures(ctx, "grievances", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "hr.grievance.write", "grievances", row.ID, nil, before, row)
}
