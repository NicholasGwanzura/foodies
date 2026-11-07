// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListItems implements the listItems operation.
func (s *Server) ListItems(ctx context.Context, req apigen.ListItemsRequestObject) (apigen.ListItemsResponseObject, error) {
	return listPage[apigen.ListItems200JSONResponse](ctx, req.Params, nil, s.service.ListItem)
}

// GetItem implements the getItem operation.
func (s *Server) GetItem(ctx context.Context, req apigen.GetItemRequestObject) (apigen.GetItemResponseObject, error) {
	row, err := s.service.GetItem(ctx, req.ItemId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Item](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetItem200JSONResponse{Body: body, Headers: apigen.GetItem200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateItem implements the createItem operation.
func (s *Server) CreateItem(ctx context.Context, req apigen.CreateItemRequestObject) (apigen.CreateItemResponseObject, error) {
	in, err := input[db.CreateItemParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateItem(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Item](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateItem201JSONResponse{Body: body, Headers: apigen.CreateItem201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateItem implements the updateItem operation.
func (s *Server) UpdateItem(ctx context.Context, req apigen.UpdateItemRequestObject) (apigen.UpdateItemResponseObject, error) {
	in, err := input[db.UpdateItemParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ItemId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateItem(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Item](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateItem200JSONResponse{Body: body, Headers: apigen.UpdateItem200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListStockBalances implements the listStockBalances operation.
func (s *Server) ListStockBalances(ctx context.Context, req apigen.ListStockBalancesRequestObject) (apigen.ListStockBalancesResponseObject, error) {
	return listPage[apigen.ListStockBalances200JSONResponse](ctx, req.Params, nil, s.service.ListStockBalance)
}

// GetStockBalance implements the getStockBalance operation.
func (s *Server) GetStockBalance(ctx context.Context, req apigen.GetStockBalanceRequestObject) (apigen.GetStockBalanceResponseObject, error) {
	row, err := s.service.GetStockBalance(ctx, req.StockBalanceId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockBalance](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetStockBalance200JSONResponse{Body: body, Headers: apigen.GetStockBalance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListStockRequisitions implements the listStockRequisitions operation.
func (s *Server) ListStockRequisitions(ctx context.Context, req apigen.ListStockRequisitionsRequestObject) (apigen.ListStockRequisitionsResponseObject, error) {
	return listPage[apigen.ListStockRequisitions200JSONResponse](ctx, req.Params, nil, s.service.ListStockRequisition)
}

// GetStockRequisition implements the getStockRequisition operation.
func (s *Server) GetStockRequisition(ctx context.Context, req apigen.GetStockRequisitionRequestObject) (apigen.GetStockRequisitionResponseObject, error) {
	row, err := s.service.GetStockRequisition(ctx, req.StockRequisitionId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.StockRequisitionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockRequisition](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetStockRequisition200JSONResponse{Body: body, Headers: apigen.GetStockRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateStockRequisition implements the createStockRequisition operation.
func (s *Server) CreateStockRequisition(ctx context.Context, req apigen.CreateStockRequisitionRequestObject) (apigen.CreateStockRequisitionResponseObject, error) {
	in, err := input[db.CreateStockRequisitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateStockRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertStockRequisitionLinesParams](items, "stock_requisition_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceStockRequisitionLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.StockRequisitionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockRequisition](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateStockRequisition201JSONResponse{Body: body, Headers: apigen.CreateStockRequisition201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateStockRequisition implements the updateStockRequisition operation.
func (s *Server) UpdateStockRequisition(ctx context.Context, req apigen.UpdateStockRequisitionRequestObject) (apigen.UpdateStockRequisitionResponseObject, error) {
	in, err := input[db.UpdateStockRequisitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.StockRequisitionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateStockRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertStockRequisitionLinesParams](items, "stock_requisition_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceStockRequisitionLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.StockRequisitionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockRequisition](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateStockRequisition200JSONResponse{Body: body, Headers: apigen.UpdateStockRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitStockRequisition implements the submitStockRequisition operation.
func (s *Server) SubmitStockRequisition(ctx context.Context, req apigen.SubmitStockRequisitionRequestObject) (apigen.SubmitStockRequisitionResponseObject, error) {
	in, err := input[db.ActSubmitStockRequisitionParams](struct{}{}, map[string]any{"id": req.StockRequisitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitStockRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitStockRequisition200JSONResponse{Body: body, Headers: apigen.SubmitStockRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// IssueStockRequisition implements the issueStockRequisition operation.
func (s *Server) IssueStockRequisition(ctx context.Context, req apigen.IssueStockRequisitionRequestObject) (apigen.IssueStockRequisitionResponseObject, error) {
	in, err := input[db.ActIssueStockRequisitionParams](struct{}{}, map[string]any{"id": req.StockRequisitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.IssueStockRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.IssueStockRequisition200JSONResponse{Body: body, Headers: apigen.IssueStockRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AcknowledgeStockRequisition implements the acknowledgeStockRequisition operation.
func (s *Server) AcknowledgeStockRequisition(ctx context.Context, req apigen.AcknowledgeStockRequisitionRequestObject) (apigen.AcknowledgeStockRequisitionResponseObject, error) {
	in, err := input[db.ActAcknowledgeStockRequisitionParams](struct{}{}, map[string]any{"id": req.StockRequisitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AcknowledgeStockRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.AcknowledgeStockRequisition200JSONResponse{Body: body, Headers: apigen.AcknowledgeStockRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListStockTransfers implements the listStockTransfers operation.
func (s *Server) ListStockTransfers(ctx context.Context, req apigen.ListStockTransfersRequestObject) (apigen.ListStockTransfersResponseObject, error) {
	return listPage[apigen.ListStockTransfers200JSONResponse](ctx, req.Params, nil, s.service.ListStockTransfer)
}

// GetStockTransfer implements the getStockTransfer operation.
func (s *Server) GetStockTransfer(ctx context.Context, req apigen.GetStockTransferRequestObject) (apigen.GetStockTransferResponseObject, error) {
	row, err := s.service.GetStockTransfer(ctx, req.StockTransferId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.StockTransferChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockTransfer](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetStockTransfer200JSONResponse{Body: body, Headers: apigen.GetStockTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateStockTransfer implements the createStockTransfer operation.
func (s *Server) CreateStockTransfer(ctx context.Context, req apigen.CreateStockTransferRequestObject) (apigen.CreateStockTransferResponseObject, error) {
	in, err := input[db.CreateStockTransferParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateStockTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertStockTransferLinesParams](items, "stock_transfer_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceStockTransferLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.StockTransferChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockTransfer](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateStockTransfer201JSONResponse{Body: body, Headers: apigen.CreateStockTransfer201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateStockTransfer implements the updateStockTransfer operation.
func (s *Server) UpdateStockTransfer(ctx context.Context, req apigen.UpdateStockTransferRequestObject) (apigen.UpdateStockTransferResponseObject, error) {
	in, err := input[db.UpdateStockTransferParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.StockTransferId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateStockTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertStockTransferLinesParams](items, "stock_transfer_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceStockTransferLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.StockTransferChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockTransfer](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateStockTransfer200JSONResponse{Body: body, Headers: apigen.UpdateStockTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AcceptStockTransfer implements the acceptStockTransfer operation.
func (s *Server) AcceptStockTransfer(ctx context.Context, req apigen.AcceptStockTransferRequestObject) (apigen.AcceptStockTransferResponseObject, error) {
	in, err := input[db.ActAcceptStockTransferParams](req.Body, map[string]any{"id": req.StockTransferId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AcceptStockTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.AcceptStockTransfer200JSONResponse{Body: body, Headers: apigen.AcceptStockTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RejectStockTransfer implements the rejectStockTransfer operation.
func (s *Server) RejectStockTransfer(ctx context.Context, req apigen.RejectStockTransferRequestObject) (apigen.RejectStockTransferResponseObject, error) {
	in, err := input[db.ActRejectStockTransferParams](req.Body, map[string]any{"id": req.StockTransferId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RejectStockTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.RejectStockTransfer200JSONResponse{Body: body, Headers: apigen.RejectStockTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DispatchStockTransfer implements the dispatchStockTransfer operation.
func (s *Server) DispatchStockTransfer(ctx context.Context, req apigen.DispatchStockTransferRequestObject) (apigen.DispatchStockTransferResponseObject, error) {
	in, err := input[db.ActDispatchStockTransferParams](struct{}{}, map[string]any{"id": req.StockTransferId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.DispatchStockTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.DispatchStockTransfer200JSONResponse{Body: body, Headers: apigen.DispatchStockTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReceiveStockTransfer implements the receiveStockTransfer operation.
func (s *Server) ReceiveStockTransfer(ctx context.Context, req apigen.ReceiveStockTransferRequestObject) (apigen.ReceiveStockTransferResponseObject, error) {
	in, err := input[db.ActReceiveStockTransferParams](struct{}{}, map[string]any{"id": req.StockTransferId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReceiveStockTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReceiveStockTransfer200JSONResponse{Body: body, Headers: apigen.ReceiveStockTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseStockTransfer implements the closeStockTransfer operation.
func (s *Server) CloseStockTransfer(ctx context.Context, req apigen.CloseStockTransferRequestObject) (apigen.CloseStockTransferResponseObject, error) {
	in, err := input[db.ActCloseStockTransferParams](struct{}{}, map[string]any{"id": req.StockTransferId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseStockTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseStockTransfer200JSONResponse{Body: body, Headers: apigen.CloseStockTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListStockReturns implements the listStockReturns operation.
func (s *Server) ListStockReturns(ctx context.Context, req apigen.ListStockReturnsRequestObject) (apigen.ListStockReturnsResponseObject, error) {
	return listPage[apigen.ListStockReturns200JSONResponse](ctx, req.Params, nil, s.service.ListStockReturn)
}

// GetStockReturn implements the getStockReturn operation.
func (s *Server) GetStockReturn(ctx context.Context, req apigen.GetStockReturnRequestObject) (apigen.GetStockReturnResponseObject, error) {
	row, err := s.service.GetStockReturn(ctx, req.StockReturnId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockReturn](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetStockReturn200JSONResponse{Body: body, Headers: apigen.GetStockReturn200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateStockReturn implements the createStockReturn operation.
func (s *Server) CreateStockReturn(ctx context.Context, req apigen.CreateStockReturnRequestObject) (apigen.CreateStockReturnResponseObject, error) {
	in, err := input[db.CreateStockReturnParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateStockReturn(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockReturn](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateStockReturn201JSONResponse{Body: body, Headers: apigen.CreateStockReturn201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateStockReturn implements the updateStockReturn operation.
func (s *Server) UpdateStockReturn(ctx context.Context, req apigen.UpdateStockReturnRequestObject) (apigen.UpdateStockReturnResponseObject, error) {
	in, err := input[db.UpdateStockReturnParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.StockReturnId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateStockReturn(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockReturn](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateStockReturn200JSONResponse{Body: body, Headers: apigen.UpdateStockReturn200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseStockReturn implements the closeStockReturn operation.
func (s *Server) CloseStockReturn(ctx context.Context, req apigen.CloseStockReturnRequestObject) (apigen.CloseStockReturnResponseObject, error) {
	in, err := input[db.ActCloseStockReturnParams](struct{}{}, map[string]any{"id": req.StockReturnId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseStockReturn(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockReturn](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseStockReturn200JSONResponse{Body: body, Headers: apigen.CloseStockReturn200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListStocktakes implements the listStocktakes operation.
func (s *Server) ListStocktakes(ctx context.Context, req apigen.ListStocktakesRequestObject) (apigen.ListStocktakesResponseObject, error) {
	return listPage[apigen.ListStocktakes200JSONResponse](ctx, req.Params, nil, s.service.ListStocktake)
}

// GetStocktake implements the getStocktake operation.
func (s *Server) GetStocktake(ctx context.Context, req apigen.GetStocktakeRequestObject) (apigen.GetStocktakeResponseObject, error) {
	row, err := s.service.GetStocktake(ctx, req.StocktakeId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.StocktakeChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Stocktake](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetStocktake200JSONResponse{Body: body, Headers: apigen.GetStocktake200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateStocktake implements the createStocktake operation.
func (s *Server) CreateStocktake(ctx context.Context, req apigen.CreateStocktakeRequestObject) (apigen.CreateStocktakeResponseObject, error) {
	in, err := input[db.CreateStocktakeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateStocktake(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertStocktakeLinesParams](items, "stocktake_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceStocktakeLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.StocktakeChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Stocktake](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateStocktake201JSONResponse{Body: body, Headers: apigen.CreateStocktake201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateStocktake implements the updateStocktake operation.
func (s *Server) UpdateStocktake(ctx context.Context, req apigen.UpdateStocktakeRequestObject) (apigen.UpdateStocktakeResponseObject, error) {
	in, err := input[db.UpdateStocktakeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.StocktakeId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateStocktake(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertStocktakeLinesParams](items, "stocktake_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceStocktakeLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.StocktakeChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Stocktake](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateStocktake200JSONResponse{Body: body, Headers: apigen.UpdateStocktake200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitStocktake implements the submitStocktake operation.
func (s *Server) SubmitStocktake(ctx context.Context, req apigen.SubmitStocktakeRequestObject) (apigen.SubmitStocktakeResponseObject, error) {
	in, err := input[db.ActSubmitStocktakeParams](struct{}{}, map[string]any{"id": req.StocktakeId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitStocktake(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Stocktake](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitStocktake200JSONResponse{Body: body, Headers: apigen.SubmitStocktake200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ApproveStocktake implements the approveStocktake operation.
func (s *Server) ApproveStocktake(ctx context.Context, req apigen.ApproveStocktakeRequestObject) (apigen.ApproveStocktakeResponseObject, error) {
	in, err := input[db.ActApproveStocktakeParams](req.Body, map[string]any{"id": req.StocktakeId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ApproveStocktake(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Stocktake](row)
	if err != nil {
		return nil, err
	}
	return apigen.ApproveStocktake200JSONResponse{Body: body, Headers: apigen.ApproveStocktake200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RequestSkipStocktake implements the requestSkipStocktake operation.
func (s *Server) RequestSkipStocktake(ctx context.Context, req apigen.RequestSkipStocktakeRequestObject) (apigen.RequestSkipStocktakeResponseObject, error) {
	in, err := input[db.ActRequestSkipStocktakeParams](req.Body, map[string]any{"id": req.StocktakeId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RequestSkipStocktake(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Stocktake](row)
	if err != nil {
		return nil, err
	}
	return apigen.RequestSkipStocktake200JSONResponse{Body: body, Headers: apigen.RequestSkipStocktake200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListStockVariances implements the listStockVariances operation.
func (s *Server) ListStockVariances(ctx context.Context, req apigen.ListStockVariancesRequestObject) (apigen.ListStockVariancesResponseObject, error) {
	return listPage[apigen.ListStockVariances200JSONResponse](ctx, req.Params, nil, s.service.ListStockVariance)
}

// GetStockVariance implements the getStockVariance operation.
func (s *Server) GetStockVariance(ctx context.Context, req apigen.GetStockVarianceRequestObject) (apigen.GetStockVarianceResponseObject, error) {
	row, err := s.service.GetStockVariance(ctx, req.StockVarianceId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetStockVariance200JSONResponse{Body: body, Headers: apigen.GetStockVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateStockVariance implements the updateStockVariance operation.
func (s *Server) UpdateStockVariance(ctx context.Context, req apigen.UpdateStockVarianceRequestObject) (apigen.UpdateStockVarianceResponseObject, error) {
	in, err := input[db.UpdateStockVarianceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.StockVarianceId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateStockVariance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateStockVariance200JSONResponse{Body: body, Headers: apigen.UpdateStockVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseStockVariance implements the closeStockVariance operation.
func (s *Server) CloseStockVariance(ctx context.Context, req apigen.CloseStockVarianceRequestObject) (apigen.CloseStockVarianceResponseObject, error) {
	in, err := input[db.ActCloseStockVarianceParams](req.Body, map[string]any{"id": req.StockVarianceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseStockVariance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseStockVariance200JSONResponse{Body: body, Headers: apigen.CloseStockVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListStockAdjustments implements the listStockAdjustments operation.
func (s *Server) ListStockAdjustments(ctx context.Context, req apigen.ListStockAdjustmentsRequestObject) (apigen.ListStockAdjustmentsResponseObject, error) {
	return listPage[apigen.ListStockAdjustments200JSONResponse](ctx, req.Params, nil, s.service.ListStockAdjustment)
}

// GetStockAdjustment implements the getStockAdjustment operation.
func (s *Server) GetStockAdjustment(ctx context.Context, req apigen.GetStockAdjustmentRequestObject) (apigen.GetStockAdjustmentResponseObject, error) {
	row, err := s.service.GetStockAdjustment(ctx, req.StockAdjustmentId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockAdjustment](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetStockAdjustment200JSONResponse{Body: body, Headers: apigen.GetStockAdjustment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateStockAdjustment implements the createStockAdjustment operation.
func (s *Server) CreateStockAdjustment(ctx context.Context, req apigen.CreateStockAdjustmentRequestObject) (apigen.CreateStockAdjustmentResponseObject, error) {
	in, err := input[db.CreateStockAdjustmentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateStockAdjustment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockAdjustment](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateStockAdjustment201JSONResponse{Body: body, Headers: apigen.CreateStockAdjustment201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateStockAdjustment implements the updateStockAdjustment operation.
func (s *Server) UpdateStockAdjustment(ctx context.Context, req apigen.UpdateStockAdjustmentRequestObject) (apigen.UpdateStockAdjustmentResponseObject, error) {
	in, err := input[db.UpdateStockAdjustmentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.StockAdjustmentId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateStockAdjustment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockAdjustment](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateStockAdjustment200JSONResponse{Body: body, Headers: apigen.UpdateStockAdjustment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitStockAdjustment implements the submitStockAdjustment operation.
func (s *Server) SubmitStockAdjustment(ctx context.Context, req apigen.SubmitStockAdjustmentRequestObject) (apigen.SubmitStockAdjustmentResponseObject, error) {
	in, err := input[db.ActSubmitStockAdjustmentParams](struct{}{}, map[string]any{"id": req.StockAdjustmentId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitStockAdjustment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StockAdjustment](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitStockAdjustment200JSONResponse{Body: body, Headers: apigen.SubmitStockAdjustment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListWasteRecords implements the listWasteRecords operation.
func (s *Server) ListWasteRecords(ctx context.Context, req apigen.ListWasteRecordsRequestObject) (apigen.ListWasteRecordsResponseObject, error) {
	return listPage[apigen.ListWasteRecords200JSONResponse](ctx, req.Params, nil, s.service.ListWasteRecord)
}

// GetWasteRecord implements the getWasteRecord operation.
func (s *Server) GetWasteRecord(ctx context.Context, req apigen.GetWasteRecordRequestObject) (apigen.GetWasteRecordResponseObject, error) {
	row, err := s.service.GetWasteRecord(ctx, req.WasteRecordId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WasteRecord](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetWasteRecord200JSONResponse{Body: body, Headers: apigen.GetWasteRecord200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateWasteRecord implements the createWasteRecord operation.
func (s *Server) CreateWasteRecord(ctx context.Context, req apigen.CreateWasteRecordRequestObject) (apigen.CreateWasteRecordResponseObject, error) {
	in, err := input[db.CreateWasteRecordParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateWasteRecord(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WasteRecord](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateWasteRecord201JSONResponse{Body: body, Headers: apigen.CreateWasteRecord201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateWasteRecord implements the updateWasteRecord operation.
func (s *Server) UpdateWasteRecord(ctx context.Context, req apigen.UpdateWasteRecordRequestObject) (apigen.UpdateWasteRecordResponseObject, error) {
	in, err := input[db.UpdateWasteRecordParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.WasteRecordId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateWasteRecord(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WasteRecord](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateWasteRecord200JSONResponse{Body: body, Headers: apigen.UpdateWasteRecord200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReviewWasteRecord implements the reviewWasteRecord operation.
func (s *Server) ReviewWasteRecord(ctx context.Context, req apigen.ReviewWasteRecordRequestObject) (apigen.ReviewWasteRecordResponseObject, error) {
	in, err := input[db.ActReviewWasteRecordParams](req.Body, map[string]any{"id": req.WasteRecordId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReviewWasteRecord(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WasteRecord](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReviewWasteRecord200JSONResponse{Body: body, Headers: apigen.ReviewWasteRecord200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDisposals implements the listDisposals operation.
func (s *Server) ListDisposals(ctx context.Context, req apigen.ListDisposalsRequestObject) (apigen.ListDisposalsResponseObject, error) {
	return listPage[apigen.ListDisposals200JSONResponse](ctx, req.Params, nil, s.service.ListDisposal)
}

// GetDisposal implements the getDisposal operation.
func (s *Server) GetDisposal(ctx context.Context, req apigen.GetDisposalRequestObject) (apigen.GetDisposalResponseObject, error) {
	row, err := s.service.GetDisposal(ctx, req.DisposalId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Disposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDisposal200JSONResponse{Body: body, Headers: apigen.GetDisposal200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDisposal implements the createDisposal operation.
func (s *Server) CreateDisposal(ctx context.Context, req apigen.CreateDisposalRequestObject) (apigen.CreateDisposalResponseObject, error) {
	in, err := input[db.CreateDisposalParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDisposal(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Disposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDisposal201JSONResponse{Body: body, Headers: apigen.CreateDisposal201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDisposal implements the updateDisposal operation.
func (s *Server) UpdateDisposal(ctx context.Context, req apigen.UpdateDisposalRequestObject) (apigen.UpdateDisposalResponseObject, error) {
	in, err := input[db.UpdateDisposalParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.DisposalId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDisposal(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Disposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDisposal200JSONResponse{Body: body, Headers: apigen.UpdateDisposal200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitDisposal implements the submitDisposal operation.
func (s *Server) SubmitDisposal(ctx context.Context, req apigen.SubmitDisposalRequestObject) (apigen.SubmitDisposalResponseObject, error) {
	in, err := input[db.ActSubmitDisposalParams](struct{}{}, map[string]any{"id": req.DisposalId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitDisposal(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Disposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitDisposal200JSONResponse{Body: body, Headers: apigen.SubmitDisposal200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RecordDisposalDisposal implements the recordDisposalDisposal operation.
func (s *Server) RecordDisposalDisposal(ctx context.Context, req apigen.RecordDisposalDisposalRequestObject) (apigen.RecordDisposalDisposalResponseObject, error) {
	in, err := input[db.ActRecordDisposalDisposalParams](req.Body, map[string]any{"id": req.DisposalId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RecordDisposalDisposal(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Disposal](row)
	if err != nil {
		return nil, err
	}
	return apigen.RecordDisposalDisposal200JSONResponse{Body: body, Headers: apigen.RecordDisposalDisposal200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRecipes implements the listRecipes operation.
func (s *Server) ListRecipes(ctx context.Context, req apigen.ListRecipesRequestObject) (apigen.ListRecipesResponseObject, error) {
	return listPage[apigen.ListRecipes200JSONResponse](ctx, req.Params, nil, s.service.ListRecipe)
}

// GetRecipe implements the getRecipe operation.
func (s *Server) GetRecipe(ctx context.Context, req apigen.GetRecipeRequestObject) (apigen.GetRecipeResponseObject, error) {
	row, err := s.service.GetRecipe(ctx, req.RecipeId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Recipe](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRecipe200JSONResponse{Body: body, Headers: apigen.GetRecipe200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRecipe implements the createRecipe operation.
func (s *Server) CreateRecipe(ctx context.Context, req apigen.CreateRecipeRequestObject) (apigen.CreateRecipeResponseObject, error) {
	in, err := input[db.CreateRecipeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRecipe(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Recipe](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRecipe201JSONResponse{Body: body, Headers: apigen.CreateRecipe201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRecipe implements the updateRecipe operation.
func (s *Server) UpdateRecipe(ctx context.Context, req apigen.UpdateRecipeRequestObject) (apigen.UpdateRecipeResponseObject, error) {
	in, err := input[db.UpdateRecipeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.RecipeId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRecipe(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Recipe](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRecipe200JSONResponse{Body: body, Headers: apigen.UpdateRecipe200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRecipeVersions implements the listRecipeVersions operation.
func (s *Server) ListRecipeVersions(ctx context.Context, req apigen.ListRecipeVersionsRequestObject) (apigen.ListRecipeVersionsResponseObject, error) {
	return listPage[apigen.ListRecipeVersions200JSONResponse](ctx, req.Params, map[string]any{"recipe_id": req.RecipeId}, s.service.ListRecipeVersion)
}

// GetRecipeVersion implements the getRecipeVersion operation.
func (s *Server) GetRecipeVersion(ctx context.Context, req apigen.GetRecipeVersionRequestObject) (apigen.GetRecipeVersionResponseObject, error) {
	row, err := s.service.GetRecipeVersion(ctx, req.RecipeVersionId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.RecipeVersionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecipeVersion](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetRecipeVersion200JSONResponse{Body: body, Headers: apigen.GetRecipeVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRecipeVersion implements the createRecipeVersion operation.
func (s *Server) CreateRecipeVersion(ctx context.Context, req apigen.CreateRecipeVersionRequestObject) (apigen.CreateRecipeVersionResponseObject, error) {
	in, err := input[db.CreateRecipeVersionParams](req.Body, map[string]any{"recipe_id": req.RecipeId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRecipeVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["ingredients"]; ok {
		rows, err := childInputs[db.InsertRecipeIngredientsParams](items, "recipe_version_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceRecipeIngredients(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.RecipeVersionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecipeVersion](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRecipeVersion201JSONResponse{Body: body, Headers: apigen.CreateRecipeVersion201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRecipeVersion implements the updateRecipeVersion operation.
func (s *Server) UpdateRecipeVersion(ctx context.Context, req apigen.UpdateRecipeVersionRequestObject) (apigen.UpdateRecipeVersionResponseObject, error) {
	in, err := input[db.UpdateRecipeVersionParams](req.Body, map[string]any{"recipe_id": req.RecipeId})
	if err != nil {
		return nil, err
	}
	in.ID = req.RecipeVersionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRecipeVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["ingredients"]; ok {
		rows, err := childInputs[db.InsertRecipeIngredientsParams](items, "recipe_version_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceRecipeIngredients(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.RecipeVersionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecipeVersion](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRecipeVersion200JSONResponse{Body: body, Headers: apigen.UpdateRecipeVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitRecipeVersion implements the submitRecipeVersion operation.
func (s *Server) SubmitRecipeVersion(ctx context.Context, req apigen.SubmitRecipeVersionRequestObject) (apigen.SubmitRecipeVersionResponseObject, error) {
	in, err := input[db.ActSubmitRecipeVersionParams](struct{}{}, map[string]any{"id": req.RecipeVersionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitRecipeVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecipeVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitRecipeVersion200JSONResponse{Body: body, Headers: apigen.SubmitRecipeVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListKitchenOrders implements the listKitchenOrders operation.
func (s *Server) ListKitchenOrders(ctx context.Context, req apigen.ListKitchenOrdersRequestObject) (apigen.ListKitchenOrdersResponseObject, error) {
	return listPage[apigen.ListKitchenOrders200JSONResponse](ctx, req.Params, nil, s.service.ListKitchenOrder)
}

// GetKitchenOrder implements the getKitchenOrder operation.
func (s *Server) GetKitchenOrder(ctx context.Context, req apigen.GetKitchenOrderRequestObject) (apigen.GetKitchenOrderResponseObject, error) {
	row, err := s.service.GetKitchenOrder(ctx, req.KitchenOrderId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.KitchenOrderChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KitchenOrder](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetKitchenOrder200JSONResponse{Body: body, Headers: apigen.GetKitchenOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateKitchenOrder implements the createKitchenOrder operation.
func (s *Server) CreateKitchenOrder(ctx context.Context, req apigen.CreateKitchenOrderRequestObject) (apigen.CreateKitchenOrderResponseObject, error) {
	in, err := input[db.CreateKitchenOrderParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateKitchenOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertKitchenOrderLinesParams](items, "kitchen_order_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceKitchenOrderLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.KitchenOrderChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KitchenOrder](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateKitchenOrder201JSONResponse{Body: body, Headers: apigen.CreateKitchenOrder201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateKitchenOrder implements the updateKitchenOrder operation.
func (s *Server) UpdateKitchenOrder(ctx context.Context, req apigen.UpdateKitchenOrderRequestObject) (apigen.UpdateKitchenOrderResponseObject, error) {
	in, err := input[db.UpdateKitchenOrderParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.KitchenOrderId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateKitchenOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertKitchenOrderLinesParams](items, "kitchen_order_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceKitchenOrderLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.KitchenOrderChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KitchenOrder](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateKitchenOrder200JSONResponse{Body: body, Headers: apigen.UpdateKitchenOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitKitchenOrder implements the submitKitchenOrder operation.
func (s *Server) SubmitKitchenOrder(ctx context.Context, req apigen.SubmitKitchenOrderRequestObject) (apigen.SubmitKitchenOrderResponseObject, error) {
	in, err := input[db.ActSubmitKitchenOrderParams](struct{}{}, map[string]any{"id": req.KitchenOrderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitKitchenOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KitchenOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitKitchenOrder200JSONResponse{Body: body, Headers: apigen.SubmitKitchenOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AcknowledgeLateKitchenOrder implements the acknowledgeLateKitchenOrder operation.
func (s *Server) AcknowledgeLateKitchenOrder(ctx context.Context, req apigen.AcknowledgeLateKitchenOrderRequestObject) (apigen.AcknowledgeLateKitchenOrderResponseObject, error) {
	in, err := input[db.ActAcknowledgeLateKitchenOrderParams](req.Body, map[string]any{"id": req.KitchenOrderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AcknowledgeLateKitchenOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KitchenOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.AcknowledgeLateKitchenOrder200JSONResponse{Body: body, Headers: apigen.AcknowledgeLateKitchenOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ConfirmKitchenOrder implements the confirmKitchenOrder operation.
func (s *Server) ConfirmKitchenOrder(ctx context.Context, req apigen.ConfirmKitchenOrderRequestObject) (apigen.ConfirmKitchenOrderResponseObject, error) {
	in, err := input[db.ActConfirmKitchenOrderParams](struct{}{}, map[string]any{"id": req.KitchenOrderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ConfirmKitchenOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KitchenOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.ConfirmKitchenOrder200JSONResponse{Body: body, Headers: apigen.ConfirmKitchenOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListProductionPlans implements the listProductionPlans operation.
func (s *Server) ListProductionPlans(ctx context.Context, req apigen.ListProductionPlansRequestObject) (apigen.ListProductionPlansResponseObject, error) {
	return listPage[apigen.ListProductionPlans200JSONResponse](ctx, req.Params, nil, s.service.ListProductionPlan)
}

// GetProductionPlan implements the getProductionPlan operation.
func (s *Server) GetProductionPlan(ctx context.Context, req apigen.GetProductionPlanRequestObject) (apigen.GetProductionPlanResponseObject, error) {
	row, err := s.service.GetProductionPlan(ctx, req.ProductionPlanId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.ProductionPlanChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductionPlan](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetProductionPlan200JSONResponse{Body: body, Headers: apigen.GetProductionPlan200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateProductionPlan implements the createProductionPlan operation.
func (s *Server) CreateProductionPlan(ctx context.Context, req apigen.CreateProductionPlanRequestObject) (apigen.CreateProductionPlanResponseObject, error) {
	in, err := input[db.CreateProductionPlanParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateProductionPlan(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertProductionPlanLinesParams](items, "production_plan_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceProductionPlanLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.ProductionPlanChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductionPlan](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateProductionPlan201JSONResponse{Body: body, Headers: apigen.CreateProductionPlan201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateProductionPlan implements the updateProductionPlan operation.
func (s *Server) UpdateProductionPlan(ctx context.Context, req apigen.UpdateProductionPlanRequestObject) (apigen.UpdateProductionPlanResponseObject, error) {
	in, err := input[db.UpdateProductionPlanParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ProductionPlanId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateProductionPlan(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertProductionPlanLinesParams](items, "production_plan_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceProductionPlanLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.ProductionPlanChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductionPlan](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateProductionPlan200JSONResponse{Body: body, Headers: apigen.UpdateProductionPlan200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ConfirmProductionPlan implements the confirmProductionPlan operation.
func (s *Server) ConfirmProductionPlan(ctx context.Context, req apigen.ConfirmProductionPlanRequestObject) (apigen.ConfirmProductionPlanResponseObject, error) {
	in, err := input[db.ActConfirmProductionPlanParams](struct{}{}, map[string]any{"id": req.ProductionPlanId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ConfirmProductionPlan(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductionPlan](row)
	if err != nil {
		return nil, err
	}
	return apigen.ConfirmProductionPlan200JSONResponse{Body: body, Headers: apigen.ConfirmProductionPlan200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListBatches implements the listBatches operation.
func (s *Server) ListBatches(ctx context.Context, req apigen.ListBatchesRequestObject) (apigen.ListBatchesResponseObject, error) {
	return listPage[apigen.ListBatches200JSONResponse](ctx, req.Params, nil, s.service.ListBatch)
}

// GetBatch implements the getBatch operation.
func (s *Server) GetBatch(ctx context.Context, req apigen.GetBatchRequestObject) (apigen.GetBatchResponseObject, error) {
	row, err := s.service.GetBatch(ctx, req.BatchId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.BatchChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Batch](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetBatch200JSONResponse{Body: body, Headers: apigen.GetBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateBatch implements the createBatch operation.
func (s *Server) CreateBatch(ctx context.Context, req apigen.CreateBatchRequestObject) (apigen.CreateBatchResponseObject, error) {
	in, err := input[db.CreateBatchParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["ingredients"]; ok {
		rows, err := childInputs[db.InsertBatchIngredientsParams](items, "batch_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceBatchIngredients(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.BatchChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Batch](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateBatch201JSONResponse{Body: body, Headers: apigen.CreateBatch201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateBatch implements the updateBatch operation.
func (s *Server) UpdateBatch(ctx context.Context, req apigen.UpdateBatchRequestObject) (apigen.UpdateBatchResponseObject, error) {
	in, err := input[db.UpdateBatchParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.BatchId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["ingredients"]; ok {
		rows, err := childInputs[db.InsertBatchIngredientsParams](items, "batch_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceBatchIngredients(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.BatchChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Batch](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateBatch200JSONResponse{Body: body, Headers: apigen.UpdateBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteBatch implements the completeBatch operation.
func (s *Server) CompleteBatch(ctx context.Context, req apigen.CompleteBatchRequestObject) (apigen.CompleteBatchResponseObject, error) {
	in, err := input[db.ActCompleteBatchParams](struct{}{}, map[string]any{"id": req.BatchId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Batch](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteBatch200JSONResponse{Body: body, Headers: apigen.CompleteBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReleaseBatch implements the releaseBatch operation.
func (s *Server) ReleaseBatch(ctx context.Context, req apigen.ReleaseBatchRequestObject) (apigen.ReleaseBatchResponseObject, error) {
	in, err := input[db.ActReleaseBatchParams](req.Body, map[string]any{"id": req.BatchId})
	if err != nil {
		return nil, err
	}
	if err := s.service.RequireSigned(ctx, "batches", req.BatchId, &req.Body.SignatureId); err != nil {
		return nil, err
	}
	row, err := s.service.ReleaseBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Batch](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReleaseBatch200JSONResponse{Body: body, Headers: apigen.ReleaseBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDispatchNotes implements the listDispatchNotes operation.
func (s *Server) ListDispatchNotes(ctx context.Context, req apigen.ListDispatchNotesRequestObject) (apigen.ListDispatchNotesResponseObject, error) {
	return listPage[apigen.ListDispatchNotes200JSONResponse](ctx, req.Params, nil, s.service.ListDispatchNote)
}

// GetDispatchNote implements the getDispatchNote operation.
func (s *Server) GetDispatchNote(ctx context.Context, req apigen.GetDispatchNoteRequestObject) (apigen.GetDispatchNoteResponseObject, error) {
	row, err := s.service.GetDispatchNote(ctx, req.DispatchNoteId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.DispatchNoteChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispatchNote](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetDispatchNote200JSONResponse{Body: body, Headers: apigen.GetDispatchNote200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDispatchNote implements the createDispatchNote operation.
func (s *Server) CreateDispatchNote(ctx context.Context, req apigen.CreateDispatchNoteRequestObject) (apigen.CreateDispatchNoteResponseObject, error) {
	in, err := input[db.CreateDispatchNoteParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDispatchNote(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertDispatchNoteLinesParams](items, "dispatch_note_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceDispatchNoteLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.DispatchNoteChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispatchNote](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDispatchNote201JSONResponse{Body: body, Headers: apigen.CreateDispatchNote201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDispatchNote implements the updateDispatchNote operation.
func (s *Server) UpdateDispatchNote(ctx context.Context, req apigen.UpdateDispatchNoteRequestObject) (apigen.UpdateDispatchNoteResponseObject, error) {
	in, err := input[db.UpdateDispatchNoteParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.DispatchNoteId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDispatchNote(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertDispatchNoteLinesParams](items, "dispatch_note_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceDispatchNoteLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.DispatchNoteChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispatchNote](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDispatchNote200JSONResponse{Body: body, Headers: apigen.UpdateDispatchNote200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// VerifyDispatchNote implements the verifyDispatchNote operation.
func (s *Server) VerifyDispatchNote(ctx context.Context, req apigen.VerifyDispatchNoteRequestObject) (apigen.VerifyDispatchNoteResponseObject, error) {
	in, err := input[db.ActVerifyDispatchNoteParams](struct{}{}, map[string]any{"id": req.DispatchNoteId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.VerifyDispatchNote(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispatchNote](row)
	if err != nil {
		return nil, err
	}
	return apigen.VerifyDispatchNote200JSONResponse{Body: body, Headers: apigen.VerifyDispatchNote200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SecurityClearDispatchNote implements the securityClearDispatchNote operation.
func (s *Server) SecurityClearDispatchNote(ctx context.Context, req apigen.SecurityClearDispatchNoteRequestObject) (apigen.SecurityClearDispatchNoteResponseObject, error) {
	in, err := input[db.ActSecurityClearDispatchNoteParams](req.Body, map[string]any{"id": req.DispatchNoteId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SecurityClearDispatchNote(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispatchNote](row)
	if err != nil {
		return nil, err
	}
	return apigen.SecurityClearDispatchNote200JSONResponse{Body: body, Headers: apigen.SecurityClearDispatchNote200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ConfirmDeliveryDispatchNote implements the confirmDeliveryDispatchNote operation.
func (s *Server) ConfirmDeliveryDispatchNote(ctx context.Context, req apigen.ConfirmDeliveryDispatchNoteRequestObject) (apigen.ConfirmDeliveryDispatchNoteResponseObject, error) {
	in, err := input[db.ActConfirmDeliveryDispatchNoteParams](struct{}{}, map[string]any{"id": req.DispatchNoteId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ConfirmDeliveryDispatchNote(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispatchNote](row)
	if err != nil {
		return nil, err
	}
	return apigen.ConfirmDeliveryDispatchNote200JSONResponse{Body: body, Headers: apigen.ConfirmDeliveryDispatchNote200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RaiseDiscrepancyDispatchNote implements the raiseDiscrepancyDispatchNote operation.
func (s *Server) RaiseDiscrepancyDispatchNote(ctx context.Context, req apigen.RaiseDiscrepancyDispatchNoteRequestObject) (apigen.RaiseDiscrepancyDispatchNoteResponseObject, error) {
	in, err := input[db.ActRaiseDiscrepancyDispatchNoteParams](req.Body, map[string]any{"id": req.DispatchNoteId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RaiseDiscrepancyDispatchNote(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispatchNote](row)
	if err != nil {
		return nil, err
	}
	return apigen.RaiseDiscrepancyDispatchNote200JSONResponse{Body: body, Headers: apigen.RaiseDiscrepancyDispatchNote200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListProducts implements the listProducts operation.
func (s *Server) ListProducts(ctx context.Context, req apigen.ListProductsRequestObject) (apigen.ListProductsResponseObject, error) {
	return listPage[apigen.ListProducts200JSONResponse](ctx, req.Params, nil, s.service.ListProduct)
}

// GetProduct implements the getProduct operation.
func (s *Server) GetProduct(ctx context.Context, req apigen.GetProductRequestObject) (apigen.GetProductResponseObject, error) {
	row, err := s.service.GetProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Product](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetProduct200JSONResponse{Body: body, Headers: apigen.GetProduct200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateProduct implements the createProduct operation.
func (s *Server) CreateProduct(ctx context.Context, req apigen.CreateProductRequestObject) (apigen.CreateProductResponseObject, error) {
	in, err := input[db.CreateProductParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateProduct(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Product](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateProduct201JSONResponse{Body: body, Headers: apigen.CreateProduct201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateProduct implements the updateProduct operation.
func (s *Server) UpdateProduct(ctx context.Context, req apigen.UpdateProductRequestObject) (apigen.UpdateProductResponseObject, error) {
	in, err := input[db.UpdateProductParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ProductId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateProduct(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Product](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateProduct200JSONResponse{Body: body, Headers: apigen.UpdateProduct200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListProductChanges implements the listProductChanges operation.
func (s *Server) ListProductChanges(ctx context.Context, req apigen.ListProductChangesRequestObject) (apigen.ListProductChangesResponseObject, error) {
	return listPage[apigen.ListProductChanges200JSONResponse](ctx, req.Params, nil, s.service.ListProductChange)
}

// GetProductChange implements the getProductChange operation.
func (s *Server) GetProductChange(ctx context.Context, req apigen.GetProductChangeRequestObject) (apigen.GetProductChangeResponseObject, error) {
	row, err := s.service.GetProductChange(ctx, req.ProductChangeId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductChange](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetProductChange200JSONResponse{Body: body, Headers: apigen.GetProductChange200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateProductChange implements the createProductChange operation.
func (s *Server) CreateProductChange(ctx context.Context, req apigen.CreateProductChangeRequestObject) (apigen.CreateProductChangeResponseObject, error) {
	in, err := input[db.CreateProductChangeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateProductChange(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductChange](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateProductChange201JSONResponse{Body: body, Headers: apigen.CreateProductChange201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateProductChange implements the updateProductChange operation.
func (s *Server) UpdateProductChange(ctx context.Context, req apigen.UpdateProductChangeRequestObject) (apigen.UpdateProductChangeResponseObject, error) {
	in, err := input[db.UpdateProductChangeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ProductChangeId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateProductChange(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductChange](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateProductChange200JSONResponse{Body: body, Headers: apigen.UpdateProductChange200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitProductChange implements the submitProductChange operation.
func (s *Server) SubmitProductChange(ctx context.Context, req apigen.SubmitProductChangeRequestObject) (apigen.SubmitProductChangeResponseObject, error) {
	in, err := input[db.ActSubmitProductChangeParams](struct{}{}, map[string]any{"id": req.ProductChangeId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitProductChange(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductChange](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitProductChange200JSONResponse{Body: body, Headers: apigen.SubmitProductChange200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCostSheets implements the listCostSheets operation.
func (s *Server) ListCostSheets(ctx context.Context, req apigen.ListCostSheetsRequestObject) (apigen.ListCostSheetsResponseObject, error) {
	return listPage[apigen.ListCostSheets200JSONResponse](ctx, req.Params, nil, s.service.ListCostSheet)
}

// GetCostSheet implements the getCostSheet operation.
func (s *Server) GetCostSheet(ctx context.Context, req apigen.GetCostSheetRequestObject) (apigen.GetCostSheetResponseObject, error) {
	row, err := s.service.GetCostSheet(ctx, req.CostSheetId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CostSheet](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCostSheet200JSONResponse{Body: body, Headers: apigen.GetCostSheet200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCostSheet implements the createCostSheet operation.
func (s *Server) CreateCostSheet(ctx context.Context, req apigen.CreateCostSheetRequestObject) (apigen.CreateCostSheetResponseObject, error) {
	in, err := input[db.CreateCostSheetParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCostSheet(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CostSheet](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCostSheet201JSONResponse{Body: body, Headers: apigen.CreateCostSheet201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCostSheet implements the updateCostSheet operation.
func (s *Server) UpdateCostSheet(ctx context.Context, req apigen.UpdateCostSheetRequestObject) (apigen.UpdateCostSheetResponseObject, error) {
	in, err := input[db.UpdateCostSheetParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CostSheetId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCostSheet(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CostSheet](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCostSheet200JSONResponse{Body: body, Headers: apigen.UpdateCostSheet200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitCostSheet implements the submitCostSheet operation.
func (s *Server) SubmitCostSheet(ctx context.Context, req apigen.SubmitCostSheetRequestObject) (apigen.SubmitCostSheetResponseObject, error) {
	in, err := input[db.ActSubmitCostSheetParams](struct{}{}, map[string]any{"id": req.CostSheetId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitCostSheet(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CostSheet](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitCostSheet200JSONResponse{Body: body, Headers: apigen.SubmitCostSheet200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListNpdProjects implements the listNpdProjects operation.
func (s *Server) ListNpdProjects(ctx context.Context, req apigen.ListNpdProjectsRequestObject) (apigen.ListNpdProjectsResponseObject, error) {
	return listPage[apigen.ListNpdProjects200JSONResponse](ctx, req.Params, nil, s.service.ListNpdProject)
}

// GetNpdProject implements the getNpdProject operation.
func (s *Server) GetNpdProject(ctx context.Context, req apigen.GetNpdProjectRequestObject) (apigen.GetNpdProjectResponseObject, error) {
	row, err := s.service.GetNpdProject(ctx, req.NpdProjectId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.NpdProject](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetNpdProject200JSONResponse{Body: body, Headers: apigen.GetNpdProject200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateNpdProject implements the createNpdProject operation.
func (s *Server) CreateNpdProject(ctx context.Context, req apigen.CreateNpdProjectRequestObject) (apigen.CreateNpdProjectResponseObject, error) {
	in, err := input[db.CreateNpdProjectParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateNpdProject(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.NpdProject](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateNpdProject201JSONResponse{Body: body, Headers: apigen.CreateNpdProject201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateNpdProject implements the updateNpdProject operation.
func (s *Server) UpdateNpdProject(ctx context.Context, req apigen.UpdateNpdProjectRequestObject) (apigen.UpdateNpdProjectResponseObject, error) {
	in, err := input[db.UpdateNpdProjectParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.NpdProjectId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateNpdProject(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.NpdProject](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateNpdProject200JSONResponse{Body: body, Headers: apigen.UpdateNpdProject200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AdvanceGateNpdProject implements the advanceGateNpdProject operation.
func (s *Server) AdvanceGateNpdProject(ctx context.Context, req apigen.AdvanceGateNpdProjectRequestObject) (apigen.AdvanceGateNpdProjectResponseObject, error) {
	in, err := input[db.ActAdvanceGateNpdProjectParams](req.Body, map[string]any{"id": req.NpdProjectId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AdvanceGateNpdProject(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.NpdProject](row)
	if err != nil {
		return nil, err
	}
	return apigen.AdvanceGateNpdProject200JSONResponse{Body: body, Headers: apigen.AdvanceGateNpdProject200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
