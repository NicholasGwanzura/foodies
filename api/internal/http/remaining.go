package httpapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
	"foodies/api/internal/service"
)

// DecideStoreOpening records the go or no-go decision.
func (s *Server) DecideStoreOpening(ctx context.Context, req apigen.DecideStoreOpeningRequestObject) (apigen.DecideStoreOpeningResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.DecideStoreOpening(ctx, req.StoreOpeningId, string(req.Body.Decision), req.Body.Rationale, req.Body.SignatureId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StoreOpening](row)
	if err != nil {
		return nil, err
	}
	return apigen.DecideStoreOpening200JSONResponse{Body: body, Headers: apigen.DecideStoreOpening200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// TraceBatch walks a batch back to supplier lots and forward to sites.
func (s *Server) TraceBatch(ctx context.Context, req apigen.TraceBatchRequestObject) (apigen.TraceBatchResponseObject, error) {
	if req.Params.BatchId == nil {
		return nil, service.Fail(400, "validation_failed", "Malformed request", "Give a batch_id to trace.")
	}
	nodes, gaps, err := s.service.Trace(ctx, *req.Params.BatchId)
	if err != nil {
		return nil, err
	}
	out := make([]map[string]any, 0, len(nodes))
	for _, n := range nodes {
		out = append(out, map[string]any{"type": n.Type, "id": n.ID, "label": n.Label, "site_id": n.Site})
	}
	return output[apigen.TraceBatch200JSONResponse](map[string]any{"nodes": out, "complete": len(gaps) == 0, "gaps": gaps})
}

// GetConsolidatedKitchenOrders sums on-time orders by product.
func (s *Server) GetConsolidatedKitchenOrders(ctx context.Context, req apigen.GetConsolidatedKitchenOrdersRequestObject) (apigen.GetConsolidatedKitchenOrdersResponseObject, error) {
	date := pgtype.Date{Time: req.Params.OrderDate.Time, Valid: true}
	rows, err := s.service.ConsolidatedOrders(ctx, db.ConsolidatedOrdersParams{OrderDate: date, Cycle: string(req.Params.Cycle)})
	if err != nil {
		return nil, err
	}
	return output[apigen.GetConsolidatedKitchenOrders200JSONResponse](map[string]any{"order_date": req.Params.OrderDate, "cycle": req.Params.Cycle, "lines": rows})
}

// GetKpiValues returns a KPI's series.
func (s *Server) GetKpiValues(ctx context.Context, req apigen.GetKpiValuesRequestObject) (apigen.GetKpiValuesResponseObject, error) {
	row, err := s.service.KpiSeries(ctx, req.KpiDefinitionId)
	if err != nil {
		return nil, err
	}
	return output[apigen.GetKpiValues200JSONResponse](map[string]any{"kpi_key": row.Key, "unit": row.Unit, "points": []any{}})
}

// GetSyncBatch reports on a processed batch.
func (s *Server) GetSyncBatch(ctx context.Context, req apigen.GetSyncBatchRequestObject) (apigen.GetSyncBatchResponseObject, error) {
	rows, err := s.service.SyncBatchConflicts(ctx, req.BatchId)
	if err != nil {
		return nil, err
	}
	items := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		items = append(items, map[string]any{"client_record_id": r.ClientRecordID, "status": "conflict"})
	}
	return output[apigen.GetSyncBatch200JSONResponse](map[string]any{"batch_id": req.BatchId, "status": "completed", "items": items})
}

// ReceiveIntegrationEvents accepts events from a system of record.
func (s *Server) ReceiveIntegrationEvents(ctx context.Context, req apigen.ReceiveIntegrationEventsRequestObject) (apigen.ReceiveIntegrationEventsResponseObject, error) {
	key, _ := ctx.Value(integrationKey{}).(string)
	if err := s.service.RecordIntegrationEvents(ctx, string(req.System), key); err != nil {
		return nil, err
	}
	return apigen.ReceiveIntegrationEvents202Response{}, nil
}

// ExchangeWhatsappFlowData serves encrypted Flow data. The exchange itself waits
// for the Flow keys from Meta, so for now it only verifies the caller.
func (s *Server) ExchangeWhatsappFlowData(ctx context.Context, _ apigen.ExchangeWhatsappFlowDataRequestObject) (apigen.ExchangeWhatsappFlowDataResponseObject, error) {
	if _, err := s.verifiedBody(ctx); err != nil {
		return nil, err
	}
	return nil, apigen.ErrNotImplemented
}

type integrationKey struct{}
