// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListStoreOpenings implements the listStoreOpenings operation.
func (s *Server) ListStoreOpenings(ctx context.Context, req apigen.ListStoreOpeningsRequestObject) (apigen.ListStoreOpeningsResponseObject, error) {
	return listPage[apigen.ListStoreOpenings200JSONResponse](ctx, req.Params, nil, s.service.ListStoreOpening)
}

// GetStoreOpening implements the getStoreOpening operation.
func (s *Server) GetStoreOpening(ctx context.Context, req apigen.GetStoreOpeningRequestObject) (apigen.GetStoreOpeningResponseObject, error) {
	row, err := s.service.GetStoreOpening(ctx, req.StoreOpeningId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StoreOpening](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetStoreOpening200JSONResponse{Body: body, Headers: apigen.GetStoreOpening200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateStoreOpening implements the createStoreOpening operation.
func (s *Server) CreateStoreOpening(ctx context.Context, req apigen.CreateStoreOpeningRequestObject) (apigen.CreateStoreOpeningResponseObject, error) {
	in, err := input[db.CreateStoreOpeningParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateStoreOpening(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StoreOpening](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateStoreOpening201JSONResponse{Body: body, Headers: apigen.CreateStoreOpening201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateStoreOpening implements the updateStoreOpening operation.
func (s *Server) UpdateStoreOpening(ctx context.Context, req apigen.UpdateStoreOpeningRequestObject) (apigen.UpdateStoreOpeningResponseObject, error) {
	in, err := input[db.UpdateStoreOpeningParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.StoreOpeningId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateStoreOpening(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StoreOpening](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateStoreOpening200JSONResponse{Body: body, Headers: apigen.UpdateStoreOpening200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RecordOpenedStoreOpening implements the recordOpenedStoreOpening operation.
func (s *Server) RecordOpenedStoreOpening(ctx context.Context, req apigen.RecordOpenedStoreOpeningRequestObject) (apigen.RecordOpenedStoreOpeningResponseObject, error) {
	in, err := input[db.ActRecordOpenedStoreOpeningParams](req.Body, map[string]any{"id": req.StoreOpeningId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RecordOpenedStoreOpening(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StoreOpening](row)
	if err != nil {
		return nil, err
	}
	return apigen.RecordOpenedStoreOpening200JSONResponse{Body: body, Headers: apigen.RecordOpenedStoreOpening200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListOccurrenceEntries implements the listOccurrenceEntries operation.
func (s *Server) ListOccurrenceEntries(ctx context.Context, req apigen.ListOccurrenceEntriesRequestObject) (apigen.ListOccurrenceEntriesResponseObject, error) {
	return listPage[apigen.ListOccurrenceEntries200JSONResponse](ctx, req.Params, nil, s.service.ListOccurrenceEntry)
}

// GetOccurrenceEntry implements the getOccurrenceEntry operation.
func (s *Server) GetOccurrenceEntry(ctx context.Context, req apigen.GetOccurrenceEntryRequestObject) (apigen.GetOccurrenceEntryResponseObject, error) {
	row, err := s.service.GetOccurrenceEntry(ctx, req.OccurrenceEntryId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.OccurrenceEntry](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetOccurrenceEntry200JSONResponse{Body: body, Headers: apigen.GetOccurrenceEntry200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateOccurrenceEntry implements the createOccurrenceEntry operation.
func (s *Server) CreateOccurrenceEntry(ctx context.Context, req apigen.CreateOccurrenceEntryRequestObject) (apigen.CreateOccurrenceEntryResponseObject, error) {
	in, err := input[db.CreateOccurrenceEntryParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateOccurrenceEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.OccurrenceEntry](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateOccurrenceEntry201JSONResponse{Body: body, Headers: apigen.CreateOccurrenceEntry201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPatrolRoutes implements the listPatrolRoutes operation.
func (s *Server) ListPatrolRoutes(ctx context.Context, req apigen.ListPatrolRoutesRequestObject) (apigen.ListPatrolRoutesResponseObject, error) {
	return listPage[apigen.ListPatrolRoutes200JSONResponse](ctx, req.Params, nil, s.service.ListPatrolRoute)
}

// GetPatrolRoute implements the getPatrolRoute operation.
func (s *Server) GetPatrolRoute(ctx context.Context, req apigen.GetPatrolRouteRequestObject) (apigen.GetPatrolRouteResponseObject, error) {
	row, err := s.service.GetPatrolRoute(ctx, req.PatrolRouteId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.PatrolRouteChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PatrolRoute](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetPatrolRoute200JSONResponse{Body: body, Headers: apigen.GetPatrolRoute200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePatrolRoute implements the createPatrolRoute operation.
func (s *Server) CreatePatrolRoute(ctx context.Context, req apigen.CreatePatrolRouteRequestObject) (apigen.CreatePatrolRouteResponseObject, error) {
	in, err := input[db.CreatePatrolRouteParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePatrolRoute(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["checkpoints"]; ok {
		rows, err := childInputs[db.InsertPatrolCheckpointsParams](items, "patrol_route_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePatrolCheckpoints(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PatrolRouteChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PatrolRoute](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePatrolRoute201JSONResponse{Body: body, Headers: apigen.CreatePatrolRoute201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePatrolRoute implements the updatePatrolRoute operation.
func (s *Server) UpdatePatrolRoute(ctx context.Context, req apigen.UpdatePatrolRouteRequestObject) (apigen.UpdatePatrolRouteResponseObject, error) {
	in, err := input[db.UpdatePatrolRouteParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PatrolRouteId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePatrolRoute(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["checkpoints"]; ok {
		rows, err := childInputs[db.InsertPatrolCheckpointsParams](items, "patrol_route_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePatrolCheckpoints(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PatrolRouteChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PatrolRoute](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePatrolRoute200JSONResponse{Body: body, Headers: apigen.UpdatePatrolRoute200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPatrolRounds implements the listPatrolRounds operation.
func (s *Server) ListPatrolRounds(ctx context.Context, req apigen.ListPatrolRoundsRequestObject) (apigen.ListPatrolRoundsResponseObject, error) {
	return listPage[apigen.ListPatrolRounds200JSONResponse](ctx, req.Params, nil, s.service.ListPatrolRound)
}

// GetPatrolRound implements the getPatrolRound operation.
func (s *Server) GetPatrolRound(ctx context.Context, req apigen.GetPatrolRoundRequestObject) (apigen.GetPatrolRoundResponseObject, error) {
	row, err := s.service.GetPatrolRound(ctx, req.PatrolRoundId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.PatrolRoundChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PatrolRound](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetPatrolRound200JSONResponse{Body: body, Headers: apigen.GetPatrolRound200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePatrolRound implements the createPatrolRound operation.
func (s *Server) CreatePatrolRound(ctx context.Context, req apigen.CreatePatrolRoundRequestObject) (apigen.CreatePatrolRoundResponseObject, error) {
	in, err := input[db.CreatePatrolRoundParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePatrolRound(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["visits"]; ok {
		rows, err := childInputs[db.InsertPatrolVisitsParams](items, "patrol_round_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePatrolVisits(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PatrolRoundChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PatrolRound](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePatrolRound201JSONResponse{Body: body, Headers: apigen.CreatePatrolRound201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePatrolRound implements the updatePatrolRound operation.
func (s *Server) UpdatePatrolRound(ctx context.Context, req apigen.UpdatePatrolRoundRequestObject) (apigen.UpdatePatrolRoundResponseObject, error) {
	in, err := input[db.UpdatePatrolRoundParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PatrolRoundId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePatrolRound(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["visits"]; ok {
		rows, err := childInputs[db.InsertPatrolVisitsParams](items, "patrol_round_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePatrolVisits(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PatrolRoundChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PatrolRound](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePatrolRound200JSONResponse{Body: body, Headers: apigen.UpdatePatrolRound200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompletePatrolRound implements the completePatrolRound operation.
func (s *Server) CompletePatrolRound(ctx context.Context, req apigen.CompletePatrolRoundRequestObject) (apigen.CompletePatrolRoundResponseObject, error) {
	in, err := input[db.ActCompletePatrolRoundParams](struct{}{}, map[string]any{"id": req.PatrolRoundId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompletePatrolRound(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PatrolRound](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompletePatrolRound200JSONResponse{Body: body, Headers: apigen.CompletePatrolRound200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCameras implements the listCameras operation.
func (s *Server) ListCameras(ctx context.Context, req apigen.ListCamerasRequestObject) (apigen.ListCamerasResponseObject, error) {
	return listPage[apigen.ListCameras200JSONResponse](ctx, req.Params, nil, s.service.ListCamera)
}

// GetCamera implements the getCamera operation.
func (s *Server) GetCamera(ctx context.Context, req apigen.GetCameraRequestObject) (apigen.GetCameraResponseObject, error) {
	row, err := s.service.GetCamera(ctx, req.CameraId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Camera](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCamera200JSONResponse{Body: body, Headers: apigen.GetCamera200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCamera implements the createCamera operation.
func (s *Server) CreateCamera(ctx context.Context, req apigen.CreateCameraRequestObject) (apigen.CreateCameraResponseObject, error) {
	in, err := input[db.CreateCameraParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCamera(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Camera](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCamera201JSONResponse{Body: body, Headers: apigen.CreateCamera201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCamera implements the updateCamera operation.
func (s *Server) UpdateCamera(ctx context.Context, req apigen.UpdateCameraRequestObject) (apigen.UpdateCameraResponseObject, error) {
	in, err := input[db.UpdateCameraParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CameraId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCamera(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Camera](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCamera200JSONResponse{Body: body, Headers: apigen.UpdateCamera200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCctvOutages implements the listCctvOutages operation.
func (s *Server) ListCctvOutages(ctx context.Context, req apigen.ListCctvOutagesRequestObject) (apigen.ListCctvOutagesResponseObject, error) {
	return listPage[apigen.ListCctvOutages200JSONResponse](ctx, req.Params, nil, s.service.ListCctvOutage)
}

// GetCctvOutage implements the getCctvOutage operation.
func (s *Server) GetCctvOutage(ctx context.Context, req apigen.GetCctvOutageRequestObject) (apigen.GetCctvOutageResponseObject, error) {
	row, err := s.service.GetCctvOutage(ctx, req.CctvOutageId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CctvOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCctvOutage200JSONResponse{Body: body, Headers: apigen.GetCctvOutage200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCctvOutage implements the createCctvOutage operation.
func (s *Server) CreateCctvOutage(ctx context.Context, req apigen.CreateCctvOutageRequestObject) (apigen.CreateCctvOutageResponseObject, error) {
	in, err := input[db.CreateCctvOutageParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCctvOutage(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CctvOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCctvOutage201JSONResponse{Body: body, Headers: apigen.CreateCctvOutage201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCctvOutage implements the updateCctvOutage operation.
func (s *Server) UpdateCctvOutage(ctx context.Context, req apigen.UpdateCctvOutageRequestObject) (apigen.UpdateCctvOutageResponseObject, error) {
	in, err := input[db.UpdateCctvOutageParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CctvOutageId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCctvOutage(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CctvOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCctvOutage200JSONResponse{Body: body, Headers: apigen.UpdateCctvOutage200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RestoreCctvOutage implements the restoreCctvOutage operation.
func (s *Server) RestoreCctvOutage(ctx context.Context, req apigen.RestoreCctvOutageRequestObject) (apigen.RestoreCctvOutageResponseObject, error) {
	in, err := input[db.ActRestoreCctvOutageParams](req.Body, map[string]any{"id": req.CctvOutageId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RestoreCctvOutage(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CctvOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.RestoreCctvOutage200JSONResponse{Body: body, Headers: apigen.RestoreCctvOutage200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListFootageRequests implements the listFootageRequests operation.
func (s *Server) ListFootageRequests(ctx context.Context, req apigen.ListFootageRequestsRequestObject) (apigen.ListFootageRequestsResponseObject, error) {
	return listPage[apigen.ListFootageRequests200JSONResponse](ctx, req.Params, nil, s.service.ListFootageRequest)
}

// GetFootageRequest implements the getFootageRequest operation.
func (s *Server) GetFootageRequest(ctx context.Context, req apigen.GetFootageRequestRequestObject) (apigen.GetFootageRequestResponseObject, error) {
	row, err := s.service.GetFootageRequest(ctx, req.FootageRequestId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FootageRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetFootageRequest200JSONResponse{Body: body, Headers: apigen.GetFootageRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateFootageRequest implements the createFootageRequest operation.
func (s *Server) CreateFootageRequest(ctx context.Context, req apigen.CreateFootageRequestRequestObject) (apigen.CreateFootageRequestResponseObject, error) {
	in, err := input[db.CreateFootageRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateFootageRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FootageRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateFootageRequest201JSONResponse{Body: body, Headers: apigen.CreateFootageRequest201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateFootageRequest implements the updateFootageRequest operation.
func (s *Server) UpdateFootageRequest(ctx context.Context, req apigen.UpdateFootageRequestRequestObject) (apigen.UpdateFootageRequestResponseObject, error) {
	in, err := input[db.UpdateFootageRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.FootageRequestId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateFootageRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FootageRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateFootageRequest200JSONResponse{Body: body, Headers: apigen.UpdateFootageRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitFootageRequest implements the submitFootageRequest operation.
func (s *Server) SubmitFootageRequest(ctx context.Context, req apigen.SubmitFootageRequestRequestObject) (apigen.SubmitFootageRequestResponseObject, error) {
	in, err := input[db.ActSubmitFootageRequestParams](struct{}{}, map[string]any{"id": req.FootageRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitFootageRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FootageRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitFootageRequest200JSONResponse{Body: body, Headers: apigen.SubmitFootageRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RecordExportFootageRequest implements the recordExportFootageRequest operation.
func (s *Server) RecordExportFootageRequest(ctx context.Context, req apigen.RecordExportFootageRequestRequestObject) (apigen.RecordExportFootageRequestResponseObject, error) {
	in, err := input[db.ActRecordExportFootageRequestParams](req.Body, map[string]any{"id": req.FootageRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RecordExportFootageRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FootageRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.RecordExportFootageRequest200JSONResponse{Body: body, Headers: apigen.RecordExportFootageRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AddCustodyEntryFootageRequest implements the addCustodyEntryFootageRequest operation.
func (s *Server) AddCustodyEntryFootageRequest(ctx context.Context, req apigen.AddCustodyEntryFootageRequestRequestObject) (apigen.AddCustodyEntryFootageRequestResponseObject, error) {
	in, err := input[db.ActAddCustodyEntryFootageRequestParams](req.Body, map[string]any{"id": req.FootageRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AddCustodyEntryFootageRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FootageRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.AddCustodyEntryFootageRequest200JSONResponse{Body: body, Headers: apigen.AddCustodyEntryFootageRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseFootageRequest implements the closeFootageRequest operation.
func (s *Server) CloseFootageRequest(ctx context.Context, req apigen.CloseFootageRequestRequestObject) (apigen.CloseFootageRequestResponseObject, error) {
	in, err := input[db.ActCloseFootageRequestParams](struct{}{}, map[string]any{"id": req.FootageRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseFootageRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FootageRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseFootageRequest200JSONResponse{Body: body, Headers: apigen.CloseFootageRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListExternalVisits implements the listExternalVisits operation.
func (s *Server) ListExternalVisits(ctx context.Context, req apigen.ListExternalVisitsRequestObject) (apigen.ListExternalVisitsResponseObject, error) {
	return listPage[apigen.ListExternalVisits200JSONResponse](ctx, req.Params, nil, s.service.ListExternalVisit)
}

// GetExternalVisit implements the getExternalVisit operation.
func (s *Server) GetExternalVisit(ctx context.Context, req apigen.GetExternalVisitRequestObject) (apigen.GetExternalVisitResponseObject, error) {
	row, err := s.service.GetExternalVisit(ctx, req.ExternalVisitId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExternalVisit](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetExternalVisit200JSONResponse{Body: body, Headers: apigen.GetExternalVisit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateExternalVisit implements the createExternalVisit operation.
func (s *Server) CreateExternalVisit(ctx context.Context, req apigen.CreateExternalVisitRequestObject) (apigen.CreateExternalVisitResponseObject, error) {
	in, err := input[db.CreateExternalVisitParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateExternalVisit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExternalVisit](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateExternalVisit201JSONResponse{Body: body, Headers: apigen.CreateExternalVisit201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateExternalVisit implements the updateExternalVisit operation.
func (s *Server) UpdateExternalVisit(ctx context.Context, req apigen.UpdateExternalVisitRequestObject) (apigen.UpdateExternalVisitResponseObject, error) {
	in, err := input[db.UpdateExternalVisitParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ExternalVisitId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateExternalVisit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExternalVisit](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateExternalVisit200JSONResponse{Body: body, Headers: apigen.UpdateExternalVisit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitExternalVisit implements the submitExternalVisit operation.
func (s *Server) SubmitExternalVisit(ctx context.Context, req apigen.SubmitExternalVisitRequestObject) (apigen.SubmitExternalVisitResponseObject, error) {
	in, err := input[db.ActSubmitExternalVisitParams](struct{}{}, map[string]any{"id": req.ExternalVisitId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitExternalVisit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExternalVisit](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitExternalVisit200JSONResponse{Body: body, Headers: apigen.SubmitExternalVisit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseExternalVisit implements the closeExternalVisit operation.
func (s *Server) CloseExternalVisit(ctx context.Context, req apigen.CloseExternalVisitRequestObject) (apigen.CloseExternalVisitResponseObject, error) {
	in, err := input[db.ActCloseExternalVisitParams](req.Body, map[string]any{"id": req.ExternalVisitId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseExternalVisit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExternalVisit](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseExternalVisit200JSONResponse{Body: body, Headers: apigen.CloseExternalVisit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
