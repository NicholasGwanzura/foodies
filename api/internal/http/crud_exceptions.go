// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListNcrs implements the listNcrs operation.
func (s *Server) ListNcrs(ctx context.Context, req apigen.ListNcrsRequestObject) (apigen.ListNcrsResponseObject, error) {
	return listPage[apigen.ListNcrs200JSONResponse](ctx, req.Params, nil, s.service.ListNcr)
}

// GetNcr implements the getNcr operation.
func (s *Server) GetNcr(ctx context.Context, req apigen.GetNcrRequestObject) (apigen.GetNcrResponseObject, error) {
	row, err := s.service.GetNcr(ctx, req.NcrId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Ncr](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetNcr200JSONResponse{Body: body, Headers: apigen.GetNcr200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateNcr implements the createNcr operation.
func (s *Server) CreateNcr(ctx context.Context, req apigen.CreateNcrRequestObject) (apigen.CreateNcrResponseObject, error) {
	in, err := input[db.CreateNcrParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateNcr(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Ncr](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateNcr201JSONResponse{Body: body, Headers: apigen.CreateNcr201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateNcr implements the updateNcr operation.
func (s *Server) UpdateNcr(ctx context.Context, req apigen.UpdateNcrRequestObject) (apigen.UpdateNcrResponseObject, error) {
	in, err := input[db.UpdateNcrParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.NcrId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateNcr(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Ncr](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateNcr200JSONResponse{Body: body, Headers: apigen.UpdateNcr200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ContainNcr implements the containNcr operation.
func (s *Server) ContainNcr(ctx context.Context, req apigen.ContainNcrRequestObject) (apigen.ContainNcrResponseObject, error) {
	in, err := input[db.ActContainNcrParams](req.Body, map[string]any{"id": req.NcrId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ContainNcr(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Ncr](row)
	if err != nil {
		return nil, err
	}
	return apigen.ContainNcr200JSONResponse{Body: body, Headers: apigen.ContainNcr200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// VerifyCloseNcr implements the verifyCloseNcr operation.
func (s *Server) VerifyCloseNcr(ctx context.Context, req apigen.VerifyCloseNcrRequestObject) (apigen.VerifyCloseNcrResponseObject, error) {
	in, err := input[db.ActVerifyCloseNcrParams](req.Body, map[string]any{"id": req.NcrId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.VerifyCloseNcr(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Ncr](row)
	if err != nil {
		return nil, err
	}
	return apigen.VerifyCloseNcr200JSONResponse{Body: body, Headers: apigen.VerifyCloseNcr200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListHolds implements the listHolds operation.
func (s *Server) ListHolds(ctx context.Context, req apigen.ListHoldsRequestObject) (apigen.ListHoldsResponseObject, error) {
	return listPage[apigen.ListHolds200JSONResponse](ctx, req.Params, nil, s.service.ListHold)
}

// GetHold implements the getHold operation.
func (s *Server) GetHold(ctx context.Context, req apigen.GetHoldRequestObject) (apigen.GetHoldResponseObject, error) {
	row, err := s.service.GetHold(ctx, req.HoldId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Hold](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetHold200JSONResponse{Body: body, Headers: apigen.GetHold200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateHold implements the createHold operation.
func (s *Server) CreateHold(ctx context.Context, req apigen.CreateHoldRequestObject) (apigen.CreateHoldResponseObject, error) {
	in, err := input[db.CreateHoldParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateHold(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Hold](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateHold201JSONResponse{Body: body, Headers: apigen.CreateHold201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateHold implements the updateHold operation.
func (s *Server) UpdateHold(ctx context.Context, req apigen.UpdateHoldRequestObject) (apigen.UpdateHoldResponseObject, error) {
	in, err := input[db.UpdateHoldParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.HoldId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateHold(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Hold](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateHold200JSONResponse{Body: body, Headers: apigen.UpdateHold200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReleaseHold implements the releaseHold operation.
func (s *Server) ReleaseHold(ctx context.Context, req apigen.ReleaseHoldRequestObject) (apigen.ReleaseHoldResponseObject, error) {
	in, err := input[db.ActReleaseHoldParams](req.Body, map[string]any{"id": req.HoldId})
	if err != nil {
		return nil, err
	}
	if err := s.service.RequireSigned(ctx, "holds", req.HoldId, &req.Body.SignatureId); err != nil {
		return nil, err
	}
	row, err := s.service.ReleaseHold(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Hold](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReleaseHold200JSONResponse{Body: body, Headers: apigen.ReleaseHold200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDispositions implements the listDispositions operation.
func (s *Server) ListDispositions(ctx context.Context, req apigen.ListDispositionsRequestObject) (apigen.ListDispositionsResponseObject, error) {
	return listPage[apigen.ListDispositions200JSONResponse](ctx, req.Params, map[string]any{"hold_id": req.HoldId}, s.service.ListDisposition)
}

// GetDisposition implements the getDisposition operation.
func (s *Server) GetDisposition(ctx context.Context, req apigen.GetDispositionRequestObject) (apigen.GetDispositionResponseObject, error) {
	row, err := s.service.GetDisposition(ctx, req.DispositionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Disposition](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDisposition200JSONResponse{Body: body, Headers: apigen.GetDisposition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDisposition implements the createDisposition operation.
func (s *Server) CreateDisposition(ctx context.Context, req apigen.CreateDispositionRequestObject) (apigen.CreateDispositionResponseObject, error) {
	in, err := input[db.CreateDispositionParams](req.Body, map[string]any{"hold_id": req.HoldId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDisposition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Disposition](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDisposition201JSONResponse{Body: body, Headers: apigen.CreateDisposition201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitDisposition implements the submitDisposition operation.
func (s *Server) SubmitDisposition(ctx context.Context, req apigen.SubmitDispositionRequestObject) (apigen.SubmitDispositionResponseObject, error) {
	in, err := input[db.ActSubmitDispositionParams](struct{}{}, map[string]any{"id": req.DispositionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitDisposition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Disposition](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitDisposition200JSONResponse{Body: body, Headers: apigen.SubmitDisposition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ExecuteDisposition implements the executeDisposition operation.
func (s *Server) ExecuteDisposition(ctx context.Context, req apigen.ExecuteDispositionRequestObject) (apigen.ExecuteDispositionResponseObject, error) {
	in, err := input[db.ActExecuteDispositionParams](struct{}{}, map[string]any{"id": req.DispositionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ExecuteDisposition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Disposition](row)
	if err != nil {
		return nil, err
	}
	return apigen.ExecuteDisposition200JSONResponse{Body: body, Headers: apigen.ExecuteDisposition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRecalls implements the listRecalls operation.
func (s *Server) ListRecalls(ctx context.Context, req apigen.ListRecallsRequestObject) (apigen.ListRecallsResponseObject, error) {
	return listPage[apigen.ListRecalls200JSONResponse](ctx, req.Params, nil, s.service.ListRecall)
}

// GetRecall implements the getRecall operation.
func (s *Server) GetRecall(ctx context.Context, req apigen.GetRecallRequestObject) (apigen.GetRecallResponseObject, error) {
	row, err := s.service.GetRecall(ctx, req.RecallId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Recall](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRecall200JSONResponse{Body: body, Headers: apigen.GetRecall200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRecall implements the createRecall operation.
func (s *Server) CreateRecall(ctx context.Context, req apigen.CreateRecallRequestObject) (apigen.CreateRecallResponseObject, error) {
	in, err := input[db.CreateRecallParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRecall(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Recall](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRecall201JSONResponse{Body: body, Headers: apigen.CreateRecall201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRecall implements the updateRecall operation.
func (s *Server) UpdateRecall(ctx context.Context, req apigen.UpdateRecallRequestObject) (apigen.UpdateRecallResponseObject, error) {
	in, err := input[db.UpdateRecallParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.RecallId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRecall(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Recall](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRecall200JSONResponse{Body: body, Headers: apigen.UpdateRecall200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AssessRecall implements the assessRecall operation.
func (s *Server) AssessRecall(ctx context.Context, req apigen.AssessRecallRequestObject) (apigen.AssessRecallResponseObject, error) {
	in, err := input[db.ActAssessRecallParams](req.Body, map[string]any{"id": req.RecallId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AssessRecall(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Recall](row)
	if err != nil {
		return nil, err
	}
	return apigen.AssessRecall200JSONResponse{Body: body, Headers: apigen.AssessRecall200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DecideRecall implements the decideRecall operation.
func (s *Server) DecideRecall(ctx context.Context, req apigen.DecideRecallRequestObject) (apigen.DecideRecallResponseObject, error) {
	in, err := input[db.ActDecideRecallParams](req.Body, map[string]any{"id": req.RecallId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.DecideRecall(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Recall](row)
	if err != nil {
		return nil, err
	}
	return apigen.DecideRecall200JSONResponse{Body: body, Headers: apigen.DecideRecall200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseRecall implements the closeRecall operation.
func (s *Server) CloseRecall(ctx context.Context, req apigen.CloseRecallRequestObject) (apigen.CloseRecallResponseObject, error) {
	in, err := input[db.ActCloseRecallParams](req.Body, map[string]any{"id": req.RecallId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseRecall(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Recall](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseRecall200JSONResponse{Body: body, Headers: apigen.CloseRecall200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRecallCommunications implements the listRecallCommunications operation.
func (s *Server) ListRecallCommunications(ctx context.Context, req apigen.ListRecallCommunicationsRequestObject) (apigen.ListRecallCommunicationsResponseObject, error) {
	return listPage[apigen.ListRecallCommunications200JSONResponse](ctx, req.Params, map[string]any{"recall_id": req.RecallId}, s.service.ListRecallCommunication)
}

// GetRecallCommunication implements the getRecallCommunication operation.
func (s *Server) GetRecallCommunication(ctx context.Context, req apigen.GetRecallCommunicationRequestObject) (apigen.GetRecallCommunicationResponseObject, error) {
	row, err := s.service.GetRecallCommunication(ctx, req.RecallCommunicationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallCommunication](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRecallCommunication200JSONResponse{Body: body, Headers: apigen.GetRecallCommunication200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRecallCommunication implements the createRecallCommunication operation.
func (s *Server) CreateRecallCommunication(ctx context.Context, req apigen.CreateRecallCommunicationRequestObject) (apigen.CreateRecallCommunicationResponseObject, error) {
	in, err := input[db.CreateRecallCommunicationParams](req.Body, map[string]any{"recall_id": req.RecallId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRecallCommunication(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallCommunication](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRecallCommunication201JSONResponse{Body: body, Headers: apigen.CreateRecallCommunication201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitRecallCommunication implements the submitRecallCommunication operation.
func (s *Server) SubmitRecallCommunication(ctx context.Context, req apigen.SubmitRecallCommunicationRequestObject) (apigen.SubmitRecallCommunicationResponseObject, error) {
	in, err := input[db.ActSubmitRecallCommunicationParams](struct{}{}, map[string]any{"id": req.RecallCommunicationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitRecallCommunication(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallCommunication](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitRecallCommunication200JSONResponse{Body: body, Headers: apigen.SubmitRecallCommunication200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// MarkSentRecallCommunication implements the markSentRecallCommunication operation.
func (s *Server) MarkSentRecallCommunication(ctx context.Context, req apigen.MarkSentRecallCommunicationRequestObject) (apigen.MarkSentRecallCommunicationResponseObject, error) {
	in, err := input[db.ActMarkSentRecallCommunicationParams](struct{}{}, map[string]any{"id": req.RecallCommunicationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.MarkSentRecallCommunication(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallCommunication](row)
	if err != nil {
		return nil, err
	}
	return apigen.MarkSentRecallCommunication200JSONResponse{Body: body, Headers: apigen.MarkSentRecallCommunication200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AcknowledgeRecallCommunication implements the acknowledgeRecallCommunication operation.
func (s *Server) AcknowledgeRecallCommunication(ctx context.Context, req apigen.AcknowledgeRecallCommunicationRequestObject) (apigen.AcknowledgeRecallCommunicationResponseObject, error) {
	in, err := input[db.ActAcknowledgeRecallCommunicationParams](struct{}{}, map[string]any{"id": req.RecallCommunicationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AcknowledgeRecallCommunication(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallCommunication](row)
	if err != nil {
		return nil, err
	}
	return apigen.AcknowledgeRecallCommunication200JSONResponse{Body: body, Headers: apigen.AcknowledgeRecallCommunication200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRecallRetrievals implements the listRecallRetrievals operation.
func (s *Server) ListRecallRetrievals(ctx context.Context, req apigen.ListRecallRetrievalsRequestObject) (apigen.ListRecallRetrievalsResponseObject, error) {
	return listPage[apigen.ListRecallRetrievals200JSONResponse](ctx, req.Params, map[string]any{"recall_id": req.RecallId}, s.service.ListRecallRetrieval)
}

// GetRecallRetrieval implements the getRecallRetrieval operation.
func (s *Server) GetRecallRetrieval(ctx context.Context, req apigen.GetRecallRetrievalRequestObject) (apigen.GetRecallRetrievalResponseObject, error) {
	row, err := s.service.GetRecallRetrieval(ctx, req.RecallRetrievalId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallRetrieval](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRecallRetrieval200JSONResponse{Body: body, Headers: apigen.GetRecallRetrieval200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRecallRetrieval implements the createRecallRetrieval operation.
func (s *Server) CreateRecallRetrieval(ctx context.Context, req apigen.CreateRecallRetrievalRequestObject) (apigen.CreateRecallRetrievalResponseObject, error) {
	in, err := input[db.CreateRecallRetrievalParams](req.Body, map[string]any{"recall_id": req.RecallId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRecallRetrieval(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallRetrieval](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRecallRetrieval201JSONResponse{Body: body, Headers: apigen.CreateRecallRetrieval201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRecallRetrieval implements the updateRecallRetrieval operation.
func (s *Server) UpdateRecallRetrieval(ctx context.Context, req apigen.UpdateRecallRetrievalRequestObject) (apigen.UpdateRecallRetrievalResponseObject, error) {
	in, err := input[db.UpdateRecallRetrievalParams](req.Body, map[string]any{"recall_id": req.RecallId})
	if err != nil {
		return nil, err
	}
	in.ID = req.RecallRetrievalId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRecallRetrieval(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallRetrieval](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRecallRetrieval200JSONResponse{Body: body, Headers: apigen.UpdateRecallRetrieval200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRecallReconciliations implements the listRecallReconciliations operation.
func (s *Server) ListRecallReconciliations(ctx context.Context, req apigen.ListRecallReconciliationsRequestObject) (apigen.ListRecallReconciliationsResponseObject, error) {
	return listPage[apigen.ListRecallReconciliations200JSONResponse](ctx, req.Params, map[string]any{"recall_id": req.RecallId}, s.service.ListRecallReconciliation)
}

// GetRecallReconciliation implements the getRecallReconciliation operation.
func (s *Server) GetRecallReconciliation(ctx context.Context, req apigen.GetRecallReconciliationRequestObject) (apigen.GetRecallReconciliationResponseObject, error) {
	row, err := s.service.GetRecallReconciliation(ctx, req.RecallReconciliationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRecallReconciliation200JSONResponse{Body: body, Headers: apigen.GetRecallReconciliation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRecallReconciliation implements the createRecallReconciliation operation.
func (s *Server) CreateRecallReconciliation(ctx context.Context, req apigen.CreateRecallReconciliationRequestObject) (apigen.CreateRecallReconciliationResponseObject, error) {
	in, err := input[db.CreateRecallReconciliationParams](req.Body, map[string]any{"recall_id": req.RecallId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRecallReconciliation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRecallReconciliation201JSONResponse{Body: body, Headers: apigen.CreateRecallReconciliation201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRecallReconciliation implements the updateRecallReconciliation operation.
func (s *Server) UpdateRecallReconciliation(ctx context.Context, req apigen.UpdateRecallReconciliationRequestObject) (apigen.UpdateRecallReconciliationResponseObject, error) {
	in, err := input[db.UpdateRecallReconciliationParams](req.Body, map[string]any{"recall_id": req.RecallId})
	if err != nil {
		return nil, err
	}
	in.ID = req.RecallReconciliationId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRecallReconciliation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRecallReconciliation200JSONResponse{Body: body, Headers: apigen.UpdateRecallReconciliation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReviewRecallReconciliation implements the reviewRecallReconciliation operation.
func (s *Server) ReviewRecallReconciliation(ctx context.Context, req apigen.ReviewRecallReconciliationRequestObject) (apigen.ReviewRecallReconciliationResponseObject, error) {
	in, err := input[db.ActReviewRecallReconciliationParams](req.Body, map[string]any{"id": req.RecallReconciliationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReviewRecallReconciliation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecallReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReviewRecallReconciliation200JSONResponse{Body: body, Headers: apigen.ReviewRecallReconciliation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListIncidents implements the listIncidents operation.
func (s *Server) ListIncidents(ctx context.Context, req apigen.ListIncidentsRequestObject) (apigen.ListIncidentsResponseObject, error) {
	return listPage[apigen.ListIncidents200JSONResponse](ctx, req.Params, nil, s.service.ListIncident)
}

// GetIncident implements the getIncident operation.
func (s *Server) GetIncident(ctx context.Context, req apigen.GetIncidentRequestObject) (apigen.GetIncidentResponseObject, error) {
	row, err := s.service.GetIncident(ctx, req.IncidentId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Incident](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetIncident200JSONResponse{Body: body, Headers: apigen.GetIncident200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateIncident implements the createIncident operation.
func (s *Server) CreateIncident(ctx context.Context, req apigen.CreateIncidentRequestObject) (apigen.CreateIncidentResponseObject, error) {
	in, err := input[db.CreateIncidentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateIncident(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Incident](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateIncident201JSONResponse{Body: body, Headers: apigen.CreateIncident201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateIncident implements the updateIncident operation.
func (s *Server) UpdateIncident(ctx context.Context, req apigen.UpdateIncidentRequestObject) (apigen.UpdateIncidentResponseObject, error) {
	in, err := input[db.UpdateIncidentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.IncidentId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateIncident(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Incident](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateIncident200JSONResponse{Body: body, Headers: apigen.UpdateIncident200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseIncident implements the closeIncident operation.
func (s *Server) CloseIncident(ctx context.Context, req apigen.CloseIncidentRequestObject) (apigen.CloseIncidentResponseObject, error) {
	in, err := input[db.ActCloseIncidentParams](req.Body, map[string]any{"id": req.IncidentId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseIncident(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Incident](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseIncident200JSONResponse{Body: body, Headers: apigen.CloseIncident200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListComplaints implements the listComplaints operation.
func (s *Server) ListComplaints(ctx context.Context, req apigen.ListComplaintsRequestObject) (apigen.ListComplaintsResponseObject, error) {
	return listPage[apigen.ListComplaints200JSONResponse](ctx, req.Params, nil, s.service.ListComplaint)
}

// GetComplaint implements the getComplaint operation.
func (s *Server) GetComplaint(ctx context.Context, req apigen.GetComplaintRequestObject) (apigen.GetComplaintResponseObject, error) {
	row, err := s.service.GetComplaint(ctx, req.ComplaintId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Complaint](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetComplaint200JSONResponse{Body: body, Headers: apigen.GetComplaint200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateComplaint implements the createComplaint operation.
func (s *Server) CreateComplaint(ctx context.Context, req apigen.CreateComplaintRequestObject) (apigen.CreateComplaintResponseObject, error) {
	in, err := input[db.CreateComplaintParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateComplaint(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Complaint](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateComplaint201JSONResponse{Body: body, Headers: apigen.CreateComplaint201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateComplaint implements the updateComplaint operation.
func (s *Server) UpdateComplaint(ctx context.Context, req apigen.UpdateComplaintRequestObject) (apigen.UpdateComplaintResponseObject, error) {
	in, err := input[db.UpdateComplaintParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ComplaintId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateComplaint(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Complaint](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateComplaint200JSONResponse{Body: body, Headers: apigen.UpdateComplaint200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// TriageComplaint implements the triageComplaint operation.
func (s *Server) TriageComplaint(ctx context.Context, req apigen.TriageComplaintRequestObject) (apigen.TriageComplaintResponseObject, error) {
	in, err := input[db.ActTriageComplaintParams](req.Body, map[string]any{"id": req.ComplaintId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.TriageComplaint(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Complaint](row)
	if err != nil {
		return nil, err
	}
	return apigen.TriageComplaint200JSONResponse{Body: body, Headers: apigen.TriageComplaint200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ResolveComplaint implements the resolveComplaint operation.
func (s *Server) ResolveComplaint(ctx context.Context, req apigen.ResolveComplaintRequestObject) (apigen.ResolveComplaintResponseObject, error) {
	in, err := input[db.ActResolveComplaintParams](req.Body, map[string]any{"id": req.ComplaintId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ResolveComplaint(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Complaint](row)
	if err != nil {
		return nil, err
	}
	return apigen.ResolveComplaint200JSONResponse{Body: body, Headers: apigen.ResolveComplaint200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseComplaint implements the closeComplaint operation.
func (s *Server) CloseComplaint(ctx context.Context, req apigen.CloseComplaintRequestObject) (apigen.CloseComplaintResponseObject, error) {
	in, err := input[db.ActCloseComplaintParams](struct{}{}, map[string]any{"id": req.ComplaintId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseComplaint(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Complaint](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseComplaint200JSONResponse{Body: body, Headers: apigen.CloseComplaint200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListServiceRecoveries implements the listServiceRecoveries operation.
func (s *Server) ListServiceRecoveries(ctx context.Context, req apigen.ListServiceRecoveriesRequestObject) (apigen.ListServiceRecoveriesResponseObject, error) {
	return listPage[apigen.ListServiceRecoveries200JSONResponse](ctx, req.Params, map[string]any{"complaint_id": req.ComplaintId}, s.service.ListServiceRecovery)
}

// GetServiceRecovery implements the getServiceRecovery operation.
func (s *Server) GetServiceRecovery(ctx context.Context, req apigen.GetServiceRecoveryRequestObject) (apigen.GetServiceRecoveryResponseObject, error) {
	row, err := s.service.GetServiceRecovery(ctx, req.ServiceRecoveryId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ServiceRecovery](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetServiceRecovery200JSONResponse{Body: body, Headers: apigen.GetServiceRecovery200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateServiceRecovery implements the createServiceRecovery operation.
func (s *Server) CreateServiceRecovery(ctx context.Context, req apigen.CreateServiceRecoveryRequestObject) (apigen.CreateServiceRecoveryResponseObject, error) {
	in, err := input[db.CreateServiceRecoveryParams](req.Body, map[string]any{"complaint_id": req.ComplaintId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateServiceRecovery(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ServiceRecovery](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateServiceRecovery201JSONResponse{Body: body, Headers: apigen.CreateServiceRecovery201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitServiceRecovery implements the submitServiceRecovery operation.
func (s *Server) SubmitServiceRecovery(ctx context.Context, req apigen.SubmitServiceRecoveryRequestObject) (apigen.SubmitServiceRecoveryResponseObject, error) {
	in, err := input[db.ActSubmitServiceRecoveryParams](struct{}{}, map[string]any{"id": req.ServiceRecoveryId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitServiceRecovery(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ServiceRecovery](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitServiceRecovery200JSONResponse{Body: body, Headers: apigen.SubmitServiceRecovery200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ExecuteServiceRecovery implements the executeServiceRecovery operation.
func (s *Server) ExecuteServiceRecovery(ctx context.Context, req apigen.ExecuteServiceRecoveryRequestObject) (apigen.ExecuteServiceRecoveryResponseObject, error) {
	in, err := input[db.ActExecuteServiceRecoveryParams](req.Body, map[string]any{"id": req.ServiceRecoveryId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ExecuteServiceRecovery(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ServiceRecovery](row)
	if err != nil {
		return nil, err
	}
	return apigen.ExecuteServiceRecovery200JSONResponse{Body: body, Headers: apigen.ExecuteServiceRecovery200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPaymentIncidents implements the listPaymentIncidents operation.
func (s *Server) ListPaymentIncidents(ctx context.Context, req apigen.ListPaymentIncidentsRequestObject) (apigen.ListPaymentIncidentsResponseObject, error) {
	return listPage[apigen.ListPaymentIncidents200JSONResponse](ctx, req.Params, nil, s.service.ListPaymentIncident)
}

// GetPaymentIncident implements the getPaymentIncident operation.
func (s *Server) GetPaymentIncident(ctx context.Context, req apigen.GetPaymentIncidentRequestObject) (apigen.GetPaymentIncidentResponseObject, error) {
	row, err := s.service.GetPaymentIncident(ctx, req.PaymentIncidentId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.PaymentIncidentChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentIncident](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetPaymentIncident200JSONResponse{Body: body, Headers: apigen.GetPaymentIncident200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePaymentIncident implements the createPaymentIncident operation.
func (s *Server) CreatePaymentIncident(ctx context.Context, req apigen.CreatePaymentIncidentRequestObject) (apigen.CreatePaymentIncidentResponseObject, error) {
	in, err := input[db.CreatePaymentIncidentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePaymentIncident(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["updates"]; ok {
		rows, err := childInputs[db.InsertPaymentIncidentUpdatesParams](items, "payment_incident_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePaymentIncidentUpdates(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PaymentIncidentChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentIncident](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePaymentIncident201JSONResponse{Body: body, Headers: apigen.CreatePaymentIncident201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePaymentIncident implements the updatePaymentIncident operation.
func (s *Server) UpdatePaymentIncident(ctx context.Context, req apigen.UpdatePaymentIncidentRequestObject) (apigen.UpdatePaymentIncidentResponseObject, error) {
	in, err := input[db.UpdatePaymentIncidentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PaymentIncidentId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePaymentIncident(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["updates"]; ok {
		rows, err := childInputs[db.InsertPaymentIncidentUpdatesParams](items, "payment_incident_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePaymentIncidentUpdates(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PaymentIncidentChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentIncident](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePaymentIncident200JSONResponse{Body: body, Headers: apigen.UpdatePaymentIncident200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// VerifyStatusPaymentIncident implements the verifyStatusPaymentIncident operation.
func (s *Server) VerifyStatusPaymentIncident(ctx context.Context, req apigen.VerifyStatusPaymentIncidentRequestObject) (apigen.VerifyStatusPaymentIncidentResponseObject, error) {
	in, err := input[db.ActVerifyStatusPaymentIncidentParams](req.Body, map[string]any{"id": req.PaymentIncidentId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.VerifyStatusPaymentIncident(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentIncident](row)
	if err != nil {
		return nil, err
	}
	return apigen.VerifyStatusPaymentIncident200JSONResponse{Body: body, Headers: apigen.VerifyStatusPaymentIncident200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ProposeCorrectionPaymentIncident implements the proposeCorrectionPaymentIncident operation.
func (s *Server) ProposeCorrectionPaymentIncident(ctx context.Context, req apigen.ProposeCorrectionPaymentIncidentRequestObject) (apigen.ProposeCorrectionPaymentIncidentResponseObject, error) {
	in, err := input[db.ActProposeCorrectionPaymentIncidentParams](req.Body, map[string]any{"id": req.PaymentIncidentId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ProposeCorrectionPaymentIncident(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentIncident](row)
	if err != nil {
		return nil, err
	}
	return apigen.ProposeCorrectionPaymentIncident200JSONResponse{Body: body, Headers: apigen.ProposeCorrectionPaymentIncident200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReconcilePaymentIncident implements the reconcilePaymentIncident operation.
func (s *Server) ReconcilePaymentIncident(ctx context.Context, req apigen.ReconcilePaymentIncidentRequestObject) (apigen.ReconcilePaymentIncidentResponseObject, error) {
	in, err := input[db.ActReconcilePaymentIncidentParams](req.Body, map[string]any{"id": req.PaymentIncidentId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReconcilePaymentIncident(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentIncident](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReconcilePaymentIncident200JSONResponse{Body: body, Headers: apigen.ReconcilePaymentIncident200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ClosePaymentIncident implements the closePaymentIncident operation.
func (s *Server) ClosePaymentIncident(ctx context.Context, req apigen.ClosePaymentIncidentRequestObject) (apigen.ClosePaymentIncidentResponseObject, error) {
	in, err := input[db.ActClosePaymentIncidentParams](struct{}{}, map[string]any{"id": req.PaymentIncidentId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ClosePaymentIncident(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentIncident](row)
	if err != nil {
		return nil, err
	}
	return apigen.ClosePaymentIncident200JSONResponse{Body: body, Headers: apigen.ClosePaymentIncident200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPosOutages implements the listPosOutages operation.
func (s *Server) ListPosOutages(ctx context.Context, req apigen.ListPosOutagesRequestObject) (apigen.ListPosOutagesResponseObject, error) {
	return listPage[apigen.ListPosOutages200JSONResponse](ctx, req.Params, nil, s.service.ListPosOutage)
}

// GetPosOutage implements the getPosOutage operation.
func (s *Server) GetPosOutage(ctx context.Context, req apigen.GetPosOutageRequestObject) (apigen.GetPosOutageResponseObject, error) {
	row, err := s.service.GetPosOutage(ctx, req.PosOutageId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PosOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPosOutage200JSONResponse{Body: body, Headers: apigen.GetPosOutage200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePosOutage implements the createPosOutage operation.
func (s *Server) CreatePosOutage(ctx context.Context, req apigen.CreatePosOutageRequestObject) (apigen.CreatePosOutageResponseObject, error) {
	in, err := input[db.CreatePosOutageParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePosOutage(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PosOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePosOutage201JSONResponse{Body: body, Headers: apigen.CreatePosOutage201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePosOutage implements the updatePosOutage operation.
func (s *Server) UpdatePosOutage(ctx context.Context, req apigen.UpdatePosOutageRequestObject) (apigen.UpdatePosOutageResponseObject, error) {
	in, err := input[db.UpdatePosOutageParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PosOutageId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePosOutage(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PosOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePosOutage200JSONResponse{Body: body, Headers: apigen.UpdatePosOutage200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitPosOutage implements the submitPosOutage operation.
func (s *Server) SubmitPosOutage(ctx context.Context, req apigen.SubmitPosOutageRequestObject) (apigen.SubmitPosOutageResponseObject, error) {
	in, err := input[db.ActSubmitPosOutageParams](struct{}{}, map[string]any{"id": req.PosOutageId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitPosOutage(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PosOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitPosOutage200JSONResponse{Body: body, Headers: apigen.SubmitPosOutage200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RestorePosOutage implements the restorePosOutage operation.
func (s *Server) RestorePosOutage(ctx context.Context, req apigen.RestorePosOutageRequestObject) (apigen.RestorePosOutageResponseObject, error) {
	in, err := input[db.ActRestorePosOutageParams](req.Body, map[string]any{"id": req.PosOutageId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RestorePosOutage(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PosOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.RestorePosOutage200JSONResponse{Body: body, Headers: apigen.RestorePosOutage200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ClosePosOutage implements the closePosOutage operation.
func (s *Server) ClosePosOutage(ctx context.Context, req apigen.ClosePosOutageRequestObject) (apigen.ClosePosOutageResponseObject, error) {
	in, err := input[db.ActClosePosOutageParams](req.Body, map[string]any{"id": req.PosOutageId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ClosePosOutage(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PosOutage](row)
	if err != nil {
		return nil, err
	}
	return apigen.ClosePosOutage200JSONResponse{Body: body, Headers: apigen.ClosePosOutage200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListReceiptBooks implements the listReceiptBooks operation.
func (s *Server) ListReceiptBooks(ctx context.Context, req apigen.ListReceiptBooksRequestObject) (apigen.ListReceiptBooksResponseObject, error) {
	return listPage[apigen.ListReceiptBooks200JSONResponse](ctx, req.Params, nil, s.service.ListReceiptBook)
}

// GetReceiptBook implements the getReceiptBook operation.
func (s *Server) GetReceiptBook(ctx context.Context, req apigen.GetReceiptBookRequestObject) (apigen.GetReceiptBookResponseObject, error) {
	row, err := s.service.GetReceiptBook(ctx, req.ReceiptBookId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReceiptBook](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetReceiptBook200JSONResponse{Body: body, Headers: apigen.GetReceiptBook200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateReceiptBook implements the createReceiptBook operation.
func (s *Server) CreateReceiptBook(ctx context.Context, req apigen.CreateReceiptBookRequestObject) (apigen.CreateReceiptBookResponseObject, error) {
	in, err := input[db.CreateReceiptBookParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateReceiptBook(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReceiptBook](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateReceiptBook201JSONResponse{Body: body, Headers: apigen.CreateReceiptBook201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateReceiptBook implements the updateReceiptBook operation.
func (s *Server) UpdateReceiptBook(ctx context.Context, req apigen.UpdateReceiptBookRequestObject) (apigen.UpdateReceiptBookResponseObject, error) {
	in, err := input[db.UpdateReceiptBookParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ReceiptBookId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateReceiptBook(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReceiptBook](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateReceiptBook200JSONResponse{Body: body, Headers: apigen.UpdateReceiptBook200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// IssueReceiptBook implements the issueReceiptBook operation.
func (s *Server) IssueReceiptBook(ctx context.Context, req apigen.IssueReceiptBookRequestObject) (apigen.IssueReceiptBookResponseObject, error) {
	in, err := input[db.ActIssueReceiptBookParams](req.Body, map[string]any{"id": req.ReceiptBookId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.IssueReceiptBook(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReceiptBook](row)
	if err != nil {
		return nil, err
	}
	return apigen.IssueReceiptBook200JSONResponse{Body: body, Headers: apigen.IssueReceiptBook200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReturnReceiptBook implements the returnReceiptBook operation.
func (s *Server) ReturnReceiptBook(ctx context.Context, req apigen.ReturnReceiptBookRequestObject) (apigen.ReturnReceiptBookResponseObject, error) {
	in, err := input[db.ActReturnReceiptBookParams](req.Body, map[string]any{"id": req.ReceiptBookId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReturnReceiptBook(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReceiptBook](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReturnReceiptBook200JSONResponse{Body: body, Headers: apigen.ReturnReceiptBook200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListManualReceipts implements the listManualReceipts operation.
func (s *Server) ListManualReceipts(ctx context.Context, req apigen.ListManualReceiptsRequestObject) (apigen.ListManualReceiptsResponseObject, error) {
	return listPage[apigen.ListManualReceipts200JSONResponse](ctx, req.Params, map[string]any{"pos_outage_id": req.PosOutageId}, s.service.ListManualReceipt)
}

// GetManualReceipt implements the getManualReceipt operation.
func (s *Server) GetManualReceipt(ctx context.Context, req apigen.GetManualReceiptRequestObject) (apigen.GetManualReceiptResponseObject, error) {
	row, err := s.service.GetManualReceipt(ctx, req.ManualReceiptId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ManualReceipt](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetManualReceipt200JSONResponse{Body: body, Headers: apigen.GetManualReceipt200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateManualReceipt implements the createManualReceipt operation.
func (s *Server) CreateManualReceipt(ctx context.Context, req apigen.CreateManualReceiptRequestObject) (apigen.CreateManualReceiptResponseObject, error) {
	in, err := input[db.CreateManualReceiptParams](req.Body, map[string]any{"pos_outage_id": req.PosOutageId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateManualReceipt(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ManualReceipt](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateManualReceipt201JSONResponse{Body: body, Headers: apigen.CreateManualReceipt201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CancelManualReceipt implements the cancelManualReceipt operation.
func (s *Server) CancelManualReceipt(ctx context.Context, req apigen.CancelManualReceiptRequestObject) (apigen.CancelManualReceiptResponseObject, error) {
	in, err := input[db.ActCancelManualReceiptParams](req.Body, map[string]any{"id": req.ManualReceiptId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CancelManualReceipt(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ManualReceipt](row)
	if err != nil {
		return nil, err
	}
	return apigen.CancelManualReceipt200JSONResponse{Body: body, Headers: apigen.CancelManualReceipt200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// BackCaptureManualReceipt implements the backCaptureManualReceipt operation.
func (s *Server) BackCaptureManualReceipt(ctx context.Context, req apigen.BackCaptureManualReceiptRequestObject) (apigen.BackCaptureManualReceiptResponseObject, error) {
	in, err := input[db.ActBackCaptureManualReceiptParams](req.Body, map[string]any{"id": req.ManualReceiptId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.BackCaptureManualReceipt(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ManualReceipt](row)
	if err != nil {
		return nil, err
	}
	return apigen.BackCaptureManualReceipt200JSONResponse{Body: body, Headers: apigen.BackCaptureManualReceipt200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListEquipmentFaults implements the listEquipmentFaults operation.
func (s *Server) ListEquipmentFaults(ctx context.Context, req apigen.ListEquipmentFaultsRequestObject) (apigen.ListEquipmentFaultsResponseObject, error) {
	return listPage[apigen.ListEquipmentFaults200JSONResponse](ctx, req.Params, nil, s.service.ListEquipmentFault)
}

// GetEquipmentFault implements the getEquipmentFault operation.
func (s *Server) GetEquipmentFault(ctx context.Context, req apigen.GetEquipmentFaultRequestObject) (apigen.GetEquipmentFaultResponseObject, error) {
	row, err := s.service.GetEquipmentFault(ctx, req.EquipmentFaultId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EquipmentFault](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetEquipmentFault200JSONResponse{Body: body, Headers: apigen.GetEquipmentFault200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateEquipmentFault implements the createEquipmentFault operation.
func (s *Server) CreateEquipmentFault(ctx context.Context, req apigen.CreateEquipmentFaultRequestObject) (apigen.CreateEquipmentFaultResponseObject, error) {
	in, err := input[db.CreateEquipmentFaultParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateEquipmentFault(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EquipmentFault](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateEquipmentFault201JSONResponse{Body: body, Headers: apigen.CreateEquipmentFault201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateEquipmentFault implements the updateEquipmentFault operation.
func (s *Server) UpdateEquipmentFault(ctx context.Context, req apigen.UpdateEquipmentFaultRequestObject) (apigen.UpdateEquipmentFaultResponseObject, error) {
	in, err := input[db.UpdateEquipmentFaultParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.EquipmentFaultId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateEquipmentFault(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EquipmentFault](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateEquipmentFault200JSONResponse{Body: body, Headers: apigen.UpdateEquipmentFault200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ConvertEquipmentFault implements the convertEquipmentFault operation.
func (s *Server) ConvertEquipmentFault(ctx context.Context, req apigen.ConvertEquipmentFaultRequestObject) (apigen.ConvertEquipmentFaultResponseObject, error) {
	in, err := input[db.ActConvertEquipmentFaultParams](req.Body, map[string]any{"id": req.EquipmentFaultId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ConvertEquipmentFault(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EquipmentFault](row)
	if err != nil {
		return nil, err
	}
	return apigen.ConvertEquipmentFault200JSONResponse{Body: body, Headers: apigen.ConvertEquipmentFault200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListWorkOrders implements the listWorkOrders operation.
func (s *Server) ListWorkOrders(ctx context.Context, req apigen.ListWorkOrdersRequestObject) (apigen.ListWorkOrdersResponseObject, error) {
	return listPage[apigen.ListWorkOrders200JSONResponse](ctx, req.Params, nil, s.service.ListWorkOrder)
}

// GetWorkOrder implements the getWorkOrder operation.
func (s *Server) GetWorkOrder(ctx context.Context, req apigen.GetWorkOrderRequestObject) (apigen.GetWorkOrderResponseObject, error) {
	row, err := s.service.GetWorkOrder(ctx, req.WorkOrderId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetWorkOrder200JSONResponse{Body: body, Headers: apigen.GetWorkOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateWorkOrder implements the createWorkOrder operation.
func (s *Server) CreateWorkOrder(ctx context.Context, req apigen.CreateWorkOrderRequestObject) (apigen.CreateWorkOrderResponseObject, error) {
	in, err := input[db.CreateWorkOrderParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateWorkOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateWorkOrder201JSONResponse{Body: body, Headers: apigen.CreateWorkOrder201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateWorkOrder implements the updateWorkOrder operation.
func (s *Server) UpdateWorkOrder(ctx context.Context, req apigen.UpdateWorkOrderRequestObject) (apigen.UpdateWorkOrderResponseObject, error) {
	in, err := input[db.UpdateWorkOrderParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.WorkOrderId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateWorkOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateWorkOrder200JSONResponse{Body: body, Headers: apigen.UpdateWorkOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AssignWorkOrder implements the assignWorkOrder operation.
func (s *Server) AssignWorkOrder(ctx context.Context, req apigen.AssignWorkOrderRequestObject) (apigen.AssignWorkOrderResponseObject, error) {
	in, err := input[db.ActAssignWorkOrderParams](req.Body, map[string]any{"id": req.WorkOrderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AssignWorkOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.AssignWorkOrder200JSONResponse{Body: body, Headers: apigen.AssignWorkOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// StartWorkOrder implements the startWorkOrder operation.
func (s *Server) StartWorkOrder(ctx context.Context, req apigen.StartWorkOrderRequestObject) (apigen.StartWorkOrderResponseObject, error) {
	in, err := input[db.ActStartWorkOrderParams](struct{}{}, map[string]any{"id": req.WorkOrderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.StartWorkOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.StartWorkOrder200JSONResponse{Body: body, Headers: apigen.StartWorkOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteWorkOrder implements the completeWorkOrder operation.
func (s *Server) CompleteWorkOrder(ctx context.Context, req apigen.CompleteWorkOrderRequestObject) (apigen.CompleteWorkOrderResponseObject, error) {
	in, err := input[db.ActCompleteWorkOrderParams](req.Body, map[string]any{"id": req.WorkOrderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteWorkOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteWorkOrder200JSONResponse{Body: body, Headers: apigen.CompleteWorkOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ClearWorkOrder implements the clearWorkOrder operation.
func (s *Server) ClearWorkOrder(ctx context.Context, req apigen.ClearWorkOrderRequestObject) (apigen.ClearWorkOrderResponseObject, error) {
	in, err := input[db.ActClearWorkOrderParams](req.Body, map[string]any{"id": req.WorkOrderId})
	if err != nil {
		return nil, err
	}
	if err := s.service.RequireSigned(ctx, "work_orders", req.WorkOrderId, &req.Body.SignatureId); err != nil {
		return nil, err
	}
	row, err := s.service.ClearWorkOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.ClearWorkOrder200JSONResponse{Body: body, Headers: apigen.ClearWorkOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListMaintenancePlans implements the listMaintenancePlans operation.
func (s *Server) ListMaintenancePlans(ctx context.Context, req apigen.ListMaintenancePlansRequestObject) (apigen.ListMaintenancePlansResponseObject, error) {
	return listPage[apigen.ListMaintenancePlans200JSONResponse](ctx, req.Params, nil, s.service.ListMaintenancePlan)
}

// GetMaintenancePlan implements the getMaintenancePlan operation.
func (s *Server) GetMaintenancePlan(ctx context.Context, req apigen.GetMaintenancePlanRequestObject) (apigen.GetMaintenancePlanResponseObject, error) {
	row, err := s.service.GetMaintenancePlan(ctx, req.MaintenancePlanId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.MaintenancePlan](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetMaintenancePlan200JSONResponse{Body: body, Headers: apigen.GetMaintenancePlan200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateMaintenancePlan implements the createMaintenancePlan operation.
func (s *Server) CreateMaintenancePlan(ctx context.Context, req apigen.CreateMaintenancePlanRequestObject) (apigen.CreateMaintenancePlanResponseObject, error) {
	in, err := input[db.CreateMaintenancePlanParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateMaintenancePlan(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.MaintenancePlan](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateMaintenancePlan201JSONResponse{Body: body, Headers: apigen.CreateMaintenancePlan201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateMaintenancePlan implements the updateMaintenancePlan operation.
func (s *Server) UpdateMaintenancePlan(ctx context.Context, req apigen.UpdateMaintenancePlanRequestObject) (apigen.UpdateMaintenancePlanResponseObject, error) {
	in, err := input[db.UpdateMaintenancePlanParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.MaintenancePlanId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateMaintenancePlan(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.MaintenancePlan](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateMaintenancePlan200JSONResponse{Body: body, Headers: apigen.UpdateMaintenancePlan200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
