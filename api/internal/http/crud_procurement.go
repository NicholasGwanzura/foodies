// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListSuppliers implements the listSuppliers operation.
func (s *Server) ListSuppliers(ctx context.Context, req apigen.ListSuppliersRequestObject) (apigen.ListSuppliersResponseObject, error) {
	return listPage[apigen.ListSuppliers200JSONResponse](ctx, req.Params, nil, s.service.ListSupplier)
}

// GetSupplier implements the getSupplier operation.
func (s *Server) GetSupplier(ctx context.Context, req apigen.GetSupplierRequestObject) (apigen.GetSupplierResponseObject, error) {
	row, err := s.service.GetSupplier(ctx, req.SupplierId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Supplier](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSupplier200JSONResponse{Body: body, Headers: apigen.GetSupplier200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateSupplier implements the createSupplier operation.
func (s *Server) CreateSupplier(ctx context.Context, req apigen.CreateSupplierRequestObject) (apigen.CreateSupplierResponseObject, error) {
	in, err := input[db.CreateSupplierParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSupplier(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Supplier](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSupplier201JSONResponse{Body: body, Headers: apigen.CreateSupplier201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateSupplier implements the updateSupplier operation.
func (s *Server) UpdateSupplier(ctx context.Context, req apigen.UpdateSupplierRequestObject) (apigen.UpdateSupplierResponseObject, error) {
	in, err := input[db.UpdateSupplierParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.SupplierId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateSupplier(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Supplier](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateSupplier200JSONResponse{Body: body, Headers: apigen.UpdateSupplier200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitSupplier implements the submitSupplier operation.
func (s *Server) SubmitSupplier(ctx context.Context, req apigen.SubmitSupplierRequestObject) (apigen.SubmitSupplierResponseObject, error) {
	in, err := input[db.ActSubmitSupplierParams](struct{}{}, map[string]any{"id": req.SupplierId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitSupplier(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Supplier](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitSupplier200JSONResponse{Body: body, Headers: apigen.SubmitSupplier200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SuspendSupplier implements the suspendSupplier operation.
func (s *Server) SuspendSupplier(ctx context.Context, req apigen.SuspendSupplierRequestObject) (apigen.SuspendSupplierResponseObject, error) {
	in, err := input[db.ActSuspendSupplierParams](req.Body, map[string]any{"id": req.SupplierId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SuspendSupplier(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Supplier](row)
	if err != nil {
		return nil, err
	}
	return apigen.SuspendSupplier200JSONResponse{Body: body, Headers: apigen.SuspendSupplier200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReinstateSupplier implements the reinstateSupplier operation.
func (s *Server) ReinstateSupplier(ctx context.Context, req apigen.ReinstateSupplierRequestObject) (apigen.ReinstateSupplierResponseObject, error) {
	in, err := input[db.ActReinstateSupplierParams](struct{}{}, map[string]any{"id": req.SupplierId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReinstateSupplier(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Supplier](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReinstateSupplier200JSONResponse{Body: body, Headers: apigen.ReinstateSupplier200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListSupplierAmendments implements the listSupplierAmendments operation.
func (s *Server) ListSupplierAmendments(ctx context.Context, req apigen.ListSupplierAmendmentsRequestObject) (apigen.ListSupplierAmendmentsResponseObject, error) {
	return listPage[apigen.ListSupplierAmendments200JSONResponse](ctx, req.Params, map[string]any{"supplier_id": req.SupplierId}, s.service.ListSupplierAmendment)
}

// GetSupplierAmendment implements the getSupplierAmendment operation.
func (s *Server) GetSupplierAmendment(ctx context.Context, req apigen.GetSupplierAmendmentRequestObject) (apigen.GetSupplierAmendmentResponseObject, error) {
	row, err := s.service.GetSupplierAmendment(ctx, req.SupplierAmendmentId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierAmendment](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSupplierAmendment200JSONResponse{Body: body, Headers: apigen.GetSupplierAmendment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateSupplierAmendment implements the createSupplierAmendment operation.
func (s *Server) CreateSupplierAmendment(ctx context.Context, req apigen.CreateSupplierAmendmentRequestObject) (apigen.CreateSupplierAmendmentResponseObject, error) {
	in, err := input[db.CreateSupplierAmendmentParams](req.Body, map[string]any{"supplier_id": req.SupplierId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSupplierAmendment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierAmendment](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSupplierAmendment201JSONResponse{Body: body, Headers: apigen.CreateSupplierAmendment201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitSupplierAmendment implements the submitSupplierAmendment operation.
func (s *Server) SubmitSupplierAmendment(ctx context.Context, req apigen.SubmitSupplierAmendmentRequestObject) (apigen.SubmitSupplierAmendmentResponseObject, error) {
	in, err := input[db.ActSubmitSupplierAmendmentParams](struct{}{}, map[string]any{"id": req.SupplierAmendmentId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitSupplierAmendment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierAmendment](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitSupplierAmendment200JSONResponse{Body: body, Headers: apigen.SubmitSupplierAmendment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPurchaseRequisitions implements the listPurchaseRequisitions operation.
func (s *Server) ListPurchaseRequisitions(ctx context.Context, req apigen.ListPurchaseRequisitionsRequestObject) (apigen.ListPurchaseRequisitionsResponseObject, error) {
	return listPage[apigen.ListPurchaseRequisitions200JSONResponse](ctx, req.Params, nil, s.service.ListPurchaseRequisition)
}

// GetPurchaseRequisition implements the getPurchaseRequisition operation.
func (s *Server) GetPurchaseRequisition(ctx context.Context, req apigen.GetPurchaseRequisitionRequestObject) (apigen.GetPurchaseRequisitionResponseObject, error) {
	row, err := s.service.GetPurchaseRequisition(ctx, req.PurchaseRequisitionId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.PurchaseRequisitionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseRequisition](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetPurchaseRequisition200JSONResponse{Body: body, Headers: apigen.GetPurchaseRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePurchaseRequisition implements the createPurchaseRequisition operation.
func (s *Server) CreatePurchaseRequisition(ctx context.Context, req apigen.CreatePurchaseRequisitionRequestObject) (apigen.CreatePurchaseRequisitionResponseObject, error) {
	in, err := input[db.CreatePurchaseRequisitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePurchaseRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertPurchaseRequisitionLinesParams](items, "purchase_requisition_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePurchaseRequisitionLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PurchaseRequisitionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseRequisition](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePurchaseRequisition201JSONResponse{Body: body, Headers: apigen.CreatePurchaseRequisition201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePurchaseRequisition implements the updatePurchaseRequisition operation.
func (s *Server) UpdatePurchaseRequisition(ctx context.Context, req apigen.UpdatePurchaseRequisitionRequestObject) (apigen.UpdatePurchaseRequisitionResponseObject, error) {
	in, err := input[db.UpdatePurchaseRequisitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PurchaseRequisitionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePurchaseRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertPurchaseRequisitionLinesParams](items, "purchase_requisition_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePurchaseRequisitionLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PurchaseRequisitionChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseRequisition](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePurchaseRequisition200JSONResponse{Body: body, Headers: apigen.UpdatePurchaseRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitPurchaseRequisition implements the submitPurchaseRequisition operation.
func (s *Server) SubmitPurchaseRequisition(ctx context.Context, req apigen.SubmitPurchaseRequisitionRequestObject) (apigen.SubmitPurchaseRequisitionResponseObject, error) {
	in, err := input[db.ActSubmitPurchaseRequisitionParams](struct{}{}, map[string]any{"id": req.PurchaseRequisitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitPurchaseRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitPurchaseRequisition200JSONResponse{Body: body, Headers: apigen.SubmitPurchaseRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ValidatePurchaseRequisition implements the validatePurchaseRequisition operation.
func (s *Server) ValidatePurchaseRequisition(ctx context.Context, req apigen.ValidatePurchaseRequisitionRequestObject) (apigen.ValidatePurchaseRequisitionResponseObject, error) {
	in, err := input[db.ActValidatePurchaseRequisitionParams](struct{}{}, map[string]any{"id": req.PurchaseRequisitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ValidatePurchaseRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.ValidatePurchaseRequisition200JSONResponse{Body: body, Headers: apigen.ValidatePurchaseRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReturnPurchaseRequisition implements the returnPurchaseRequisition operation.
func (s *Server) ReturnPurchaseRequisition(ctx context.Context, req apigen.ReturnPurchaseRequisitionRequestObject) (apigen.ReturnPurchaseRequisitionResponseObject, error) {
	in, err := input[db.ActReturnPurchaseRequisitionParams](req.Body, map[string]any{"id": req.PurchaseRequisitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReturnPurchaseRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReturnPurchaseRequisition200JSONResponse{Body: body, Headers: apigen.ReturnPurchaseRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListQuotations implements the listQuotations operation.
func (s *Server) ListQuotations(ctx context.Context, req apigen.ListQuotationsRequestObject) (apigen.ListQuotationsResponseObject, error) {
	return listPage[apigen.ListQuotations200JSONResponse](ctx, req.Params, map[string]any{"purchase_requisition_id": req.PurchaseRequisitionId}, s.service.ListQuotation)
}

// GetQuotation implements the getQuotation operation.
func (s *Server) GetQuotation(ctx context.Context, req apigen.GetQuotationRequestObject) (apigen.GetQuotationResponseObject, error) {
	row, err := s.service.GetQuotation(ctx, req.QuotationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Quotation](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetQuotation200JSONResponse{Body: body, Headers: apigen.GetQuotation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateQuotation implements the createQuotation operation.
func (s *Server) CreateQuotation(ctx context.Context, req apigen.CreateQuotationRequestObject) (apigen.CreateQuotationResponseObject, error) {
	in, err := input[db.CreateQuotationParams](req.Body, map[string]any{"purchase_requisition_id": req.PurchaseRequisitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateQuotation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Quotation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateQuotation201JSONResponse{Body: body, Headers: apigen.CreateQuotation201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateQuotation implements the updateQuotation operation.
func (s *Server) UpdateQuotation(ctx context.Context, req apigen.UpdateQuotationRequestObject) (apigen.UpdateQuotationResponseObject, error) {
	in, err := input[db.UpdateQuotationParams](req.Body, map[string]any{"purchase_requisition_id": req.PurchaseRequisitionId})
	if err != nil {
		return nil, err
	}
	in.ID = req.QuotationId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateQuotation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Quotation](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateQuotation200JSONResponse{Body: body, Headers: apigen.UpdateQuotation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DeleteQuotation implements the deleteQuotation operation.
func (s *Server) DeleteQuotation(ctx context.Context, req apigen.DeleteQuotationRequestObject) (apigen.DeleteQuotationResponseObject, error) {
	if err := s.service.DeleteQuotation(ctx, req.QuotationId); err != nil {
		return nil, err
	}
	return apigen.DeleteQuotation204Response{}, nil
}

// ListPurchaseOrders implements the listPurchaseOrders operation.
func (s *Server) ListPurchaseOrders(ctx context.Context, req apigen.ListPurchaseOrdersRequestObject) (apigen.ListPurchaseOrdersResponseObject, error) {
	return listPage[apigen.ListPurchaseOrders200JSONResponse](ctx, req.Params, nil, s.service.ListPurchaseOrder)
}

// GetPurchaseOrder implements the getPurchaseOrder operation.
func (s *Server) GetPurchaseOrder(ctx context.Context, req apigen.GetPurchaseOrderRequestObject) (apigen.GetPurchaseOrderResponseObject, error) {
	row, err := s.service.GetPurchaseOrder(ctx, req.PurchaseOrderId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.PurchaseOrderChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseOrder](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetPurchaseOrder200JSONResponse{Body: body, Headers: apigen.GetPurchaseOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePurchaseOrder implements the createPurchaseOrder operation.
func (s *Server) CreatePurchaseOrder(ctx context.Context, req apigen.CreatePurchaseOrderRequestObject) (apigen.CreatePurchaseOrderResponseObject, error) {
	in, err := input[db.CreatePurchaseOrderParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePurchaseOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertPurchaseOrderLinesParams](items, "purchase_order_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePurchaseOrderLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PurchaseOrderChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseOrder](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePurchaseOrder201JSONResponse{Body: body, Headers: apigen.CreatePurchaseOrder201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePurchaseOrder implements the updatePurchaseOrder operation.
func (s *Server) UpdatePurchaseOrder(ctx context.Context, req apigen.UpdatePurchaseOrderRequestObject) (apigen.UpdatePurchaseOrderResponseObject, error) {
	in, err := input[db.UpdatePurchaseOrderParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PurchaseOrderId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePurchaseOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertPurchaseOrderLinesParams](items, "purchase_order_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplacePurchaseOrderLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.PurchaseOrderChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseOrder](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePurchaseOrder200JSONResponse{Body: body, Headers: apigen.UpdatePurchaseOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitPurchaseOrder implements the submitPurchaseOrder operation.
func (s *Server) SubmitPurchaseOrder(ctx context.Context, req apigen.SubmitPurchaseOrderRequestObject) (apigen.SubmitPurchaseOrderResponseObject, error) {
	in, err := input[db.ActSubmitPurchaseOrderParams](struct{}{}, map[string]any{"id": req.PurchaseOrderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitPurchaseOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitPurchaseOrder200JSONResponse{Body: body, Headers: apigen.SubmitPurchaseOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// IssuePurchaseOrder implements the issuePurchaseOrder operation.
func (s *Server) IssuePurchaseOrder(ctx context.Context, req apigen.IssuePurchaseOrderRequestObject) (apigen.IssuePurchaseOrderResponseObject, error) {
	in, err := input[db.ActIssuePurchaseOrderParams](struct{}{}, map[string]any{"id": req.PurchaseOrderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.IssuePurchaseOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.IssuePurchaseOrder200JSONResponse{Body: body, Headers: apigen.IssuePurchaseOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CancelPurchaseOrder implements the cancelPurchaseOrder operation.
func (s *Server) CancelPurchaseOrder(ctx context.Context, req apigen.CancelPurchaseOrderRequestObject) (apigen.CancelPurchaseOrderResponseObject, error) {
	in, err := input[db.ActCancelPurchaseOrderParams](req.Body, map[string]any{"id": req.PurchaseOrderId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CancelPurchaseOrder(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PurchaseOrder](row)
	if err != nil {
		return nil, err
	}
	return apigen.CancelPurchaseOrder200JSONResponse{Body: body, Headers: apigen.CancelPurchaseOrder200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListGoodsReceipts implements the listGoodsReceipts operation.
func (s *Server) ListGoodsReceipts(ctx context.Context, req apigen.ListGoodsReceiptsRequestObject) (apigen.ListGoodsReceiptsResponseObject, error) {
	return listPage[apigen.ListGoodsReceipts200JSONResponse](ctx, req.Params, nil, s.service.ListGoodsReceipt)
}

// GetGoodsReceipt implements the getGoodsReceipt operation.
func (s *Server) GetGoodsReceipt(ctx context.Context, req apigen.GetGoodsReceiptRequestObject) (apigen.GetGoodsReceiptResponseObject, error) {
	row, err := s.service.GetGoodsReceipt(ctx, req.GoodsReceiptId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.GoodsReceiptChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.GoodsReceipt](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetGoodsReceipt200JSONResponse{Body: body, Headers: apigen.GetGoodsReceipt200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateGoodsReceipt implements the createGoodsReceipt operation.
func (s *Server) CreateGoodsReceipt(ctx context.Context, req apigen.CreateGoodsReceiptRequestObject) (apigen.CreateGoodsReceiptResponseObject, error) {
	in, err := input[db.CreateGoodsReceiptParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateGoodsReceipt(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertGoodsReceiptLinesParams](items, "goods_receipt_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceGoodsReceiptLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.GoodsReceiptChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.GoodsReceipt](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateGoodsReceipt201JSONResponse{Body: body, Headers: apigen.CreateGoodsReceipt201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateGoodsReceipt implements the updateGoodsReceipt operation.
func (s *Server) UpdateGoodsReceipt(ctx context.Context, req apigen.UpdateGoodsReceiptRequestObject) (apigen.UpdateGoodsReceiptResponseObject, error) {
	in, err := input[db.UpdateGoodsReceiptParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.GoodsReceiptId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateGoodsReceipt(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertGoodsReceiptLinesParams](items, "goods_receipt_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceGoodsReceiptLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.GoodsReceiptChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.GoodsReceipt](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateGoodsReceipt200JSONResponse{Body: body, Headers: apigen.UpdateGoodsReceipt200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitGoodsReceipt implements the submitGoodsReceipt operation.
func (s *Server) SubmitGoodsReceipt(ctx context.Context, req apigen.SubmitGoodsReceiptRequestObject) (apigen.SubmitGoodsReceiptResponseObject, error) {
	in, err := input[db.ActSubmitGoodsReceiptParams](struct{}{}, map[string]any{"id": req.GoodsReceiptId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitGoodsReceipt(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.GoodsReceipt](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitGoodsReceipt200JSONResponse{Body: body, Headers: apigen.SubmitGoodsReceipt200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDeliveryDiscrepancies implements the listDeliveryDiscrepancies operation.
func (s *Server) ListDeliveryDiscrepancies(ctx context.Context, req apigen.ListDeliveryDiscrepanciesRequestObject) (apigen.ListDeliveryDiscrepanciesResponseObject, error) {
	return listPage[apigen.ListDeliveryDiscrepancies200JSONResponse](ctx, req.Params, nil, s.service.ListDeliveryDiscrepancy)
}

// GetDeliveryDiscrepancy implements the getDeliveryDiscrepancy operation.
func (s *Server) GetDeliveryDiscrepancy(ctx context.Context, req apigen.GetDeliveryDiscrepancyRequestObject) (apigen.GetDeliveryDiscrepancyResponseObject, error) {
	row, err := s.service.GetDeliveryDiscrepancy(ctx, req.DeliveryDiscrepancyId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DeliveryDiscrepancy](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDeliveryDiscrepancy200JSONResponse{Body: body, Headers: apigen.GetDeliveryDiscrepancy200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDeliveryDiscrepancy implements the createDeliveryDiscrepancy operation.
func (s *Server) CreateDeliveryDiscrepancy(ctx context.Context, req apigen.CreateDeliveryDiscrepancyRequestObject) (apigen.CreateDeliveryDiscrepancyResponseObject, error) {
	in, err := input[db.CreateDeliveryDiscrepancyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDeliveryDiscrepancy(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DeliveryDiscrepancy](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDeliveryDiscrepancy201JSONResponse{Body: body, Headers: apigen.CreateDeliveryDiscrepancy201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDeliveryDiscrepancy implements the updateDeliveryDiscrepancy operation.
func (s *Server) UpdateDeliveryDiscrepancy(ctx context.Context, req apigen.UpdateDeliveryDiscrepancyRequestObject) (apigen.UpdateDeliveryDiscrepancyResponseObject, error) {
	in, err := input[db.UpdateDeliveryDiscrepancyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.DeliveryDiscrepancyId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDeliveryDiscrepancy(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DeliveryDiscrepancy](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDeliveryDiscrepancy200JSONResponse{Body: body, Headers: apigen.UpdateDeliveryDiscrepancy200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseDeliveryDiscrepancy implements the closeDeliveryDiscrepancy operation.
func (s *Server) CloseDeliveryDiscrepancy(ctx context.Context, req apigen.CloseDeliveryDiscrepancyRequestObject) (apigen.CloseDeliveryDiscrepancyResponseObject, error) {
	in, err := input[db.ActCloseDeliveryDiscrepancyParams](req.Body, map[string]any{"id": req.DeliveryDiscrepancyId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseDeliveryDiscrepancy(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DeliveryDiscrepancy](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseDeliveryDiscrepancy200JSONResponse{Body: body, Headers: apigen.CloseDeliveryDiscrepancy200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListSupplierInvoices implements the listSupplierInvoices operation.
func (s *Server) ListSupplierInvoices(ctx context.Context, req apigen.ListSupplierInvoicesRequestObject) (apigen.ListSupplierInvoicesResponseObject, error) {
	return listPage[apigen.ListSupplierInvoices200JSONResponse](ctx, req.Params, nil, s.service.ListSupplierInvoice)
}

// GetSupplierInvoice implements the getSupplierInvoice operation.
func (s *Server) GetSupplierInvoice(ctx context.Context, req apigen.GetSupplierInvoiceRequestObject) (apigen.GetSupplierInvoiceResponseObject, error) {
	row, err := s.service.GetSupplierInvoice(ctx, req.SupplierInvoiceId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.SupplierInvoiceChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierInvoice](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetSupplierInvoice200JSONResponse{Body: body, Headers: apigen.GetSupplierInvoice200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateSupplierInvoice implements the createSupplierInvoice operation.
func (s *Server) CreateSupplierInvoice(ctx context.Context, req apigen.CreateSupplierInvoiceRequestObject) (apigen.CreateSupplierInvoiceResponseObject, error) {
	in, err := input[db.CreateSupplierInvoiceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSupplierInvoice(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertSupplierInvoiceLinesParams](items, "supplier_invoice_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceSupplierInvoiceLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.SupplierInvoiceChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierInvoice](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSupplierInvoice201JSONResponse{Body: body, Headers: apigen.CreateSupplierInvoice201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateSupplierInvoice implements the updateSupplierInvoice operation.
func (s *Server) UpdateSupplierInvoice(ctx context.Context, req apigen.UpdateSupplierInvoiceRequestObject) (apigen.UpdateSupplierInvoiceResponseObject, error) {
	in, err := input[db.UpdateSupplierInvoiceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.SupplierInvoiceId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateSupplierInvoice(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertSupplierInvoiceLinesParams](items, "supplier_invoice_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceSupplierInvoiceLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.SupplierInvoiceChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierInvoice](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateSupplierInvoice200JSONResponse{Body: body, Headers: apigen.UpdateSupplierInvoice200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// MatchSupplierInvoice implements the matchSupplierInvoice operation.
func (s *Server) MatchSupplierInvoice(ctx context.Context, req apigen.MatchSupplierInvoiceRequestObject) (apigen.MatchSupplierInvoiceResponseObject, error) {
	in, err := input[db.ActMatchSupplierInvoiceParams](struct{}{}, map[string]any{"id": req.SupplierInvoiceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.MatchSupplierInvoice(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierInvoice](row)
	if err != nil {
		return nil, err
	}
	return apigen.MatchSupplierInvoice200JSONResponse{Body: body, Headers: apigen.MatchSupplierInvoice200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// HoldSupplierInvoice implements the holdSupplierInvoice operation.
func (s *Server) HoldSupplierInvoice(ctx context.Context, req apigen.HoldSupplierInvoiceRequestObject) (apigen.HoldSupplierInvoiceResponseObject, error) {
	in, err := input[db.ActHoldSupplierInvoiceParams](req.Body, map[string]any{"id": req.SupplierInvoiceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.HoldSupplierInvoice(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierInvoice](row)
	if err != nil {
		return nil, err
	}
	return apigen.HoldSupplierInvoice200JSONResponse{Body: body, Headers: apigen.HoldSupplierInvoice200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ApproveExceptionSupplierInvoice implements the approveExceptionSupplierInvoice operation.
func (s *Server) ApproveExceptionSupplierInvoice(ctx context.Context, req apigen.ApproveExceptionSupplierInvoiceRequestObject) (apigen.ApproveExceptionSupplierInvoiceResponseObject, error) {
	in, err := input[db.ActApproveExceptionSupplierInvoiceParams](req.Body, map[string]any{"id": req.SupplierInvoiceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ApproveExceptionSupplierInvoice(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierInvoice](row)
	if err != nil {
		return nil, err
	}
	return apigen.ApproveExceptionSupplierInvoice200JSONResponse{Body: body, Headers: apigen.ApproveExceptionSupplierInvoice200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// PostSupplierInvoice implements the postSupplierInvoice operation.
func (s *Server) PostSupplierInvoice(ctx context.Context, req apigen.PostSupplierInvoiceRequestObject) (apigen.PostSupplierInvoiceResponseObject, error) {
	in, err := input[db.ActPostSupplierInvoiceParams](struct{}{}, map[string]any{"id": req.SupplierInvoiceId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.PostSupplierInvoice(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SupplierInvoice](row)
	if err != nil {
		return nil, err
	}
	return apigen.PostSupplierInvoice200JSONResponse{Body: body, Headers: apigen.PostSupplierInvoice200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPaymentRequisitions implements the listPaymentRequisitions operation.
func (s *Server) ListPaymentRequisitions(ctx context.Context, req apigen.ListPaymentRequisitionsRequestObject) (apigen.ListPaymentRequisitionsResponseObject, error) {
	return listPage[apigen.ListPaymentRequisitions200JSONResponse](ctx, req.Params, nil, s.service.ListPaymentRequisition)
}

// GetPaymentRequisition implements the getPaymentRequisition operation.
func (s *Server) GetPaymentRequisition(ctx context.Context, req apigen.GetPaymentRequisitionRequestObject) (apigen.GetPaymentRequisitionResponseObject, error) {
	row, err := s.service.GetPaymentRequisition(ctx, req.PaymentRequisitionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPaymentRequisition200JSONResponse{Body: body, Headers: apigen.GetPaymentRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePaymentRequisition implements the createPaymentRequisition operation.
func (s *Server) CreatePaymentRequisition(ctx context.Context, req apigen.CreatePaymentRequisitionRequestObject) (apigen.CreatePaymentRequisitionResponseObject, error) {
	in, err := input[db.CreatePaymentRequisitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePaymentRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePaymentRequisition201JSONResponse{Body: body, Headers: apigen.CreatePaymentRequisition201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePaymentRequisition implements the updatePaymentRequisition operation.
func (s *Server) UpdatePaymentRequisition(ctx context.Context, req apigen.UpdatePaymentRequisitionRequestObject) (apigen.UpdatePaymentRequisitionResponseObject, error) {
	in, err := input[db.UpdatePaymentRequisitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PaymentRequisitionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePaymentRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePaymentRequisition200JSONResponse{Body: body, Headers: apigen.UpdatePaymentRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitPaymentRequisition implements the submitPaymentRequisition operation.
func (s *Server) SubmitPaymentRequisition(ctx context.Context, req apigen.SubmitPaymentRequisitionRequestObject) (apigen.SubmitPaymentRequisitionResponseObject, error) {
	in, err := input[db.ActSubmitPaymentRequisitionParams](struct{}{}, map[string]any{"id": req.PaymentRequisitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitPaymentRequisition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentRequisition](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitPaymentRequisition200JSONResponse{Body: body, Headers: apigen.SubmitPaymentRequisition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPaymentBatches implements the listPaymentBatches operation.
func (s *Server) ListPaymentBatches(ctx context.Context, req apigen.ListPaymentBatchesRequestObject) (apigen.ListPaymentBatchesResponseObject, error) {
	return listPage[apigen.ListPaymentBatches200JSONResponse](ctx, req.Params, nil, s.service.ListPaymentBatch)
}

// GetPaymentBatch implements the getPaymentBatch operation.
func (s *Server) GetPaymentBatch(ctx context.Context, req apigen.GetPaymentBatchRequestObject) (apigen.GetPaymentBatchResponseObject, error) {
	row, err := s.service.GetPaymentBatch(ctx, req.PaymentBatchId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentBatch](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPaymentBatch200JSONResponse{Body: body, Headers: apigen.GetPaymentBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePaymentBatch implements the createPaymentBatch operation.
func (s *Server) CreatePaymentBatch(ctx context.Context, req apigen.CreatePaymentBatchRequestObject) (apigen.CreatePaymentBatchResponseObject, error) {
	in, err := input[db.CreatePaymentBatchParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePaymentBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentBatch](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePaymentBatch201JSONResponse{Body: body, Headers: apigen.CreatePaymentBatch201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePaymentBatch implements the updatePaymentBatch operation.
func (s *Server) UpdatePaymentBatch(ctx context.Context, req apigen.UpdatePaymentBatchRequestObject) (apigen.UpdatePaymentBatchResponseObject, error) {
	in, err := input[db.UpdatePaymentBatchParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PaymentBatchId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePaymentBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentBatch](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePaymentBatch200JSONResponse{Body: body, Headers: apigen.UpdatePaymentBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitPaymentBatch implements the submitPaymentBatch operation.
func (s *Server) SubmitPaymentBatch(ctx context.Context, req apigen.SubmitPaymentBatchRequestObject) (apigen.SubmitPaymentBatchResponseObject, error) {
	in, err := input[db.ActSubmitPaymentBatchParams](struct{}{}, map[string]any{"id": req.PaymentBatchId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitPaymentBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentBatch](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitPaymentBatch200JSONResponse{Body: body, Headers: apigen.SubmitPaymentBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RecordExecutionPaymentBatch implements the recordExecutionPaymentBatch operation.
func (s *Server) RecordExecutionPaymentBatch(ctx context.Context, req apigen.RecordExecutionPaymentBatchRequestObject) (apigen.RecordExecutionPaymentBatchResponseObject, error) {
	in, err := input[db.ActRecordExecutionPaymentBatchParams](req.Body, map[string]any{"id": req.PaymentBatchId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RecordExecutionPaymentBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PaymentBatch](row)
	if err != nil {
		return nil, err
	}
	return apigen.RecordExecutionPaymentBatch200JSONResponse{Body: body, Headers: apigen.RecordExecutionPaymentBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListEmergencyProcurements implements the listEmergencyProcurements operation.
func (s *Server) ListEmergencyProcurements(ctx context.Context, req apigen.ListEmergencyProcurementsRequestObject) (apigen.ListEmergencyProcurementsResponseObject, error) {
	return listPage[apigen.ListEmergencyProcurements200JSONResponse](ctx, req.Params, nil, s.service.ListEmergencyProcurement)
}

// GetEmergencyProcurement implements the getEmergencyProcurement operation.
func (s *Server) GetEmergencyProcurement(ctx context.Context, req apigen.GetEmergencyProcurementRequestObject) (apigen.GetEmergencyProcurementResponseObject, error) {
	row, err := s.service.GetEmergencyProcurement(ctx, req.EmergencyProcurementId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmergencyProcurement](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetEmergencyProcurement200JSONResponse{Body: body, Headers: apigen.GetEmergencyProcurement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateEmergencyProcurement implements the createEmergencyProcurement operation.
func (s *Server) CreateEmergencyProcurement(ctx context.Context, req apigen.CreateEmergencyProcurementRequestObject) (apigen.CreateEmergencyProcurementResponseObject, error) {
	in, err := input[db.CreateEmergencyProcurementParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateEmergencyProcurement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmergencyProcurement](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateEmergencyProcurement201JSONResponse{Body: body, Headers: apigen.CreateEmergencyProcurement201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateEmergencyProcurement implements the updateEmergencyProcurement operation.
func (s *Server) UpdateEmergencyProcurement(ctx context.Context, req apigen.UpdateEmergencyProcurementRequestObject) (apigen.UpdateEmergencyProcurementResponseObject, error) {
	in, err := input[db.UpdateEmergencyProcurementParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.EmergencyProcurementId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateEmergencyProcurement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmergencyProcurement](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateEmergencyProcurement200JSONResponse{Body: body, Headers: apigen.UpdateEmergencyProcurement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitEmergencyProcurement implements the submitEmergencyProcurement operation.
func (s *Server) SubmitEmergencyProcurement(ctx context.Context, req apigen.SubmitEmergencyProcurementRequestObject) (apigen.SubmitEmergencyProcurementResponseObject, error) {
	in, err := input[db.ActSubmitEmergencyProcurementParams](struct{}{}, map[string]any{"id": req.EmergencyProcurementId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitEmergencyProcurement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmergencyProcurement](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitEmergencyProcurement200JSONResponse{Body: body, Headers: apigen.SubmitEmergencyProcurement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReviewEmergencyProcurement implements the reviewEmergencyProcurement operation.
func (s *Server) ReviewEmergencyProcurement(ctx context.Context, req apigen.ReviewEmergencyProcurementRequestObject) (apigen.ReviewEmergencyProcurementResponseObject, error) {
	in, err := input[db.ActReviewEmergencyProcurementParams](req.Body, map[string]any{"id": req.EmergencyProcurementId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReviewEmergencyProcurement(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmergencyProcurement](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReviewEmergencyProcurement200JSONResponse{Body: body, Headers: apigen.ReviewEmergencyProcurement200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListUnauthorizedPurchases implements the listUnauthorizedPurchases operation.
func (s *Server) ListUnauthorizedPurchases(ctx context.Context, req apigen.ListUnauthorizedPurchasesRequestObject) (apigen.ListUnauthorizedPurchasesResponseObject, error) {
	return listPage[apigen.ListUnauthorizedPurchases200JSONResponse](ctx, req.Params, nil, s.service.ListUnauthorizedPurchase)
}

// GetUnauthorizedPurchase implements the getUnauthorizedPurchase operation.
func (s *Server) GetUnauthorizedPurchase(ctx context.Context, req apigen.GetUnauthorizedPurchaseRequestObject) (apigen.GetUnauthorizedPurchaseResponseObject, error) {
	row, err := s.service.GetUnauthorizedPurchase(ctx, req.UnauthorizedPurchaseId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.UnauthorizedPurchase](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetUnauthorizedPurchase200JSONResponse{Body: body, Headers: apigen.GetUnauthorizedPurchase200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateUnauthorizedPurchase implements the createUnauthorizedPurchase operation.
func (s *Server) CreateUnauthorizedPurchase(ctx context.Context, req apigen.CreateUnauthorizedPurchaseRequestObject) (apigen.CreateUnauthorizedPurchaseResponseObject, error) {
	in, err := input[db.CreateUnauthorizedPurchaseParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateUnauthorizedPurchase(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.UnauthorizedPurchase](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateUnauthorizedPurchase201JSONResponse{Body: body, Headers: apigen.CreateUnauthorizedPurchase201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateUnauthorizedPurchase implements the updateUnauthorizedPurchase operation.
func (s *Server) UpdateUnauthorizedPurchase(ctx context.Context, req apigen.UpdateUnauthorizedPurchaseRequestObject) (apigen.UpdateUnauthorizedPurchaseResponseObject, error) {
	in, err := input[db.UpdateUnauthorizedPurchaseParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.UnauthorizedPurchaseId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateUnauthorizedPurchase(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.UnauthorizedPurchase](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateUnauthorizedPurchase200JSONResponse{Body: body, Headers: apigen.UpdateUnauthorizedPurchase200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CloseUnauthorizedPurchase implements the closeUnauthorizedPurchase operation.
func (s *Server) CloseUnauthorizedPurchase(ctx context.Context, req apigen.CloseUnauthorizedPurchaseRequestObject) (apigen.CloseUnauthorizedPurchaseResponseObject, error) {
	in, err := input[db.ActCloseUnauthorizedPurchaseParams](req.Body, map[string]any{"id": req.UnauthorizedPurchaseId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CloseUnauthorizedPurchase(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.UnauthorizedPurchase](row)
	if err != nil {
		return nil, err
	}
	return apigen.CloseUnauthorizedPurchase200JSONResponse{Body: body, Headers: apigen.CloseUnauthorizedPurchase200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListBudgets implements the listBudgets operation.
func (s *Server) ListBudgets(ctx context.Context, req apigen.ListBudgetsRequestObject) (apigen.ListBudgetsResponseObject, error) {
	return listPage[apigen.ListBudgets200JSONResponse](ctx, req.Params, nil, s.service.ListBudget)
}

// GetBudget implements the getBudget operation.
func (s *Server) GetBudget(ctx context.Context, req apigen.GetBudgetRequestObject) (apigen.GetBudgetResponseObject, error) {
	row, err := s.service.GetBudget(ctx, req.BudgetId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.BudgetChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Budget](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetBudget200JSONResponse{Body: body, Headers: apigen.GetBudget200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateBudget implements the createBudget operation.
func (s *Server) CreateBudget(ctx context.Context, req apigen.CreateBudgetRequestObject) (apigen.CreateBudgetResponseObject, error) {
	in, err := input[db.CreateBudgetParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateBudget(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertBudgetLinesParams](items, "budget_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceBudgetLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.BudgetChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Budget](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateBudget201JSONResponse{Body: body, Headers: apigen.CreateBudget201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateBudget implements the updateBudget operation.
func (s *Server) UpdateBudget(ctx context.Context, req apigen.UpdateBudgetRequestObject) (apigen.UpdateBudgetResponseObject, error) {
	in, err := input[db.UpdateBudgetParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.BudgetId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateBudget(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["lines"]; ok {
		rows, err := childInputs[db.InsertBudgetLinesParams](items, "budget_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceBudgetLines(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.BudgetChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Budget](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateBudget200JSONResponse{Body: body, Headers: apigen.UpdateBudget200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitBudget implements the submitBudget operation.
func (s *Server) SubmitBudget(ctx context.Context, req apigen.SubmitBudgetRequestObject) (apigen.SubmitBudgetResponseObject, error) {
	in, err := input[db.ActSubmitBudgetParams](struct{}{}, map[string]any{"id": req.BudgetId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitBudget(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Budget](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitBudget200JSONResponse{Body: body, Headers: apigen.SubmitBudget200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListBudgetVariances implements the listBudgetVariances operation.
func (s *Server) ListBudgetVariances(ctx context.Context, req apigen.ListBudgetVariancesRequestObject) (apigen.ListBudgetVariancesResponseObject, error) {
	return listPage[apigen.ListBudgetVariances200JSONResponse](ctx, req.Params, map[string]any{"budget_id": req.BudgetId}, s.service.ListBudgetVariance)
}

// GetBudgetVariance implements the getBudgetVariance operation.
func (s *Server) GetBudgetVariance(ctx context.Context, req apigen.GetBudgetVarianceRequestObject) (apigen.GetBudgetVarianceResponseObject, error) {
	row, err := s.service.GetBudgetVariance(ctx, req.BudgetVarianceId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BudgetVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetBudgetVariance200JSONResponse{Body: body, Headers: apigen.GetBudgetVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateBudgetVariance implements the updateBudgetVariance operation.
func (s *Server) UpdateBudgetVariance(ctx context.Context, req apigen.UpdateBudgetVarianceRequestObject) (apigen.UpdateBudgetVarianceResponseObject, error) {
	in, err := input[db.UpdateBudgetVarianceParams](req.Body, map[string]any{"budget_id": req.BudgetId})
	if err != nil {
		return nil, err
	}
	in.ID = req.BudgetVarianceId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateBudgetVariance(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BudgetVariance](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateBudgetVariance200JSONResponse{Body: body, Headers: apigen.UpdateBudgetVariance200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListBudgetReallocations implements the listBudgetReallocations operation.
func (s *Server) ListBudgetReallocations(ctx context.Context, req apigen.ListBudgetReallocationsRequestObject) (apigen.ListBudgetReallocationsResponseObject, error) {
	return listPage[apigen.ListBudgetReallocations200JSONResponse](ctx, req.Params, nil, s.service.ListBudgetReallocation)
}

// GetBudgetReallocation implements the getBudgetReallocation operation.
func (s *Server) GetBudgetReallocation(ctx context.Context, req apigen.GetBudgetReallocationRequestObject) (apigen.GetBudgetReallocationResponseObject, error) {
	row, err := s.service.GetBudgetReallocation(ctx, req.BudgetReallocationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BudgetReallocation](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetBudgetReallocation200JSONResponse{Body: body, Headers: apigen.GetBudgetReallocation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateBudgetReallocation implements the createBudgetReallocation operation.
func (s *Server) CreateBudgetReallocation(ctx context.Context, req apigen.CreateBudgetReallocationRequestObject) (apigen.CreateBudgetReallocationResponseObject, error) {
	in, err := input[db.CreateBudgetReallocationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateBudgetReallocation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BudgetReallocation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateBudgetReallocation201JSONResponse{Body: body, Headers: apigen.CreateBudgetReallocation201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitBudgetReallocation implements the submitBudgetReallocation operation.
func (s *Server) SubmitBudgetReallocation(ctx context.Context, req apigen.SubmitBudgetReallocationRequestObject) (apigen.SubmitBudgetReallocationResponseObject, error) {
	in, err := input[db.ActSubmitBudgetReallocationParams](struct{}{}, map[string]any{"id": req.BudgetReallocationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitBudgetReallocation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.BudgetReallocation](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitBudgetReallocation200JSONResponse{Body: body, Headers: apigen.SubmitBudgetReallocation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
