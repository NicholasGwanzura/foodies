// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListFloatCertificates implements the listFloatCertificates operation.
func (s *Server) ListFloatCertificates(ctx context.Context, req apigen.ListFloatCertificatesRequestObject) (apigen.ListFloatCertificatesResponseObject, error) {
	return listPage[apigen.ListFloatCertificates200JSONResponse](ctx, req.Params, nil, s.service.ListFloatCertificate)
}

// GetFloatCertificate implements the getFloatCertificate operation.
func (s *Server) GetFloatCertificate(ctx context.Context, req apigen.GetFloatCertificateRequestObject) (apigen.GetFloatCertificateResponseObject, error) {
	row, err := s.service.GetFloatCertificate(ctx, req.FloatCertificateId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatCertificate](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetFloatCertificate200JSONResponse{Body: body, Headers: apigen.GetFloatCertificate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateFloatCertificate implements the createFloatCertificate operation.
func (s *Server) CreateFloatCertificate(ctx context.Context, req apigen.CreateFloatCertificateRequestObject) (apigen.CreateFloatCertificateResponseObject, error) {
	in, err := input[db.CreateFloatCertificateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateFloatCertificate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatCertificate](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateFloatCertificate201JSONResponse{Body: body, Headers: apigen.CreateFloatCertificate201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitFloatCertificate implements the submitFloatCertificate operation.
func (s *Server) SubmitFloatCertificate(ctx context.Context, req apigen.SubmitFloatCertificateRequestObject) (apigen.SubmitFloatCertificateResponseObject, error) {
	in, err := input[db.ActSubmitFloatCertificateParams](struct{}{}, map[string]any{"id": req.FloatCertificateId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitFloatCertificate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatCertificate](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitFloatCertificate200JSONResponse{Body: body, Headers: apigen.SubmitFloatCertificate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListFloatIssues implements the listFloatIssues operation.
func (s *Server) ListFloatIssues(ctx context.Context, req apigen.ListFloatIssuesRequestObject) (apigen.ListFloatIssuesResponseObject, error) {
	return listPage[apigen.ListFloatIssues200JSONResponse](ctx, req.Params, nil, s.service.ListFloatIssue)
}

// GetFloatIssue implements the getFloatIssue operation.
func (s *Server) GetFloatIssue(ctx context.Context, req apigen.GetFloatIssueRequestObject) (apigen.GetFloatIssueResponseObject, error) {
	row, err := s.service.GetFloatIssue(ctx, req.FloatIssueId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatIssue](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetFloatIssue200JSONResponse{Body: body, Headers: apigen.GetFloatIssue200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateFloatIssue implements the createFloatIssue operation.
func (s *Server) CreateFloatIssue(ctx context.Context, req apigen.CreateFloatIssueRequestObject) (apigen.CreateFloatIssueResponseObject, error) {
	in, err := input[db.CreateFloatIssueParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateFloatIssue(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatIssue](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateFloatIssue201JSONResponse{Body: body, Headers: apigen.CreateFloatIssue201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AcknowledgeFloatIssue implements the acknowledgeFloatIssue operation.
func (s *Server) AcknowledgeFloatIssue(ctx context.Context, req apigen.AcknowledgeFloatIssueRequestObject) (apigen.AcknowledgeFloatIssueResponseObject, error) {
	in, err := input[db.ActAcknowledgeFloatIssueParams](struct{}{}, map[string]any{"id": req.FloatIssueId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AcknowledgeFloatIssue(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatIssue](row)
	if err != nil {
		return nil, err
	}
	return apigen.AcknowledgeFloatIssue200JSONResponse{Body: body, Headers: apigen.AcknowledgeFloatIssue200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReturnFloatIssue implements the returnFloatIssue operation.
func (s *Server) ReturnFloatIssue(ctx context.Context, req apigen.ReturnFloatIssueRequestObject) (apigen.ReturnFloatIssueResponseObject, error) {
	in, err := input[db.ActReturnFloatIssueParams](req.Body, map[string]any{"id": req.FloatIssueId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReturnFloatIssue(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatIssue](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReturnFloatIssue200JSONResponse{Body: body, Headers: apigen.ReturnFloatIssue200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListFloatExceptions implements the listFloatExceptions operation.
func (s *Server) ListFloatExceptions(ctx context.Context, req apigen.ListFloatExceptionsRequestObject) (apigen.ListFloatExceptionsResponseObject, error) {
	return listPage[apigen.ListFloatExceptions200JSONResponse](ctx, req.Params, nil, s.service.ListFloatException)
}

// GetFloatException implements the getFloatException operation.
func (s *Server) GetFloatException(ctx context.Context, req apigen.GetFloatExceptionRequestObject) (apigen.GetFloatExceptionResponseObject, error) {
	row, err := s.service.GetFloatException(ctx, req.FloatExceptionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatException](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetFloatException200JSONResponse{Body: body, Headers: apigen.GetFloatException200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateFloatException implements the createFloatException operation.
func (s *Server) CreateFloatException(ctx context.Context, req apigen.CreateFloatExceptionRequestObject) (apigen.CreateFloatExceptionResponseObject, error) {
	in, err := input[db.CreateFloatExceptionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateFloatException(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatException](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateFloatException201JSONResponse{Body: body, Headers: apigen.CreateFloatException201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitFloatException implements the submitFloatException operation.
func (s *Server) SubmitFloatException(ctx context.Context, req apigen.SubmitFloatExceptionRequestObject) (apigen.SubmitFloatExceptionResponseObject, error) {
	in, err := input[db.ActSubmitFloatExceptionParams](struct{}{}, map[string]any{"id": req.FloatExceptionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitFloatException(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatException](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitFloatException200JSONResponse{Body: body, Headers: apigen.SubmitFloatException200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReimburseFloatException implements the reimburseFloatException operation.
func (s *Server) ReimburseFloatException(ctx context.Context, req apigen.ReimburseFloatExceptionRequestObject) (apigen.ReimburseFloatExceptionResponseObject, error) {
	in, err := input[db.ActReimburseFloatExceptionParams](req.Body, map[string]any{"id": req.FloatExceptionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReimburseFloatException(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FloatException](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReimburseFloatException200JSONResponse{Body: body, Headers: apigen.ReimburseFloatException200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListTillFlushes implements the listTillFlushes operation.
func (s *Server) ListTillFlushes(ctx context.Context, req apigen.ListTillFlushesRequestObject) (apigen.ListTillFlushesResponseObject, error) {
	return listPage[apigen.ListTillFlushes200JSONResponse](ctx, req.Params, nil, s.service.ListTillFlush)
}

// GetTillFlush implements the getTillFlush operation.
func (s *Server) GetTillFlush(ctx context.Context, req apigen.GetTillFlushRequestObject) (apigen.GetTillFlushResponseObject, error) {
	row, err := s.service.GetTillFlush(ctx, req.TillFlushId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TillFlush](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetTillFlush200JSONResponse{Body: body, Headers: apigen.GetTillFlush200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateTillFlush implements the createTillFlush operation.
func (s *Server) CreateTillFlush(ctx context.Context, req apigen.CreateTillFlushRequestObject) (apigen.CreateTillFlushResponseObject, error) {
	in, err := input[db.CreateTillFlushParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateTillFlush(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TillFlush](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateTillFlush201JSONResponse{Body: body, Headers: apigen.CreateTillFlush201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateTillFlush implements the updateTillFlush operation.
func (s *Server) UpdateTillFlush(ctx context.Context, req apigen.UpdateTillFlushRequestObject) (apigen.UpdateTillFlushResponseObject, error) {
	in, err := input[db.UpdateTillFlushParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.TillFlushId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateTillFlush(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TillFlush](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateTillFlush200JSONResponse{Body: body, Headers: apigen.UpdateTillFlush200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitTillFlush implements the submitTillFlush operation.
func (s *Server) SubmitTillFlush(ctx context.Context, req apigen.SubmitTillFlushRequestObject) (apigen.SubmitTillFlushResponseObject, error) {
	in, err := input[db.ActSubmitTillFlushParams](struct{}{}, map[string]any{"id": req.TillFlushId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitTillFlush(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TillFlush](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitTillFlush200JSONResponse{Body: body, Headers: apigen.SubmitTillFlush200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCashUps implements the listCashUps operation.
func (s *Server) ListCashUps(ctx context.Context, req apigen.ListCashUpsRequestObject) (apigen.ListCashUpsResponseObject, error) {
	return listPage[apigen.ListCashUps200JSONResponse](ctx, req.Params, nil, s.service.ListCashUp)
}

// GetCashUp implements the getCashUp operation.
func (s *Server) GetCashUp(ctx context.Context, req apigen.GetCashUpRequestObject) (apigen.GetCashUpResponseObject, error) {
	row, err := s.service.GetCashUp(ctx, req.CashUpId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.CashUpChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashUp](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetCashUp200JSONResponse{Body: body, Headers: apigen.GetCashUp200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCashUp implements the createCashUp operation.
func (s *Server) CreateCashUp(ctx context.Context, req apigen.CreateCashUpRequestObject) (apigen.CreateCashUpResponseObject, error) {
	in, err := input[db.CreateCashUpParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCashUp(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["tenders"]; ok {
		rows, err := childInputs[db.InsertCashUpTendersParams](items, "cash_up_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceCashUpTenders(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.CashUpChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashUp](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCashUp201JSONResponse{Body: body, Headers: apigen.CreateCashUp201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCashUp implements the updateCashUp operation.
func (s *Server) UpdateCashUp(ctx context.Context, req apigen.UpdateCashUpRequestObject) (apigen.UpdateCashUpResponseObject, error) {
	in, err := input[db.UpdateCashUpParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CashUpId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCashUp(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["tenders"]; ok {
		rows, err := childInputs[db.InsertCashUpTendersParams](items, "cash_up_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceCashUpTenders(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.CashUpChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashUp](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCashUp200JSONResponse{Body: body, Headers: apigen.UpdateCashUp200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitCashUp implements the submitCashUp operation.
func (s *Server) SubmitCashUp(ctx context.Context, req apigen.SubmitCashUpRequestObject) (apigen.SubmitCashUpResponseObject, error) {
	in, err := input[db.ActSubmitCashUpParams](struct{}{}, map[string]any{"id": req.CashUpId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitCashUp(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashUp](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitCashUp200JSONResponse{Body: body, Headers: apigen.SubmitCashUp200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReturnCashUp implements the returnCashUp operation.
func (s *Server) ReturnCashUp(ctx context.Context, req apigen.ReturnCashUpRequestObject) (apigen.ReturnCashUpResponseObject, error) {
	in, err := input[db.ActReturnCashUpParams](req.Body, map[string]any{"id": req.CashUpId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReturnCashUp(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashUp](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReturnCashUp200JSONResponse{Body: body, Headers: apigen.ReturnCashUp200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCashVariances implements the listCashVariances operation.
func (s *Server) ListCashVariances(ctx context.Context, req apigen.ListCashVariancesRequestObject) (apigen.ListCashVariancesResponseObject, error) {
	return listPage[apigen.ListCashVariances200JSONResponse](ctx, req.Params, nil, s.service.ListCashVariance)
}

// GetCashVariance implements the getCashVariance operation.
func (s *Server) GetCashVariance(ctx context.Context, req apigen.GetCashVarianceRequestObject) (apigen.GetCashVarianceResponseObject, error) {
	row, err := s.service.GetCashVariance(ctx, req.CashVarianceId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCashVariance200JSONResponse{Body: body, Headers: apigen.GetCashVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCashVariance implements the updateCashVariance operation.
func (s *Server) UpdateCashVariance(ctx context.Context, req apigen.UpdateCashVarianceRequestObject) (apigen.UpdateCashVarianceResponseObject, error) {
	in, err := input[db.UpdateCashVarianceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CashVarianceId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCashVariance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCashVariance200JSONResponse{Body: body, Headers: apigen.UpdateCashVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RecordRecoveryCashVariance implements the recordRecoveryCashVariance operation.
func (s *Server) RecordRecoveryCashVariance(ctx context.Context, req apigen.RecordRecoveryCashVarianceRequestObject) (apigen.RecordRecoveryCashVarianceResponseObject, error) {
	in, err := input[db.ActRecordRecoveryCashVarianceParams](req.Body, map[string]any{"id": req.CashVarianceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RecordRecoveryCashVariance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.RecordRecoveryCashVariance200JSONResponse{Body: body, Headers: apigen.RecordRecoveryCashVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReferHrCashVariance implements the referHrCashVariance operation.
func (s *Server) ReferHrCashVariance(ctx context.Context, req apigen.ReferHrCashVarianceRequestObject) (apigen.ReferHrCashVarianceResponseObject, error) {
	in, err := input[db.ActReferHrCashVarianceParams](req.Body, map[string]any{"id": req.CashVarianceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReferHrCashVariance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReferHrCashVariance200JSONResponse{Body: body, Headers: apigen.ReferHrCashVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseCashVariance implements the closeCashVariance operation.
func (s *Server) CloseCashVariance(ctx context.Context, req apigen.CloseCashVarianceRequestObject) (apigen.CloseCashVarianceResponseObject, error) {
	in, err := input[db.ActCloseCashVarianceParams](req.Body, map[string]any{"id": req.CashVarianceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseCashVariance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CashVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseCashVariance200JSONResponse{Body: body, Headers: apigen.CloseCashVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCitProviders implements the listCitProviders operation.
func (s *Server) ListCitProviders(ctx context.Context, req apigen.ListCitProvidersRequestObject) (apigen.ListCitProvidersResponseObject, error) {
	return listPage[apigen.ListCitProviders200JSONResponse](ctx, req.Params, nil, s.service.ListCitProvider)
}

// GetCitProvider implements the getCitProvider operation.
func (s *Server) GetCitProvider(ctx context.Context, req apigen.GetCitProviderRequestObject) (apigen.GetCitProviderResponseObject, error) {
	row, err := s.service.GetCitProvider(ctx, req.CitProviderId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitProvider](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCitProvider200JSONResponse{Body: body, Headers: apigen.GetCitProvider200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCitProvider implements the createCitProvider operation.
func (s *Server) CreateCitProvider(ctx context.Context, req apigen.CreateCitProviderRequestObject) (apigen.CreateCitProviderResponseObject, error) {
	in, err := input[db.CreateCitProviderParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCitProvider(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitProvider](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCitProvider201JSONResponse{Body: body, Headers: apigen.CreateCitProvider201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCitProvider implements the updateCitProvider operation.
func (s *Server) UpdateCitProvider(ctx context.Context, req apigen.UpdateCitProviderRequestObject) (apigen.UpdateCitProviderResponseObject, error) {
	in, err := input[db.UpdateCitProviderParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CitProviderId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCitProvider(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitProvider](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCitProvider200JSONResponse{Body: body, Headers: apigen.UpdateCitProvider200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCitPersons implements the listCitPersons operation.
func (s *Server) ListCitPersons(ctx context.Context, req apigen.ListCitPersonsRequestObject) (apigen.ListCitPersonsResponseObject, error) {
	return listPage[apigen.ListCitPersons200JSONResponse](ctx, req.Params, map[string]any{"cit_provider_id": req.CitProviderId}, s.service.ListCitPerson)
}

// CreateCitPerson implements the createCitPerson operation.
func (s *Server) CreateCitPerson(ctx context.Context, req apigen.CreateCitPersonRequestObject) (apigen.CreateCitPersonResponseObject, error) {
	in, err := input[db.CreateCitPersonParams](req.Body, map[string]any{"cit_provider_id": req.CitProviderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCitPerson(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitPerson](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCitPerson201JSONResponse{Body: body, Headers: apigen.CreateCitPerson201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCitPerson implements the updateCitPerson operation.
func (s *Server) UpdateCitPerson(ctx context.Context, req apigen.UpdateCitPersonRequestObject) (apigen.UpdateCitPersonResponseObject, error) {
	in, err := input[db.UpdateCitPersonParams](req.Body, map[string]any{"cit_provider_id": req.CitProviderId})
	if err != nil {
		return nil, err
	}
	in.ID = req.CitPersonId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCitPerson(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitPerson](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCitPerson200JSONResponse{Body: body, Headers: apigen.UpdateCitPerson200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCitTransits implements the listCitTransits operation.
func (s *Server) ListCitTransits(ctx context.Context, req apigen.ListCitTransitsRequestObject) (apigen.ListCitTransitsResponseObject, error) {
	return listPage[apigen.ListCitTransits200JSONResponse](ctx, req.Params, nil, s.service.ListCitTransit)
}

// GetCitTransit implements the getCitTransit operation.
func (s *Server) GetCitTransit(ctx context.Context, req apigen.GetCitTransitRequestObject) (apigen.GetCitTransitResponseObject, error) {
	row, err := s.service.GetCitTransit(ctx, req.CitTransitId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.CitTransitChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitTransit](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetCitTransit200JSONResponse{Body: body, Headers: apigen.GetCitTransit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCitTransit implements the createCitTransit operation.
func (s *Server) CreateCitTransit(ctx context.Context, req apigen.CreateCitTransitRequestObject) (apigen.CreateCitTransitResponseObject, error) {
	in, err := input[db.CreateCitTransitParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCitTransit(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["bags"]; ok {
		rows, err := childInputs[db.InsertCitBagsParams](items, "cit_transit_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceCitBags(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.CitTransitChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitTransit](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCitTransit201JSONResponse{Body: body, Headers: apigen.CreateCitTransit201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCitTransit implements the updateCitTransit operation.
func (s *Server) UpdateCitTransit(ctx context.Context, req apigen.UpdateCitTransitRequestObject) (apigen.UpdateCitTransitResponseObject, error) {
	in, err := input[db.UpdateCitTransitParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CitTransitId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCitTransit(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["bags"]; ok {
		rows, err := childInputs[db.InsertCitBagsParams](items, "cit_transit_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceCitBags(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.CitTransitChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitTransit](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCitTransit200JSONResponse{Body: body, Headers: apigen.UpdateCitTransit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SealCitTransit implements the sealCitTransit operation.
func (s *Server) SealCitTransit(ctx context.Context, req apigen.SealCitTransitRequestObject) (apigen.SealCitTransitResponseObject, error) {
	in, err := input[db.ActSealCitTransitParams](struct{}{}, map[string]any{"id": req.CitTransitId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SealCitTransit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitTransit](row)
	if err != nil {
		return nil, err
	}
	return apigen.SealCitTransit200JSONResponse{Body: body, Headers: apigen.SealCitTransit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RaiseDiscrepancyCitTransit implements the raiseDiscrepancyCitTransit operation.
func (s *Server) RaiseDiscrepancyCitTransit(ctx context.Context, req apigen.RaiseDiscrepancyCitTransitRequestObject) (apigen.RaiseDiscrepancyCitTransitResponseObject, error) {
	in, err := input[db.ActRaiseDiscrepancyCitTransitParams](req.Body, map[string]any{"id": req.CitTransitId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RaiseDiscrepancyCitTransit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitTransit](row)
	if err != nil {
		return nil, err
	}
	return apigen.RaiseDiscrepancyCitTransit200JSONResponse{Body: body, Headers: apigen.RaiseDiscrepancyCitTransit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListBankDeposits implements the listBankDeposits operation.
func (s *Server) ListBankDeposits(ctx context.Context, req apigen.ListBankDepositsRequestObject) (apigen.ListBankDepositsResponseObject, error) {
	return listPage[apigen.ListBankDeposits200JSONResponse](ctx, req.Params, nil, s.service.ListBankDeposit)
}

// GetBankDeposit implements the getBankDeposit operation.
func (s *Server) GetBankDeposit(ctx context.Context, req apigen.GetBankDepositRequestObject) (apigen.GetBankDepositResponseObject, error) {
	row, err := s.service.GetBankDeposit(ctx, req.BankDepositId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BankDeposit](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetBankDeposit200JSONResponse{Body: body, Headers: apigen.GetBankDeposit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateBankDeposit implements the createBankDeposit operation.
func (s *Server) CreateBankDeposit(ctx context.Context, req apigen.CreateBankDepositRequestObject) (apigen.CreateBankDepositResponseObject, error) {
	in, err := input[db.CreateBankDepositParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateBankDeposit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BankDeposit](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateBankDeposit201JSONResponse{Body: body, Headers: apigen.CreateBankDeposit201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateBankDeposit implements the updateBankDeposit operation.
func (s *Server) UpdateBankDeposit(ctx context.Context, req apigen.UpdateBankDepositRequestObject) (apigen.UpdateBankDepositResponseObject, error) {
	in, err := input[db.UpdateBankDepositParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.BankDepositId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateBankDeposit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BankDeposit](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateBankDeposit200JSONResponse{Body: body, Headers: apigen.UpdateBankDeposit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitProofBankDeposit implements the submitProofBankDeposit operation.
func (s *Server) SubmitProofBankDeposit(ctx context.Context, req apigen.SubmitProofBankDepositRequestObject) (apigen.SubmitProofBankDepositResponseObject, error) {
	in, err := input[db.ActSubmitProofBankDepositParams](req.Body, map[string]any{"id": req.BankDepositId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitProofBankDeposit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BankDeposit](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitProofBankDeposit200JSONResponse{Body: body, Headers: apigen.SubmitProofBankDeposit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// VerifyBankDeposit implements the verifyBankDeposit operation.
func (s *Server) VerifyBankDeposit(ctx context.Context, req apigen.VerifyBankDepositRequestObject) (apigen.VerifyBankDepositResponseObject, error) {
	in, err := input[db.ActVerifyBankDepositParams](struct{}{}, map[string]any{"id": req.BankDepositId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.VerifyBankDeposit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BankDeposit](row)
	if err != nil {
		return nil, err
	}
	return apigen.VerifyBankDeposit200JSONResponse{Body: body, Headers: apigen.VerifyBankDeposit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCitReconciliations implements the listCitReconciliations operation.
func (s *Server) ListCitReconciliations(ctx context.Context, req apigen.ListCitReconciliationsRequestObject) (apigen.ListCitReconciliationsResponseObject, error) {
	return listPage[apigen.ListCitReconciliations200JSONResponse](ctx, req.Params, nil, s.service.ListCitReconciliation)
}

// GetCitReconciliation implements the getCitReconciliation operation.
func (s *Server) GetCitReconciliation(ctx context.Context, req apigen.GetCitReconciliationRequestObject) (apigen.GetCitReconciliationResponseObject, error) {
	row, err := s.service.GetCitReconciliation(ctx, req.CitReconciliationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCitReconciliation200JSONResponse{Body: body, Headers: apigen.GetCitReconciliation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCitReconciliation implements the createCitReconciliation operation.
func (s *Server) CreateCitReconciliation(ctx context.Context, req apigen.CreateCitReconciliationRequestObject) (apigen.CreateCitReconciliationResponseObject, error) {
	in, err := input[db.CreateCitReconciliationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCitReconciliation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCitReconciliation201JSONResponse{Body: body, Headers: apigen.CreateCitReconciliation201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCitReconciliation implements the updateCitReconciliation operation.
func (s *Server) UpdateCitReconciliation(ctx context.Context, req apigen.UpdateCitReconciliationRequestObject) (apigen.UpdateCitReconciliationResponseObject, error) {
	in, err := input[db.UpdateCitReconciliationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CitReconciliationId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCitReconciliation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCitReconciliation200JSONResponse{Body: body, Headers: apigen.UpdateCitReconciliation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteCitReconciliation implements the completeCitReconciliation operation.
func (s *Server) CompleteCitReconciliation(ctx context.Context, req apigen.CompleteCitReconciliationRequestObject) (apigen.CompleteCitReconciliationResponseObject, error) {
	in, err := input[db.ActCompleteCitReconciliationParams](struct{}{}, map[string]any{"id": req.CitReconciliationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteCitReconciliation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CitReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteCitReconciliation200JSONResponse{Body: body, Headers: apigen.CompleteCitReconciliation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListSurpriseCounts implements the listSurpriseCounts operation.
func (s *Server) ListSurpriseCounts(ctx context.Context, req apigen.ListSurpriseCountsRequestObject) (apigen.ListSurpriseCountsResponseObject, error) {
	return listPage[apigen.ListSurpriseCounts200JSONResponse](ctx, req.Params, nil, s.service.ListSurpriseCount)
}

// GetSurpriseCount implements the getSurpriseCount operation.
func (s *Server) GetSurpriseCount(ctx context.Context, req apigen.GetSurpriseCountRequestObject) (apigen.GetSurpriseCountResponseObject, error) {
	row, err := s.service.GetSurpriseCount(ctx, req.SurpriseCountId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SurpriseCount](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSurpriseCount200JSONResponse{Body: body, Headers: apigen.GetSurpriseCount200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateSurpriseCount implements the createSurpriseCount operation.
func (s *Server) CreateSurpriseCount(ctx context.Context, req apigen.CreateSurpriseCountRequestObject) (apigen.CreateSurpriseCountResponseObject, error) {
	in, err := input[db.CreateSurpriseCountParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSurpriseCount(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SurpriseCount](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSurpriseCount201JSONResponse{Body: body, Headers: apigen.CreateSurpriseCount201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateSurpriseCount implements the updateSurpriseCount operation.
func (s *Server) UpdateSurpriseCount(ctx context.Context, req apigen.UpdateSurpriseCountRequestObject) (apigen.UpdateSurpriseCountResponseObject, error) {
	in, err := input[db.UpdateSurpriseCountParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.SurpriseCountId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateSurpriseCount(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SurpriseCount](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateSurpriseCount200JSONResponse{Body: body, Headers: apigen.UpdateSurpriseCount200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteSurpriseCount implements the completeSurpriseCount operation.
func (s *Server) CompleteSurpriseCount(ctx context.Context, req apigen.CompleteSurpriseCountRequestObject) (apigen.CompleteSurpriseCountResponseObject, error) {
	in, err := input[db.ActCompleteSurpriseCountParams](struct{}{}, map[string]any{"id": req.SurpriseCountId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteSurpriseCount(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SurpriseCount](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteSurpriseCount200JSONResponse{Body: body, Headers: apigen.CompleteSurpriseCount200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPettyCashAccounts implements the listPettyCashAccounts operation.
func (s *Server) ListPettyCashAccounts(ctx context.Context, req apigen.ListPettyCashAccountsRequestObject) (apigen.ListPettyCashAccountsResponseObject, error) {
	return listPage[apigen.ListPettyCashAccounts200JSONResponse](ctx, req.Params, nil, s.service.ListPettyCashAccount)
}

// GetPettyCashAccount implements the getPettyCashAccount operation.
func (s *Server) GetPettyCashAccount(ctx context.Context, req apigen.GetPettyCashAccountRequestObject) (apigen.GetPettyCashAccountResponseObject, error) {
	row, err := s.service.GetPettyCashAccount(ctx, req.PettyCashAccountId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashAccount](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPettyCashAccount200JSONResponse{Body: body, Headers: apigen.GetPettyCashAccount200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePettyCashAccount implements the createPettyCashAccount operation.
func (s *Server) CreatePettyCashAccount(ctx context.Context, req apigen.CreatePettyCashAccountRequestObject) (apigen.CreatePettyCashAccountResponseObject, error) {
	in, err := input[db.CreatePettyCashAccountParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePettyCashAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashAccount](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePettyCashAccount201JSONResponse{Body: body, Headers: apigen.CreatePettyCashAccount201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePettyCashAccount implements the updatePettyCashAccount operation.
func (s *Server) UpdatePettyCashAccount(ctx context.Context, req apigen.UpdatePettyCashAccountRequestObject) (apigen.UpdatePettyCashAccountResponseObject, error) {
	in, err := input[db.UpdatePettyCashAccountParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PettyCashAccountId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePettyCashAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashAccount](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePettyCashAccount200JSONResponse{Body: body, Headers: apigen.UpdatePettyCashAccount200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPettyCashVouchers implements the listPettyCashVouchers operation.
func (s *Server) ListPettyCashVouchers(ctx context.Context, req apigen.ListPettyCashVouchersRequestObject) (apigen.ListPettyCashVouchersResponseObject, error) {
	return listPage[apigen.ListPettyCashVouchers200JSONResponse](ctx, req.Params, nil, s.service.ListPettyCashVoucher)
}

// GetPettyCashVoucher implements the getPettyCashVoucher operation.
func (s *Server) GetPettyCashVoucher(ctx context.Context, req apigen.GetPettyCashVoucherRequestObject) (apigen.GetPettyCashVoucherResponseObject, error) {
	row, err := s.service.GetPettyCashVoucher(ctx, req.PettyCashVoucherId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashVoucher](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPettyCashVoucher200JSONResponse{Body: body, Headers: apigen.GetPettyCashVoucher200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePettyCashVoucher implements the createPettyCashVoucher operation.
func (s *Server) CreatePettyCashVoucher(ctx context.Context, req apigen.CreatePettyCashVoucherRequestObject) (apigen.CreatePettyCashVoucherResponseObject, error) {
	in, err := input[db.CreatePettyCashVoucherParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePettyCashVoucher(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashVoucher](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePettyCashVoucher201JSONResponse{Body: body, Headers: apigen.CreatePettyCashVoucher201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePettyCashVoucher implements the updatePettyCashVoucher operation.
func (s *Server) UpdatePettyCashVoucher(ctx context.Context, req apigen.UpdatePettyCashVoucherRequestObject) (apigen.UpdatePettyCashVoucherResponseObject, error) {
	in, err := input[db.UpdatePettyCashVoucherParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PettyCashVoucherId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePettyCashVoucher(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashVoucher](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePettyCashVoucher200JSONResponse{Body: body, Headers: apigen.UpdatePettyCashVoucher200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitPettyCashVoucher implements the submitPettyCashVoucher operation.
func (s *Server) SubmitPettyCashVoucher(ctx context.Context, req apigen.SubmitPettyCashVoucherRequestObject) (apigen.SubmitPettyCashVoucherResponseObject, error) {
	in, err := input[db.ActSubmitPettyCashVoucherParams](struct{}{}, map[string]any{"id": req.PettyCashVoucherId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitPettyCashVoucher(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashVoucher](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitPettyCashVoucher200JSONResponse{Body: body, Headers: apigen.SubmitPettyCashVoucher200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// VerifyPettyCashVoucher implements the verifyPettyCashVoucher operation.
func (s *Server) VerifyPettyCashVoucher(ctx context.Context, req apigen.VerifyPettyCashVoucherRequestObject) (apigen.VerifyPettyCashVoucherResponseObject, error) {
	in, err := input[db.ActVerifyPettyCashVoucherParams](struct{}{}, map[string]any{"id": req.PettyCashVoucherId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.VerifyPettyCashVoucher(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashVoucher](row)
	if err != nil {
		return nil, err
	}
	return apigen.VerifyPettyCashVoucher200JSONResponse{Body: body, Headers: apigen.VerifyPettyCashVoucher200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AttachReceiptPettyCashVoucher implements the attachReceiptPettyCashVoucher operation.
func (s *Server) AttachReceiptPettyCashVoucher(ctx context.Context, req apigen.AttachReceiptPettyCashVoucherRequestObject) (apigen.AttachReceiptPettyCashVoucherResponseObject, error) {
	in, err := input[db.ActAttachReceiptPettyCashVoucherParams](req.Body, map[string]any{"id": req.PettyCashVoucherId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AttachReceiptPettyCashVoucher(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashVoucher](row)
	if err != nil {
		return nil, err
	}
	return apigen.AttachReceiptPettyCashVoucher200JSONResponse{Body: body, Headers: apigen.AttachReceiptPettyCashVoucher200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPettyCashHandovers implements the listPettyCashHandovers operation.
func (s *Server) ListPettyCashHandovers(ctx context.Context, req apigen.ListPettyCashHandoversRequestObject) (apigen.ListPettyCashHandoversResponseObject, error) {
	return listPage[apigen.ListPettyCashHandovers200JSONResponse](ctx, req.Params, nil, s.service.ListPettyCashHandover)
}

// GetPettyCashHandover implements the getPettyCashHandover operation.
func (s *Server) GetPettyCashHandover(ctx context.Context, req apigen.GetPettyCashHandoverRequestObject) (apigen.GetPettyCashHandoverResponseObject, error) {
	row, err := s.service.GetPettyCashHandover(ctx, req.PettyCashHandoverId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashHandover](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPettyCashHandover200JSONResponse{Body: body, Headers: apigen.GetPettyCashHandover200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePettyCashHandover implements the createPettyCashHandover operation.
func (s *Server) CreatePettyCashHandover(ctx context.Context, req apigen.CreatePettyCashHandoverRequestObject) (apigen.CreatePettyCashHandoverResponseObject, error) {
	in, err := input[db.CreatePettyCashHandoverParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePettyCashHandover(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashHandover](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePettyCashHandover201JSONResponse{Body: body, Headers: apigen.CreatePettyCashHandover201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitPettyCashHandover implements the submitPettyCashHandover operation.
func (s *Server) SubmitPettyCashHandover(ctx context.Context, req apigen.SubmitPettyCashHandoverRequestObject) (apigen.SubmitPettyCashHandoverResponseObject, error) {
	in, err := input[db.ActSubmitPettyCashHandoverParams](struct{}{}, map[string]any{"id": req.PettyCashHandoverId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitPettyCashHandover(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashHandover](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitPettyCashHandover200JSONResponse{Body: body, Headers: apigen.SubmitPettyCashHandover200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPettyCashReconciliations implements the listPettyCashReconciliations operation.
func (s *Server) ListPettyCashReconciliations(ctx context.Context, req apigen.ListPettyCashReconciliationsRequestObject) (apigen.ListPettyCashReconciliationsResponseObject, error) {
	return listPage[apigen.ListPettyCashReconciliations200JSONResponse](ctx, req.Params, nil, s.service.ListPettyCashReconciliation)
}

// GetPettyCashReconciliation implements the getPettyCashReconciliation operation.
func (s *Server) GetPettyCashReconciliation(ctx context.Context, req apigen.GetPettyCashReconciliationRequestObject) (apigen.GetPettyCashReconciliationResponseObject, error) {
	row, err := s.service.GetPettyCashReconciliation(ctx, req.PettyCashReconciliationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPettyCashReconciliation200JSONResponse{Body: body, Headers: apigen.GetPettyCashReconciliation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePettyCashReconciliation implements the createPettyCashReconciliation operation.
func (s *Server) CreatePettyCashReconciliation(ctx context.Context, req apigen.CreatePettyCashReconciliationRequestObject) (apigen.CreatePettyCashReconciliationResponseObject, error) {
	in, err := input[db.CreatePettyCashReconciliationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePettyCashReconciliation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePettyCashReconciliation201JSONResponse{Body: body, Headers: apigen.CreatePettyCashReconciliation201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitPettyCashReconciliation implements the submitPettyCashReconciliation operation.
func (s *Server) SubmitPettyCashReconciliation(ctx context.Context, req apigen.SubmitPettyCashReconciliationRequestObject) (apigen.SubmitPettyCashReconciliationResponseObject, error) {
	in, err := input[db.ActSubmitPettyCashReconciliationParams](struct{}{}, map[string]any{"id": req.PettyCashReconciliationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitPettyCashReconciliation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitPettyCashReconciliation200JSONResponse{Body: body, Headers: apigen.SubmitPettyCashReconciliation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReviewPettyCashReconciliation implements the reviewPettyCashReconciliation operation.
func (s *Server) ReviewPettyCashReconciliation(ctx context.Context, req apigen.ReviewPettyCashReconciliationRequestObject) (apigen.ReviewPettyCashReconciliationResponseObject, error) {
	in, err := input[db.ActReviewPettyCashReconciliationParams](req.Body, map[string]any{"id": req.PettyCashReconciliationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReviewPettyCashReconciliation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashReconciliation](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReviewPettyCashReconciliation200JSONResponse{Body: body, Headers: apigen.ReviewPettyCashReconciliation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPettyCashReimbursements implements the listPettyCashReimbursements operation.
func (s *Server) ListPettyCashReimbursements(ctx context.Context, req apigen.ListPettyCashReimbursementsRequestObject) (apigen.ListPettyCashReimbursementsResponseObject, error) {
	return listPage[apigen.ListPettyCashReimbursements200JSONResponse](ctx, req.Params, nil, s.service.ListPettyCashReimbursement)
}

// GetPettyCashReimbursement implements the getPettyCashReimbursement operation.
func (s *Server) GetPettyCashReimbursement(ctx context.Context, req apigen.GetPettyCashReimbursementRequestObject) (apigen.GetPettyCashReimbursementResponseObject, error) {
	row, err := s.service.GetPettyCashReimbursement(ctx, req.PettyCashReimbursementId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashReimbursement](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPettyCashReimbursement200JSONResponse{Body: body, Headers: apigen.GetPettyCashReimbursement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePettyCashReimbursement implements the createPettyCashReimbursement operation.
func (s *Server) CreatePettyCashReimbursement(ctx context.Context, req apigen.CreatePettyCashReimbursementRequestObject) (apigen.CreatePettyCashReimbursementResponseObject, error) {
	in, err := input[db.CreatePettyCashReimbursementParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePettyCashReimbursement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashReimbursement](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePettyCashReimbursement201JSONResponse{Body: body, Headers: apigen.CreatePettyCashReimbursement201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitPettyCashReimbursement implements the submitPettyCashReimbursement operation.
func (s *Server) SubmitPettyCashReimbursement(ctx context.Context, req apigen.SubmitPettyCashReimbursementRequestObject) (apigen.SubmitPettyCashReimbursementResponseObject, error) {
	in, err := input[db.ActSubmitPettyCashReimbursementParams](struct{}{}, map[string]any{"id": req.PettyCashReimbursementId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitPettyCashReimbursement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashReimbursement](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitPettyCashReimbursement200JSONResponse{Body: body, Headers: apigen.SubmitPettyCashReimbursement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RecordPaymentPettyCashReimbursement implements the recordPaymentPettyCashReimbursement operation.
func (s *Server) RecordPaymentPettyCashReimbursement(ctx context.Context, req apigen.RecordPaymentPettyCashReimbursementRequestObject) (apigen.RecordPaymentPettyCashReimbursementResponseObject, error) {
	in, err := input[db.ActRecordPaymentPettyCashReimbursementParams](req.Body, map[string]any{"id": req.PettyCashReimbursementId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RecordPaymentPettyCashReimbursement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashReimbursement](row)
	if err != nil {
		return nil, err
	}
	return apigen.RecordPaymentPettyCashReimbursement200JSONResponse{Body: body, Headers: apigen.RecordPaymentPettyCashReimbursement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
