package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

// ApproveCashUp approves a reviewed cash-up.
func (s *Server) ApproveCashUp(ctx context.Context, req apigen.ApproveCashUpRequestObject) (apigen.ApproveCashUpResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.ApproveCashUp(ctx, db.ActApproveCashUpParams{ID: req.CashUpId}, req.Body.SignatureId)
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
	return apigen.ApproveCashUp200JSONResponse{Body: body, Headers: apigen.ApproveCashUp200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

func citBody(row db.CitTransit) (apigen.CitTransit, *string, error) {
	body, err := output[apigen.CitTransit](row)
	return body, etag(row.RowVersion), err
}

// VerifyIdentityCitTransit matches the crew.
func (s *Server) VerifyIdentityCitTransit(ctx context.Context, req apigen.VerifyIdentityCitTransitRequestObject) (apigen.VerifyIdentityCitTransitResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.VerifyIdentityCitTransit(ctx, req.CitTransitId, req.Body.CitPersonId)
	if err != nil {
		return nil, err
	}
	body, tag, err := citBody(row)
	if err != nil {
		return nil, err
	}
	return apigen.VerifyIdentityCitTransit200JSONResponse{Body: body, Headers: apigen.VerifyIdentityCitTransit200ResponseHeaders{ETag: tag}}, nil
}

// HandOverCitTransit records the handover.
func (s *Server) HandOverCitTransit(ctx context.Context, req apigen.HandOverCitTransitRequestObject) (apigen.HandOverCitTransitResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.HandOverCitTransit(ctx, req.CitTransitId, req.Body.CitReceiptNumber, req.Body.ArmouredVehicleConfirmed)
	if err != nil {
		return nil, err
	}
	body, tag, err := citBody(row)
	if err != nil {
		return nil, err
	}
	return apigen.HandOverCitTransit200JSONResponse{Body: body, Headers: apigen.HandOverCitTransit200ResponseHeaders{ETag: tag}}, nil
}

// ReceiveVaultCitTransit checks the amount at the vault.
func (s *Server) ReceiveVaultCitTransit(ctx context.Context, req apigen.ReceiveVaultCitTransitRequestObject) (apigen.ReceiveVaultCitTransitResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.ReceiveVaultCitTransit(ctx, req.CitTransitId, req.Body.Amount.Amount, req.Body.Amount.Currency)
	if err != nil {
		return nil, err
	}
	body, tag, err := citBody(row)
	if err != nil {
		return nil, err
	}
	return apigen.ReceiveVaultCitTransit200JSONResponse{Body: body, Headers: apigen.ReceiveVaultCitTransit200ResponseHeaders{ETag: tag}}, nil
}

// IssuePettyCashVoucher pays out an approved voucher.
func (s *Server) IssuePettyCashVoucher(ctx context.Context, req apigen.IssuePettyCashVoucherRequestObject) (apigen.IssuePettyCashVoucherResponseObject, error) {
	if err := checkRequired(req.Body); err != nil {
		return nil, err
	}
	row, err := s.service.IssuePettyCashVoucher(ctx, req.PettyCashVoucherId, req.Body.IssuedAmount.Amount, req.Body.IssuedAmount.Currency)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PettyCashVoucher](row)
	if err != nil {
		return nil, err
	}
	return apigen.IssuePettyCashVoucher200JSONResponse{Body: body, Headers: apigen.IssuePettyCashVoucher200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
