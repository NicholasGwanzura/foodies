// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListItem returns a page of items.
func (s *Service) ListItem(ctx context.Context, p db.ListItemParams) ([]db.Item, error) {
	return wrap2(s.q(ctx).ListItem(ctx, p))("list items")
}

// GetItem returns one row of items.
func (s *Service) GetItem(ctx context.Context, id uuid.UUID) (db.Item, error) {
	row, err := s.q(ctx).GetItem(ctx, id)
	if err != nil {
		return row, notFound(err, "items")
	}
	return row, nil
}

// CreateItem adds one row to items.
func (s *Service) CreateItem(ctx context.Context, p db.CreateItemParams) (db.Item, error) {
	var zero db.Item
	if err := s.guard(ctx, "CreateItem", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateItem(ctx, p)
	if err != nil {
		return zero, wrap(err, "create items")
	}
	if err := s.audit(ctx, "stock.item.write", "items", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateItem", row)
}

// UpdateItem changes one row of items.
func (s *Service) UpdateItem(ctx context.Context, p db.UpdateItemParams) (db.Item, error) {
	before, err := s.GetItem(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateItem(ctx, p)
	if err != nil {
		return before, wrap(err, "update items")
	}
	if err := s.voidSignatures(ctx, "items", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "stock.item.write", "items", row.ID, nil, before, row)
}

// ListStockBalance returns a page of stock balances.
func (s *Service) ListStockBalance(ctx context.Context, p db.ListStockBalanceParams) ([]db.StockBalance, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListStockBalance(ctx, p))("list stock balances")
}

// GetStockBalance returns one row of stock balances.
func (s *Service) GetStockBalance(ctx context.Context, id uuid.UUID) (db.StockBalance, error) {
	row, err := s.q(ctx).GetStockBalance(ctx, id)
	if err != nil {
		return row, notFound(err, "stock balances")
	}
	return row, inScope(ctx, row.SiteID)
}

// ListStockRequisition returns a page of stock requisitions.
func (s *Service) ListStockRequisition(ctx context.Context, p db.ListStockRequisitionParams) ([]db.StockRequisition, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListStockRequisition(ctx, p))("list stock requisitions")
}

// GetStockRequisition returns one row of stock requisitions.
func (s *Service) GetStockRequisition(ctx context.Context, id uuid.UUID) (db.StockRequisition, error) {
	row, err := s.q(ctx).GetStockRequisition(ctx, id)
	if err != nil {
		return row, notFound(err, "stock requisitions")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateStockRequisition adds one row to stock requisitions.
func (s *Service) CreateStockRequisition(ctx context.Context, p db.CreateStockRequisitionParams) (db.StockRequisition, error) {
	var zero db.StockRequisition
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateStockRequisition", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "SRQ", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateStockRequisition(ctx, p)
	if err != nil {
		return zero, wrap(err, "create stock requisitions")
	}
	if err := s.audit(ctx, "stock.stock_requisition.write", "stock_requisitions", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateStockRequisition", row)
}

// UpdateStockRequisition changes one row of stock requisitions.
func (s *Service) UpdateStockRequisition(ctx context.Context, p db.UpdateStockRequisitionParams) (db.StockRequisition, error) {
	before, err := s.GetStockRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateStockRequisition(ctx, p)
	if err != nil {
		return before, wrap(err, "update stock requisitions")
	}
	if err := s.voidSignatures(ctx, "stock_requisitions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "stock.stock_requisition.write", "stock_requisitions", row.ID, &row.SiteID, before, row)
}

// SubmitStockRequisition is the submit action on stock requisitions.
func (s *Service) SubmitStockRequisition(ctx context.Context, p db.ActSubmitStockRequisitionParams) (db.StockRequisition, error) {
	before, err := s.GetStockRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_requisition.submit", "stock_requisitions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitStockRequisition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitStockRequisition(ctx, p)
	if err != nil {
		return before, stateError(err, "submit stock requisitions")
	}
	if err := s.audit(ctx, "stock.stock_requisition.submit", "stock_requisitions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitStockRequisition", row)
}

// IssueStockRequisition is the issue stock action on stock requisitions.
func (s *Service) IssueStockRequisition(ctx context.Context, p db.ActIssueStockRequisitionParams) (db.StockRequisition, error) {
	before, err := s.GetStockRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_requisition.issue", "stock_requisitions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "IssueStockRequisition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "stock_requisitions", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActIssueStockRequisition(ctx, p)
	if err != nil {
		return before, stateError(err, "issue stock requisitions")
	}
	if err := s.audit(ctx, "stock.stock_requisition.issue", "stock_requisitions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "IssueStockRequisition", row)
}

// AcknowledgeStockRequisition is the recipient acknowledges action on stock requisitions.
func (s *Service) AcknowledgeStockRequisition(ctx context.Context, p db.ActAcknowledgeStockRequisitionParams) (db.StockRequisition, error) {
	before, err := s.GetStockRequisition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_requisition.acknowledge", "stock_requisitions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "AcknowledgeStockRequisition", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAcknowledgeStockRequisition(ctx, p)
	if err != nil {
		return before, stateError(err, "acknowledge stock requisitions")
	}
	if err := s.audit(ctx, "stock.stock_requisition.acknowledge", "stock_requisitions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "AcknowledgeStockRequisition", row)
}

// StockRequisitionChildren returns the rows of a record's child tables, by API field.
func (s *Service) StockRequisitionChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListStockRequisitionLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list stock_requisition_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplaceStockRequisitionLines swaps a record's stock requisition lines for the given rows.
func (s *Service) ReplaceStockRequisitionLines(ctx context.Context, id uuid.UUID, rows []db.InsertStockRequisitionLinesParams) error {
	if err := s.q(ctx).DeleteStockRequisitionLines(ctx, id); err != nil {
		return wrap(err, "clear stock_requisition_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertStockRequisitionLines(ctx, row); err != nil {
			return wrap(err, "add stock_requisition_lines")
		}
	}
	return nil
}

// ListStockTransfer returns a page of stock transfers.
func (s *Service) ListStockTransfer(ctx context.Context, p db.ListStockTransferParams) ([]db.StockTransfer, error) {
	return wrap2(s.q(ctx).ListStockTransfer(ctx, p))("list stock transfers")
}

// GetStockTransfer returns one row of stock transfers.
func (s *Service) GetStockTransfer(ctx context.Context, id uuid.UUID) (db.StockTransfer, error) {
	row, err := s.q(ctx).GetStockTransfer(ctx, id)
	if err != nil {
		return row, notFound(err, "stock transfers")
	}
	return row, nil
}

// CreateStockTransfer adds one row to stock transfers.
func (s *Service) CreateStockTransfer(ctx context.Context, p db.CreateStockTransferParams) (db.StockTransfer, error) {
	var zero db.StockTransfer
	if err := s.guard(ctx, "CreateStockTransfer", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "TRF", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateStockTransfer(ctx, p)
	if err != nil {
		return zero, wrap(err, "create stock transfers")
	}
	if err := s.audit(ctx, "stock.stock_transfer.write", "stock_transfers", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateStockTransfer", row)
}

// UpdateStockTransfer changes one row of stock transfers.
func (s *Service) UpdateStockTransfer(ctx context.Context, p db.UpdateStockTransferParams) (db.StockTransfer, error) {
	before, err := s.GetStockTransfer(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateStockTransfer(ctx, p)
	if err != nil {
		return before, wrap(err, "update stock transfers")
	}
	if err := s.voidSignatures(ctx, "stock_transfers", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "stock.stock_transfer.write", "stock_transfers", row.ID, nil, before, row)
}

// AcceptStockTransfer is the stock custodian accepts action on stock transfers.
func (s *Service) AcceptStockTransfer(ctx context.Context, p db.ActAcceptStockTransferParams) (db.StockTransfer, error) {
	before, err := s.GetStockTransfer(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_transfer.accept", "stock_transfers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "AcceptStockTransfer", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "stock_transfers", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActAcceptStockTransfer(ctx, p)
	if err != nil {
		return before, stateError(err, "accept stock transfers")
	}
	if err := s.audit(ctx, "stock.stock_transfer.accept", "stock_transfers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "AcceptStockTransfer", row)
}

// RejectStockTransfer is the reject action on stock transfers.
func (s *Service) RejectStockTransfer(ctx context.Context, p db.ActRejectStockTransferParams) (db.StockTransfer, error) {
	before, err := s.GetStockTransfer(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_transfer.reject", "stock_transfers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RejectStockTransfer", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRejectStockTransfer(ctx, p)
	if err != nil {
		return before, stateError(err, "reject stock transfers")
	}
	if err := s.audit(ctx, "stock.stock_transfer.reject", "stock_transfers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RejectStockTransfer", row)
}

// DispatchStockTransfer is the dispatch with security check action on stock transfers.
func (s *Service) DispatchStockTransfer(ctx context.Context, p db.ActDispatchStockTransferParams) (db.StockTransfer, error) {
	before, err := s.GetStockTransfer(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_transfer.dispatch", "stock_transfers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "DispatchStockTransfer", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActDispatchStockTransfer(ctx, p)
	if err != nil {
		return before, stateError(err, "dispatch stock transfers")
	}
	if err := s.audit(ctx, "stock.stock_transfer.dispatch", "stock_transfers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "DispatchStockTransfer", row)
}

// ReceiveStockTransfer is the receive and count action on stock transfers.
func (s *Service) ReceiveStockTransfer(ctx context.Context, p db.ActReceiveStockTransferParams) (db.StockTransfer, error) {
	before, err := s.GetStockTransfer(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_transfer.receive", "stock_transfers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReceiveStockTransfer", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReceiveStockTransfer(ctx, p)
	if err != nil {
		return before, stateError(err, "receive stock transfers")
	}
	if err := s.audit(ctx, "stock.stock_transfer.receive", "stock_transfers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReceiveStockTransfer", row)
}

// CloseStockTransfer is the close action on stock transfers.
func (s *Service) CloseStockTransfer(ctx context.Context, p db.ActCloseStockTransferParams) (db.StockTransfer, error) {
	before, err := s.GetStockTransfer(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_transfer.close", "stock_transfers", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseStockTransfer", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "stock_transfers", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseStockTransfer(ctx, p)
	if err != nil {
		return before, stateError(err, "close stock transfers")
	}
	if err := s.audit(ctx, "stock.stock_transfer.close", "stock_transfers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseStockTransfer", row)
}

// StockTransferChildren returns the rows of a record's child tables, by API field.
func (s *Service) StockTransferChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListStockTransferLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list stock_transfer_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplaceStockTransferLines swaps a record's stock transfer lines for the given rows.
func (s *Service) ReplaceStockTransferLines(ctx context.Context, id uuid.UUID, rows []db.InsertStockTransferLinesParams) error {
	if err := s.q(ctx).DeleteStockTransferLines(ctx, id); err != nil {
		return wrap(err, "clear stock_transfer_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertStockTransferLines(ctx, row); err != nil {
			return wrap(err, "add stock_transfer_lines")
		}
	}
	return nil
}

// ListStockReturn returns a page of stock returns.
func (s *Service) ListStockReturn(ctx context.Context, p db.ListStockReturnParams) ([]db.StockReturn, error) {
	return wrap2(s.q(ctx).ListStockReturn(ctx, p))("list stock returns")
}

// GetStockReturn returns one row of stock returns.
func (s *Service) GetStockReturn(ctx context.Context, id uuid.UUID) (db.StockReturn, error) {
	row, err := s.q(ctx).GetStockReturn(ctx, id)
	if err != nil {
		return row, notFound(err, "stock returns")
	}
	return row, nil
}

// CreateStockReturn adds one row to stock returns.
func (s *Service) CreateStockReturn(ctx context.Context, p db.CreateStockReturnParams) (db.StockReturn, error) {
	var zero db.StockReturn
	if err := s.guard(ctx, "CreateStockReturn", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "RTN", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateStockReturn(ctx, p)
	if err != nil {
		return zero, wrap(err, "create stock returns")
	}
	if err := s.audit(ctx, "stock.stock_return.write", "stock_returns", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateStockReturn", row)
}

// UpdateStockReturn changes one row of stock returns.
func (s *Service) UpdateStockReturn(ctx context.Context, p db.UpdateStockReturnParams) (db.StockReturn, error) {
	before, err := s.GetStockReturn(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateStockReturn(ctx, p)
	if err != nil {
		return before, wrap(err, "update stock returns")
	}
	if err := s.voidSignatures(ctx, "stock_returns", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "stock.stock_return.write", "stock_returns", row.ID, nil, before, row)
}

// CloseStockReturn is the close action on stock returns.
func (s *Service) CloseStockReturn(ctx context.Context, p db.ActCloseStockReturnParams) (db.StockReturn, error) {
	before, err := s.GetStockReturn(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_return.close", "stock_returns", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseStockReturn", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "stock_returns", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseStockReturn(ctx, p)
	if err != nil {
		return before, stateError(err, "close stock returns")
	}
	if err := s.audit(ctx, "stock.stock_return.close", "stock_returns", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseStockReturn", row)
}

// ListStocktake returns a page of stocktakes.
func (s *Service) ListStocktake(ctx context.Context, p db.ListStocktakeParams) ([]db.Stocktake, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListStocktake(ctx, p))("list stocktakes")
}

// GetStocktake returns one row of stocktakes.
func (s *Service) GetStocktake(ctx context.Context, id uuid.UUID) (db.Stocktake, error) {
	row, err := s.q(ctx).GetStocktake(ctx, id)
	if err != nil {
		return row, notFound(err, "stocktakes")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateStocktake adds one row to stocktakes.
func (s *Service) CreateStocktake(ctx context.Context, p db.CreateStocktakeParams) (db.Stocktake, error) {
	var zero db.Stocktake
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateStocktake", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "STK", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateStocktake(ctx, p)
	if err != nil {
		return zero, wrap(err, "create stocktakes")
	}
	if err := s.audit(ctx, "stock.stocktake.write", "stocktakes", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateStocktake", row)
}

// UpdateStocktake changes one row of stocktakes.
func (s *Service) UpdateStocktake(ctx context.Context, p db.UpdateStocktakeParams) (db.Stocktake, error) {
	before, err := s.GetStocktake(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateStocktake(ctx, p)
	if err != nil {
		return before, wrap(err, "update stocktakes")
	}
	if err := s.voidSignatures(ctx, "stocktakes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "stock.stocktake.write", "stocktakes", row.ID, &row.SiteID, before, row)
}

// SubmitStocktake is the submit action on stocktakes.
func (s *Service) SubmitStocktake(ctx context.Context, p db.ActSubmitStocktakeParams) (db.Stocktake, error) {
	before, err := s.GetStocktake(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stocktake.submit", "stocktakes", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitStocktake", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitStocktake(ctx, p)
	if err != nil {
		return before, stateError(err, "submit stocktakes")
	}
	if err := s.audit(ctx, "stock.stocktake.submit", "stocktakes", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitStocktake", row)
}

// ApproveStocktake is the approve action on stocktakes.
func (s *Service) ApproveStocktake(ctx context.Context, p db.ActApproveStocktakeParams) (db.Stocktake, error) {
	before, err := s.GetStocktake(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stocktake.approve", "stocktakes", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ApproveStocktake", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "stocktakes", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActApproveStocktake(ctx, p)
	if err != nil {
		return before, stateError(err, "approve stocktakes")
	}
	if err := s.audit(ctx, "stock.stocktake.approve", "stocktakes", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ApproveStocktake", row)
}

// RequestSkipStocktake is the request approval to skip action on stocktakes.
func (s *Service) RequestSkipStocktake(ctx context.Context, p db.ActRequestSkipStocktakeParams) (db.Stocktake, error) {
	before, err := s.GetStocktake(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stocktake.request_skip", "stocktakes", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RequestSkipStocktake", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRequestSkipStocktake(ctx, p)
	if err != nil {
		return before, stateError(err, "request skip stocktakes")
	}
	if err := s.audit(ctx, "stock.stocktake.request_skip", "stocktakes", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RequestSkipStocktake", row)
}

// StocktakeChildren returns the rows of a record's child tables, by API field.
func (s *Service) StocktakeChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListStocktakeLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list stocktake_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ListStockVariance returns a page of stock variances.
func (s *Service) ListStockVariance(ctx context.Context, p db.ListStockVarianceParams) ([]db.StockVariance, error) {
	return wrap2(s.q(ctx).ListStockVariance(ctx, p))("list stock variances")
}

// GetStockVariance returns one row of stock variances.
func (s *Service) GetStockVariance(ctx context.Context, id uuid.UUID) (db.StockVariance, error) {
	row, err := s.q(ctx).GetStockVariance(ctx, id)
	if err != nil {
		return row, notFound(err, "stock variances")
	}
	return row, nil
}

// UpdateStockVariance changes one row of stock variances.
func (s *Service) UpdateStockVariance(ctx context.Context, p db.UpdateStockVarianceParams) (db.StockVariance, error) {
	before, err := s.GetStockVariance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateStockVariance(ctx, p)
	if err != nil {
		return before, wrap(err, "update stock variances")
	}
	if err := s.voidSignatures(ctx, "stock_variances", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "stock.stock_variance.write", "stock_variances", row.ID, nil, before, row)
}

// CloseStockVariance is the close action on stock variances.
func (s *Service) CloseStockVariance(ctx context.Context, p db.ActCloseStockVarianceParams) (db.StockVariance, error) {
	before, err := s.GetStockVariance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_variance.close", "stock_variances", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CloseStockVariance", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "stock_variances", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseStockVariance(ctx, p)
	if err != nil {
		return before, stateError(err, "close stock variances")
	}
	if err := s.audit(ctx, "stock.stock_variance.close", "stock_variances", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CloseStockVariance", row)
}

// ListStockAdjustment returns a page of stock adjustments.
func (s *Service) ListStockAdjustment(ctx context.Context, p db.ListStockAdjustmentParams) ([]db.StockAdjustment, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListStockAdjustment(ctx, p))("list stock adjustments")
}

// GetStockAdjustment returns one row of stock adjustments.
func (s *Service) GetStockAdjustment(ctx context.Context, id uuid.UUID) (db.StockAdjustment, error) {
	row, err := s.q(ctx).GetStockAdjustment(ctx, id)
	if err != nil {
		return row, notFound(err, "stock adjustments")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateStockAdjustment adds one row to stock adjustments.
func (s *Service) CreateStockAdjustment(ctx context.Context, p db.CreateStockAdjustmentParams) (db.StockAdjustment, error) {
	var zero db.StockAdjustment
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateStockAdjustment", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "ADJ", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateStockAdjustment(ctx, p)
	if err != nil {
		return zero, wrap(err, "create stock adjustments")
	}
	if err := s.audit(ctx, "stock.stock_adjustment.write", "stock_adjustments", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateStockAdjustment", row)
}

// UpdateStockAdjustment changes one row of stock adjustments.
func (s *Service) UpdateStockAdjustment(ctx context.Context, p db.UpdateStockAdjustmentParams) (db.StockAdjustment, error) {
	before, err := s.GetStockAdjustment(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateStockAdjustment(ctx, p)
	if err != nil {
		return before, wrap(err, "update stock adjustments")
	}
	if err := s.voidSignatures(ctx, "stock_adjustments", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "stock.stock_adjustment.write", "stock_adjustments", row.ID, &row.SiteID, before, row)
}

// SubmitStockAdjustment is the submit action on stock adjustments.
func (s *Service) SubmitStockAdjustment(ctx context.Context, p db.ActSubmitStockAdjustmentParams) (db.StockAdjustment, error) {
	before, err := s.GetStockAdjustment(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.stock_adjustment.submit", "stock_adjustments", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitStockAdjustment", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitStockAdjustment(ctx, p)
	if err != nil {
		return before, stateError(err, "submit stock adjustments")
	}
	if err := s.audit(ctx, "stock.stock_adjustment.submit", "stock_adjustments", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "stock_adjustments", row.ID, &row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitStockAdjustment", row)
}

// ListWasteRecord returns a page of waste records.
func (s *Service) ListWasteRecord(ctx context.Context, p db.ListWasteRecordParams) ([]db.WasteRecord, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListWasteRecord(ctx, p))("list waste records")
}

// GetWasteRecord returns one row of waste records.
func (s *Service) GetWasteRecord(ctx context.Context, id uuid.UUID) (db.WasteRecord, error) {
	row, err := s.q(ctx).GetWasteRecord(ctx, id)
	if err != nil {
		return row, notFound(err, "waste records")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateWasteRecord adds one row to waste records.
func (s *Service) CreateWasteRecord(ctx context.Context, p db.CreateWasteRecordParams) (db.WasteRecord, error) {
	var zero db.WasteRecord
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateWasteRecord", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "WST", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateWasteRecord(ctx, p)
	if err != nil {
		return zero, wrap(err, "create waste records")
	}
	if err := s.audit(ctx, "stock.waste_record.write", "waste_records", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateWasteRecord", row)
}

// UpdateWasteRecord changes one row of waste records.
func (s *Service) UpdateWasteRecord(ctx context.Context, p db.UpdateWasteRecordParams) (db.WasteRecord, error) {
	before, err := s.GetWasteRecord(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateWasteRecord(ctx, p)
	if err != nil {
		return before, wrap(err, "update waste records")
	}
	if err := s.voidSignatures(ctx, "waste_records", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "stock.waste_record.write", "waste_records", row.ID, &row.SiteID, before, row)
}

// ReviewWasteRecord is the supervisor review action on waste records.
func (s *Service) ReviewWasteRecord(ctx context.Context, p db.ActReviewWasteRecordParams) (db.WasteRecord, error) {
	before, err := s.GetWasteRecord(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.waste_record.review", "waste_records", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReviewWasteRecord", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "waste_records", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActReviewWasteRecord(ctx, p)
	if err != nil {
		return before, stateError(err, "review waste records")
	}
	if err := s.audit(ctx, "stock.waste_record.review", "waste_records", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReviewWasteRecord", row)
}

// ListDisposal returns a page of disposals.
func (s *Service) ListDisposal(ctx context.Context, p db.ListDisposalParams) ([]db.Disposal, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListDisposal(ctx, p))("list disposals")
}

// GetDisposal returns one row of disposals.
func (s *Service) GetDisposal(ctx context.Context, id uuid.UUID) (db.Disposal, error) {
	row, err := s.q(ctx).GetDisposal(ctx, id)
	if err != nil {
		return row, notFound(err, "disposals")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateDisposal adds one row to disposals.
func (s *Service) CreateDisposal(ctx context.Context, p db.CreateDisposalParams) (db.Disposal, error) {
	var zero db.Disposal
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateDisposal", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "DSP", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDisposal(ctx, p)
	if err != nil {
		return zero, wrap(err, "create disposals")
	}
	if err := s.audit(ctx, "stock.disposal.write", "disposals", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDisposal", row)
}

// UpdateDisposal changes one row of disposals.
func (s *Service) UpdateDisposal(ctx context.Context, p db.UpdateDisposalParams) (db.Disposal, error) {
	before, err := s.GetDisposal(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDisposal(ctx, p)
	if err != nil {
		return before, wrap(err, "update disposals")
	}
	if err := s.voidSignatures(ctx, "disposals", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "stock.disposal.write", "disposals", row.ID, &row.SiteID, before, row)
}

// SubmitDisposal is the submit action on disposals.
func (s *Service) SubmitDisposal(ctx context.Context, p db.ActSubmitDisposalParams) (db.Disposal, error) {
	before, err := s.GetDisposal(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.disposal.submit", "disposals", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitDisposal", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitDisposal(ctx, p)
	if err != nil {
		return before, stateError(err, "submit disposals")
	}
	if err := s.audit(ctx, "stock.disposal.submit", "disposals", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "disposals", row.ID, &row.SiteID, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitDisposal", row)
}

// RecordDisposalDisposal is the record disposal done action on disposals.
func (s *Service) RecordDisposalDisposal(ctx context.Context, p db.ActRecordDisposalDisposalParams) (db.Disposal, error) {
	before, err := s.GetDisposal(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "stock.disposal.record_disposal", "disposals", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RecordDisposalDisposal", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRecordDisposalDisposal(ctx, p)
	if err != nil {
		return before, stateError(err, "record disposal disposals")
	}
	if err := s.audit(ctx, "stock.disposal.record_disposal", "disposals", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RecordDisposalDisposal", row)
}

// ListRecipe returns a page of recipes.
func (s *Service) ListRecipe(ctx context.Context, p db.ListRecipeParams) ([]db.Recipe, error) {
	return wrap2(s.q(ctx).ListRecipe(ctx, p))("list recipes")
}

// GetRecipe returns one row of recipes.
func (s *Service) GetRecipe(ctx context.Context, id uuid.UUID) (db.Recipe, error) {
	row, err := s.q(ctx).GetRecipe(ctx, id)
	if err != nil {
		return row, notFound(err, "recipes")
	}
	return row, nil
}

// CreateRecipe adds one row to recipes.
func (s *Service) CreateRecipe(ctx context.Context, p db.CreateRecipeParams) (db.Recipe, error) {
	var zero db.Recipe
	if err := s.guard(ctx, "CreateRecipe", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRecipe(ctx, p)
	if err != nil {
		return zero, wrap(err, "create recipes")
	}
	if err := s.audit(ctx, "production.recipe.write", "recipes", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateRecipe", row)
}

// UpdateRecipe changes one row of recipes.
func (s *Service) UpdateRecipe(ctx context.Context, p db.UpdateRecipeParams) (db.Recipe, error) {
	before, err := s.GetRecipe(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRecipe(ctx, p)
	if err != nil {
		return before, wrap(err, "update recipes")
	}
	if err := s.voidSignatures(ctx, "recipes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "production.recipe.write", "recipes", row.ID, nil, before, row)
}

// ListRecipeVersion returns a page of recipe versions.
func (s *Service) ListRecipeVersion(ctx context.Context, p db.ListRecipeVersionParams) ([]db.RecipeVersion, error) {
	return wrap2(s.q(ctx).ListRecipeVersion(ctx, p))("list recipe versions")
}

// GetRecipeVersion returns one row of recipe versions.
func (s *Service) GetRecipeVersion(ctx context.Context, id uuid.UUID) (db.RecipeVersion, error) {
	row, err := s.q(ctx).GetRecipeVersion(ctx, id)
	if err != nil {
		return row, notFound(err, "recipe versions")
	}
	return row, nil
}

// CreateRecipeVersion adds one row to recipe versions.
func (s *Service) CreateRecipeVersion(ctx context.Context, p db.CreateRecipeVersionParams) (db.RecipeVersion, error) {
	var zero db.RecipeVersion
	if err := s.guard(ctx, "CreateRecipeVersion", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRecipeVersion(ctx, p)
	if err != nil {
		return zero, wrap(err, "create recipe versions")
	}
	if err := s.audit(ctx, "production.recipe_version.write", "recipe_versions", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateRecipeVersion", row)
}

// UpdateRecipeVersion changes one row of recipe versions.
func (s *Service) UpdateRecipeVersion(ctx context.Context, p db.UpdateRecipeVersionParams) (db.RecipeVersion, error) {
	before, err := s.GetRecipeVersion(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRecipeVersion(ctx, p)
	if err != nil {
		return before, wrap(err, "update recipe versions")
	}
	if err := s.voidSignatures(ctx, "recipe_versions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "production.recipe_version.write", "recipe_versions", row.ID, nil, before, row)
}

// SubmitRecipeVersion is the submit action on recipe versions.
func (s *Service) SubmitRecipeVersion(ctx context.Context, p db.ActSubmitRecipeVersionParams) (db.RecipeVersion, error) {
	before, err := s.GetRecipeVersion(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.recipe_version.submit", "recipe_versions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitRecipeVersion", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitRecipeVersion(ctx, p)
	if err != nil {
		return before, stateError(err, "submit recipe versions")
	}
	if err := s.audit(ctx, "production.recipe_version.submit", "recipe_versions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "recipe_versions", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitRecipeVersion", row)
}

// RecipeVersionChildren returns the rows of a record's child tables, by API field.
func (s *Service) RecipeVersionChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListRecipeIngredients(ctx, id)
		if err != nil {
			return nil, wrap(err, "list recipe_ingredients")
		}
		out["ingredients"] = childRows(rows)
	}
	return out, nil
}

// ReplaceRecipeIngredients swaps a record's recipe ingredients for the given rows.
func (s *Service) ReplaceRecipeIngredients(ctx context.Context, id uuid.UUID, rows []db.InsertRecipeIngredientsParams) error {
	if err := s.q(ctx).DeleteRecipeIngredients(ctx, id); err != nil {
		return wrap(err, "clear recipe_ingredients")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertRecipeIngredients(ctx, row); err != nil {
			return wrap(err, "add recipe_ingredients")
		}
	}
	return nil
}

// ListKitchenOrder returns a page of kitchen orders.
func (s *Service) ListKitchenOrder(ctx context.Context, p db.ListKitchenOrderParams) ([]db.KitchenOrder, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListKitchenOrder(ctx, p))("list kitchen orders")
}

// GetKitchenOrder returns one row of kitchen orders.
func (s *Service) GetKitchenOrder(ctx context.Context, id uuid.UUID) (db.KitchenOrder, error) {
	row, err := s.q(ctx).GetKitchenOrder(ctx, id)
	if err != nil {
		return row, notFound(err, "kitchen orders")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateKitchenOrder adds one row to kitchen orders.
func (s *Service) CreateKitchenOrder(ctx context.Context, p db.CreateKitchenOrderParams) (db.KitchenOrder, error) {
	var zero db.KitchenOrder
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateKitchenOrder", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "CKO", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateKitchenOrder(ctx, p)
	if err != nil {
		return zero, wrap(err, "create kitchen orders")
	}
	if err := s.audit(ctx, "production.kitchen_order.write", "kitchen_orders", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateKitchenOrder", row)
}

// UpdateKitchenOrder changes one row of kitchen orders.
func (s *Service) UpdateKitchenOrder(ctx context.Context, p db.UpdateKitchenOrderParams) (db.KitchenOrder, error) {
	before, err := s.GetKitchenOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateKitchenOrder(ctx, p)
	if err != nil {
		return before, wrap(err, "update kitchen orders")
	}
	if err := s.voidSignatures(ctx, "kitchen_orders", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "production.kitchen_order.write", "kitchen_orders", row.ID, &row.SiteID, before, row)
}

// AcknowledgeLateKitchenOrder is the acknowledge a late order action on kitchen orders.
func (s *Service) AcknowledgeLateKitchenOrder(ctx context.Context, p db.ActAcknowledgeLateKitchenOrderParams) (db.KitchenOrder, error) {
	before, err := s.GetKitchenOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.kitchen_order.acknowledge_late", "kitchen_orders", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "AcknowledgeLateKitchenOrder", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAcknowledgeLateKitchenOrder(ctx, p)
	if err != nil {
		return before, stateError(err, "acknowledge late kitchen orders")
	}
	if err := s.audit(ctx, "production.kitchen_order.acknowledge_late", "kitchen_orders", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "AcknowledgeLateKitchenOrder", row)
}

// ConfirmKitchenOrder is the central kitchen confirms flagged quantities action on kitchen orders.
func (s *Service) ConfirmKitchenOrder(ctx context.Context, p db.ActConfirmKitchenOrderParams) (db.KitchenOrder, error) {
	before, err := s.GetKitchenOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.kitchen_order.confirm", "kitchen_orders", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ConfirmKitchenOrder", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActConfirmKitchenOrder(ctx, p)
	if err != nil {
		return before, stateError(err, "confirm kitchen orders")
	}
	if err := s.audit(ctx, "production.kitchen_order.confirm", "kitchen_orders", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ConfirmKitchenOrder", row)
}

// KitchenOrderChildren returns the rows of a record's child tables, by API field.
func (s *Service) KitchenOrderChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListKitchenOrderLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list kitchen_order_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ListProductionPlan returns a page of production plans.
func (s *Service) ListProductionPlan(ctx context.Context, p db.ListProductionPlanParams) ([]db.ProductionPlan, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListProductionPlan(ctx, p))("list production plans")
}

// GetProductionPlan returns one row of production plans.
func (s *Service) GetProductionPlan(ctx context.Context, id uuid.UUID) (db.ProductionPlan, error) {
	row, err := s.q(ctx).GetProductionPlan(ctx, id)
	if err != nil {
		return row, notFound(err, "production plans")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateProductionPlan adds one row to production plans.
func (s *Service) CreateProductionPlan(ctx context.Context, p db.CreateProductionPlanParams) (db.ProductionPlan, error) {
	var zero db.ProductionPlan
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateProductionPlan", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "PPL", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateProductionPlan(ctx, p)
	if err != nil {
		return zero, wrap(err, "create production plans")
	}
	if err := s.audit(ctx, "production.production_plan.write", "production_plans", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateProductionPlan", row)
}

// UpdateProductionPlan changes one row of production plans.
func (s *Service) UpdateProductionPlan(ctx context.Context, p db.UpdateProductionPlanParams) (db.ProductionPlan, error) {
	before, err := s.GetProductionPlan(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateProductionPlan(ctx, p)
	if err != nil {
		return before, wrap(err, "update production plans")
	}
	if err := s.voidSignatures(ctx, "production_plans", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "production.production_plan.write", "production_plans", row.ID, &row.SiteID, before, row)
}

// ConfirmProductionPlan is the confirm the plan action on production plans.
func (s *Service) ConfirmProductionPlan(ctx context.Context, p db.ActConfirmProductionPlanParams) (db.ProductionPlan, error) {
	before, err := s.GetProductionPlan(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.production_plan.confirm", "production_plans", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ConfirmProductionPlan", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActConfirmProductionPlan(ctx, p)
	if err != nil {
		return before, stateError(err, "confirm production plans")
	}
	if err := s.audit(ctx, "production.production_plan.confirm", "production_plans", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ConfirmProductionPlan", row)
}

// ProductionPlanChildren returns the rows of a record's child tables, by API field.
func (s *Service) ProductionPlanChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListProductionPlanLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list production_plan_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplaceProductionPlanLines swaps a record's production plan lines for the given rows.
func (s *Service) ReplaceProductionPlanLines(ctx context.Context, id uuid.UUID, rows []db.InsertProductionPlanLinesParams) error {
	if err := s.q(ctx).DeleteProductionPlanLines(ctx, id); err != nil {
		return wrap(err, "clear production_plan_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertProductionPlanLines(ctx, row); err != nil {
			return wrap(err, "add production_plan_lines")
		}
	}
	return nil
}

// ListBatch returns a page of batches.
func (s *Service) ListBatch(ctx context.Context, p db.ListBatchParams) ([]db.Batch, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListBatch(ctx, p))("list batches")
}

// GetBatch returns one row of batches.
func (s *Service) GetBatch(ctx context.Context, id uuid.UUID) (db.Batch, error) {
	row, err := s.q(ctx).GetBatch(ctx, id)
	if err != nil {
		return row, notFound(err, "batches")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateBatch adds one row to batches.
func (s *Service) CreateBatch(ctx context.Context, p db.CreateBatchParams) (db.Batch, error) {
	var zero db.Batch
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateBatch", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateBatch(ctx, p)
	if err != nil {
		return zero, wrap(err, "create batches")
	}
	if err := s.audit(ctx, "production.batch.write", "batches", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateBatch", row)
}

// UpdateBatch changes one row of batches.
func (s *Service) UpdateBatch(ctx context.Context, p db.UpdateBatchParams) (db.Batch, error) {
	before, err := s.GetBatch(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateBatch(ctx, p)
	if err != nil {
		return before, wrap(err, "update batches")
	}
	if err := s.voidSignatures(ctx, "batches", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "production.batch.write", "batches", row.ID, &row.SiteID, before, row)
}

// CompleteBatch is the record output and complete action on batches.
func (s *Service) CompleteBatch(ctx context.Context, p db.ActCompleteBatchParams) (db.Batch, error) {
	before, err := s.GetBatch(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.batch.complete", "batches", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CompleteBatch", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "batches", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompleteBatch(ctx, p)
	if err != nil {
		return before, stateError(err, "complete batches")
	}
	if err := s.audit(ctx, "production.batch.complete", "batches", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CompleteBatch", row)
}

// ReleaseBatch is the qc release action on batches.
func (s *Service) ReleaseBatch(ctx context.Context, p db.ActReleaseBatchParams) (db.Batch, error) {
	before, err := s.GetBatch(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.batch.release", "batches", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReleaseBatch", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReleaseBatch(ctx, p)
	if err != nil {
		return before, stateError(err, "release batches")
	}
	if err := s.audit(ctx, "production.batch.release", "batches", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReleaseBatch", row)
}

// BatchChildren returns the rows of a record's child tables, by API field.
func (s *Service) BatchChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListBatchIngredients(ctx, id)
		if err != nil {
			return nil, wrap(err, "list batch_ingredients")
		}
		out["ingredients"] = childRows(rows)
	}
	return out, nil
}

// ReplaceBatchIngredients swaps a record's batch ingredients for the given rows.
func (s *Service) ReplaceBatchIngredients(ctx context.Context, id uuid.UUID, rows []db.InsertBatchIngredientsParams) error {
	if err := s.q(ctx).DeleteBatchIngredients(ctx, id); err != nil {
		return wrap(err, "clear batch_ingredients")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertBatchIngredients(ctx, row); err != nil {
			return wrap(err, "add batch_ingredients")
		}
	}
	return nil
}

// ListDispatchNote returns a page of dispatch notes.
func (s *Service) ListDispatchNote(ctx context.Context, p db.ListDispatchNoteParams) ([]db.DispatchNote, error) {
	return wrap2(s.q(ctx).ListDispatchNote(ctx, p))("list dispatch notes")
}

// GetDispatchNote returns one row of dispatch notes.
func (s *Service) GetDispatchNote(ctx context.Context, id uuid.UUID) (db.DispatchNote, error) {
	row, err := s.q(ctx).GetDispatchNote(ctx, id)
	if err != nil {
		return row, notFound(err, "dispatch notes")
	}
	return row, nil
}

// CreateDispatchNote adds one row to dispatch notes.
func (s *Service) CreateDispatchNote(ctx context.Context, p db.CreateDispatchNoteParams) (db.DispatchNote, error) {
	var zero db.DispatchNote
	if err := s.guard(ctx, "CreateDispatchNote", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "DN", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDispatchNote(ctx, p)
	if err != nil {
		return zero, wrap(err, "create dispatch notes")
	}
	if err := s.audit(ctx, "production.dispatch_note.write", "dispatch_notes", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDispatchNote", row)
}

// UpdateDispatchNote changes one row of dispatch notes.
func (s *Service) UpdateDispatchNote(ctx context.Context, p db.UpdateDispatchNoteParams) (db.DispatchNote, error) {
	before, err := s.GetDispatchNote(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDispatchNote(ctx, p)
	if err != nil {
		return before, wrap(err, "update dispatch notes")
	}
	if err := s.voidSignatures(ctx, "dispatch_notes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "production.dispatch_note.write", "dispatch_notes", row.ID, nil, before, row)
}

// VerifyDispatchNote is the supervisor verifies quantities action on dispatch notes.
func (s *Service) VerifyDispatchNote(ctx context.Context, p db.ActVerifyDispatchNoteParams) (db.DispatchNote, error) {
	before, err := s.GetDispatchNote(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.dispatch_note.verify", "dispatch_notes", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "VerifyDispatchNote", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "dispatch_notes", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActVerifyDispatchNote(ctx, p)
	if err != nil {
		return before, stateError(err, "verify dispatch notes")
	}
	if err := s.audit(ctx, "production.dispatch_note.verify", "dispatch_notes", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "VerifyDispatchNote", row)
}

// SecurityClearDispatchNote is the security matches items to the d-note action on dispatch notes.
func (s *Service) SecurityClearDispatchNote(ctx context.Context, p db.ActSecurityClearDispatchNoteParams) (db.DispatchNote, error) {
	before, err := s.GetDispatchNote(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.dispatch_note.security_clear", "dispatch_notes", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SecurityClearDispatchNote", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSecurityClearDispatchNote(ctx, p)
	if err != nil {
		return before, stateError(err, "security clear dispatch notes")
	}
	if err := s.audit(ctx, "production.dispatch_note.security_clear", "dispatch_notes", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SecurityClearDispatchNote", row)
}

// ConfirmDeliveryDispatchNote is the receiving supervisor confirms action on dispatch notes.
func (s *Service) ConfirmDeliveryDispatchNote(ctx context.Context, p db.ActConfirmDeliveryDispatchNoteParams) (db.DispatchNote, error) {
	before, err := s.GetDispatchNote(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.dispatch_note.confirm_delivery", "dispatch_notes", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ConfirmDeliveryDispatchNote", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActConfirmDeliveryDispatchNote(ctx, p)
	if err != nil {
		return before, stateError(err, "confirm delivery dispatch notes")
	}
	if err := s.audit(ctx, "production.dispatch_note.confirm_delivery", "dispatch_notes", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ConfirmDeliveryDispatchNote", row)
}

// RaiseDiscrepancyDispatchNote is the record a shortage, rejection or temperature issue action on dispatch notes.
func (s *Service) RaiseDiscrepancyDispatchNote(ctx context.Context, p db.ActRaiseDiscrepancyDispatchNoteParams) (db.DispatchNote, error) {
	before, err := s.GetDispatchNote(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "production.dispatch_note.raise_discrepancy", "dispatch_notes", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RaiseDiscrepancyDispatchNote", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRaiseDiscrepancyDispatchNote(ctx, p)
	if err != nil {
		return before, stateError(err, "raise discrepancy dispatch notes")
	}
	if err := s.audit(ctx, "production.dispatch_note.raise_discrepancy", "dispatch_notes", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RaiseDiscrepancyDispatchNote", row)
}

// DispatchNoteChildren returns the rows of a record's child tables, by API field.
func (s *Service) DispatchNoteChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListDispatchNoteLines(ctx, id)
		if err != nil {
			return nil, wrap(err, "list dispatch_note_lines")
		}
		out["lines"] = childRows(rows)
	}
	return out, nil
}

// ReplaceDispatchNoteLines swaps a record's dispatch note lines for the given rows.
func (s *Service) ReplaceDispatchNoteLines(ctx context.Context, id uuid.UUID, rows []db.InsertDispatchNoteLinesParams) error {
	if err := s.q(ctx).DeleteDispatchNoteLines(ctx, id); err != nil {
		return wrap(err, "clear dispatch_note_lines")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertDispatchNoteLines(ctx, row); err != nil {
			return wrap(err, "add dispatch_note_lines")
		}
	}
	return nil
}

// ListProduct returns a page of products.
func (s *Service) ListProduct(ctx context.Context, p db.ListProductParams) ([]db.Product, error) {
	return wrap2(s.q(ctx).ListProduct(ctx, p))("list products")
}

// GetProduct returns one row of products.
func (s *Service) GetProduct(ctx context.Context, id uuid.UUID) (db.Product, error) {
	row, err := s.q(ctx).GetProduct(ctx, id)
	if err != nil {
		return row, notFound(err, "products")
	}
	return row, nil
}

// CreateProduct adds one row to products.
func (s *Service) CreateProduct(ctx context.Context, p db.CreateProductParams) (db.Product, error) {
	var zero db.Product
	if err := s.guard(ctx, "CreateProduct", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateProduct(ctx, p)
	if err != nil {
		return zero, wrap(err, "create products")
	}
	if err := s.audit(ctx, "products.product.write", "products", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateProduct", row)
}

// UpdateProduct changes one row of products.
func (s *Service) UpdateProduct(ctx context.Context, p db.UpdateProductParams) (db.Product, error) {
	before, err := s.GetProduct(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateProduct(ctx, p)
	if err != nil {
		return before, wrap(err, "update products")
	}
	if err := s.voidSignatures(ctx, "products", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "products.product.write", "products", row.ID, nil, before, row)
}

// ListProductChange returns a page of product changes.
func (s *Service) ListProductChange(ctx context.Context, p db.ListProductChangeParams) ([]db.ProductChange, error) {
	return wrap2(s.q(ctx).ListProductChange(ctx, p))("list product changes")
}

// GetProductChange returns one row of product changes.
func (s *Service) GetProductChange(ctx context.Context, id uuid.UUID) (db.ProductChange, error) {
	row, err := s.q(ctx).GetProductChange(ctx, id)
	if err != nil {
		return row, notFound(err, "product changes")
	}
	return row, nil
}

// CreateProductChange adds one row to product changes.
func (s *Service) CreateProductChange(ctx context.Context, p db.CreateProductChangeParams) (db.ProductChange, error) {
	var zero db.ProductChange
	if err := s.guard(ctx, "CreateProductChange", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "PCG", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateProductChange(ctx, p)
	if err != nil {
		return zero, wrap(err, "create product changes")
	}
	if err := s.audit(ctx, "products.product_change.write", "product_changes", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateProductChange", row)
}

// UpdateProductChange changes one row of product changes.
func (s *Service) UpdateProductChange(ctx context.Context, p db.UpdateProductChangeParams) (db.ProductChange, error) {
	before, err := s.GetProductChange(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateProductChange(ctx, p)
	if err != nil {
		return before, wrap(err, "update product changes")
	}
	if err := s.voidSignatures(ctx, "product_changes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "products.product_change.write", "product_changes", row.ID, nil, before, row)
}

// SubmitProductChange is the submit action on product changes.
func (s *Service) SubmitProductChange(ctx context.Context, p db.ActSubmitProductChangeParams) (db.ProductChange, error) {
	before, err := s.GetProductChange(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "products.product_change.submit", "product_changes", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitProductChange", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitProductChange(ctx, p)
	if err != nil {
		return before, stateError(err, "submit product changes")
	}
	if err := s.audit(ctx, "products.product_change.submit", "product_changes", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "product_changes", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitProductChange", row)
}

// ListCostSheet returns a page of cost sheets.
func (s *Service) ListCostSheet(ctx context.Context, p db.ListCostSheetParams) ([]db.CostSheet, error) {
	return wrap2(s.q(ctx).ListCostSheet(ctx, p))("list cost sheets")
}

// GetCostSheet returns one row of cost sheets.
func (s *Service) GetCostSheet(ctx context.Context, id uuid.UUID) (db.CostSheet, error) {
	row, err := s.q(ctx).GetCostSheet(ctx, id)
	if err != nil {
		return row, notFound(err, "cost sheets")
	}
	return row, nil
}

// CreateCostSheet adds one row to cost sheets.
func (s *Service) CreateCostSheet(ctx context.Context, p db.CreateCostSheetParams) (db.CostSheet, error) {
	var zero db.CostSheet
	if err := s.guard(ctx, "CreateCostSheet", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "CST", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCostSheet(ctx, p)
	if err != nil {
		return zero, wrap(err, "create cost sheets")
	}
	if err := s.audit(ctx, "products.cost_sheet.write", "cost_sheets", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateCostSheet", row)
}

// UpdateCostSheet changes one row of cost sheets.
func (s *Service) UpdateCostSheet(ctx context.Context, p db.UpdateCostSheetParams) (db.CostSheet, error) {
	before, err := s.GetCostSheet(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCostSheet(ctx, p)
	if err != nil {
		return before, wrap(err, "update cost sheets")
	}
	if err := s.voidSignatures(ctx, "cost_sheets", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "products.cost_sheet.write", "cost_sheets", row.ID, nil, before, row)
}

// SubmitCostSheet is the submit action on cost sheets.
func (s *Service) SubmitCostSheet(ctx context.Context, p db.ActSubmitCostSheetParams) (db.CostSheet, error) {
	before, err := s.GetCostSheet(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "products.cost_sheet.submit", "cost_sheets", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitCostSheet", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitCostSheet(ctx, p)
	if err != nil {
		return before, stateError(err, "submit cost sheets")
	}
	if err := s.audit(ctx, "products.cost_sheet.submit", "cost_sheets", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "cost_sheets", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitCostSheet", row)
}

// ListNpdProject returns a page of npd projects.
func (s *Service) ListNpdProject(ctx context.Context, p db.ListNpdProjectParams) ([]db.NpdProject, error) {
	return wrap2(s.q(ctx).ListNpdProject(ctx, p))("list npd projects")
}

// GetNpdProject returns one row of npd projects.
func (s *Service) GetNpdProject(ctx context.Context, id uuid.UUID) (db.NpdProject, error) {
	row, err := s.q(ctx).GetNpdProject(ctx, id)
	if err != nil {
		return row, notFound(err, "npd projects")
	}
	return row, nil
}

// CreateNpdProject adds one row to npd projects.
func (s *Service) CreateNpdProject(ctx context.Context, p db.CreateNpdProjectParams) (db.NpdProject, error) {
	var zero db.NpdProject
	if err := s.guard(ctx, "CreateNpdProject", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "NPD", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateNpdProject(ctx, p)
	if err != nil {
		return zero, wrap(err, "create npd projects")
	}
	if err := s.audit(ctx, "products.npd_project.write", "npd_projects", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateNpdProject", row)
}

// UpdateNpdProject changes one row of npd projects.
func (s *Service) UpdateNpdProject(ctx context.Context, p db.UpdateNpdProjectParams) (db.NpdProject, error) {
	before, err := s.GetNpdProject(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateNpdProject(ctx, p)
	if err != nil {
		return before, wrap(err, "update npd projects")
	}
	if err := s.voidSignatures(ctx, "npd_projects", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "products.npd_project.write", "npd_projects", row.ID, nil, before, row)
}

// AdvanceGateNpdProject is the record gate decision for approval action on npd projects.
func (s *Service) AdvanceGateNpdProject(ctx context.Context, p db.ActAdvanceGateNpdProjectParams) (db.NpdProject, error) {
	before, err := s.GetNpdProject(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "products.npd_project.advance_gate", "npd_projects", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "AdvanceGateNpdProject", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAdvanceGateNpdProject(ctx, p)
	if err != nil {
		return before, stateError(err, "advance gate npd projects")
	}
	if err := s.audit(ctx, "products.npd_project.advance_gate", "npd_projects", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "AdvanceGateNpdProject", row)
}
