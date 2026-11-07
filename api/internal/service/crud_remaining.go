// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListStoreOpening returns a page of store openings.
func (s *Service) ListStoreOpening(ctx context.Context, p db.ListStoreOpeningParams) ([]db.StoreOpening, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListStoreOpening(ctx, p))("list store openings")
}

// GetStoreOpening returns one row of store openings.
func (s *Service) GetStoreOpening(ctx context.Context, id uuid.UUID) (db.StoreOpening, error) {
	row, err := s.q(ctx).GetStoreOpening(ctx, id)
	if err != nil {
		return row, notFound(err, "store openings")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateStoreOpening adds one row to store openings.
func (s *Service) CreateStoreOpening(ctx context.Context, p db.CreateStoreOpeningParams) (db.StoreOpening, error) {
	var zero db.StoreOpening
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateStoreOpening", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "OPN", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateStoreOpening(ctx, p)
	if err != nil {
		return zero, wrap(err, "create store openings")
	}
	if err := s.audit(ctx, "store_ops.store_opening.write", "store_openings", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateStoreOpening", row)
}

// UpdateStoreOpening changes one row of store openings.
func (s *Service) UpdateStoreOpening(ctx context.Context, p db.UpdateStoreOpeningParams) (db.StoreOpening, error) {
	before, err := s.GetStoreOpening(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateStoreOpening(ctx, p)
	if err != nil {
		return before, wrap(err, "update store openings")
	}
	if err := s.voidSignatures(ctx, "store_openings", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "store_ops.store_opening.write", "store_openings", row.ID, &row.SiteID, before, row)
}

// RecordOpenedStoreOpening is the record actual opening action on store openings.
func (s *Service) RecordOpenedStoreOpening(ctx context.Context, p db.ActRecordOpenedStoreOpeningParams) (db.StoreOpening, error) {
	before, err := s.GetStoreOpening(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "store_ops.store_opening.record_opened", "store_openings", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RecordOpenedStoreOpening", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRecordOpenedStoreOpening(ctx, p)
	if err != nil {
		return before, stateError(err, "record opened store openings")
	}
	if err := s.audit(ctx, "store_ops.store_opening.record_opened", "store_openings", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RecordOpenedStoreOpening", row)
}

// ListOccurrenceEntry returns a page of occurrence entries.
func (s *Service) ListOccurrenceEntry(ctx context.Context, p db.ListOccurrenceEntryParams) ([]db.OccurrenceEntry, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListOccurrenceEntry(ctx, p))("list occurrence entries")
}

// GetOccurrenceEntry returns one row of occurrence entries.
func (s *Service) GetOccurrenceEntry(ctx context.Context, id uuid.UUID) (db.OccurrenceEntry, error) {
	row, err := s.q(ctx).GetOccurrenceEntry(ctx, id)
	if err != nil {
		return row, notFound(err, "occurrence entries")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateOccurrenceEntry adds one row to occurrence entries.
func (s *Service) CreateOccurrenceEntry(ctx context.Context, p db.CreateOccurrenceEntryParams) (db.OccurrenceEntry, error) {
	var zero db.OccurrenceEntry
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateOccurrenceEntry", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "OB", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateOccurrenceEntry(ctx, p)
	if err != nil {
		return zero, wrap(err, "create occurrence entries")
	}
	if err := s.audit(ctx, "security.occurrence_entry.write", "occurrence_entries", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateOccurrenceEntry", row)
}

// ListPatrolRoute returns a page of patrol routes.
func (s *Service) ListPatrolRoute(ctx context.Context, p db.ListPatrolRouteParams) ([]db.PatrolRoute, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListPatrolRoute(ctx, p))("list patrol routes")
}

// GetPatrolRoute returns one row of patrol routes.
func (s *Service) GetPatrolRoute(ctx context.Context, id uuid.UUID) (db.PatrolRoute, error) {
	row, err := s.q(ctx).GetPatrolRoute(ctx, id)
	if err != nil {
		return row, notFound(err, "patrol routes")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreatePatrolRoute adds one row to patrol routes.
func (s *Service) CreatePatrolRoute(ctx context.Context, p db.CreatePatrolRouteParams) (db.PatrolRoute, error) {
	var zero db.PatrolRoute
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreatePatrolRoute", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePatrolRoute(ctx, p)
	if err != nil {
		return zero, wrap(err, "create patrol routes")
	}
	if err := s.audit(ctx, "security.patrol_route.write", "patrol_routes", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreatePatrolRoute", row)
}

// UpdatePatrolRoute changes one row of patrol routes.
func (s *Service) UpdatePatrolRoute(ctx context.Context, p db.UpdatePatrolRouteParams) (db.PatrolRoute, error) {
	before, err := s.GetPatrolRoute(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePatrolRoute(ctx, p)
	if err != nil {
		return before, wrap(err, "update patrol routes")
	}
	if err := s.voidSignatures(ctx, "patrol_routes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "security.patrol_route.write", "patrol_routes", row.ID, &row.SiteID, before, row)
}

// PatrolRouteChildren returns the rows of a record's child tables, by API field.
func (s *Service) PatrolRouteChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListPatrolCheckpoints(ctx, id)
		if err != nil {
			return nil, wrap(err, "list patrol_checkpoints")
		}
		out["checkpoints"] = childRows(rows)
	}
	return out, nil
}

// ReplacePatrolCheckpoints swaps a record's patrol checkpoints for the given rows.
func (s *Service) ReplacePatrolCheckpoints(ctx context.Context, id uuid.UUID, rows []db.InsertPatrolCheckpointsParams) error {
	if err := s.q(ctx).DeletePatrolCheckpoints(ctx, id); err != nil {
		return wrap(err, "clear patrol_checkpoints")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertPatrolCheckpoints(ctx, row); err != nil {
			return wrap(err, "add patrol_checkpoints")
		}
	}
	return nil
}

// ListPatrolRound returns a page of patrol rounds.
func (s *Service) ListPatrolRound(ctx context.Context, p db.ListPatrolRoundParams) ([]db.PatrolRound, error) {
	return wrap2(s.q(ctx).ListPatrolRound(ctx, p))("list patrol rounds")
}

// GetPatrolRound returns one row of patrol rounds.
func (s *Service) GetPatrolRound(ctx context.Context, id uuid.UUID) (db.PatrolRound, error) {
	row, err := s.q(ctx).GetPatrolRound(ctx, id)
	if err != nil {
		return row, notFound(err, "patrol rounds")
	}
	return row, nil
}

// CreatePatrolRound adds one row to patrol rounds.
func (s *Service) CreatePatrolRound(ctx context.Context, p db.CreatePatrolRoundParams) (db.PatrolRound, error) {
	var zero db.PatrolRound
	if err := s.guard(ctx, "CreatePatrolRound", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "PTR", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePatrolRound(ctx, p)
	if err != nil {
		return zero, wrap(err, "create patrol rounds")
	}
	if err := s.audit(ctx, "security.patrol_round.write", "patrol_rounds", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreatePatrolRound", row)
}

// UpdatePatrolRound changes one row of patrol rounds.
func (s *Service) UpdatePatrolRound(ctx context.Context, p db.UpdatePatrolRoundParams) (db.PatrolRound, error) {
	before, err := s.GetPatrolRound(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePatrolRound(ctx, p)
	if err != nil {
		return before, wrap(err, "update patrol rounds")
	}
	if err := s.voidSignatures(ctx, "patrol_rounds", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "security.patrol_round.write", "patrol_rounds", row.ID, nil, before, row)
}

// CompletePatrolRound is the complete the round action on patrol rounds.
func (s *Service) CompletePatrolRound(ctx context.Context, p db.ActCompletePatrolRoundParams) (db.PatrolRound, error) {
	before, err := s.GetPatrolRound(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "security.patrol_round.complete", "patrol_rounds", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CompletePatrolRound", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "patrol_rounds", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompletePatrolRound(ctx, p)
	if err != nil {
		return before, stateError(err, "complete patrol rounds")
	}
	if err := s.audit(ctx, "security.patrol_round.complete", "patrol_rounds", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CompletePatrolRound", row)
}

// PatrolRoundChildren returns the rows of a record's child tables, by API field.
func (s *Service) PatrolRoundChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListPatrolVisits(ctx, id)
		if err != nil {
			return nil, wrap(err, "list patrol_visits")
		}
		out["visits"] = childRows(rows)
	}
	return out, nil
}

// ReplacePatrolVisits swaps a record's patrol visits for the given rows.
func (s *Service) ReplacePatrolVisits(ctx context.Context, id uuid.UUID, rows []db.InsertPatrolVisitsParams) error {
	if err := s.q(ctx).DeletePatrolVisits(ctx, id); err != nil {
		return wrap(err, "clear patrol_visits")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertPatrolVisits(ctx, row); err != nil {
			return wrap(err, "add patrol_visits")
		}
	}
	return nil
}

// ListCamera returns a page of cameras.
func (s *Service) ListCamera(ctx context.Context, p db.ListCameraParams) ([]db.Camera, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListCamera(ctx, p))("list cameras")
}

// GetCamera returns one row of cameras.
func (s *Service) GetCamera(ctx context.Context, id uuid.UUID) (db.Camera, error) {
	row, err := s.q(ctx).GetCamera(ctx, id)
	if err != nil {
		return row, notFound(err, "cameras")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateCamera adds one row to cameras.
func (s *Service) CreateCamera(ctx context.Context, p db.CreateCameraParams) (db.Camera, error) {
	var zero db.Camera
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateCamera", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCamera(ctx, p)
	if err != nil {
		return zero, wrap(err, "create cameras")
	}
	if err := s.audit(ctx, "cctv.camera.write", "cameras", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateCamera", row)
}

// UpdateCamera changes one row of cameras.
func (s *Service) UpdateCamera(ctx context.Context, p db.UpdateCameraParams) (db.Camera, error) {
	before, err := s.GetCamera(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCamera(ctx, p)
	if err != nil {
		return before, wrap(err, "update cameras")
	}
	if err := s.voidSignatures(ctx, "cameras", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cctv.camera.write", "cameras", row.ID, &row.SiteID, before, row)
}

// ListCctvOutage returns a page of cctv outages.
func (s *Service) ListCctvOutage(ctx context.Context, p db.ListCctvOutageParams) ([]db.CctvOutage, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListCctvOutage(ctx, p))("list cctv outages")
}

// GetCctvOutage returns one row of cctv outages.
func (s *Service) GetCctvOutage(ctx context.Context, id uuid.UUID) (db.CctvOutage, error) {
	row, err := s.q(ctx).GetCctvOutage(ctx, id)
	if err != nil {
		return row, notFound(err, "cctv outages")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateCctvOutage adds one row to cctv outages.
func (s *Service) CreateCctvOutage(ctx context.Context, p db.CreateCctvOutageParams) (db.CctvOutage, error) {
	var zero db.CctvOutage
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateCctvOutage", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "CCO", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCctvOutage(ctx, p)
	if err != nil {
		return zero, wrap(err, "create cctv outages")
	}
	if err := s.audit(ctx, "cctv.cctv_outage.write", "cctv_outages", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateCctvOutage", row)
}

// UpdateCctvOutage changes one row of cctv outages.
func (s *Service) UpdateCctvOutage(ctx context.Context, p db.UpdateCctvOutageParams) (db.CctvOutage, error) {
	before, err := s.GetCctvOutage(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCctvOutage(ctx, p)
	if err != nil {
		return before, wrap(err, "update cctv outages")
	}
	if err := s.voidSignatures(ctx, "cctv_outages", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cctv.cctv_outage.write", "cctv_outages", row.ID, &row.SiteID, before, row)
}

// RestoreCctvOutage is the record restoration action on cctv outages.
func (s *Service) RestoreCctvOutage(ctx context.Context, p db.ActRestoreCctvOutageParams) (db.CctvOutage, error) {
	before, err := s.GetCctvOutage(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cctv.cctv_outage.restore", "cctv_outages", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RestoreCctvOutage", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRestoreCctvOutage(ctx, p)
	if err != nil {
		return before, stateError(err, "restore cctv outages")
	}
	if err := s.audit(ctx, "cctv.cctv_outage.restore", "cctv_outages", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RestoreCctvOutage", row)
}

// ListFootageRequest returns a page of footage requests.
func (s *Service) ListFootageRequest(ctx context.Context, p db.ListFootageRequestParams) ([]db.FootageRequest, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListFootageRequest(ctx, p))("list footage requests")
}

// GetFootageRequest returns one row of footage requests.
func (s *Service) GetFootageRequest(ctx context.Context, id uuid.UUID) (db.FootageRequest, error) {
	row, err := s.q(ctx).GetFootageRequest(ctx, id)
	if err != nil {
		return row, notFound(err, "footage requests")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateFootageRequest adds one row to footage requests.
func (s *Service) CreateFootageRequest(ctx context.Context, p db.CreateFootageRequestParams) (db.FootageRequest, error) {
	var zero db.FootageRequest
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateFootageRequest", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "CTV", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateFootageRequest(ctx, p)
	if err != nil {
		return zero, wrap(err, "create footage requests")
	}
	if err := s.audit(ctx, "cctv.footage_request.write", "footage_requests", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateFootageRequest", row)
}

// UpdateFootageRequest changes one row of footage requests.
func (s *Service) UpdateFootageRequest(ctx context.Context, p db.UpdateFootageRequestParams) (db.FootageRequest, error) {
	before, err := s.GetFootageRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateFootageRequest(ctx, p)
	if err != nil {
		return before, wrap(err, "update footage requests")
	}
	if err := s.voidSignatures(ctx, "footage_requests", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cctv.footage_request.write", "footage_requests", row.ID, &row.SiteID, before, row)
}

// SubmitFootageRequest is the submit action on footage requests.
func (s *Service) SubmitFootageRequest(ctx context.Context, p db.ActSubmitFootageRequestParams) (db.FootageRequest, error) {
	before, err := s.GetFootageRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cctv.footage_request.submit", "footage_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitFootageRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitFootageRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "submit footage requests")
	}
	if err := s.audit(ctx, "cctv.footage_request.submit", "footage_requests", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "footage_requests", row.ID, &row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitFootageRequest", row)
}

// RecordExportFootageRequest is the record export details action on footage requests.
func (s *Service) RecordExportFootageRequest(ctx context.Context, p db.ActRecordExportFootageRequestParams) (db.FootageRequest, error) {
	before, err := s.GetFootageRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cctv.footage_request.record_export", "footage_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RecordExportFootageRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRecordExportFootageRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "record export footage requests")
	}
	if err := s.audit(ctx, "cctv.footage_request.record_export", "footage_requests", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RecordExportFootageRequest", row)
}

// AddCustodyEntryFootageRequest is the add a chain-of-custody entry action on footage requests.
func (s *Service) AddCustodyEntryFootageRequest(ctx context.Context, p db.ActAddCustodyEntryFootageRequestParams) (db.FootageRequest, error) {
	before, err := s.GetFootageRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cctv.footage_request.add_custody_entry", "footage_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "AddCustodyEntryFootageRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAddCustodyEntryFootageRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "add custody entry footage requests")
	}
	if err := s.audit(ctx, "cctv.footage_request.add_custody_entry", "footage_requests", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "AddCustodyEntryFootageRequest", row)
}

// CloseFootageRequest is the close action on footage requests.
func (s *Service) CloseFootageRequest(ctx context.Context, p db.ActCloseFootageRequestParams) (db.FootageRequest, error) {
	before, err := s.GetFootageRequest(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cctv.footage_request.close", "footage_requests", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseFootageRequest", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "footage_requests", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseFootageRequest(ctx, p)
	if err != nil {
		return before, stateError(err, "close footage requests")
	}
	if err := s.audit(ctx, "cctv.footage_request.close", "footage_requests", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseFootageRequest", row)
}

// ListExternalVisit returns a page of external visits.
func (s *Service) ListExternalVisit(ctx context.Context, p db.ListExternalVisitParams) ([]db.ExternalVisit, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListExternalVisit(ctx, p))("list external visits")
}

// GetExternalVisit returns one row of external visits.
func (s *Service) GetExternalVisit(ctx context.Context, id uuid.UUID) (db.ExternalVisit, error) {
	row, err := s.q(ctx).GetExternalVisit(ctx, id)
	if err != nil {
		return row, notFound(err, "external visits")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateExternalVisit adds one row to external visits.
func (s *Service) CreateExternalVisit(ctx context.Context, p db.CreateExternalVisitParams) (db.ExternalVisit, error) {
	var zero db.ExternalVisit
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateExternalVisit", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "VIS", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateExternalVisit(ctx, p)
	if err != nil {
		return zero, wrap(err, "create external visits")
	}
	if err := s.audit(ctx, "external_visits.external_visit.write", "external_visits", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateExternalVisit", row)
}

// UpdateExternalVisit changes one row of external visits.
func (s *Service) UpdateExternalVisit(ctx context.Context, p db.UpdateExternalVisitParams) (db.ExternalVisit, error) {
	before, err := s.GetExternalVisit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateExternalVisit(ctx, p)
	if err != nil {
		return before, wrap(err, "update external visits")
	}
	if err := s.voidSignatures(ctx, "external_visits", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "external_visits.external_visit.write", "external_visits", row.ID, &row.SiteID, before, row)
}

// SubmitExternalVisit is the submit action on external visits.
func (s *Service) SubmitExternalVisit(ctx context.Context, p db.ActSubmitExternalVisitParams) (db.ExternalVisit, error) {
	before, err := s.GetExternalVisit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "external_visits.external_visit.submit", "external_visits", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitExternalVisit", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitExternalVisit(ctx, p)
	if err != nil {
		return before, stateError(err, "submit external visits")
	}
	if err := s.audit(ctx, "external_visits.external_visit.submit", "external_visits", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "external_visits", row.ID, &row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitExternalVisit", row)
}

// CloseExternalVisit is the close out action on external visits.
func (s *Service) CloseExternalVisit(ctx context.Context, p db.ActCloseExternalVisitParams) (db.ExternalVisit, error) {
	before, err := s.GetExternalVisit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "external_visits.external_visit.close", "external_visits", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseExternalVisit", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "external_visits", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseExternalVisit(ctx, p)
	if err != nil {
		return before, stateError(err, "close external visits")
	}
	if err := s.audit(ctx, "external_visits.external_visit.close", "external_visits", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseExternalVisit", row)
}
