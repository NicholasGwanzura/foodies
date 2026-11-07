// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

// ListComplexes implements the listComplexes operation.
func (s *Server) ListComplexes(ctx context.Context, req apigen.ListComplexesRequestObject) (apigen.ListComplexesResponseObject, error) {
	return listPage[apigen.ListComplexes200JSONResponse](ctx, req.Params, nil, s.service.ListComplex)
}

// GetComplex implements the getComplex operation.
func (s *Server) GetComplex(ctx context.Context, req apigen.GetComplexRequestObject) (apigen.GetComplexResponseObject, error) {
	row, err := s.service.GetComplex(ctx, req.ComplexId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Complex](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetComplex200JSONResponse{Body: body, Headers: apigen.GetComplex200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateComplex implements the createComplex operation.
func (s *Server) CreateComplex(ctx context.Context, req apigen.CreateComplexRequestObject) (apigen.CreateComplexResponseObject, error) {
	in, err := input[db.CreateComplexParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateComplex(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Complex](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateComplex201JSONResponse{Body: body, Headers: apigen.CreateComplex201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateComplex implements the updateComplex operation.
func (s *Server) UpdateComplex(ctx context.Context, req apigen.UpdateComplexRequestObject) (apigen.UpdateComplexResponseObject, error) {
	in, err := input[db.UpdateComplexParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ComplexId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateComplex(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Complex](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateComplex200JSONResponse{Body: body, Headers: apigen.UpdateComplex200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListSites implements the listSites operation.
func (s *Server) ListSites(ctx context.Context, req apigen.ListSitesRequestObject) (apigen.ListSitesResponseObject, error) {
	return listPage[apigen.ListSites200JSONResponse](ctx, req.Params, nil, s.service.ListSite)
}

// GetSite implements the getSite operation.
func (s *Server) GetSite(ctx context.Context, req apigen.GetSiteRequestObject) (apigen.GetSiteResponseObject, error) {
	row, err := s.service.GetSite(ctx, req.SiteId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Site](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSite200JSONResponse{Body: body, Headers: apigen.GetSite200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateSite implements the createSite operation.
func (s *Server) CreateSite(ctx context.Context, req apigen.CreateSiteRequestObject) (apigen.CreateSiteResponseObject, error) {
	in, err := input[db.CreateSiteParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSite(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Site](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSite201JSONResponse{Body: body, Headers: apigen.CreateSite201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateSite implements the updateSite operation.
func (s *Server) UpdateSite(ctx context.Context, req apigen.UpdateSiteRequestObject) (apigen.UpdateSiteResponseObject, error) {
	in, err := input[db.UpdateSiteParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.SiteId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateSite(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Site](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateSite200JSONResponse{Body: body, Headers: apigen.UpdateSite200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDepartments implements the listDepartments operation.
func (s *Server) ListDepartments(ctx context.Context, req apigen.ListDepartmentsRequestObject) (apigen.ListDepartmentsResponseObject, error) {
	return listPage[apigen.ListDepartments200JSONResponse](ctx, req.Params, nil, s.service.ListDepartment)
}

// GetDepartment implements the getDepartment operation.
func (s *Server) GetDepartment(ctx context.Context, req apigen.GetDepartmentRequestObject) (apigen.GetDepartmentResponseObject, error) {
	row, err := s.service.GetDepartment(ctx, req.DepartmentId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Department](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDepartment200JSONResponse{Body: body, Headers: apigen.GetDepartment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDepartment implements the createDepartment operation.
func (s *Server) CreateDepartment(ctx context.Context, req apigen.CreateDepartmentRequestObject) (apigen.CreateDepartmentResponseObject, error) {
	in, err := input[db.CreateDepartmentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDepartment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Department](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDepartment201JSONResponse{Body: body, Headers: apigen.CreateDepartment201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDepartment implements the updateDepartment operation.
func (s *Server) UpdateDepartment(ctx context.Context, req apigen.UpdateDepartmentRequestObject) (apigen.UpdateDepartmentResponseObject, error) {
	in, err := input[db.UpdateDepartmentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.DepartmentId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDepartment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Department](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDepartment200JSONResponse{Body: body, Headers: apigen.UpdateDepartment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCostCentres implements the listCostCentres operation.
func (s *Server) ListCostCentres(ctx context.Context, req apigen.ListCostCentresRequestObject) (apigen.ListCostCentresResponseObject, error) {
	return listPage[apigen.ListCostCentres200JSONResponse](ctx, req.Params, nil, s.service.ListCostCentre)
}

// GetCostCentre implements the getCostCentre operation.
func (s *Server) GetCostCentre(ctx context.Context, req apigen.GetCostCentreRequestObject) (apigen.GetCostCentreResponseObject, error) {
	row, err := s.service.GetCostCentre(ctx, req.CostCentreId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CostCentre](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCostCentre200JSONResponse{Body: body, Headers: apigen.GetCostCentre200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCostCentre implements the createCostCentre operation.
func (s *Server) CreateCostCentre(ctx context.Context, req apigen.CreateCostCentreRequestObject) (apigen.CreateCostCentreResponseObject, error) {
	in, err := input[db.CreateCostCentreParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCostCentre(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CostCentre](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCostCentre201JSONResponse{Body: body, Headers: apigen.CreateCostCentre201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCostCentre implements the updateCostCentre operation.
func (s *Server) UpdateCostCentre(ctx context.Context, req apigen.UpdateCostCentreRequestObject) (apigen.UpdateCostCentreResponseObject, error) {
	in, err := input[db.UpdateCostCentreParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CostCentreId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCostCentre(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.CostCentre](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCostCentre200JSONResponse{Body: body, Headers: apigen.UpdateCostCentre200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListProductionLines implements the listProductionLines operation.
func (s *Server) ListProductionLines(ctx context.Context, req apigen.ListProductionLinesRequestObject) (apigen.ListProductionLinesResponseObject, error) {
	return listPage[apigen.ListProductionLines200JSONResponse](ctx, req.Params, nil, s.service.ListProductionLine)
}

// GetProductionLine implements the getProductionLine operation.
func (s *Server) GetProductionLine(ctx context.Context, req apigen.GetProductionLineRequestObject) (apigen.GetProductionLineResponseObject, error) {
	row, err := s.service.GetProductionLine(ctx, req.ProductionLineId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductionLine](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetProductionLine200JSONResponse{Body: body, Headers: apigen.GetProductionLine200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateProductionLine implements the createProductionLine operation.
func (s *Server) CreateProductionLine(ctx context.Context, req apigen.CreateProductionLineRequestObject) (apigen.CreateProductionLineResponseObject, error) {
	in, err := input[db.CreateProductionLineParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateProductionLine(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductionLine](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateProductionLine201JSONResponse{Body: body, Headers: apigen.CreateProductionLine201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateProductionLine implements the updateProductionLine operation.
func (s *Server) UpdateProductionLine(ctx context.Context, req apigen.UpdateProductionLineRequestObject) (apigen.UpdateProductionLineResponseObject, error) {
	in, err := input[db.UpdateProductionLineParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ProductionLineId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateProductionLine(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ProductionLine](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateProductionLine200JSONResponse{Body: body, Headers: apigen.UpdateProductionLine200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListTills implements the listTills operation.
func (s *Server) ListTills(ctx context.Context, req apigen.ListTillsRequestObject) (apigen.ListTillsResponseObject, error) {
	return listPage[apigen.ListTills200JSONResponse](ctx, req.Params, nil, s.service.ListTill)
}

// GetTill implements the getTill operation.
func (s *Server) GetTill(ctx context.Context, req apigen.GetTillRequestObject) (apigen.GetTillResponseObject, error) {
	row, err := s.service.GetTill(ctx, req.TillId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Till](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetTill200JSONResponse{Body: body, Headers: apigen.GetTill200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateTill implements the createTill operation.
func (s *Server) CreateTill(ctx context.Context, req apigen.CreateTillRequestObject) (apigen.CreateTillResponseObject, error) {
	in, err := input[db.CreateTillParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateTill(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Till](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateTill201JSONResponse{Body: body, Headers: apigen.CreateTill201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateTill implements the updateTill operation.
func (s *Server) UpdateTill(ctx context.Context, req apigen.UpdateTillRequestObject) (apigen.UpdateTillResponseObject, error) {
	in, err := input[db.UpdateTillParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.TillId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateTill(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Till](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateTill200JSONResponse{Body: body, Headers: apigen.UpdateTill200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListSafes implements the listSafes operation.
func (s *Server) ListSafes(ctx context.Context, req apigen.ListSafesRequestObject) (apigen.ListSafesResponseObject, error) {
	return listPage[apigen.ListSafes200JSONResponse](ctx, req.Params, nil, s.service.ListSafe)
}

// GetSafe implements the getSafe operation.
func (s *Server) GetSafe(ctx context.Context, req apigen.GetSafeRequestObject) (apigen.GetSafeResponseObject, error) {
	row, err := s.service.GetSafe(ctx, req.SafeId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Safe](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSafe200JSONResponse{Body: body, Headers: apigen.GetSafe200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateSafe implements the createSafe operation.
func (s *Server) CreateSafe(ctx context.Context, req apigen.CreateSafeRequestObject) (apigen.CreateSafeResponseObject, error) {
	in, err := input[db.CreateSafeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSafe(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Safe](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSafe201JSONResponse{Body: body, Headers: apigen.CreateSafe201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateSafe implements the updateSafe operation.
func (s *Server) UpdateSafe(ctx context.Context, req apigen.UpdateSafeRequestObject) (apigen.UpdateSafeResponseObject, error) {
	in, err := input[db.UpdateSafeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.SafeId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateSafe(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Safe](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateSafe200JSONResponse{Body: body, Headers: apigen.UpdateSafe200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListStorageLocations implements the listStorageLocations operation.
func (s *Server) ListStorageLocations(ctx context.Context, req apigen.ListStorageLocationsRequestObject) (apigen.ListStorageLocationsResponseObject, error) {
	return listPage[apigen.ListStorageLocations200JSONResponse](ctx, req.Params, nil, s.service.ListStorageLocation)
}

// GetStorageLocation implements the getStorageLocation operation.
func (s *Server) GetStorageLocation(ctx context.Context, req apigen.GetStorageLocationRequestObject) (apigen.GetStorageLocationResponseObject, error) {
	row, err := s.service.GetStorageLocation(ctx, req.StorageLocationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StorageLocation](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetStorageLocation200JSONResponse{Body: body, Headers: apigen.GetStorageLocation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateStorageLocation implements the createStorageLocation operation.
func (s *Server) CreateStorageLocation(ctx context.Context, req apigen.CreateStorageLocationRequestObject) (apigen.CreateStorageLocationResponseObject, error) {
	in, err := input[db.CreateStorageLocationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateStorageLocation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StorageLocation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateStorageLocation201JSONResponse{Body: body, Headers: apigen.CreateStorageLocation201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateStorageLocation implements the updateStorageLocation operation.
func (s *Server) UpdateStorageLocation(ctx context.Context, req apigen.UpdateStorageLocationRequestObject) (apigen.UpdateStorageLocationResponseObject, error) {
	in, err := input[db.UpdateStorageLocationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.StorageLocationId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateStorageLocation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.StorageLocation](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateStorageLocation200JSONResponse{Body: body, Headers: apigen.UpdateStorageLocation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListCurrencies implements the listCurrencies operation.
func (s *Server) ListCurrencies(ctx context.Context, req apigen.ListCurrenciesRequestObject) (apigen.ListCurrenciesResponseObject, error) {
	return listPage[apigen.ListCurrencies200JSONResponse](ctx, req.Params, nil, s.service.ListCurrency)
}

// GetCurrency implements the getCurrency operation.
func (s *Server) GetCurrency(ctx context.Context, req apigen.GetCurrencyRequestObject) (apigen.GetCurrencyResponseObject, error) {
	row, err := s.service.GetCurrency(ctx, req.CurrencyId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Currency](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetCurrency200JSONResponse{Body: body, Headers: apigen.GetCurrency200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateCurrency implements the createCurrency operation.
func (s *Server) CreateCurrency(ctx context.Context, req apigen.CreateCurrencyRequestObject) (apigen.CreateCurrencyResponseObject, error) {
	in, err := input[db.CreateCurrencyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateCurrency(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Currency](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateCurrency201JSONResponse{Body: body, Headers: apigen.CreateCurrency201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateCurrency implements the updateCurrency operation.
func (s *Server) UpdateCurrency(ctx context.Context, req apigen.UpdateCurrencyRequestObject) (apigen.UpdateCurrencyResponseObject, error) {
	in, err := input[db.UpdateCurrencyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.CurrencyId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateCurrency(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Currency](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateCurrency200JSONResponse{Body: body, Headers: apigen.UpdateCurrency200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListFxRates implements the listFxRates operation.
func (s *Server) ListFxRates(ctx context.Context, req apigen.ListFxRatesRequestObject) (apigen.ListFxRatesResponseObject, error) {
	return listPage[apigen.ListFxRates200JSONResponse](ctx, req.Params, nil, s.service.ListFxRate)
}

// GetFxRate implements the getFxRate operation.
func (s *Server) GetFxRate(ctx context.Context, req apigen.GetFxRateRequestObject) (apigen.GetFxRateResponseObject, error) {
	row, err := s.service.GetFxRate(ctx, req.FxRateId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FxRate](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetFxRate200JSONResponse{Body: body, Headers: apigen.GetFxRate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateFxRate implements the createFxRate operation.
func (s *Server) CreateFxRate(ctx context.Context, req apigen.CreateFxRateRequestObject) (apigen.CreateFxRateResponseObject, error) {
	in, err := input[db.CreateFxRateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateFxRate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FxRate](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateFxRate201JSONResponse{Body: body, Headers: apigen.CreateFxRate201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateFxRate implements the updateFxRate operation.
func (s *Server) UpdateFxRate(ctx context.Context, req apigen.UpdateFxRateRequestObject) (apigen.UpdateFxRateResponseObject, error) {
	in, err := input[db.UpdateFxRateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.FxRateId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateFxRate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FxRate](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateFxRate200JSONResponse{Body: body, Headers: apigen.UpdateFxRate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPublicHolidays implements the listPublicHolidays operation.
func (s *Server) ListPublicHolidays(ctx context.Context, req apigen.ListPublicHolidaysRequestObject) (apigen.ListPublicHolidaysResponseObject, error) {
	return listPage[apigen.ListPublicHolidays200JSONResponse](ctx, req.Params, nil, s.service.ListPublicHoliday)
}

// GetPublicHoliday implements the getPublicHoliday operation.
func (s *Server) GetPublicHoliday(ctx context.Context, req apigen.GetPublicHolidayRequestObject) (apigen.GetPublicHolidayResponseObject, error) {
	row, err := s.service.GetPublicHoliday(ctx, req.PublicHolidayId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PublicHoliday](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPublicHoliday200JSONResponse{Body: body, Headers: apigen.GetPublicHoliday200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreatePublicHoliday implements the createPublicHoliday operation.
func (s *Server) CreatePublicHoliday(ctx context.Context, req apigen.CreatePublicHolidayRequestObject) (apigen.CreatePublicHolidayResponseObject, error) {
	in, err := input[db.CreatePublicHolidayParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreatePublicHoliday(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PublicHoliday](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreatePublicHoliday201JSONResponse{Body: body, Headers: apigen.CreatePublicHoliday201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdatePublicHoliday implements the updatePublicHoliday operation.
func (s *Server) UpdatePublicHoliday(ctx context.Context, req apigen.UpdatePublicHolidayRequestObject) (apigen.UpdatePublicHolidayResponseObject, error) {
	in, err := input[db.UpdatePublicHolidayParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.PublicHolidayId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdatePublicHoliday(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PublicHoliday](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdatePublicHoliday200JSONResponse{Body: body, Headers: apigen.UpdatePublicHoliday200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DeletePublicHoliday implements the deletePublicHoliday operation.
func (s *Server) DeletePublicHoliday(ctx context.Context, req apigen.DeletePublicHolidayRequestObject) (apigen.DeletePublicHolidayResponseObject, error) {
	if err := s.service.DeletePublicHoliday(ctx, req.PublicHolidayId); err != nil {
		return nil, err
	}
	return apigen.DeletePublicHoliday204Response{}, nil
}

// ListEmployees implements the listEmployees operation.
func (s *Server) ListEmployees(ctx context.Context, req apigen.ListEmployeesRequestObject) (apigen.ListEmployeesResponseObject, error) {
	return listPage[apigen.ListEmployees200JSONResponse](ctx, req.Params, nil, s.service.ListEmployee)
}

// GetEmployee implements the getEmployee operation.
func (s *Server) GetEmployee(ctx context.Context, req apigen.GetEmployeeRequestObject) (apigen.GetEmployeeResponseObject, error) {
	row, err := s.service.GetEmployee(ctx, req.EmployeeId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Employee](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetEmployee200JSONResponse{Body: body, Headers: apigen.GetEmployee200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateEmployee implements the createEmployee operation.
func (s *Server) CreateEmployee(ctx context.Context, req apigen.CreateEmployeeRequestObject) (apigen.CreateEmployeeResponseObject, error) {
	in, err := input[db.CreateEmployeeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateEmployee(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Employee](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateEmployee201JSONResponse{Body: body, Headers: apigen.CreateEmployee201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateEmployee implements the updateEmployee operation.
func (s *Server) UpdateEmployee(ctx context.Context, req apigen.UpdateEmployeeRequestObject) (apigen.UpdateEmployeeResponseObject, error) {
	in, err := input[db.UpdateEmployeeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.EmployeeId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateEmployee(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Employee](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateEmployee200JSONResponse{Body: body, Headers: apigen.UpdateEmployee200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListEmployeeStatusChanges implements the listEmployeeStatusChanges operation.
func (s *Server) ListEmployeeStatusChanges(ctx context.Context, req apigen.ListEmployeeStatusChangesRequestObject) (apigen.ListEmployeeStatusChangesResponseObject, error) {
	return listPage[apigen.ListEmployeeStatusChanges200JSONResponse](ctx, req.Params, map[string]any{"employee_id": req.EmployeeId}, s.service.ListEmployeeStatusChange)
}

// CreateEmployeeStatusChange implements the createEmployeeStatusChange operation.
func (s *Server) CreateEmployeeStatusChange(ctx context.Context, req apigen.CreateEmployeeStatusChangeRequestObject) (apigen.CreateEmployeeStatusChangeResponseObject, error) {
	in, err := input[db.CreateEmployeeStatusChangeParams](req.Body, map[string]any{"employee_id": req.EmployeeId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateEmployeeStatusChange(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmployeeStatusChange](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateEmployeeStatusChange201JSONResponse{Body: body, Headers: apigen.CreateEmployeeStatusChange201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListEmployeeTransfers implements the listEmployeeTransfers operation.
func (s *Server) ListEmployeeTransfers(ctx context.Context, req apigen.ListEmployeeTransfersRequestObject) (apigen.ListEmployeeTransfersResponseObject, error) {
	return listPage[apigen.ListEmployeeTransfers200JSONResponse](ctx, req.Params, map[string]any{"employee_id": req.EmployeeId}, s.service.ListEmployeeTransfer)
}

// GetEmployeeTransfer implements the getEmployeeTransfer operation.
func (s *Server) GetEmployeeTransfer(ctx context.Context, req apigen.GetEmployeeTransferRequestObject) (apigen.GetEmployeeTransferResponseObject, error) {
	row, err := s.service.GetEmployeeTransfer(ctx, req.EmployeeTransferId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmployeeTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetEmployeeTransfer200JSONResponse{Body: body, Headers: apigen.GetEmployeeTransfer200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateEmployeeTransfer implements the createEmployeeTransfer operation.
func (s *Server) CreateEmployeeTransfer(ctx context.Context, req apigen.CreateEmployeeTransferRequestObject) (apigen.CreateEmployeeTransferResponseObject, error) {
	in, err := input[db.CreateEmployeeTransferParams](req.Body, map[string]any{"employee_id": req.EmployeeId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateEmployeeTransfer(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmployeeTransfer](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateEmployeeTransfer201JSONResponse{Body: body, Headers: apigen.CreateEmployeeTransfer201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListEmployeeDocuments implements the listEmployeeDocuments operation.
func (s *Server) ListEmployeeDocuments(ctx context.Context, req apigen.ListEmployeeDocumentsRequestObject) (apigen.ListEmployeeDocumentsResponseObject, error) {
	return listPage[apigen.ListEmployeeDocuments200JSONResponse](ctx, req.Params, map[string]any{"employee_id": req.EmployeeId}, s.service.ListEmployeeDocument)
}

// GetEmployeeDocument implements the getEmployeeDocument operation.
func (s *Server) GetEmployeeDocument(ctx context.Context, req apigen.GetEmployeeDocumentRequestObject) (apigen.GetEmployeeDocumentResponseObject, error) {
	row, err := s.service.GetEmployeeDocument(ctx, req.EmployeeDocumentId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmployeeDocument](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetEmployeeDocument200JSONResponse{Body: body, Headers: apigen.GetEmployeeDocument200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateEmployeeDocument implements the createEmployeeDocument operation.
func (s *Server) CreateEmployeeDocument(ctx context.Context, req apigen.CreateEmployeeDocumentRequestObject) (apigen.CreateEmployeeDocumentResponseObject, error) {
	in, err := input[db.CreateEmployeeDocumentParams](req.Body, map[string]any{"employee_id": req.EmployeeId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateEmployeeDocument(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmployeeDocument](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateEmployeeDocument201JSONResponse{Body: body, Headers: apigen.CreateEmployeeDocument201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DeleteEmployeeDocument implements the deleteEmployeeDocument operation.
func (s *Server) DeleteEmployeeDocument(ctx context.Context, req apigen.DeleteEmployeeDocumentRequestObject) (apigen.DeleteEmployeeDocumentResponseObject, error) {
	if err := s.service.DeleteEmployeeDocument(ctx, req.EmployeeDocumentId); err != nil {
		return nil, err
	}
	return apigen.DeleteEmployeeDocument204Response{}, nil
}

// ListEmergencyContacts implements the listEmergencyContacts operation.
func (s *Server) ListEmergencyContacts(ctx context.Context, req apigen.ListEmergencyContactsRequestObject) (apigen.ListEmergencyContactsResponseObject, error) {
	return listPage[apigen.ListEmergencyContacts200JSONResponse](ctx, req.Params, map[string]any{"employee_id": req.EmployeeId}, s.service.ListEmergencyContact)
}

// CreateEmergencyContact implements the createEmergencyContact operation.
func (s *Server) CreateEmergencyContact(ctx context.Context, req apigen.CreateEmergencyContactRequestObject) (apigen.CreateEmergencyContactResponseObject, error) {
	in, err := input[db.CreateEmergencyContactParams](req.Body, map[string]any{"employee_id": req.EmployeeId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateEmergencyContact(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmergencyContact](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateEmergencyContact201JSONResponse{Body: body, Headers: apigen.CreateEmergencyContact201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateEmergencyContact implements the updateEmergencyContact operation.
func (s *Server) UpdateEmergencyContact(ctx context.Context, req apigen.UpdateEmergencyContactRequestObject) (apigen.UpdateEmergencyContactResponseObject, error) {
	in, err := input[db.UpdateEmergencyContactParams](req.Body, map[string]any{"employee_id": req.EmployeeId})
	if err != nil {
		return nil, err
	}
	in.ID = req.EmergencyContactId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateEmergencyContact(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EmergencyContact](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateEmergencyContact200JSONResponse{Body: body, Headers: apigen.UpdateEmergencyContact200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DeleteEmergencyContact implements the deleteEmergencyContact operation.
func (s *Server) DeleteEmergencyContact(ctx context.Context, req apigen.DeleteEmergencyContactRequestObject) (apigen.DeleteEmergencyContactResponseObject, error) {
	if err := s.service.DeleteEmergencyContact(ctx, req.EmergencyContactId); err != nil {
		return nil, err
	}
	return apigen.DeleteEmergencyContact204Response{}, nil
}

// ListUsers implements the listUsers operation.
func (s *Server) ListUsers(ctx context.Context, req apigen.ListUsersRequestObject) (apigen.ListUsersResponseObject, error) {
	return listPage[apigen.ListUsers200JSONResponse](ctx, req.Params, nil, s.service.ListUser)
}

// GetUser implements the getUser operation.
func (s *Server) GetUser(ctx context.Context, req apigen.GetUserRequestObject) (apigen.GetUserResponseObject, error) {
	row, err := s.service.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.User](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetUser200JSONResponse{Body: body, Headers: apigen.GetUser200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateUser implements the createUser operation.
func (s *Server) CreateUser(ctx context.Context, req apigen.CreateUserRequestObject) (apigen.CreateUserResponseObject, error) {
	in, err := input[db.CreateUserParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateUser(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.User](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateUser201JSONResponse{Body: body, Headers: apigen.CreateUser201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateUser implements the updateUser operation.
func (s *Server) UpdateUser(ctx context.Context, req apigen.UpdateUserRequestObject) (apigen.UpdateUserResponseObject, error) {
	in, err := input[db.UpdateUserParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.UserId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateUser(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.User](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateUser200JSONResponse{Body: body, Headers: apigen.UpdateUser200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRoles implements the listRoles operation.
func (s *Server) ListRoles(ctx context.Context, req apigen.ListRolesRequestObject) (apigen.ListRolesResponseObject, error) {
	return listPage[apigen.ListRoles200JSONResponse](ctx, req.Params, nil, s.service.ListRole)
}

// GetRole implements the getRole operation.
func (s *Server) GetRole(ctx context.Context, req apigen.GetRoleRequestObject) (apigen.GetRoleResponseObject, error) {
	row, err := s.service.GetRole(ctx, req.RoleId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Role](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRole200JSONResponse{Body: body, Headers: apigen.GetRole200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRole implements the createRole operation.
func (s *Server) CreateRole(ctx context.Context, req apigen.CreateRoleRequestObject) (apigen.CreateRoleResponseObject, error) {
	in, err := input[db.CreateRoleParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRole(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Role](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRole201JSONResponse{Body: body, Headers: apigen.CreateRole201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRole implements the updateRole operation.
func (s *Server) UpdateRole(ctx context.Context, req apigen.UpdateRoleRequestObject) (apigen.UpdateRoleResponseObject, error) {
	in, err := input[db.UpdateRoleParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.RoleId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRole(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Role](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRole200JSONResponse{Body: body, Headers: apigen.UpdateRole200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DeleteRole implements the deleteRole operation.
func (s *Server) DeleteRole(ctx context.Context, req apigen.DeleteRoleRequestObject) (apigen.DeleteRoleResponseObject, error) {
	if err := s.service.DeleteRole(ctx, req.RoleId); err != nil {
		return nil, err
	}
	return apigen.DeleteRole204Response{}, nil
}

// ListRoleAssignments implements the listRoleAssignments operation.
func (s *Server) ListRoleAssignments(ctx context.Context, req apigen.ListRoleAssignmentsRequestObject) (apigen.ListRoleAssignmentsResponseObject, error) {
	return listPage[apigen.ListRoleAssignments200JSONResponse](ctx, req.Params, nil, s.service.ListRoleAssignment)
}

// GetRoleAssignment implements the getRoleAssignment operation.
func (s *Server) GetRoleAssignment(ctx context.Context, req apigen.GetRoleAssignmentRequestObject) (apigen.GetRoleAssignmentResponseObject, error) {
	row, err := s.service.GetRoleAssignment(ctx, req.RoleAssignmentId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RoleAssignment](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRoleAssignment200JSONResponse{Body: body, Headers: apigen.GetRoleAssignment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRoleAssignment implements the createRoleAssignment operation.
func (s *Server) CreateRoleAssignment(ctx context.Context, req apigen.CreateRoleAssignmentRequestObject) (apigen.CreateRoleAssignmentResponseObject, error) {
	in, err := input[db.CreateRoleAssignmentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRoleAssignment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RoleAssignment](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRoleAssignment201JSONResponse{Body: body, Headers: apigen.CreateRoleAssignment201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRoleAssignment implements the updateRoleAssignment operation.
func (s *Server) UpdateRoleAssignment(ctx context.Context, req apigen.UpdateRoleAssignmentRequestObject) (apigen.UpdateRoleAssignmentResponseObject, error) {
	in, err := input[db.UpdateRoleAssignmentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.RoleAssignmentId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRoleAssignment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RoleAssignment](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRoleAssignment200JSONResponse{Body: body, Headers: apigen.UpdateRoleAssignment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DeleteRoleAssignment implements the deleteRoleAssignment operation.
func (s *Server) DeleteRoleAssignment(ctx context.Context, req apigen.DeleteRoleAssignmentRequestObject) (apigen.DeleteRoleAssignmentResponseObject, error) {
	if err := s.service.DeleteRoleAssignment(ctx, req.RoleAssignmentId); err != nil {
		return nil, err
	}
	return apigen.DeleteRoleAssignment204Response{}, nil
}

// ListDelegations implements the listDelegations operation.
func (s *Server) ListDelegations(ctx context.Context, req apigen.ListDelegationsRequestObject) (apigen.ListDelegationsResponseObject, error) {
	return listPage[apigen.ListDelegations200JSONResponse](ctx, req.Params, nil, s.service.ListDelegation)
}

// GetDelegation implements the getDelegation operation.
func (s *Server) GetDelegation(ctx context.Context, req apigen.GetDelegationRequestObject) (apigen.GetDelegationResponseObject, error) {
	row, err := s.service.GetDelegation(ctx, req.DelegationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Delegation](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDelegation200JSONResponse{Body: body, Headers: apigen.GetDelegation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDelegation implements the createDelegation operation.
func (s *Server) CreateDelegation(ctx context.Context, req apigen.CreateDelegationRequestObject) (apigen.CreateDelegationResponseObject, error) {
	in, err := input[db.CreateDelegationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDelegation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Delegation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDelegation201JSONResponse{Body: body, Headers: apigen.CreateDelegation201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDelegation implements the updateDelegation operation.
func (s *Server) UpdateDelegation(ctx context.Context, req apigen.UpdateDelegationRequestObject) (apigen.UpdateDelegationResponseObject, error) {
	in, err := input[db.UpdateDelegationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.DelegationId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDelegation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Delegation](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDelegation200JSONResponse{Body: body, Headers: apigen.UpdateDelegation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListSodRules implements the listSodRules operation.
func (s *Server) ListSodRules(ctx context.Context, req apigen.ListSodRulesRequestObject) (apigen.ListSodRulesResponseObject, error) {
	return listPage[apigen.ListSodRules200JSONResponse](ctx, req.Params, nil, s.service.ListSodRule)
}

// GetSodRule implements the getSodRule operation.
func (s *Server) GetSodRule(ctx context.Context, req apigen.GetSodRuleRequestObject) (apigen.GetSodRuleResponseObject, error) {
	row, err := s.service.GetSodRule(ctx, req.SodRuleId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SodRule](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSodRule200JSONResponse{Body: body, Headers: apigen.GetSodRule200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateSodRule implements the createSodRule operation.
func (s *Server) CreateSodRule(ctx context.Context, req apigen.CreateSodRuleRequestObject) (apigen.CreateSodRuleResponseObject, error) {
	in, err := input[db.CreateSodRuleParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSodRule(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SodRule](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSodRule201JSONResponse{Body: body, Headers: apigen.CreateSodRule201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateSodRule implements the updateSodRule operation.
func (s *Server) UpdateSodRule(ctx context.Context, req apigen.UpdateSodRuleRequestObject) (apigen.UpdateSodRuleResponseObject, error) {
	in, err := input[db.UpdateSodRuleParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.SodRuleId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateSodRule(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SodRule](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateSodRule200JSONResponse{Body: body, Headers: apigen.UpdateSodRule200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDevices implements the listDevices operation.
func (s *Server) ListDevices(ctx context.Context, req apigen.ListDevicesRequestObject) (apigen.ListDevicesResponseObject, error) {
	return listPage[apigen.ListDevices200JSONResponse](ctx, req.Params, nil, s.service.ListDevice)
}

// GetDevice implements the getDevice operation.
func (s *Server) GetDevice(ctx context.Context, req apigen.GetDeviceRequestObject) (apigen.GetDeviceResponseObject, error) {
	row, err := s.service.GetDevice(ctx, req.DeviceId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Device](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDevice200JSONResponse{Body: body, Headers: apigen.GetDevice200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDevice implements the createDevice operation.
func (s *Server) CreateDevice(ctx context.Context, req apigen.CreateDeviceRequestObject) (apigen.CreateDeviceResponseObject, error) {
	in, err := input[db.CreateDeviceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDevice(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Device](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDevice201JSONResponse{Body: body, Headers: apigen.CreateDevice201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateDevice implements the updateDevice operation.
func (s *Server) UpdateDevice(ctx context.Context, req apigen.UpdateDeviceRequestObject) (apigen.UpdateDeviceResponseObject, error) {
	in, err := input[db.UpdateDeviceParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.DeviceId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateDevice(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Device](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateDevice200JSONResponse{Body: body, Headers: apigen.UpdateDevice200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListAuditEvents implements the listAuditEvents operation.
func (s *Server) ListAuditEvents(ctx context.Context, req apigen.ListAuditEventsRequestObject) (apigen.ListAuditEventsResponseObject, error) {
	return listPage[apigen.ListAuditEvents200JSONResponse](ctx, req.Params, nil, s.service.ListAuditEvent)
}

// GetAuditEvent implements the getAuditEvent operation.
func (s *Server) GetAuditEvent(ctx context.Context, req apigen.GetAuditEventRequestObject) (apigen.GetAuditEventResponseObject, error) {
	row, err := s.service.GetAuditEvent(ctx, req.AuditEventId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.AuditEvent](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetAuditEvent200JSONResponse{Body: body, Headers: apigen.GetAuditEvent200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
