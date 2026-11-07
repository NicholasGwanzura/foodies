// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListSupplier returns a page of suppliers.
func (s *Service) ListSupplier(ctx context.Context, p db.ListSupplierParams) ([]db.Supplier, error) {
	return wrap2(s.q(ctx).ListSupplier(ctx, p))("list suppliers")
}

// GetSupplier returns one row of suppliers.
func (s *Service) GetSupplier(ctx context.Context, id uuid.UUID) (db.Supplier, error) {
	row, err := s.q(ctx).GetSupplier(ctx, id)
	if err != nil {
		return row, notFound(err, "suppliers")
	}
	return row, nil
}

// CreateSupplier adds one row to suppliers.
func (s *Service) CreateSupplier(ctx context.Context, p db.CreateSupplierParams) (db.Supplier, error) {
	var zero db.Supplier
	if err := s.guard(ctx, "CreateSupplier", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "SUP", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateSupplier(ctx, p)
	if err != nil {
		return zero, wrap(err, "create suppliers")
	}
	if err := s.audit(ctx, "procurement.supplier.write", "suppliers", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateSupplier", row)
}

// UpdateSupplier changes one row of suppliers.
func (s *Service) UpdateSupplier(ctx context.Context, p db.UpdateSupplierParams) (db.Supplier, error) {
	before, err := s.GetSupplier(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateSupplier(ctx, p)
	if err != nil {
		return before, wrap(err, "update suppliers")
	}
	if err := s.voidSignatures(ctx, "suppliers", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.supplier.write", "suppliers", row.ID, nil, before, row)
}

// SubmitSupplier is the submit action on suppliers.
func (s *Service) SubmitSupplier(ctx context.Context, p db.ActSubmitSupplierParams) (db.Supplier, error) {
	before, err := s.GetSupplier(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.supplier.submit", "suppliers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitSupplier", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitSupplier(ctx, p)
	if err != nil {
		return before, stateError(err, "submit suppliers")
	}
	if err := s.audit(ctx, "procurement.supplier.submit", "suppliers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitSupplier", row)
}

// SuspendSupplier is the suspend action on suppliers.
func (s *Service) SuspendSupplier(ctx context.Context, p db.ActSuspendSupplierParams) (db.Supplier, error) {
	before, err := s.GetSupplier(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.supplier.suspend", "suppliers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SuspendSupplier", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSuspendSupplier(ctx, p)
	if err != nil {
		return before, stateError(err, "suspend suppliers")
	}
	if err := s.audit(ctx, "procurement.supplier.suspend", "suppliers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SuspendSupplier", row)
}

// ReinstateSupplier is the reinstate action on suppliers.
func (s *Service) ReinstateSupplier(ctx context.Context, p db.ActReinstateSupplierParams) (db.Supplier, error) {
	before, err := s.GetSupplier(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.supplier.reinstate", "suppliers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReinstateSupplier", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReinstateSupplier(ctx, p)
	if err != nil {
		return before, stateError(err, "reinstate suppliers")
	}
	if err := s.audit(ctx, "procurement.supplier.reinstate", "suppliers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReinstateSupplier", row)
}

// ListSupplierAmendment returns a page of supplier amendments.
func (s *Service) ListSupplierAmendment(ctx context.Context, p db.ListSupplierAmendmentParams) ([]db.SupplierAmendment, error) {
	return wrap2(s.q(ctx).ListSupplierAmendment(ctx, p))("list supplier amendments")
}

// GetSupplierAmendment returns one row of supplier amendments.
func (s *Service) GetSupplierAmendment(ctx context.Context, id uuid.UUID) (db.SupplierAmendment, error) {
	row, err := s.q(ctx).GetSupplierAmendment(ctx, id)
	if err != nil {
		return row, notFound(err, "supplier amendments")
	}
	return row, nil
}

// CreateSupplierAmendment adds one row to supplier amendments.
func (s *Service) CreateSupplierAmendment(ctx context.Context, p db.CreateSupplierAmendmentParams) (db.SupplierAmendment, error) {
	var zero db.SupplierAmendment
	if err := s.guard(ctx, "CreateSupplierAmendment", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateSupplierAmendment(ctx, p)
	if err != nil {
		return zero, wrap(err, "create supplier amendments")
	}
	if err := s.audit(ctx, "procurement.supplier_amendment.write", "supplier_amendments", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateSupplierAmendment", row)
}

// SubmitSupplierAmendment is the submit action on supplier amendments.
func (s *Service) SubmitSupplierAmendment(ctx context.Context, p db.ActSubmitSupplierAmendmentParams) (db.SupplierAmendment, error) {
	before, err := s.GetSupplierAmendment(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.supplier_amendment.submit", "supplier_amendments", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitSupplierAmendment", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitSupplierAmendment(ctx, p)
	if err != nil {
		return before, stateError(err, "submit supplier amendments")
	}
	if err := s.audit(ctx, "procurement.supplier_amendment.submit", "supplier_amendments", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "supplier_amendments", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitSupplierAmendment", row)
}

// ListPurchaseRequisition returns a page of purchase requisitions.
func (s *Service) ListPurchaseRequisition(ctx context.Context, p db.ListPurchaseRequisitionParams) ([]db.PurchaseRequisition, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListPurchaseRequisition(ctx, p))("list purchase requisitions")
}

// GetPurchaseRequisition returns one row of purchase requisitions.
func (s *Service) GetPurchaseRequisition(ctx context.Context, id uuid.UUID) (db.PurchaseRequisition, error) {
	row, err := s.q(ctx).GetPurchaseRequisition(ctx, id)
	if err != nil {
		return row, notFound(err, "purchase requisitions")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreatePurchaseRequisition adds one row to purchase requisitions.
func (s *Service) CreatePurchaseRequisition(ctx context.Context, p db.CreatePurchaseRequisitionParams) (db.PurchaseRequisition, error) {
	var zero db.PurchaseRequisition
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreatePurchaseRequisition", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "PR", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePurchaseRequisition(ctx, p)
	if err != nil {
		return zero, wrap(err, "create purchase requisitions")
	}
	if err := s.audit(ctx, "procurement.purchase_requisition.write", "purchase_requisitions", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreatePurchaseRequisition", row)
}

// UpdatePurchaseRequisition changes one row of purchase requisitions.
func (s *Service) UpdatePurchaseRequisition(ctx context.Context, p db.UpdatePurchaseRequisitionParams) (db.PurchaseRequisition, error) {
	before, err := s.GetPurchaseRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePurchaseRequisition(ctx, p)
	if err != nil {
		return before, wrap(err, "update purchase requisitions")
	}
	if err := s.voidSignatures(ctx, "purchase_requisitions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.purchase_requisition.write", "purchase_requisitions", row.ID, &row.SiteID, before, row)
}

// SubmitPurchaseRequisition is the submit action on purchase requisitions.
func (s *Service) SubmitPurchaseRequisition(ctx context.Context, p db.ActSubmitPurchaseRequisitionParams) (db.PurchaseRequisition, error) {
	before, err := s.GetPurchaseRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.purchase_requisition.submit", "purchase_requisitions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitPurchaseRequisition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitPurchaseRequisition(ctx, p)
	if err != nil {
		return before, stateError(err, "submit purchase requisitions")
	}
	if err := s.audit(ctx, "procurement.purchase_requisition.submit", "purchase_requisitions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "purchase_requisitions", row.ID, &row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitPurchaseRequisition", row)
}

// ValidatePurchaseRequisition is the procurement validates completeness action on purchase requisitions.
func (s *Service) ValidatePurchaseRequisition(ctx context.Context, p db.ActValidatePurchaseRequisitionParams) (db.PurchaseRequisition, error) {
	before, err := s.GetPurchaseRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.purchase_requisition.validate", "purchase_requisitions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ValidatePurchaseRequisition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActValidatePurchaseRequisition(ctx, p)
	if err != nil {
		return before, stateError(err, "validate purchase requisitions")
	}
	if err := s.audit(ctx, "procurement.purchase_requisition.validate", "purchase_requisitions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ValidatePurchaseRequisition", row)
}

// ReturnPurchaseRequisition is the return as incomplete action on purchase requisitions.
func (s *Service) ReturnPurchaseRequisition(ctx context.Context, p db.ActReturnPurchaseRequisitionParams) (db.PurchaseRequisition, error) {
	before, err := s.GetPurchaseRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.purchase_requisition.return", "purchase_requisitions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReturnPurchaseRequisition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReturnPurchaseRequisition(ctx, p)
	if err != nil {
		return before, stateError(err, "return purchase requisitions")
	}
	if err := s.audit(ctx, "procurement.purchase_requisition.return", "purchase_requisitions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReturnPurchaseRequisition", row)
}

// PurchaseRequisitionChildren returns the rows of a record's child tables, by API field.
func (s *Service) PurchaseRequisitionChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListPurchaseRequisitionLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list purchase_requisition_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplacePurchaseRequisitionLines swaps a record's purchase requisition lines for the given rows.
func (s *Service) ReplacePurchaseRequisitionLines(ctx context.Context, id uuid.UUID, rows []db.InsertPurchaseRequisitionLinesParams) error {
	if err := s.q(ctx).DeletePurchaseRequisitionLines(ctx, id); err != nil {
		return wrap(err, "clear purchase_requisition_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertPurchaseRequisitionLines(ctx, row); err != nil {
			return wrap(err, "add purchase_requisition_lines")
		}
	}
	return nil
}

// ListQuotation returns a page of quotations.
func (s *Service) ListQuotation(ctx context.Context, p db.ListQuotationParams) ([]db.Quotation, error) {
	return wrap2(s.q(ctx).ListQuotation(ctx, p))("list quotations")
}

// GetQuotation returns one row of quotations.
func (s *Service) GetQuotation(ctx context.Context, id uuid.UUID) (db.Quotation, error) {
	row, err := s.q(ctx).GetQuotation(ctx, id)
	if err != nil {
		return row, notFound(err, "quotations")
	}
	return row, nil
}

// CreateQuotation adds one row to quotations.
func (s *Service) CreateQuotation(ctx context.Context, p db.CreateQuotationParams) (db.Quotation, error) {
	var zero db.Quotation
	if err := s.guard(ctx, "CreateQuotation", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateQuotation(ctx, p)
	if err != nil {
		return zero, wrap(err, "create quotations")
	}
	if err := s.audit(ctx, "procurement.quotation.write", "quotations", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateQuotation", row)
}

// UpdateQuotation changes one row of quotations.
func (s *Service) UpdateQuotation(ctx context.Context, p db.UpdateQuotationParams) (db.Quotation, error) {
	before, err := s.GetQuotation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateQuotation(ctx, p)
	if err != nil {
		return before, wrap(err, "update quotations")
	}
	if err := s.voidSignatures(ctx, "quotations", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.quotation.write", "quotations", row.ID, nil, before, row)
}

// DeleteQuotation removes one row of quotations.
func (s *Service) DeleteQuotation(ctx context.Context, id uuid.UUID) error {
	before, err := s.GetQuotation(ctx, id)
	if err != nil {
		return err
	}
	if err := s.q(ctx).DeleteQuotation(ctx, id); err != nil {
		return wrap(err, "delete quotations")
	}
	return s.audit(ctx, "quotations.delete", "quotations", id, nil, before, nil)
}

// ListPurchaseOrder returns a page of purchase orders.
func (s *Service) ListPurchaseOrder(ctx context.Context, p db.ListPurchaseOrderParams) ([]db.PurchaseOrder, error) {
	return wrap2(s.q(ctx).ListPurchaseOrder(ctx, p))("list purchase orders")
}

// GetPurchaseOrder returns one row of purchase orders.
func (s *Service) GetPurchaseOrder(ctx context.Context, id uuid.UUID) (db.PurchaseOrder, error) {
	row, err := s.q(ctx).GetPurchaseOrder(ctx, id)
	if err != nil {
		return row, notFound(err, "purchase orders")
	}
	return row, nil
}

// CreatePurchaseOrder adds one row to purchase orders.
func (s *Service) CreatePurchaseOrder(ctx context.Context, p db.CreatePurchaseOrderParams) (db.PurchaseOrder, error) {
	var zero db.PurchaseOrder
	if err := s.guard(ctx, "CreatePurchaseOrder", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "PO", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePurchaseOrder(ctx, p)
	if err != nil {
		return zero, wrap(err, "create purchase orders")
	}
	if err := s.audit(ctx, "procurement.purchase_order.write", "purchase_orders", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreatePurchaseOrder", row)
}

// UpdatePurchaseOrder changes one row of purchase orders.
func (s *Service) UpdatePurchaseOrder(ctx context.Context, p db.UpdatePurchaseOrderParams) (db.PurchaseOrder, error) {
	before, err := s.GetPurchaseOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePurchaseOrder(ctx, p)
	if err != nil {
		return before, wrap(err, "update purchase orders")
	}
	if err := s.voidSignatures(ctx, "purchase_orders", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.purchase_order.write", "purchase_orders", row.ID, nil, before, row)
}

// SubmitPurchaseOrder is the submit action on purchase orders.
func (s *Service) SubmitPurchaseOrder(ctx context.Context, p db.ActSubmitPurchaseOrderParams) (db.PurchaseOrder, error) {
	before, err := s.GetPurchaseOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.purchase_order.submit", "purchase_orders", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitPurchaseOrder", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitPurchaseOrder(ctx, p)
	if err != nil {
		return before, stateError(err, "submit purchase orders")
	}
	if err := s.audit(ctx, "procurement.purchase_order.submit", "purchase_orders", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "purchase_orders", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitPurchaseOrder", row)
}

// IssuePurchaseOrder is the issue to supplier action on purchase orders.
func (s *Service) IssuePurchaseOrder(ctx context.Context, p db.ActIssuePurchaseOrderParams) (db.PurchaseOrder, error) {
	before, err := s.GetPurchaseOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.purchase_order.issue", "purchase_orders", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "IssuePurchaseOrder", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "purchase_orders", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActIssuePurchaseOrder(ctx, p)
	if err != nil {
		return before, stateError(err, "issue purchase orders")
	}
	if err := s.audit(ctx, "procurement.purchase_order.issue", "purchase_orders", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "IssuePurchaseOrder", row)
}

// CancelPurchaseOrder is the cancel action on purchase orders.
func (s *Service) CancelPurchaseOrder(ctx context.Context, p db.ActCancelPurchaseOrderParams) (db.PurchaseOrder, error) {
	before, err := s.GetPurchaseOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.purchase_order.cancel", "purchase_orders", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CancelPurchaseOrder", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActCancelPurchaseOrder(ctx, p)
	if err != nil {
		return before, stateError(err, "cancel purchase orders")
	}
	if err := s.audit(ctx, "procurement.purchase_order.cancel", "purchase_orders", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CancelPurchaseOrder", row)
}

// PurchaseOrderChildren returns the rows of a record's child tables, by API field.
func (s *Service) PurchaseOrderChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListPurchaseOrderLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list purchase_order_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ListGoodsReceipt returns a page of goods receipts.
func (s *Service) ListGoodsReceipt(ctx context.Context, p db.ListGoodsReceiptParams) ([]db.GoodsReceipt, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListGoodsReceipt(ctx, p))("list goods receipts")
}

// GetGoodsReceipt returns one row of goods receipts.
func (s *Service) GetGoodsReceipt(ctx context.Context, id uuid.UUID) (db.GoodsReceipt, error) {
	row, err := s.q(ctx).GetGoodsReceipt(ctx, id)
	if err != nil {
		return row, notFound(err, "goods receipts")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateGoodsReceipt adds one row to goods receipts.
func (s *Service) CreateGoodsReceipt(ctx context.Context, p db.CreateGoodsReceiptParams) (db.GoodsReceipt, error) {
	var zero db.GoodsReceipt
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateGoodsReceipt", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "GRV", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateGoodsReceipt(ctx, p)
	if err != nil {
		return zero, wrap(err, "create goods receipts")
	}
	if err := s.audit(ctx, "procurement.goods_receipt.write", "goods_receipts", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateGoodsReceipt", row)
}

// UpdateGoodsReceipt changes one row of goods receipts.
func (s *Service) UpdateGoodsReceipt(ctx context.Context, p db.UpdateGoodsReceiptParams) (db.GoodsReceipt, error) {
	before, err := s.GetGoodsReceipt(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateGoodsReceipt(ctx, p)
	if err != nil {
		return before, wrap(err, "update goods receipts")
	}
	if err := s.voidSignatures(ctx, "goods_receipts", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.goods_receipt.write", "goods_receipts", row.ID, &row.SiteID, before, row)
}

// SubmitGoodsReceipt is the submit action on goods receipts.
func (s *Service) SubmitGoodsReceipt(ctx context.Context, p db.ActSubmitGoodsReceiptParams) (db.GoodsReceipt, error) {
	before, err := s.GetGoodsReceipt(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.goods_receipt.submit", "goods_receipts", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitGoodsReceipt", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitGoodsReceipt(ctx, p)
	if err != nil {
		return before, stateError(err, "submit goods receipts")
	}
	if err := s.audit(ctx, "procurement.goods_receipt.submit", "goods_receipts", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitGoodsReceipt", row)
}

// GoodsReceiptChildren returns the rows of a record's child tables, by API field.
func (s *Service) GoodsReceiptChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListGoodsReceiptLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list goods_receipt_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplaceGoodsReceiptLines swaps a record's goods receipt lines for the given rows.
func (s *Service) ReplaceGoodsReceiptLines(ctx context.Context, id uuid.UUID, rows []db.InsertGoodsReceiptLinesParams) error {
	if err := s.q(ctx).DeleteGoodsReceiptLines(ctx, id); err != nil {
		return wrap(err, "clear goods_receipt_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertGoodsReceiptLines(ctx, row); err != nil {
			return wrap(err, "add goods_receipt_lines")
		}
	}
	return nil
}

// ListDeliveryDiscrepancy returns a page of delivery discrepancies.
func (s *Service) ListDeliveryDiscrepancy(ctx context.Context, p db.ListDeliveryDiscrepancyParams) ([]db.DeliveryDiscrepancy, error) {
	return wrap2(s.q(ctx).ListDeliveryDiscrepancy(ctx, p))("list delivery discrepancies")
}

// GetDeliveryDiscrepancy returns one row of delivery discrepancies.
func (s *Service) GetDeliveryDiscrepancy(ctx context.Context, id uuid.UUID) (db.DeliveryDiscrepancy, error) {
	row, err := s.q(ctx).GetDeliveryDiscrepancy(ctx, id)
	if err != nil {
		return row, notFound(err, "delivery discrepancies")
	}
	return row, nil
}

// CreateDeliveryDiscrepancy adds one row to delivery discrepancies.
func (s *Service) CreateDeliveryDiscrepancy(ctx context.Context, p db.CreateDeliveryDiscrepancyParams) (db.DeliveryDiscrepancy, error) {
	var zero db.DeliveryDiscrepancy
	if err := s.guard(ctx, "CreateDeliveryDiscrepancy", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "DD", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDeliveryDiscrepancy(ctx, p)
	if err != nil {
		return zero, wrap(err, "create delivery discrepancies")
	}
	if err := s.audit(ctx, "procurement.delivery_discrepancy.write", "delivery_discrepancies", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDeliveryDiscrepancy", row)
}

// UpdateDeliveryDiscrepancy changes one row of delivery discrepancies.
func (s *Service) UpdateDeliveryDiscrepancy(ctx context.Context, p db.UpdateDeliveryDiscrepancyParams) (db.DeliveryDiscrepancy, error) {
	before, err := s.GetDeliveryDiscrepancy(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDeliveryDiscrepancy(ctx, p)
	if err != nil {
		return before, wrap(err, "update delivery discrepancies")
	}
	if err := s.voidSignatures(ctx, "delivery_discrepancies", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.delivery_discrepancy.write", "delivery_discrepancies", row.ID, nil, before, row)
}

// CloseDeliveryDiscrepancy is the close action on delivery discrepancies.
func (s *Service) CloseDeliveryDiscrepancy(ctx context.Context, p db.ActCloseDeliveryDiscrepancyParams) (db.DeliveryDiscrepancy, error) {
	before, err := s.GetDeliveryDiscrepancy(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.delivery_discrepancy.close", "delivery_discrepancies", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseDeliveryDiscrepancy", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "delivery_discrepancies", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseDeliveryDiscrepancy(ctx, p)
	if err != nil {
		return before, stateError(err, "close delivery discrepancies")
	}
	if err := s.audit(ctx, "procurement.delivery_discrepancy.close", "delivery_discrepancies", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseDeliveryDiscrepancy", row)
}

// ListSupplierInvoice returns a page of supplier invoices.
func (s *Service) ListSupplierInvoice(ctx context.Context, p db.ListSupplierInvoiceParams) ([]db.SupplierInvoice, error) {
	return wrap2(s.q(ctx).ListSupplierInvoice(ctx, p))("list supplier invoices")
}

// GetSupplierInvoice returns one row of supplier invoices.
func (s *Service) GetSupplierInvoice(ctx context.Context, id uuid.UUID) (db.SupplierInvoice, error) {
	row, err := s.q(ctx).GetSupplierInvoice(ctx, id)
	if err != nil {
		return row, notFound(err, "supplier invoices")
	}
	return row, nil
}

// CreateSupplierInvoice adds one row to supplier invoices.
func (s *Service) CreateSupplierInvoice(ctx context.Context, p db.CreateSupplierInvoiceParams) (db.SupplierInvoice, error) {
	var zero db.SupplierInvoice
	if err := s.guard(ctx, "CreateSupplierInvoice", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "INV", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateSupplierInvoice(ctx, p)
	if err != nil {
		return zero, wrap(err, "create supplier invoices")
	}
	if err := s.audit(ctx, "procurement.supplier_invoice.write", "supplier_invoices", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateSupplierInvoice", row)
}

// UpdateSupplierInvoice changes one row of supplier invoices.
func (s *Service) UpdateSupplierInvoice(ctx context.Context, p db.UpdateSupplierInvoiceParams) (db.SupplierInvoice, error) {
	before, err := s.GetSupplierInvoice(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateSupplierInvoice(ctx, p)
	if err != nil {
		return before, wrap(err, "update supplier invoices")
	}
	if err := s.voidSignatures(ctx, "supplier_invoices", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.supplier_invoice.write", "supplier_invoices", row.ID, nil, before, row)
}

// HoldSupplierInvoice is the place on hold action on supplier invoices.
func (s *Service) HoldSupplierInvoice(ctx context.Context, p db.ActHoldSupplierInvoiceParams) (db.SupplierInvoice, error) {
	before, err := s.GetSupplierInvoice(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.supplier_invoice.hold", "supplier_invoices", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "HoldSupplierInvoice", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActHoldSupplierInvoice(ctx, p)
	if err != nil {
		return before, stateError(err, "hold supplier invoices")
	}
	if err := s.audit(ctx, "procurement.supplier_invoice.hold", "supplier_invoices", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "HoldSupplierInvoice", row)
}

// ApproveExceptionSupplierInvoice is the record an authorised exception action on supplier invoices.
func (s *Service) ApproveExceptionSupplierInvoice(ctx context.Context, p db.ActApproveExceptionSupplierInvoiceParams) (db.SupplierInvoice, error) {
	before, err := s.GetSupplierInvoice(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.supplier_invoice.approve_exception", "supplier_invoices", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ApproveExceptionSupplierInvoice", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActApproveExceptionSupplierInvoice(ctx, p)
	if err != nil {
		return before, stateError(err, "approve exception supplier invoices")
	}
	if err := s.audit(ctx, "procurement.supplier_invoice.approve_exception", "supplier_invoices", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ApproveExceptionSupplierInvoice", row)
}

// PostSupplierInvoice is the post to sap business one action on supplier invoices.
func (s *Service) PostSupplierInvoice(ctx context.Context, p db.ActPostSupplierInvoiceParams) (db.SupplierInvoice, error) {
	before, err := s.GetSupplierInvoice(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.supplier_invoice.post", "supplier_invoices", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "PostSupplierInvoice", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "supplier_invoices", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActPostSupplierInvoice(ctx, p)
	if err != nil {
		return before, stateError(err, "post supplier invoices")
	}
	if err := s.audit(ctx, "procurement.supplier_invoice.post", "supplier_invoices", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "PostSupplierInvoice", row)
}

// SupplierInvoiceChildren returns the rows of a record's child tables, by API field.
func (s *Service) SupplierInvoiceChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListSupplierInvoiceLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list supplier_invoice_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplaceSupplierInvoiceLines swaps a record's supplier invoice lines for the given rows.
func (s *Service) ReplaceSupplierInvoiceLines(ctx context.Context, id uuid.UUID, rows []db.InsertSupplierInvoiceLinesParams) error {
	if err := s.q(ctx).DeleteSupplierInvoiceLines(ctx, id); err != nil {
		return wrap(err, "clear supplier_invoice_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertSupplierInvoiceLines(ctx, row); err != nil {
			return wrap(err, "add supplier_invoice_lines")
		}
	}
	return nil
}

// ListPaymentRequisition returns a page of payment requisitions.
func (s *Service) ListPaymentRequisition(ctx context.Context, p db.ListPaymentRequisitionParams) ([]db.PaymentRequisition, error) {
	return wrap2(s.q(ctx).ListPaymentRequisition(ctx, p))("list payment requisitions")
}

// GetPaymentRequisition returns one row of payment requisitions.
func (s *Service) GetPaymentRequisition(ctx context.Context, id uuid.UUID) (db.PaymentRequisition, error) {
	row, err := s.q(ctx).GetPaymentRequisition(ctx, id)
	if err != nil {
		return row, notFound(err, "payment requisitions")
	}
	return row, nil
}

// CreatePaymentRequisition adds one row to payment requisitions.
func (s *Service) CreatePaymentRequisition(ctx context.Context, p db.CreatePaymentRequisitionParams) (db.PaymentRequisition, error) {
	var zero db.PaymentRequisition
	if err := s.guard(ctx, "CreatePaymentRequisition", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "PAY", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePaymentRequisition(ctx, p)
	if err != nil {
		return zero, wrap(err, "create payment requisitions")
	}
	if err := s.audit(ctx, "procurement.payment_requisition.write", "payment_requisitions", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreatePaymentRequisition", row)
}

// UpdatePaymentRequisition changes one row of payment requisitions.
func (s *Service) UpdatePaymentRequisition(ctx context.Context, p db.UpdatePaymentRequisitionParams) (db.PaymentRequisition, error) {
	before, err := s.GetPaymentRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePaymentRequisition(ctx, p)
	if err != nil {
		return before, wrap(err, "update payment requisitions")
	}
	if err := s.voidSignatures(ctx, "payment_requisitions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.payment_requisition.write", "payment_requisitions", row.ID, nil, before, row)
}

// SubmitPaymentRequisition is the submit action on payment requisitions.
func (s *Service) SubmitPaymentRequisition(ctx context.Context, p db.ActSubmitPaymentRequisitionParams) (db.PaymentRequisition, error) {
	before, err := s.GetPaymentRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.payment_requisition.submit", "payment_requisitions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitPaymentRequisition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitPaymentRequisition(ctx, p)
	if err != nil {
		return before, stateError(err, "submit payment requisitions")
	}
	if err := s.audit(ctx, "procurement.payment_requisition.submit", "payment_requisitions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "payment_requisitions", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitPaymentRequisition", row)
}

// ListPaymentBatch returns a page of payment batches.
func (s *Service) ListPaymentBatch(ctx context.Context, p db.ListPaymentBatchParams) ([]db.PaymentBatch, error) {
	return wrap2(s.q(ctx).ListPaymentBatch(ctx, p))("list payment batches")
}

// GetPaymentBatch returns one row of payment batches.
func (s *Service) GetPaymentBatch(ctx context.Context, id uuid.UUID) (db.PaymentBatch, error) {
	row, err := s.q(ctx).GetPaymentBatch(ctx, id)
	if err != nil {
		return row, notFound(err, "payment batches")
	}
	return row, nil
}

// CreatePaymentBatch adds one row to payment batches.
func (s *Service) CreatePaymentBatch(ctx context.Context, p db.CreatePaymentBatchParams) (db.PaymentBatch, error) {
	var zero db.PaymentBatch
	if err := s.guard(ctx, "CreatePaymentBatch", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "PBT", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePaymentBatch(ctx, p)
	if err != nil {
		return zero, wrap(err, "create payment batches")
	}
	if err := s.audit(ctx, "procurement.payment_batch.write", "payment_batches", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreatePaymentBatch", row)
}

// UpdatePaymentBatch changes one row of payment batches.
func (s *Service) UpdatePaymentBatch(ctx context.Context, p db.UpdatePaymentBatchParams) (db.PaymentBatch, error) {
	before, err := s.GetPaymentBatch(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePaymentBatch(ctx, p)
	if err != nil {
		return before, wrap(err, "update payment batches")
	}
	if err := s.voidSignatures(ctx, "payment_batches", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.payment_batch.write", "payment_batches", row.ID, nil, before, row)
}

// SubmitPaymentBatch is the submit action on payment batches.
func (s *Service) SubmitPaymentBatch(ctx context.Context, p db.ActSubmitPaymentBatchParams) (db.PaymentBatch, error) {
	before, err := s.GetPaymentBatch(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.payment_batch.submit", "payment_batches", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitPaymentBatch", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitPaymentBatch(ctx, p)
	if err != nil {
		return before, stateError(err, "submit payment batches")
	}
	if err := s.audit(ctx, "procurement.payment_batch.submit", "payment_batches", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "payment_batches", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitPaymentBatch", row)
}

// RecordExecutionPaymentBatch is the record payment made in the bank portal action on payment batches.
func (s *Service) RecordExecutionPaymentBatch(ctx context.Context, p db.ActRecordExecutionPaymentBatchParams) (db.PaymentBatch, error) {
	before, err := s.GetPaymentBatch(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.payment_batch.record_execution", "payment_batches", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RecordExecutionPaymentBatch", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRecordExecutionPaymentBatch(ctx, p)
	if err != nil {
		return before, stateError(err, "record execution payment batches")
	}
	if err := s.audit(ctx, "procurement.payment_batch.record_execution", "payment_batches", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RecordExecutionPaymentBatch", row)
}

// ListEmergencyProcurement returns a page of emergency procurements.
func (s *Service) ListEmergencyProcurement(ctx context.Context, p db.ListEmergencyProcurementParams) ([]db.EmergencyProcurement, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListEmergencyProcurement(ctx, p))("list emergency procurements")
}

// GetEmergencyProcurement returns one row of emergency procurements.
func (s *Service) GetEmergencyProcurement(ctx context.Context, id uuid.UUID) (db.EmergencyProcurement, error) {
	row, err := s.q(ctx).GetEmergencyProcurement(ctx, id)
	if err != nil {
		return row, notFound(err, "emergency procurements")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateEmergencyProcurement adds one row to emergency procurements.
func (s *Service) CreateEmergencyProcurement(ctx context.Context, p db.CreateEmergencyProcurementParams) (db.EmergencyProcurement, error) {
	var zero db.EmergencyProcurement
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateEmergencyProcurement", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "EPR", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateEmergencyProcurement(ctx, p)
	if err != nil {
		return zero, wrap(err, "create emergency procurements")
	}
	if err := s.audit(ctx, "procurement.emergency_procurement.write", "emergency_procurements", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateEmergencyProcurement", row)
}

// UpdateEmergencyProcurement changes one row of emergency procurements.
func (s *Service) UpdateEmergencyProcurement(ctx context.Context, p db.UpdateEmergencyProcurementParams) (db.EmergencyProcurement, error) {
	before, err := s.GetEmergencyProcurement(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateEmergencyProcurement(ctx, p)
	if err != nil {
		return before, wrap(err, "update emergency procurements")
	}
	if err := s.voidSignatures(ctx, "emergency_procurements", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.emergency_procurement.write", "emergency_procurements", row.ID, &row.SiteID, before, row)
}

// SubmitEmergencyProcurement is the submit action on emergency procurements.
func (s *Service) SubmitEmergencyProcurement(ctx context.Context, p db.ActSubmitEmergencyProcurementParams) (db.EmergencyProcurement, error) {
	before, err := s.GetEmergencyProcurement(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.emergency_procurement.submit", "emergency_procurements", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitEmergencyProcurement", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitEmergencyProcurement(ctx, p)
	if err != nil {
		return before, stateError(err, "submit emergency procurements")
	}
	if err := s.audit(ctx, "procurement.emergency_procurement.submit", "emergency_procurements", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "emergency_procurements", row.ID, &row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitEmergencyProcurement", row)
}

// ReviewEmergencyProcurement is the record post-purchase review action on emergency procurements.
func (s *Service) ReviewEmergencyProcurement(ctx context.Context, p db.ActReviewEmergencyProcurementParams) (db.EmergencyProcurement, error) {
	before, err := s.GetEmergencyProcurement(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.emergency_procurement.review", "emergency_procurements", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReviewEmergencyProcurement", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "emergency_procurements", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActReviewEmergencyProcurement(ctx, p)
	if err != nil {
		return before, stateError(err, "review emergency procurements")
	}
	if err := s.audit(ctx, "procurement.emergency_procurement.review", "emergency_procurements", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReviewEmergencyProcurement", row)
}

// ListUnauthorizedPurchase returns a page of unauthorized purchases.
func (s *Service) ListUnauthorizedPurchase(ctx context.Context, p db.ListUnauthorizedPurchaseParams) ([]db.UnauthorizedPurchase, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListUnauthorizedPurchase(ctx, p))("list unauthorized purchases")
}

// GetUnauthorizedPurchase returns one row of unauthorized purchases.
func (s *Service) GetUnauthorizedPurchase(ctx context.Context, id uuid.UUID) (db.UnauthorizedPurchase, error) {
	row, err := s.q(ctx).GetUnauthorizedPurchase(ctx, id)
	if err != nil {
		return row, notFound(err, "unauthorized purchases")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateUnauthorizedPurchase adds one row to unauthorized purchases.
func (s *Service) CreateUnauthorizedPurchase(ctx context.Context, p db.CreateUnauthorizedPurchaseParams) (db.UnauthorizedPurchase, error) {
	var zero db.UnauthorizedPurchase
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateUnauthorizedPurchase", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "UAP", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateUnauthorizedPurchase(ctx, p)
	if err != nil {
		return zero, wrap(err, "create unauthorized purchases")
	}
	if err := s.audit(ctx, "procurement.unauthorized_purchase.write", "unauthorized_purchases", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateUnauthorizedPurchase", row)
}

// UpdateUnauthorizedPurchase changes one row of unauthorized purchases.
func (s *Service) UpdateUnauthorizedPurchase(ctx context.Context, p db.UpdateUnauthorizedPurchaseParams) (db.UnauthorizedPurchase, error) {
	before, err := s.GetUnauthorizedPurchase(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateUnauthorizedPurchase(ctx, p)
	if err != nil {
		return before, wrap(err, "update unauthorized purchases")
	}
	if err := s.voidSignatures(ctx, "unauthorized_purchases", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "procurement.unauthorized_purchase.write", "unauthorized_purchases", row.ID, &row.SiteID, before, row)
}

// CloseUnauthorizedPurchase is the close action on unauthorized purchases.
func (s *Service) CloseUnauthorizedPurchase(ctx context.Context, p db.ActCloseUnauthorizedPurchaseParams) (db.UnauthorizedPurchase, error) {
	before, err := s.GetUnauthorizedPurchase(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "procurement.unauthorized_purchase.close", "unauthorized_purchases", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseUnauthorizedPurchase", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "unauthorized_purchases", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseUnauthorizedPurchase(ctx, p)
	if err != nil {
		return before, stateError(err, "close unauthorized purchases")
	}
	if err := s.audit(ctx, "procurement.unauthorized_purchase.close", "unauthorized_purchases", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseUnauthorizedPurchase", row)
}

// ListBudget returns a page of budgets.
func (s *Service) ListBudget(ctx context.Context, p db.ListBudgetParams) ([]db.Budget, error) {
	return wrap2(s.q(ctx).ListBudget(ctx, p))("list budgets")
}

// GetBudget returns one row of budgets.
func (s *Service) GetBudget(ctx context.Context, id uuid.UUID) (db.Budget, error) {
	row, err := s.q(ctx).GetBudget(ctx, id)
	if err != nil {
		return row, notFound(err, "budgets")
	}
	return row, nil
}

// CreateBudget adds one row to budgets.
func (s *Service) CreateBudget(ctx context.Context, p db.CreateBudgetParams) (db.Budget, error) {
	var zero db.Budget
	if err := s.guard(ctx, "CreateBudget", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "BUD", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateBudget(ctx, p)
	if err != nil {
		return zero, wrap(err, "create budgets")
	}
	if err := s.audit(ctx, "budgets.budget.write", "budgets", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateBudget", row)
}

// UpdateBudget changes one row of budgets.
func (s *Service) UpdateBudget(ctx context.Context, p db.UpdateBudgetParams) (db.Budget, error) {
	before, err := s.GetBudget(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateBudget(ctx, p)
	if err != nil {
		return before, wrap(err, "update budgets")
	}
	if err := s.voidSignatures(ctx, "budgets", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "budgets.budget.write", "budgets", row.ID, nil, before, row)
}

// SubmitBudget is the submit action on budgets.
func (s *Service) SubmitBudget(ctx context.Context, p db.ActSubmitBudgetParams) (db.Budget, error) {
	before, err := s.GetBudget(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "budgets.budget.submit", "budgets", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitBudget", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitBudget(ctx, p)
	if err != nil {
		return before, stateError(err, "submit budgets")
	}
	if err := s.audit(ctx, "budgets.budget.submit", "budgets", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "budgets", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitBudget", row)
}

// BudgetChildren returns the rows of a record's child tables, by API field.
func (s *Service) BudgetChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListBudgetLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list budget_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplaceBudgetLines swaps a record's budget lines for the given rows.
func (s *Service) ReplaceBudgetLines(ctx context.Context, id uuid.UUID, rows []db.InsertBudgetLinesParams) error {
	if err := s.q(ctx).DeleteBudgetLines(ctx, id); err != nil {
		return wrap(err, "clear budget_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertBudgetLines(ctx, row); err != nil {
			return wrap(err, "add budget_lines")
		}
	}
	return nil
}

// ListBudgetVariance returns a page of budget variances.
func (s *Service) ListBudgetVariance(ctx context.Context, p db.ListBudgetVarianceParams) ([]db.BudgetVariance, error) {
	return wrap2(s.q(ctx).ListBudgetVariance(ctx, p))("list budget variances")
}

// GetBudgetVariance returns one row of budget variances.
func (s *Service) GetBudgetVariance(ctx context.Context, id uuid.UUID) (db.BudgetVariance, error) {
	row, err := s.q(ctx).GetBudgetVariance(ctx, id)
	if err != nil {
		return row, notFound(err, "budget variances")
	}
	return row, nil
}

// UpdateBudgetVariance changes one row of budget variances.
func (s *Service) UpdateBudgetVariance(ctx context.Context, p db.UpdateBudgetVarianceParams) (db.BudgetVariance, error) {
	before, err := s.GetBudgetVariance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateBudgetVariance(ctx, p)
	if err != nil {
		return before, wrap(err, "update budget variances")
	}
	if err := s.voidSignatures(ctx, "budget_variances", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "budgets.budget_variance.write", "budget_variances", row.ID, nil, before, row)
}

// ListBudgetReallocation returns a page of budget reallocations.
func (s *Service) ListBudgetReallocation(ctx context.Context, p db.ListBudgetReallocationParams) ([]db.BudgetReallocation, error) {
	return wrap2(s.q(ctx).ListBudgetReallocation(ctx, p))("list budget reallocations")
}

// GetBudgetReallocation returns one row of budget reallocations.
func (s *Service) GetBudgetReallocation(ctx context.Context, id uuid.UUID) (db.BudgetReallocation, error) {
	row, err := s.q(ctx).GetBudgetReallocation(ctx, id)
	if err != nil {
		return row, notFound(err, "budget reallocations")
	}
	return row, nil
}

// CreateBudgetReallocation adds one row to budget reallocations.
func (s *Service) CreateBudgetReallocation(ctx context.Context, p db.CreateBudgetReallocationParams) (db.BudgetReallocation, error) {
	var zero db.BudgetReallocation
	if err := s.guard(ctx, "CreateBudgetReallocation", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "BRA", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateBudgetReallocation(ctx, p)
	if err != nil {
		return zero, wrap(err, "create budget reallocations")
	}
	if err := s.audit(ctx, "budgets.budget_reallocation.write", "budget_reallocations", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateBudgetReallocation", row)
}

// SubmitBudgetReallocation is the submit action on budget reallocations.
func (s *Service) SubmitBudgetReallocation(ctx context.Context, p db.ActSubmitBudgetReallocationParams) (db.BudgetReallocation, error) {
	before, err := s.GetBudgetReallocation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "budgets.budget_reallocation.submit", "budget_reallocations", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitBudgetReallocation", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitBudgetReallocation(ctx, p)
	if err != nil {
		return before, stateError(err, "submit budget reallocations")
	}
	if err := s.audit(ctx, "budgets.budget_reallocation.submit", "budget_reallocations", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "budget_reallocations", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitBudgetReallocation", row)
}
