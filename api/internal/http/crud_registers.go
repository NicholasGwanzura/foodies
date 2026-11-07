// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListAssetClasses implements the listAssetClasses operation.
func (s *Server) ListAssetClasses(ctx context.Context, req apigen.ListAssetClassesRequestObject) (apigen.ListAssetClassesResponseObject, error) {
	return listPage[apigen.ListAssetClasses200JSONResponse](ctx, req.Params, nil, s.service.ListAssetClass)
}

// GetAssetClass implements the getAssetClass operation.
func (s *Server) GetAssetClass(ctx context.Context, req apigen.GetAssetClassRequestObject) (apigen.GetAssetClassResponseObject, error) {
	row, err := s.service.GetAssetClass(ctx, req.AssetClassId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetClass](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetAssetClass200JSONResponse{Body: body, Headers: apigen.GetAssetClass200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateAssetClass implements the createAssetClass operation.
func (s *Server) CreateAssetClass(ctx context.Context, req apigen.CreateAssetClassRequestObject) (apigen.CreateAssetClassResponseObject, error) {
	in, err := input[db.CreateAssetClassParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateAssetClass(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetClass](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateAssetClass201JSONResponse{Body: body, Headers: apigen.CreateAssetClass201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateAssetClass implements the updateAssetClass operation.
func (s *Server) UpdateAssetClass(ctx context.Context, req apigen.UpdateAssetClassRequestObject) (apigen.UpdateAssetClassResponseObject, error) {
	in, err := input[db.UpdateAssetClassParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.AssetClassId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateAssetClass(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetClass](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateAssetClass200JSONResponse{Body: body, Headers: apigen.UpdateAssetClass200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListAssets implements the listAssets operation.
func (s *Server) ListAssets(ctx context.Context, req apigen.ListAssetsRequestObject) (apigen.ListAssetsResponseObject, error) {
	return listPage[apigen.ListAssets200JSONResponse](ctx, req.Params, nil, s.service.ListAsset)
}

// GetAsset implements the getAsset operation.
func (s *Server) GetAsset(ctx context.Context, req apigen.GetAssetRequestObject) (apigen.GetAssetResponseObject, error) {
	row, err := s.service.GetAsset(ctx, req.AssetId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Asset](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetAsset200JSONResponse{Body: body, Headers: apigen.GetAsset200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateAsset implements the createAsset operation.
func (s *Server) CreateAsset(ctx context.Context, req apigen.CreateAssetRequestObject) (apigen.CreateAssetResponseObject, error) {
	in, err := input[db.CreateAssetParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateAsset(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Asset](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateAsset201JSONResponse{Body: body, Headers: apigen.CreateAsset201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateAsset implements the updateAsset operation.
func (s *Server) UpdateAsset(ctx context.Context, req apigen.UpdateAssetRequestObject) (apigen.UpdateAssetResponseObject, error) {
	in, err := input[db.UpdateAssetParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.AssetId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateAsset(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Asset](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateAsset200JSONResponse{Body: body, Headers: apigen.UpdateAsset200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCapexRequests implements the listCapexRequests operation.
func (s *Server) ListCapexRequests(ctx context.Context, req apigen.ListCapexRequestsRequestObject) (apigen.ListCapexRequestsResponseObject, error) {
	return listPage[apigen.ListCapexRequests200JSONResponse](ctx, req.Params, nil, s.service.ListCapexRequest)
}

// GetCapexRequest implements the getCapexRequest operation.
func (s *Server) GetCapexRequest(ctx context.Context, req apigen.GetCapexRequestRequestObject) (apigen.GetCapexRequestResponseObject, error) {
	row, err := s.service.GetCapexRequest(ctx, req.CapexRequestId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CapexRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCapexRequest200JSONResponse{Body: body, Headers: apigen.GetCapexRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCapexRequest implements the createCapexRequest operation.
func (s *Server) CreateCapexRequest(ctx context.Context, req apigen.CreateCapexRequestRequestObject) (apigen.CreateCapexRequestResponseObject, error) {
	in, err := input[db.CreateCapexRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCapexRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CapexRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCapexRequest201JSONResponse{Body: body, Headers: apigen.CreateCapexRequest201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCapexRequest implements the updateCapexRequest operation.
func (s *Server) UpdateCapexRequest(ctx context.Context, req apigen.UpdateCapexRequestRequestObject) (apigen.UpdateCapexRequestResponseObject, error) {
	in, err := input[db.UpdateCapexRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CapexRequestId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCapexRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CapexRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCapexRequest200JSONResponse{Body: body, Headers: apigen.UpdateCapexRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitCapexRequest implements the submitCapexRequest operation.
func (s *Server) SubmitCapexRequest(ctx context.Context, req apigen.SubmitCapexRequestRequestObject) (apigen.SubmitCapexRequestResponseObject, error) {
	in, err := input[db.ActSubmitCapexRequestParams](struct{}{}, map[string]any{"id": req.CapexRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitCapexRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CapexRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitCapexRequest200JSONResponse{Body: body, Headers: apigen.SubmitCapexRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListAssetTransfers implements the listAssetTransfers operation.
func (s *Server) ListAssetTransfers(ctx context.Context, req apigen.ListAssetTransfersRequestObject) (apigen.ListAssetTransfersResponseObject, error) {
	return listPage[apigen.ListAssetTransfers200JSONResponse](ctx, req.Params, nil, s.service.ListAssetTransfer)
}

// GetAssetTransfer implements the getAssetTransfer operation.
func (s *Server) GetAssetTransfer(ctx context.Context, req apigen.GetAssetTransferRequestObject) (apigen.GetAssetTransferResponseObject, error) {
	row, err := s.service.GetAssetTransfer(ctx, req.AssetTransferId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetAssetTransfer200JSONResponse{Body: body, Headers: apigen.GetAssetTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateAssetTransfer implements the createAssetTransfer operation.
func (s *Server) CreateAssetTransfer(ctx context.Context, req apigen.CreateAssetTransferRequestObject) (apigen.CreateAssetTransferResponseObject, error) {
	in, err := input[db.CreateAssetTransferParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateAssetTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateAssetTransfer201JSONResponse{Body: body, Headers: apigen.CreateAssetTransfer201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitAssetTransfer implements the submitAssetTransfer operation.
func (s *Server) SubmitAssetTransfer(ctx context.Context, req apigen.SubmitAssetTransferRequestObject) (apigen.SubmitAssetTransferResponseObject, error) {
	in, err := input[db.ActSubmitAssetTransferParams](struct{}{}, map[string]any{"id": req.AssetTransferId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitAssetTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitAssetTransfer200JSONResponse{Body: body, Headers: apigen.SubmitAssetTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReceiveAssetTransfer implements the receiveAssetTransfer operation.
func (s *Server) ReceiveAssetTransfer(ctx context.Context, req apigen.ReceiveAssetTransferRequestObject) (apigen.ReceiveAssetTransferResponseObject, error) {
	in, err := input[db.ActReceiveAssetTransferParams](struct{}{}, map[string]any{"id": req.AssetTransferId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReceiveAssetTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReceiveAssetTransfer200JSONResponse{Body: body, Headers: apigen.ReceiveAssetTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListAssetVerifications implements the listAssetVerifications operation.
func (s *Server) ListAssetVerifications(ctx context.Context, req apigen.ListAssetVerificationsRequestObject) (apigen.ListAssetVerificationsResponseObject, error) {
	return listPage[apigen.ListAssetVerifications200JSONResponse](ctx, req.Params, nil, s.service.ListAssetVerification)
}

// GetAssetVerification implements the getAssetVerification operation.
func (s *Server) GetAssetVerification(ctx context.Context, req apigen.GetAssetVerificationRequestObject) (apigen.GetAssetVerificationResponseObject, error) {
	row, err := s.service.GetAssetVerification(ctx, req.AssetVerificationId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.AssetVerificationChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetVerification](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetAssetVerification200JSONResponse{Body: body, Headers: apigen.GetAssetVerification200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateAssetVerification implements the createAssetVerification operation.
func (s *Server) CreateAssetVerification(ctx context.Context, req apigen.CreateAssetVerificationRequestObject) (apigen.CreateAssetVerificationResponseObject, error) {
	in, err := input[db.CreateAssetVerificationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateAssetVerification(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertAssetVerificationLinesParams](items, "asset_verification_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceAssetVerificationLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.AssetVerificationChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetVerification](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateAssetVerification201JSONResponse{Body: body, Headers: apigen.CreateAssetVerification201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateAssetVerification implements the updateAssetVerification operation.
func (s *Server) UpdateAssetVerification(ctx context.Context, req apigen.UpdateAssetVerificationRequestObject) (apigen.UpdateAssetVerificationResponseObject, error) {
	in, err := input[db.UpdateAssetVerificationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.AssetVerificationId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateAssetVerification(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertAssetVerificationLinesParams](items, "asset_verification_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceAssetVerificationLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.AssetVerificationChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetVerification](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateAssetVerification200JSONResponse{Body: body, Headers: apigen.UpdateAssetVerification200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitAssetVerification implements the submitAssetVerification operation.
func (s *Server) SubmitAssetVerification(ctx context.Context, req apigen.SubmitAssetVerificationRequestObject) (apigen.SubmitAssetVerificationResponseObject, error) {
	in, err := input[db.ActSubmitAssetVerificationParams](struct{}{}, map[string]any{"id": req.AssetVerificationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitAssetVerification(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetVerification](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitAssetVerification200JSONResponse{Body: body, Headers: apigen.SubmitAssetVerification200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReviewAssetVerification implements the reviewAssetVerification operation.
func (s *Server) ReviewAssetVerification(ctx context.Context, req apigen.ReviewAssetVerificationRequestObject) (apigen.ReviewAssetVerificationResponseObject, error) {
	in, err := input[db.ActReviewAssetVerificationParams](req.Body, map[string]any{"id": req.AssetVerificationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReviewAssetVerification(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetVerification](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReviewAssetVerification200JSONResponse{Body: body, Headers: apigen.ReviewAssetVerification200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListAssetDisposals implements the listAssetDisposals operation.
func (s *Server) ListAssetDisposals(ctx context.Context, req apigen.ListAssetDisposalsRequestObject) (apigen.ListAssetDisposalsResponseObject, error) {
	return listPage[apigen.ListAssetDisposals200JSONResponse](ctx, req.Params, nil, s.service.ListAssetDisposal)
}

// GetAssetDisposal implements the getAssetDisposal operation.
func (s *Server) GetAssetDisposal(ctx context.Context, req apigen.GetAssetDisposalRequestObject) (apigen.GetAssetDisposalResponseObject, error) {
	row, err := s.service.GetAssetDisposal(ctx, req.AssetDisposalId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetDisposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetAssetDisposal200JSONResponse{Body: body, Headers: apigen.GetAssetDisposal200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateAssetDisposal implements the createAssetDisposal operation.
func (s *Server) CreateAssetDisposal(ctx context.Context, req apigen.CreateAssetDisposalRequestObject) (apigen.CreateAssetDisposalResponseObject, error) {
	in, err := input[db.CreateAssetDisposalParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateAssetDisposal(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetDisposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateAssetDisposal201JSONResponse{Body: body, Headers: apigen.CreateAssetDisposal201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateAssetDisposal implements the updateAssetDisposal operation.
func (s *Server) UpdateAssetDisposal(ctx context.Context, req apigen.UpdateAssetDisposalRequestObject) (apigen.UpdateAssetDisposalResponseObject, error) {
	in, err := input[db.UpdateAssetDisposalParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.AssetDisposalId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateAssetDisposal(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetDisposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateAssetDisposal200JSONResponse{Body: body, Headers: apigen.UpdateAssetDisposal200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitAssetDisposal implements the submitAssetDisposal operation.
func (s *Server) SubmitAssetDisposal(ctx context.Context, req apigen.SubmitAssetDisposalRequestObject) (apigen.SubmitAssetDisposalResponseObject, error) {
	in, err := input[db.ActSubmitAssetDisposalParams](struct{}{}, map[string]any{"id": req.AssetDisposalId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitAssetDisposal(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetDisposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitAssetDisposal200JSONResponse{Body: body, Headers: apigen.SubmitAssetDisposal200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ConfirmRemovalAssetDisposal implements the confirmRemovalAssetDisposal operation.
func (s *Server) ConfirmRemovalAssetDisposal(ctx context.Context, req apigen.ConfirmRemovalAssetDisposalRequestObject) (apigen.ConfirmRemovalAssetDisposalResponseObject, error) {
	in, err := input[db.ActConfirmRemovalAssetDisposalParams](struct{}{}, map[string]any{"id": req.AssetDisposalId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ConfirmRemovalAssetDisposal(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AssetDisposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.ConfirmRemovalAssetDisposal200JSONResponse{Body: body, Headers: apigen.ConfirmRemovalAssetDisposal200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListKeyItems implements the listKeyItems operation.
func (s *Server) ListKeyItems(ctx context.Context, req apigen.ListKeyItemsRequestObject) (apigen.ListKeyItemsResponseObject, error) {
	return listPage[apigen.ListKeyItems200JSONResponse](ctx, req.Params, nil, s.service.ListKeyItem)
}

// GetKeyItem implements the getKeyItem operation.
func (s *Server) GetKeyItem(ctx context.Context, req apigen.GetKeyItemRequestObject) (apigen.GetKeyItemResponseObject, error) {
	row, err := s.service.GetKeyItem(ctx, req.KeyItemId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyItem](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetKeyItem200JSONResponse{Body: body, Headers: apigen.GetKeyItem200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateKeyItem implements the createKeyItem operation.
func (s *Server) CreateKeyItem(ctx context.Context, req apigen.CreateKeyItemRequestObject) (apigen.CreateKeyItemResponseObject, error) {
	in, err := input[db.CreateKeyItemParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateKeyItem(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyItem](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateKeyItem201JSONResponse{Body: body, Headers: apigen.CreateKeyItem201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateKeyItem implements the updateKeyItem operation.
func (s *Server) UpdateKeyItem(ctx context.Context, req apigen.UpdateKeyItemRequestObject) (apigen.UpdateKeyItemResponseObject, error) {
	in, err := input[db.UpdateKeyItemParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.KeyItemId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateKeyItem(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyItem](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateKeyItem200JSONResponse{Body: body, Headers: apigen.UpdateKeyItem200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListKeyMovements implements the listKeyMovements operation.
func (s *Server) ListKeyMovements(ctx context.Context, req apigen.ListKeyMovementsRequestObject) (apigen.ListKeyMovementsResponseObject, error) {
	return listPage[apigen.ListKeyMovements200JSONResponse](ctx, req.Params, nil, s.service.ListKeyMovement)
}

// GetKeyMovement implements the getKeyMovement operation.
func (s *Server) GetKeyMovement(ctx context.Context, req apigen.GetKeyMovementRequestObject) (apigen.GetKeyMovementResponseObject, error) {
	row, err := s.service.GetKeyMovement(ctx, req.KeyMovementId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyMovement](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetKeyMovement200JSONResponse{Body: body, Headers: apigen.GetKeyMovement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateKeyMovement implements the createKeyMovement operation.
func (s *Server) CreateKeyMovement(ctx context.Context, req apigen.CreateKeyMovementRequestObject) (apigen.CreateKeyMovementResponseObject, error) {
	in, err := input[db.CreateKeyMovementParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateKeyMovement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyMovement](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateKeyMovement201JSONResponse{Body: body, Headers: apigen.CreateKeyMovement201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReturnKeyMovement implements the returnKeyMovement operation.
func (s *Server) ReturnKeyMovement(ctx context.Context, req apigen.ReturnKeyMovementRequestObject) (apigen.ReturnKeyMovementResponseObject, error) {
	in, err := input[db.ActReturnKeyMovementParams](req.Body, map[string]any{"id": req.KeyMovementId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReturnKeyMovement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyMovement](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReturnKeyMovement200JSONResponse{Body: body, Headers: apigen.ReturnKeyMovement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListKeyDuplicationRequests implements the listKeyDuplicationRequests operation.
func (s *Server) ListKeyDuplicationRequests(ctx context.Context, req apigen.ListKeyDuplicationRequestsRequestObject) (apigen.ListKeyDuplicationRequestsResponseObject, error) {
	return listPage[apigen.ListKeyDuplicationRequests200JSONResponse](ctx, req.Params, nil, s.service.ListKeyDuplicationRequest)
}

// GetKeyDuplicationRequest implements the getKeyDuplicationRequest operation.
func (s *Server) GetKeyDuplicationRequest(ctx context.Context, req apigen.GetKeyDuplicationRequestRequestObject) (apigen.GetKeyDuplicationRequestResponseObject, error) {
	row, err := s.service.GetKeyDuplicationRequest(ctx, req.KeyDuplicationRequestId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyDuplicationRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetKeyDuplicationRequest200JSONResponse{Body: body, Headers: apigen.GetKeyDuplicationRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateKeyDuplicationRequest implements the createKeyDuplicationRequest operation.
func (s *Server) CreateKeyDuplicationRequest(ctx context.Context, req apigen.CreateKeyDuplicationRequestRequestObject) (apigen.CreateKeyDuplicationRequestResponseObject, error) {
	in, err := input[db.CreateKeyDuplicationRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateKeyDuplicationRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyDuplicationRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateKeyDuplicationRequest201JSONResponse{Body: body, Headers: apigen.CreateKeyDuplicationRequest201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateKeyDuplicationRequest implements the updateKeyDuplicationRequest operation.
func (s *Server) UpdateKeyDuplicationRequest(ctx context.Context, req apigen.UpdateKeyDuplicationRequestRequestObject) (apigen.UpdateKeyDuplicationRequestResponseObject, error) {
	in, err := input[db.UpdateKeyDuplicationRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.KeyDuplicationRequestId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateKeyDuplicationRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyDuplicationRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateKeyDuplicationRequest200JSONResponse{Body: body, Headers: apigen.UpdateKeyDuplicationRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitKeyDuplicationRequest implements the submitKeyDuplicationRequest operation.
func (s *Server) SubmitKeyDuplicationRequest(ctx context.Context, req apigen.SubmitKeyDuplicationRequestRequestObject) (apigen.SubmitKeyDuplicationRequestResponseObject, error) {
	in, err := input[db.ActSubmitKeyDuplicationRequestParams](struct{}{}, map[string]any{"id": req.KeyDuplicationRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitKeyDuplicationRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyDuplicationRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitKeyDuplicationRequest200JSONResponse{Body: body, Headers: apigen.SubmitKeyDuplicationRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListKeyAudits implements the listKeyAudits operation.
func (s *Server) ListKeyAudits(ctx context.Context, req apigen.ListKeyAuditsRequestObject) (apigen.ListKeyAuditsResponseObject, error) {
	return listPage[apigen.ListKeyAudits200JSONResponse](ctx, req.Params, nil, s.service.ListKeyAudit)
}

// GetKeyAudit implements the getKeyAudit operation.
func (s *Server) GetKeyAudit(ctx context.Context, req apigen.GetKeyAuditRequestObject) (apigen.GetKeyAuditResponseObject, error) {
	row, err := s.service.GetKeyAudit(ctx, req.KeyAuditId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.KeyAuditChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyAudit](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetKeyAudit200JSONResponse{Body: body, Headers: apigen.GetKeyAudit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateKeyAudit implements the createKeyAudit operation.
func (s *Server) CreateKeyAudit(ctx context.Context, req apigen.CreateKeyAuditRequestObject) (apigen.CreateKeyAuditResponseObject, error) {
	in, err := input[db.CreateKeyAuditParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateKeyAudit(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertKeyAuditLinesParams](items, "key_audit_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceKeyAuditLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.KeyAuditChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyAudit](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateKeyAudit201JSONResponse{Body: body, Headers: apigen.CreateKeyAudit201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateKeyAudit implements the updateKeyAudit operation.
func (s *Server) UpdateKeyAudit(ctx context.Context, req apigen.UpdateKeyAuditRequestObject) (apigen.UpdateKeyAuditResponseObject, error) {
	in, err := input[db.UpdateKeyAuditParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.KeyAuditId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateKeyAudit(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertKeyAuditLinesParams](items, "key_audit_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceKeyAuditLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.KeyAuditChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyAudit](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateKeyAudit200JSONResponse{Body: body, Headers: apigen.UpdateKeyAudit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteKeyAudit implements the completeKeyAudit operation.
func (s *Server) CompleteKeyAudit(ctx context.Context, req apigen.CompleteKeyAuditRequestObject) (apigen.CompleteKeyAuditResponseObject, error) {
	in, err := input[db.ActCompleteKeyAuditParams](struct{}{}, map[string]any{"id": req.KeyAuditId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteKeyAudit(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KeyAudit](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteKeyAudit200JSONResponse{Body: body, Headers: apigen.CompleteKeyAudit200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListVisitorEntries implements the listVisitorEntries operation.
func (s *Server) ListVisitorEntries(ctx context.Context, req apigen.ListVisitorEntriesRequestObject) (apigen.ListVisitorEntriesResponseObject, error) {
	return listPage[apigen.ListVisitorEntries200JSONResponse](ctx, req.Params, nil, s.service.ListVisitorEntry)
}

// GetVisitorEntry implements the getVisitorEntry operation.
func (s *Server) GetVisitorEntry(ctx context.Context, req apigen.GetVisitorEntryRequestObject) (apigen.GetVisitorEntryResponseObject, error) {
	row, err := s.service.GetVisitorEntry(ctx, req.VisitorEntryId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.VisitorEntry](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetVisitorEntry200JSONResponse{Body: body, Headers: apigen.GetVisitorEntry200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateVisitorEntry implements the createVisitorEntry operation.
func (s *Server) CreateVisitorEntry(ctx context.Context, req apigen.CreateVisitorEntryRequestObject) (apigen.CreateVisitorEntryResponseObject, error) {
	in, err := input[db.CreateVisitorEntryParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateVisitorEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.VisitorEntry](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateVisitorEntry201JSONResponse{Body: body, Headers: apigen.CreateVisitorEntry201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateVisitorEntry implements the updateVisitorEntry operation.
func (s *Server) UpdateVisitorEntry(ctx context.Context, req apigen.UpdateVisitorEntryRequestObject) (apigen.UpdateVisitorEntryResponseObject, error) {
	in, err := input[db.UpdateVisitorEntryParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.VisitorEntryId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateVisitorEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.VisitorEntry](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateVisitorEntry200JSONResponse{Body: body, Headers: apigen.UpdateVisitorEntry200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SignOutVisitorEntry implements the signOutVisitorEntry operation.
func (s *Server) SignOutVisitorEntry(ctx context.Context, req apigen.SignOutVisitorEntryRequestObject) (apigen.SignOutVisitorEntryResponseObject, error) {
	in, err := input[db.ActSignOutVisitorEntryParams](struct{}{}, map[string]any{"id": req.VisitorEntryId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SignOutVisitorEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.VisitorEntry](row)
	if err != nil {
		return nil, err
	}
	return apigen.SignOutVisitorEntry200JSONResponse{Body: body, Headers: apigen.SignOutVisitorEntry200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListLostPropertyItems implements the listLostPropertyItems operation.
func (s *Server) ListLostPropertyItems(ctx context.Context, req apigen.ListLostPropertyItemsRequestObject) (apigen.ListLostPropertyItemsResponseObject, error) {
	return listPage[apigen.ListLostPropertyItems200JSONResponse](ctx, req.Params, nil, s.service.ListLostPropertyItem)
}

// GetLostPropertyItem implements the getLostPropertyItem operation.
func (s *Server) GetLostPropertyItem(ctx context.Context, req apigen.GetLostPropertyItemRequestObject) (apigen.GetLostPropertyItemResponseObject, error) {
	row, err := s.service.GetLostPropertyItem(ctx, req.LostPropertyItemId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LostPropertyItem](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetLostPropertyItem200JSONResponse{Body: body, Headers: apigen.GetLostPropertyItem200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateLostPropertyItem implements the createLostPropertyItem operation.
func (s *Server) CreateLostPropertyItem(ctx context.Context, req apigen.CreateLostPropertyItemRequestObject) (apigen.CreateLostPropertyItemResponseObject, error) {
	in, err := input[db.CreateLostPropertyItemParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateLostPropertyItem(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LostPropertyItem](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateLostPropertyItem201JSONResponse{Body: body, Headers: apigen.CreateLostPropertyItem201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateLostPropertyItem implements the updateLostPropertyItem operation.
func (s *Server) UpdateLostPropertyItem(ctx context.Context, req apigen.UpdateLostPropertyItemRequestObject) (apigen.UpdateLostPropertyItemResponseObject, error) {
	in, err := input[db.UpdateLostPropertyItemParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.LostPropertyItemId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateLostPropertyItem(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LostPropertyItem](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateLostPropertyItem200JSONResponse{Body: body, Headers: apigen.UpdateLostPropertyItem200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReturnLostPropertyItem implements the returnLostPropertyItem operation.
func (s *Server) ReturnLostPropertyItem(ctx context.Context, req apigen.ReturnLostPropertyItemRequestObject) (apigen.ReturnLostPropertyItemResponseObject, error) {
	in, err := input[db.ActReturnLostPropertyItemParams](req.Body, map[string]any{"id": req.LostPropertyItemId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReturnLostPropertyItem(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LostPropertyItem](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReturnLostPropertyItem200JSONResponse{Body: body, Headers: apigen.ReturnLostPropertyItem200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DisposeLostPropertyItem implements the disposeLostPropertyItem operation.
func (s *Server) DisposeLostPropertyItem(ctx context.Context, req apigen.DisposeLostPropertyItemRequestObject) (apigen.DisposeLostPropertyItemResponseObject, error) {
	in, err := input[db.ActDisposeLostPropertyItemParams](req.Body, map[string]any{"id": req.LostPropertyItemId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.DisposeLostPropertyItem(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LostPropertyItem](row)
	if err != nil {
		return nil, err
	}
	return apigen.DisposeLostPropertyItem200JSONResponse{Body: body, Headers: apigen.DisposeLostPropertyItem200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDocuments implements the listDocuments operation.
func (s *Server) ListDocuments(ctx context.Context, req apigen.ListDocumentsRequestObject) (apigen.ListDocumentsResponseObject, error) {
	return listPage[apigen.ListDocuments200JSONResponse](ctx, req.Params, nil, s.service.ListDocument)
}

// GetDocument implements the getDocument operation.
func (s *Server) GetDocument(ctx context.Context, req apigen.GetDocumentRequestObject) (apigen.GetDocumentResponseObject, error) {
	row, err := s.service.GetDocument(ctx, req.DocumentId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Document](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDocument200JSONResponse{Body: body, Headers: apigen.GetDocument200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDocument implements the createDocument operation.
func (s *Server) CreateDocument(ctx context.Context, req apigen.CreateDocumentRequestObject) (apigen.CreateDocumentResponseObject, error) {
	in, err := input[db.CreateDocumentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDocument(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Document](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDocument201JSONResponse{Body: body, Headers: apigen.CreateDocument201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDocument implements the updateDocument operation.
func (s *Server) UpdateDocument(ctx context.Context, req apigen.UpdateDocumentRequestObject) (apigen.UpdateDocumentResponseObject, error) {
	in, err := input[db.UpdateDocumentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.DocumentId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDocument(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Document](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDocument200JSONResponse{Body: body, Headers: apigen.UpdateDocument200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDocumentRevisions implements the listDocumentRevisions operation.
func (s *Server) ListDocumentRevisions(ctx context.Context, req apigen.ListDocumentRevisionsRequestObject) (apigen.ListDocumentRevisionsResponseObject, error) {
	return listPage[apigen.ListDocumentRevisions200JSONResponse](ctx, req.Params, map[string]any{"document_id": req.DocumentId}, s.service.ListDocumentRevision)
}

// GetDocumentRevision implements the getDocumentRevision operation.
func (s *Server) GetDocumentRevision(ctx context.Context, req apigen.GetDocumentRevisionRequestObject) (apigen.GetDocumentRevisionResponseObject, error) {
	row, err := s.service.GetDocumentRevision(ctx, req.DocumentRevisionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRevision](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDocumentRevision200JSONResponse{Body: body, Headers: apigen.GetDocumentRevision200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDocumentRevision implements the createDocumentRevision operation.
func (s *Server) CreateDocumentRevision(ctx context.Context, req apigen.CreateDocumentRevisionRequestObject) (apigen.CreateDocumentRevisionResponseObject, error) {
	in, err := input[db.CreateDocumentRevisionParams](req.Body, map[string]any{"document_id": req.DocumentId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDocumentRevision(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRevision](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDocumentRevision201JSONResponse{Body: body, Headers: apigen.CreateDocumentRevision201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDocumentRevision implements the updateDocumentRevision operation.
func (s *Server) UpdateDocumentRevision(ctx context.Context, req apigen.UpdateDocumentRevisionRequestObject) (apigen.UpdateDocumentRevisionResponseObject, error) {
	in, err := input[db.UpdateDocumentRevisionParams](req.Body, map[string]any{"document_id": req.DocumentId})
	if err != nil {
		return nil, err
	}
	in.ID = req.DocumentRevisionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDocumentRevision(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRevision](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDocumentRevision200JSONResponse{Body: body, Headers: apigen.UpdateDocumentRevision200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitDocumentRevision implements the submitDocumentRevision operation.
func (s *Server) SubmitDocumentRevision(ctx context.Context, req apigen.SubmitDocumentRevisionRequestObject) (apigen.SubmitDocumentRevisionResponseObject, error) {
	in, err := input[db.ActSubmitDocumentRevisionParams](struct{}{}, map[string]any{"id": req.DocumentRevisionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitDocumentRevision(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRevision](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitDocumentRevision200JSONResponse{Body: body, Headers: apigen.SubmitDocumentRevision200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// IssueDocumentRevision implements the issueDocumentRevision operation.
func (s *Server) IssueDocumentRevision(ctx context.Context, req apigen.IssueDocumentRevisionRequestObject) (apigen.IssueDocumentRevisionResponseObject, error) {
	in, err := input[db.ActIssueDocumentRevisionParams](struct{}{}, map[string]any{"id": req.DocumentRevisionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.IssueDocumentRevision(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRevision](row)
	if err != nil {
		return nil, err
	}
	return apigen.IssueDocumentRevision200JSONResponse{Body: body, Headers: apigen.IssueDocumentRevision200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// WithdrawDocumentRevision implements the withdrawDocumentRevision operation.
func (s *Server) WithdrawDocumentRevision(ctx context.Context, req apigen.WithdrawDocumentRevisionRequestObject) (apigen.WithdrawDocumentRevisionResponseObject, error) {
	in, err := input[db.ActWithdrawDocumentRevisionParams](struct{}{}, map[string]any{"id": req.DocumentRevisionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.WithdrawDocumentRevision(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRevision](row)
	if err != nil {
		return nil, err
	}
	return apigen.WithdrawDocumentRevision200JSONResponse{Body: body, Headers: apigen.WithdrawDocumentRevision200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDocumentRequests implements the listDocumentRequests operation.
func (s *Server) ListDocumentRequests(ctx context.Context, req apigen.ListDocumentRequestsRequestObject) (apigen.ListDocumentRequestsResponseObject, error) {
	return listPage[apigen.ListDocumentRequests200JSONResponse](ctx, req.Params, nil, s.service.ListDocumentRequest)
}

// GetDocumentRequest implements the getDocumentRequest operation.
func (s *Server) GetDocumentRequest(ctx context.Context, req apigen.GetDocumentRequestRequestObject) (apigen.GetDocumentRequestResponseObject, error) {
	row, err := s.service.GetDocumentRequest(ctx, req.DocumentRequestId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDocumentRequest200JSONResponse{Body: body, Headers: apigen.GetDocumentRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDocumentRequest implements the createDocumentRequest operation.
func (s *Server) CreateDocumentRequest(ctx context.Context, req apigen.CreateDocumentRequestRequestObject) (apigen.CreateDocumentRequestResponseObject, error) {
	in, err := input[db.CreateDocumentRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDocumentRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDocumentRequest201JSONResponse{Body: body, Headers: apigen.CreateDocumentRequest201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDocumentRequest implements the updateDocumentRequest operation.
func (s *Server) UpdateDocumentRequest(ctx context.Context, req apigen.UpdateDocumentRequestRequestObject) (apigen.UpdateDocumentRequestResponseObject, error) {
	in, err := input[db.UpdateDocumentRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.DocumentRequestId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDocumentRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDocumentRequest200JSONResponse{Body: body, Headers: apigen.UpdateDocumentRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseDocumentRequest implements the closeDocumentRequest operation.
func (s *Server) CloseDocumentRequest(ctx context.Context, req apigen.CloseDocumentRequestRequestObject) (apigen.CloseDocumentRequestResponseObject, error) {
	in, err := input[db.ActCloseDocumentRequestParams](struct{}{}, map[string]any{"id": req.DocumentRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseDocumentRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseDocumentRequest200JSONResponse{Body: body, Headers: apigen.CloseDocumentRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDocumentAcknowledgements implements the listDocumentAcknowledgements operation.
func (s *Server) ListDocumentAcknowledgements(ctx context.Context, req apigen.ListDocumentAcknowledgementsRequestObject) (apigen.ListDocumentAcknowledgementsResponseObject, error) {
	return listPage[apigen.ListDocumentAcknowledgements200JSONResponse](ctx, req.Params, nil, s.service.ListDocumentAcknowledgement)
}

// GetDocumentAcknowledgement implements the getDocumentAcknowledgement operation.
func (s *Server) GetDocumentAcknowledgement(ctx context.Context, req apigen.GetDocumentAcknowledgementRequestObject) (apigen.GetDocumentAcknowledgementResponseObject, error) {
	row, err := s.service.GetDocumentAcknowledgement(ctx, req.DocumentAcknowledgementId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentAcknowledgement](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDocumentAcknowledgement200JSONResponse{Body: body, Headers: apigen.GetDocumentAcknowledgement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AcknowledgeDocumentAcknowledgement implements the acknowledgeDocumentAcknowledgement operation.
func (s *Server) AcknowledgeDocumentAcknowledgement(ctx context.Context, req apigen.AcknowledgeDocumentAcknowledgementRequestObject) (apigen.AcknowledgeDocumentAcknowledgementResponseObject, error) {
	in, err := input[db.ActAcknowledgeDocumentAcknowledgementParams](req.Body, map[string]any{"id": req.DocumentAcknowledgementId})
	if err != nil {
		return nil, err
	}
	if err := s.service.RequireSigned(ctx, "document_acknowledgements", req.DocumentAcknowledgementId, &req.Body.SignatureId); err != nil {
		return nil, err
	}
	row, err := s.service.AcknowledgeDocumentAcknowledgement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DocumentAcknowledgement](row)
	if err != nil {
		return nil, err
	}
	return apigen.AcknowledgeDocumentAcknowledgement200JSONResponse{Body: body, Headers: apigen.AcknowledgeDocumentAcknowledgement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCourses implements the listCourses operation.
func (s *Server) ListCourses(ctx context.Context, req apigen.ListCoursesRequestObject) (apigen.ListCoursesResponseObject, error) {
	return listPage[apigen.ListCourses200JSONResponse](ctx, req.Params, nil, s.service.ListCourse)
}

// GetCourse implements the getCourse operation.
func (s *Server) GetCourse(ctx context.Context, req apigen.GetCourseRequestObject) (apigen.GetCourseResponseObject, error) {
	row, err := s.service.GetCourse(ctx, req.CourseId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Course](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCourse200JSONResponse{Body: body, Headers: apigen.GetCourse200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCourse implements the createCourse operation.
func (s *Server) CreateCourse(ctx context.Context, req apigen.CreateCourseRequestObject) (apigen.CreateCourseResponseObject, error) {
	in, err := input[db.CreateCourseParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCourse(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Course](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCourse201JSONResponse{Body: body, Headers: apigen.CreateCourse201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCourse implements the updateCourse operation.
func (s *Server) UpdateCourse(ctx context.Context, req apigen.UpdateCourseRequestObject) (apigen.UpdateCourseResponseObject, error) {
	in, err := input[db.UpdateCourseParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CourseId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCourse(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Course](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCourse200JSONResponse{Body: body, Headers: apigen.UpdateCourse200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListTrainingRequirements implements the listTrainingRequirements operation.
func (s *Server) ListTrainingRequirements(ctx context.Context, req apigen.ListTrainingRequirementsRequestObject) (apigen.ListTrainingRequirementsResponseObject, error) {
	return listPage[apigen.ListTrainingRequirements200JSONResponse](ctx, req.Params, nil, s.service.ListTrainingRequirement)
}

// GetTrainingRequirement implements the getTrainingRequirement operation.
func (s *Server) GetTrainingRequirement(ctx context.Context, req apigen.GetTrainingRequirementRequestObject) (apigen.GetTrainingRequirementResponseObject, error) {
	row, err := s.service.GetTrainingRequirement(ctx, req.TrainingRequirementId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingRequirement](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetTrainingRequirement200JSONResponse{Body: body, Headers: apigen.GetTrainingRequirement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateTrainingRequirement implements the createTrainingRequirement operation.
func (s *Server) CreateTrainingRequirement(ctx context.Context, req apigen.CreateTrainingRequirementRequestObject) (apigen.CreateTrainingRequirementResponseObject, error) {
	in, err := input[db.CreateTrainingRequirementParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateTrainingRequirement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingRequirement](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateTrainingRequirement201JSONResponse{Body: body, Headers: apigen.CreateTrainingRequirement201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateTrainingRequirement implements the updateTrainingRequirement operation.
func (s *Server) UpdateTrainingRequirement(ctx context.Context, req apigen.UpdateTrainingRequirementRequestObject) (apigen.UpdateTrainingRequirementResponseObject, error) {
	in, err := input[db.UpdateTrainingRequirementParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.TrainingRequirementId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateTrainingRequirement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingRequirement](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateTrainingRequirement200JSONResponse{Body: body, Headers: apigen.UpdateTrainingRequirement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DeleteTrainingRequirement implements the deleteTrainingRequirement operation.
func (s *Server) DeleteTrainingRequirement(ctx context.Context, req apigen.DeleteTrainingRequirementRequestObject) (apigen.DeleteTrainingRequirementResponseObject, error) {
	if err := s.service.DeleteTrainingRequirement(ctx, req.TrainingRequirementId); err != nil {
		return nil, err
	}
	return apigen.DeleteTrainingRequirement204Response{}, nil
}

// ListTrainingSessions implements the listTrainingSessions operation.
func (s *Server) ListTrainingSessions(ctx context.Context, req apigen.ListTrainingSessionsRequestObject) (apigen.ListTrainingSessionsResponseObject, error) {
	return listPage[apigen.ListTrainingSessions200JSONResponse](ctx, req.Params, nil, s.service.ListTrainingSession)
}

// GetTrainingSession implements the getTrainingSession operation.
func (s *Server) GetTrainingSession(ctx context.Context, req apigen.GetTrainingSessionRequestObject) (apigen.GetTrainingSessionResponseObject, error) {
	row, err := s.service.GetTrainingSession(ctx, req.TrainingSessionId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.TrainingSessionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingSession](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetTrainingSession200JSONResponse{Body: body, Headers: apigen.GetTrainingSession200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateTrainingSession implements the createTrainingSession operation.
func (s *Server) CreateTrainingSession(ctx context.Context, req apigen.CreateTrainingSessionRequestObject) (apigen.CreateTrainingSessionResponseObject, error) {
	in, err := input[db.CreateTrainingSessionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateTrainingSession(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["participants"]; ok {
		rows, err := childInputs[db.InsertTrainingParticipantsParams](items, "training_session_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceTrainingParticipants(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.TrainingSessionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingSession](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateTrainingSession201JSONResponse{Body: body, Headers: apigen.CreateTrainingSession201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateTrainingSession implements the updateTrainingSession operation.
func (s *Server) UpdateTrainingSession(ctx context.Context, req apigen.UpdateTrainingSessionRequestObject) (apigen.UpdateTrainingSessionResponseObject, error) {
	in, err := input[db.UpdateTrainingSessionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.TrainingSessionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateTrainingSession(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["participants"]; ok {
		rows, err := childInputs[db.InsertTrainingParticipantsParams](items, "training_session_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceTrainingParticipants(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.TrainingSessionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingSession](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateTrainingSession200JSONResponse{Body: body, Headers: apigen.UpdateTrainingSession200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseTrainingSession implements the closeTrainingSession operation.
func (s *Server) CloseTrainingSession(ctx context.Context, req apigen.CloseTrainingSessionRequestObject) (apigen.CloseTrainingSessionResponseObject, error) {
	in, err := input[db.ActCloseTrainingSessionParams](struct{}{}, map[string]any{"id": req.TrainingSessionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseTrainingSession(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingSession](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseTrainingSession200JSONResponse{Body: body, Headers: apigen.CloseTrainingSession200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListTrainingAssessments implements the listTrainingAssessments operation.
func (s *Server) ListTrainingAssessments(ctx context.Context, req apigen.ListTrainingAssessmentsRequestObject) (apigen.ListTrainingAssessmentsResponseObject, error) {
	return listPage[apigen.ListTrainingAssessments200JSONResponse](ctx, req.Params, nil, s.service.ListTrainingAssessment)
}

// GetTrainingAssessment implements the getTrainingAssessment operation.
func (s *Server) GetTrainingAssessment(ctx context.Context, req apigen.GetTrainingAssessmentRequestObject) (apigen.GetTrainingAssessmentResponseObject, error) {
	row, err := s.service.GetTrainingAssessment(ctx, req.TrainingAssessmentId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingAssessment](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetTrainingAssessment200JSONResponse{Body: body, Headers: apigen.GetTrainingAssessment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateTrainingAssessment implements the createTrainingAssessment operation.
func (s *Server) CreateTrainingAssessment(ctx context.Context, req apigen.CreateTrainingAssessmentRequestObject) (apigen.CreateTrainingAssessmentResponseObject, error) {
	in, err := input[db.CreateTrainingAssessmentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateTrainingAssessment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingAssessment](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateTrainingAssessment201JSONResponse{Body: body, Headers: apigen.CreateTrainingAssessment201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCompetencies implements the listCompetencies operation.
func (s *Server) ListCompetencies(ctx context.Context, req apigen.ListCompetenciesRequestObject) (apigen.ListCompetenciesResponseObject, error) {
	return listPage[apigen.ListCompetencies200JSONResponse](ctx, req.Params, nil, s.service.ListCompetency)
}

// GetCompetency implements the getCompetency operation.
func (s *Server) GetCompetency(ctx context.Context, req apigen.GetCompetencyRequestObject) (apigen.GetCompetencyResponseObject, error) {
	row, err := s.service.GetCompetency(ctx, req.CompetencyId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Competency](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCompetency200JSONResponse{Body: body, Headers: apigen.GetCompetency200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCompetency implements the createCompetency operation.
func (s *Server) CreateCompetency(ctx context.Context, req apigen.CreateCompetencyRequestObject) (apigen.CreateCompetencyResponseObject, error) {
	in, err := input[db.CreateCompetencyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCompetency(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Competency](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCompetency201JSONResponse{Body: body, Headers: apigen.CreateCompetency201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCompetency implements the updateCompetency operation.
func (s *Server) UpdateCompetency(ctx context.Context, req apigen.UpdateCompetencyRequestObject) (apigen.UpdateCompetencyResponseObject, error) {
	in, err := input[db.UpdateCompetencyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CompetencyId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCompetency(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Competency](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCompetency200JSONResponse{Body: body, Headers: apigen.UpdateCompetency200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCertificates implements the listCertificates operation.
func (s *Server) ListCertificates(ctx context.Context, req apigen.ListCertificatesRequestObject) (apigen.ListCertificatesResponseObject, error) {
	return listPage[apigen.ListCertificates200JSONResponse](ctx, req.Params, nil, s.service.ListCertificate)
}

// GetCertificate implements the getCertificate operation.
func (s *Server) GetCertificate(ctx context.Context, req apigen.GetCertificateRequestObject) (apigen.GetCertificateResponseObject, error) {
	row, err := s.service.GetCertificate(ctx, req.CertificateId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Certificate](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCertificate200JSONResponse{Body: body, Headers: apigen.GetCertificate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCertificate implements the createCertificate operation.
func (s *Server) CreateCertificate(ctx context.Context, req apigen.CreateCertificateRequestObject) (apigen.CreateCertificateResponseObject, error) {
	in, err := input[db.CreateCertificateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCertificate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Certificate](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCertificate201JSONResponse{Body: body, Headers: apigen.CreateCertificate201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCertificate implements the updateCertificate operation.
func (s *Server) UpdateCertificate(ctx context.Context, req apigen.UpdateCertificateRequestObject) (apigen.UpdateCertificateResponseObject, error) {
	in, err := input[db.UpdateCertificateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CertificateId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCertificate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Certificate](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCertificate200JSONResponse{Body: body, Headers: apigen.UpdateCertificate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListTrainingRequests implements the listTrainingRequests operation.
func (s *Server) ListTrainingRequests(ctx context.Context, req apigen.ListTrainingRequestsRequestObject) (apigen.ListTrainingRequestsResponseObject, error) {
	return listPage[apigen.ListTrainingRequests200JSONResponse](ctx, req.Params, nil, s.service.ListTrainingRequest)
}

// GetTrainingRequest implements the getTrainingRequest operation.
func (s *Server) GetTrainingRequest(ctx context.Context, req apigen.GetTrainingRequestRequestObject) (apigen.GetTrainingRequestResponseObject, error) {
	row, err := s.service.GetTrainingRequest(ctx, req.TrainingRequestId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetTrainingRequest200JSONResponse{Body: body, Headers: apigen.GetTrainingRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateTrainingRequest implements the createTrainingRequest operation.
func (s *Server) CreateTrainingRequest(ctx context.Context, req apigen.CreateTrainingRequestRequestObject) (apigen.CreateTrainingRequestResponseObject, error) {
	in, err := input[db.CreateTrainingRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateTrainingRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateTrainingRequest201JSONResponse{Body: body, Headers: apigen.CreateTrainingRequest201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateTrainingRequest implements the updateTrainingRequest operation.
func (s *Server) UpdateTrainingRequest(ctx context.Context, req apigen.UpdateTrainingRequestRequestObject) (apigen.UpdateTrainingRequestResponseObject, error) {
	in, err := input[db.UpdateTrainingRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.TrainingRequestId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateTrainingRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateTrainingRequest200JSONResponse{Body: body, Headers: apigen.UpdateTrainingRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitTrainingRequest implements the submitTrainingRequest operation.
func (s *Server) SubmitTrainingRequest(ctx context.Context, req apigen.SubmitTrainingRequestRequestObject) (apigen.SubmitTrainingRequestResponseObject, error) {
	in, err := input[db.ActSubmitTrainingRequestParams](struct{}{}, map[string]any{"id": req.TrainingRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitTrainingRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.TrainingRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitTrainingRequest200JSONResponse{Body: body, Headers: apigen.SubmitTrainingRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListLeaveTypes implements the listLeaveTypes operation.
func (s *Server) ListLeaveTypes(ctx context.Context, req apigen.ListLeaveTypesRequestObject) (apigen.ListLeaveTypesResponseObject, error) {
	return listPage[apigen.ListLeaveTypes200JSONResponse](ctx, req.Params, nil, s.service.ListLeaveType)
}

// GetLeaveType implements the getLeaveType operation.
func (s *Server) GetLeaveType(ctx context.Context, req apigen.GetLeaveTypeRequestObject) (apigen.GetLeaveTypeResponseObject, error) {
	row, err := s.service.GetLeaveType(ctx, req.LeaveTypeId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LeaveType](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetLeaveType200JSONResponse{Body: body, Headers: apigen.GetLeaveType200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateLeaveType implements the createLeaveType operation.
func (s *Server) CreateLeaveType(ctx context.Context, req apigen.CreateLeaveTypeRequestObject) (apigen.CreateLeaveTypeResponseObject, error) {
	in, err := input[db.CreateLeaveTypeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateLeaveType(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LeaveType](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateLeaveType201JSONResponse{Body: body, Headers: apigen.CreateLeaveType201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateLeaveType implements the updateLeaveType operation.
func (s *Server) UpdateLeaveType(ctx context.Context, req apigen.UpdateLeaveTypeRequestObject) (apigen.UpdateLeaveTypeResponseObject, error) {
	in, err := input[db.UpdateLeaveTypeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.LeaveTypeId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateLeaveType(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LeaveType](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateLeaveType200JSONResponse{Body: body, Headers: apigen.UpdateLeaveType200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListLeaveBalances implements the listLeaveBalances operation.
func (s *Server) ListLeaveBalances(ctx context.Context, req apigen.ListLeaveBalancesRequestObject) (apigen.ListLeaveBalancesResponseObject, error) {
	return listPage[apigen.ListLeaveBalances200JSONResponse](ctx, req.Params, nil, s.service.ListLeaveBalance)
}

// GetLeaveBalance implements the getLeaveBalance operation.
func (s *Server) GetLeaveBalance(ctx context.Context, req apigen.GetLeaveBalanceRequestObject) (apigen.GetLeaveBalanceResponseObject, error) {
	row, err := s.service.GetLeaveBalance(ctx, req.LeaveBalanceId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LeaveBalance](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetLeaveBalance200JSONResponse{Body: body, Headers: apigen.GetLeaveBalance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListLeaveRequests implements the listLeaveRequests operation.
func (s *Server) ListLeaveRequests(ctx context.Context, req apigen.ListLeaveRequestsRequestObject) (apigen.ListLeaveRequestsResponseObject, error) {
	return listPage[apigen.ListLeaveRequests200JSONResponse](ctx, req.Params, nil, s.service.ListLeaveRequest)
}

// GetLeaveRequest implements the getLeaveRequest operation.
func (s *Server) GetLeaveRequest(ctx context.Context, req apigen.GetLeaveRequestRequestObject) (apigen.GetLeaveRequestResponseObject, error) {
	row, err := s.service.GetLeaveRequest(ctx, req.LeaveRequestId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LeaveRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetLeaveRequest200JSONResponse{Body: body, Headers: apigen.GetLeaveRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateLeaveRequest implements the createLeaveRequest operation.
func (s *Server) CreateLeaveRequest(ctx context.Context, req apigen.CreateLeaveRequestRequestObject) (apigen.CreateLeaveRequestResponseObject, error) {
	in, err := input[db.CreateLeaveRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateLeaveRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LeaveRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateLeaveRequest201JSONResponse{Body: body, Headers: apigen.CreateLeaveRequest201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateLeaveRequest implements the updateLeaveRequest operation.
func (s *Server) UpdateLeaveRequest(ctx context.Context, req apigen.UpdateLeaveRequestRequestObject) (apigen.UpdateLeaveRequestResponseObject, error) {
	in, err := input[db.UpdateLeaveRequestParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.LeaveRequestId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateLeaveRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LeaveRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateLeaveRequest200JSONResponse{Body: body, Headers: apigen.UpdateLeaveRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitLeaveRequest implements the submitLeaveRequest operation.
func (s *Server) SubmitLeaveRequest(ctx context.Context, req apigen.SubmitLeaveRequestRequestObject) (apigen.SubmitLeaveRequestResponseObject, error) {
	in, err := input[db.ActSubmitLeaveRequestParams](struct{}{}, map[string]any{"id": req.LeaveRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitLeaveRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LeaveRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitLeaveRequest200JSONResponse{Body: body, Headers: apigen.SubmitLeaveRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CancelLeaveRequest implements the cancelLeaveRequest operation.
func (s *Server) CancelLeaveRequest(ctx context.Context, req apigen.CancelLeaveRequestRequestObject) (apigen.CancelLeaveRequestResponseObject, error) {
	in, err := input[db.ActCancelLeaveRequestParams](struct{}{}, map[string]any{"id": req.LeaveRequestId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CancelLeaveRequest(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LeaveRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.CancelLeaveRequest200JSONResponse{Body: body, Headers: apigen.CancelLeaveRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListExitClearances implements the listExitClearances operation.
func (s *Server) ListExitClearances(ctx context.Context, req apigen.ListExitClearancesRequestObject) (apigen.ListExitClearancesResponseObject, error) {
	return listPage[apigen.ListExitClearances200JSONResponse](ctx, req.Params, nil, s.service.ListExitClearance)
}

// GetExitClearance implements the getExitClearance operation.
func (s *Server) GetExitClearance(ctx context.Context, req apigen.GetExitClearanceRequestObject) (apigen.GetExitClearanceResponseObject, error) {
	row, err := s.service.GetExitClearance(ctx, req.ExitClearanceId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.ExitClearanceChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExitClearance](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetExitClearance200JSONResponse{Body: body, Headers: apigen.GetExitClearance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateExitClearance implements the createExitClearance operation.
func (s *Server) CreateExitClearance(ctx context.Context, req apigen.CreateExitClearanceRequestObject) (apigen.CreateExitClearanceResponseObject, error) {
	in, err := input[db.CreateExitClearanceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateExitClearance(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["sections"]; ok {
		rows, err := childInputs[db.InsertExitClearanceSectionsParams](items, "exit_clearance_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceExitClearanceSections(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.ExitClearanceChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExitClearance](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateExitClearance201JSONResponse{Body: body, Headers: apigen.CreateExitClearance201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateExitClearance implements the updateExitClearance operation.
func (s *Server) UpdateExitClearance(ctx context.Context, req apigen.UpdateExitClearanceRequestObject) (apigen.UpdateExitClearanceResponseObject, error) {
	in, err := input[db.UpdateExitClearanceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ExitClearanceId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateExitClearance(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["sections"]; ok {
		rows, err := childInputs[db.InsertExitClearanceSectionsParams](items, "exit_clearance_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceExitClearanceSections(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.ExitClearanceChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExitClearance](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateExitClearance200JSONResponse{Body: body, Headers: apigen.UpdateExitClearance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SignSectionExitClearance implements the signSectionExitClearance operation.
func (s *Server) SignSectionExitClearance(ctx context.Context, req apigen.SignSectionExitClearanceRequestObject) (apigen.SignSectionExitClearanceResponseObject, error) {
	in, err := input[db.ActSignSectionExitClearanceParams](req.Body, map[string]any{"id": req.ExitClearanceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SignSectionExitClearance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExitClearance](row)
	if err != nil {
		return nil, err
	}
	return apigen.SignSectionExitClearance200JSONResponse{Body: body, Headers: apigen.SignSectionExitClearance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteExitClearance implements the completeExitClearance operation.
func (s *Server) CompleteExitClearance(ctx context.Context, req apigen.CompleteExitClearanceRequestObject) (apigen.CompleteExitClearanceResponseObject, error) {
	in, err := input[db.ActCompleteExitClearanceParams](struct{}{}, map[string]any{"id": req.ExitClearanceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteExitClearance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExitClearance](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteExitClearance200JSONResponse{Body: body, Headers: apigen.CompleteExitClearance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRosters implements the listRosters operation.
func (s *Server) ListRosters(ctx context.Context, req apigen.ListRostersRequestObject) (apigen.ListRostersResponseObject, error) {
	return listPage[apigen.ListRosters200JSONResponse](ctx, req.Params, nil, s.service.ListRoster)
}

// GetRoster implements the getRoster operation.
func (s *Server) GetRoster(ctx context.Context, req apigen.GetRosterRequestObject) (apigen.GetRosterResponseObject, error) {
	row, err := s.service.GetRoster(ctx, req.RosterId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.RosterChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Roster](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetRoster200JSONResponse{Body: body, Headers: apigen.GetRoster200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRoster implements the createRoster operation.
func (s *Server) CreateRoster(ctx context.Context, req apigen.CreateRosterRequestObject) (apigen.CreateRosterResponseObject, error) {
	in, err := input[db.CreateRosterParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRoster(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["shifts"]; ok {
		rows, err := childInputs[db.InsertRosterShiftsParams](items, "roster_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceRosterShifts(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.RosterChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Roster](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRoster201JSONResponse{Body: body, Headers: apigen.CreateRoster201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRoster implements the updateRoster operation.
func (s *Server) UpdateRoster(ctx context.Context, req apigen.UpdateRosterRequestObject) (apigen.UpdateRosterResponseObject, error) {
	in, err := input[db.UpdateRosterParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.RosterId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRoster(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["shifts"]; ok {
		rows, err := childInputs[db.InsertRosterShiftsParams](items, "roster_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceRosterShifts(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.RosterChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Roster](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRoster200JSONResponse{Body: body, Headers: apigen.UpdateRoster200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitRoster implements the submitRoster operation.
func (s *Server) SubmitRoster(ctx context.Context, req apigen.SubmitRosterRequestObject) (apigen.SubmitRosterResponseObject, error) {
	in, err := input[db.ActSubmitRosterParams](struct{}{}, map[string]any{"id": req.RosterId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitRoster(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Roster](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitRoster200JSONResponse{Body: body, Headers: apigen.SubmitRoster200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListAttendanceRecords implements the listAttendanceRecords operation.
func (s *Server) ListAttendanceRecords(ctx context.Context, req apigen.ListAttendanceRecordsRequestObject) (apigen.ListAttendanceRecordsResponseObject, error) {
	return listPage[apigen.ListAttendanceRecords200JSONResponse](ctx, req.Params, nil, s.service.ListAttendanceRecord)
}

// GetAttendanceRecord implements the getAttendanceRecord operation.
func (s *Server) GetAttendanceRecord(ctx context.Context, req apigen.GetAttendanceRecordRequestObject) (apigen.GetAttendanceRecordResponseObject, error) {
	row, err := s.service.GetAttendanceRecord(ctx, req.AttendanceRecordId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AttendanceRecord](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetAttendanceRecord200JSONResponse{Body: body, Headers: apigen.GetAttendanceRecord200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateAttendanceRecord implements the createAttendanceRecord operation.
func (s *Server) CreateAttendanceRecord(ctx context.Context, req apigen.CreateAttendanceRecordRequestObject) (apigen.CreateAttendanceRecordResponseObject, error) {
	in, err := input[db.CreateAttendanceRecordParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateAttendanceRecord(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AttendanceRecord](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateAttendanceRecord201JSONResponse{Body: body, Headers: apigen.CreateAttendanceRecord201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateAttendanceRecord implements the updateAttendanceRecord operation.
func (s *Server) UpdateAttendanceRecord(ctx context.Context, req apigen.UpdateAttendanceRecordRequestObject) (apigen.UpdateAttendanceRecordResponseObject, error) {
	in, err := input[db.UpdateAttendanceRecordParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.AttendanceRecordId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateAttendanceRecord(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AttendanceRecord](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateAttendanceRecord200JSONResponse{Body: body, Headers: apigen.UpdateAttendanceRecord200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListUniformIssues implements the listUniformIssues operation.
func (s *Server) ListUniformIssues(ctx context.Context, req apigen.ListUniformIssuesRequestObject) (apigen.ListUniformIssuesResponseObject, error) {
	return listPage[apigen.ListUniformIssues200JSONResponse](ctx, req.Params, nil, s.service.ListUniformIssue)
}

// GetUniformIssue implements the getUniformIssue operation.
func (s *Server) GetUniformIssue(ctx context.Context, req apigen.GetUniformIssueRequestObject) (apigen.GetUniformIssueResponseObject, error) {
	row, err := s.service.GetUniformIssue(ctx, req.UniformIssueId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.UniformIssue](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetUniformIssue200JSONResponse{Body: body, Headers: apigen.GetUniformIssue200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateUniformIssue implements the createUniformIssue operation.
func (s *Server) CreateUniformIssue(ctx context.Context, req apigen.CreateUniformIssueRequestObject) (apigen.CreateUniformIssueResponseObject, error) {
	in, err := input[db.CreateUniformIssueParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateUniformIssue(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.UniformIssue](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateUniformIssue201JSONResponse{Body: body, Headers: apigen.CreateUniformIssue201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateUniformIssue implements the updateUniformIssue operation.
func (s *Server) UpdateUniformIssue(ctx context.Context, req apigen.UpdateUniformIssueRequestObject) (apigen.UpdateUniformIssueResponseObject, error) {
	in, err := input[db.UpdateUniformIssueParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.UniformIssueId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateUniformIssue(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.UniformIssue](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateUniformIssue200JSONResponse{Body: body, Headers: apigen.UpdateUniformIssue200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRecruitmentRequisitions implements the listRecruitmentRequisitions operation.
func (s *Server) ListRecruitmentRequisitions(ctx context.Context, req apigen.ListRecruitmentRequisitionsRequestObject) (apigen.ListRecruitmentRequisitionsResponseObject, error) {
	return listPage[apigen.ListRecruitmentRequisitions200JSONResponse](ctx, req.Params, nil, s.service.ListRecruitmentRequisition)
}

// GetRecruitmentRequisition implements the getRecruitmentRequisition operation.
func (s *Server) GetRecruitmentRequisition(ctx context.Context, req apigen.GetRecruitmentRequisitionRequestObject) (apigen.GetRecruitmentRequisitionResponseObject, error) {
	row, err := s.service.GetRecruitmentRequisition(ctx, req.RecruitmentRequisitionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecruitmentRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRecruitmentRequisition200JSONResponse{Body: body, Headers: apigen.GetRecruitmentRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRecruitmentRequisition implements the createRecruitmentRequisition operation.
func (s *Server) CreateRecruitmentRequisition(ctx context.Context, req apigen.CreateRecruitmentRequisitionRequestObject) (apigen.CreateRecruitmentRequisitionResponseObject, error) {
	in, err := input[db.CreateRecruitmentRequisitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRecruitmentRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecruitmentRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRecruitmentRequisition201JSONResponse{Body: body, Headers: apigen.CreateRecruitmentRequisition201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRecruitmentRequisition implements the updateRecruitmentRequisition operation.
func (s *Server) UpdateRecruitmentRequisition(ctx context.Context, req apigen.UpdateRecruitmentRequisitionRequestObject) (apigen.UpdateRecruitmentRequisitionResponseObject, error) {
	in, err := input[db.UpdateRecruitmentRequisitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.RecruitmentRequisitionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRecruitmentRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecruitmentRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRecruitmentRequisition200JSONResponse{Body: body, Headers: apigen.UpdateRecruitmentRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitRecruitmentRequisition implements the submitRecruitmentRequisition operation.
func (s *Server) SubmitRecruitmentRequisition(ctx context.Context, req apigen.SubmitRecruitmentRequisitionRequestObject) (apigen.SubmitRecruitmentRequisitionResponseObject, error) {
	in, err := input[db.ActSubmitRecruitmentRequisitionParams](struct{}{}, map[string]any{"id": req.RecruitmentRequisitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitRecruitmentRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecruitmentRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitRecruitmentRequisition200JSONResponse{Body: body, Headers: apigen.SubmitRecruitmentRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCandidates implements the listCandidates operation.
func (s *Server) ListCandidates(ctx context.Context, req apigen.ListCandidatesRequestObject) (apigen.ListCandidatesResponseObject, error) {
	return listPage[apigen.ListCandidates200JSONResponse](ctx, req.Params, nil, s.service.ListCandidate)
}

// GetCandidate implements the getCandidate operation.
func (s *Server) GetCandidate(ctx context.Context, req apigen.GetCandidateRequestObject) (apigen.GetCandidateResponseObject, error) {
	row, err := s.service.GetCandidate(ctx, req.CandidateId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Candidate](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCandidate200JSONResponse{Body: body, Headers: apigen.GetCandidate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCandidate implements the createCandidate operation.
func (s *Server) CreateCandidate(ctx context.Context, req apigen.CreateCandidateRequestObject) (apigen.CreateCandidateResponseObject, error) {
	in, err := input[db.CreateCandidateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCandidate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Candidate](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCandidate201JSONResponse{Body: body, Headers: apigen.CreateCandidate201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCandidate implements the updateCandidate operation.
func (s *Server) UpdateCandidate(ctx context.Context, req apigen.UpdateCandidateRequestObject) (apigen.UpdateCandidateResponseObject, error) {
	in, err := input[db.UpdateCandidateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CandidateId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCandidate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Candidate](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCandidate200JSONResponse{Body: body, Headers: apigen.UpdateCandidate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPerformanceReviews implements the listPerformanceReviews operation.
func (s *Server) ListPerformanceReviews(ctx context.Context, req apigen.ListPerformanceReviewsRequestObject) (apigen.ListPerformanceReviewsResponseObject, error) {
	return listPage[apigen.ListPerformanceReviews200JSONResponse](ctx, req.Params, nil, s.service.ListPerformanceReview)
}

// GetPerformanceReview implements the getPerformanceReview operation.
func (s *Server) GetPerformanceReview(ctx context.Context, req apigen.GetPerformanceReviewRequestObject) (apigen.GetPerformanceReviewResponseObject, error) {
	row, err := s.service.GetPerformanceReview(ctx, req.PerformanceReviewId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PerformanceReview](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPerformanceReview200JSONResponse{Body: body, Headers: apigen.GetPerformanceReview200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePerformanceReview implements the createPerformanceReview operation.
func (s *Server) CreatePerformanceReview(ctx context.Context, req apigen.CreatePerformanceReviewRequestObject) (apigen.CreatePerformanceReviewResponseObject, error) {
	in, err := input[db.CreatePerformanceReviewParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePerformanceReview(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PerformanceReview](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePerformanceReview201JSONResponse{Body: body, Headers: apigen.CreatePerformanceReview201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePerformanceReview implements the updatePerformanceReview operation.
func (s *Server) UpdatePerformanceReview(ctx context.Context, req apigen.UpdatePerformanceReviewRequestObject) (apigen.UpdatePerformanceReviewResponseObject, error) {
	in, err := input[db.UpdatePerformanceReviewParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PerformanceReviewId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePerformanceReview(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PerformanceReview](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePerformanceReview200JSONResponse{Body: body, Headers: apigen.UpdatePerformanceReview200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompletePerformanceReview implements the completePerformanceReview operation.
func (s *Server) CompletePerformanceReview(ctx context.Context, req apigen.CompletePerformanceReviewRequestObject) (apigen.CompletePerformanceReviewResponseObject, error) {
	in, err := input[db.ActCompletePerformanceReviewParams](struct{}{}, map[string]any{"id": req.PerformanceReviewId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompletePerformanceReview(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PerformanceReview](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompletePerformanceReview200JSONResponse{Body: body, Headers: apigen.CompletePerformanceReview200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDisciplinaryCases implements the listDisciplinaryCases operation.
func (s *Server) ListDisciplinaryCases(ctx context.Context, req apigen.ListDisciplinaryCasesRequestObject) (apigen.ListDisciplinaryCasesResponseObject, error) {
	return listPage[apigen.ListDisciplinaryCases200JSONResponse](ctx, req.Params, nil, s.service.ListDisciplinaryCase)
}

// GetDisciplinaryCase implements the getDisciplinaryCase operation.
func (s *Server) GetDisciplinaryCase(ctx context.Context, req apigen.GetDisciplinaryCaseRequestObject) (apigen.GetDisciplinaryCaseResponseObject, error) {
	row, err := s.service.GetDisciplinaryCase(ctx, req.DisciplinaryCaseId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DisciplinaryCase](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDisciplinaryCase200JSONResponse{Body: body, Headers: apigen.GetDisciplinaryCase200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDisciplinaryCase implements the createDisciplinaryCase operation.
func (s *Server) CreateDisciplinaryCase(ctx context.Context, req apigen.CreateDisciplinaryCaseRequestObject) (apigen.CreateDisciplinaryCaseResponseObject, error) {
	in, err := input[db.CreateDisciplinaryCaseParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDisciplinaryCase(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DisciplinaryCase](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDisciplinaryCase201JSONResponse{Body: body, Headers: apigen.CreateDisciplinaryCase201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDisciplinaryCase implements the updateDisciplinaryCase operation.
func (s *Server) UpdateDisciplinaryCase(ctx context.Context, req apigen.UpdateDisciplinaryCaseRequestObject) (apigen.UpdateDisciplinaryCaseResponseObject, error) {
	in, err := input[db.UpdateDisciplinaryCaseParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.DisciplinaryCaseId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDisciplinaryCase(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DisciplinaryCase](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDisciplinaryCase200JSONResponse{Body: body, Headers: apigen.UpdateDisciplinaryCase200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListGrievances implements the listGrievances operation.
func (s *Server) ListGrievances(ctx context.Context, req apigen.ListGrievancesRequestObject) (apigen.ListGrievancesResponseObject, error) {
	return listPage[apigen.ListGrievances200JSONResponse](ctx, req.Params, nil, s.service.ListGrievance)
}

// GetGrievance implements the getGrievance operation.
func (s *Server) GetGrievance(ctx context.Context, req apigen.GetGrievanceRequestObject) (apigen.GetGrievanceResponseObject, error) {
	row, err := s.service.GetGrievance(ctx, req.GrievanceId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Grievance](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetGrievance200JSONResponse{Body: body, Headers: apigen.GetGrievance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateGrievance implements the createGrievance operation.
func (s *Server) CreateGrievance(ctx context.Context, req apigen.CreateGrievanceRequestObject) (apigen.CreateGrievanceResponseObject, error) {
	in, err := input[db.CreateGrievanceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateGrievance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Grievance](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateGrievance201JSONResponse{Body: body, Headers: apigen.CreateGrievance201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateGrievance implements the updateGrievance operation.
func (s *Server) UpdateGrievance(ctx context.Context, req apigen.UpdateGrievanceRequestObject) (apigen.UpdateGrievanceResponseObject, error) {
	in, err := input[db.UpdateGrievanceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.GrievanceId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateGrievance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Grievance](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateGrievance200JSONResponse{Body: body, Headers: apigen.UpdateGrievance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
