// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListComplex returns a page of complexes.
func (s *Service) ListComplex(ctx context.Context, p db.ListComplexParams) ([]db.Complex, error) {
	return wrap2(s.q(ctx).ListComplex(ctx, p))("list complexes")
}

// GetComplex returns one row of complexes.
func (s *Service) GetComplex(ctx context.Context, id uuid.UUID) (db.Complex, error) {
	row, err := s.q(ctx).GetComplex(ctx, id)
	if err != nil {
		return row, notFound(err, "complexes")
	}
	return row, nil
}

// CreateComplex adds one row to complexes.
func (s *Service) CreateComplex(ctx context.Context, p db.CreateComplexParams) (db.Complex, error) {
	var zero db.Complex
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateComplex(ctx, p)
	if err != nil {
		return zero, wrap(err, "create complexes")
	}
	return row, s.audit(ctx, "org.complex.write", "complexes", row.ID, nil, nil, row)
}

// UpdateComplex changes one row of complexes.
func (s *Service) UpdateComplex(ctx context.Context, p db.UpdateComplexParams) (db.Complex, error) {
	before, err := s.GetComplex(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateComplex(ctx, p)
	if err != nil {
		return before, wrap(err, "update complexes")
	}
	if err := s.voidSignatures(ctx, "complexes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.complex.write", "complexes", row.ID, nil, before, row)
}

// ListSite returns a page of sites.
func (s *Service) ListSite(ctx context.Context, p db.ListSiteParams) ([]db.Site, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListSite(ctx, p))("list sites")
}

// GetSite returns one row of sites.
func (s *Service) GetSite(ctx context.Context, id uuid.UUID) (db.Site, error) {
	row, err := s.q(ctx).GetSite(ctx, id)
	if err != nil {
		return row, notFound(err, "sites")
	}
	return row, inScope(ctx, &row.ID)
}

// CreateSite adds one row to sites.
func (s *Service) CreateSite(ctx context.Context, p db.CreateSiteParams) (db.Site, error) {
	var zero db.Site
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateSite(ctx, p)
	if err != nil {
		return zero, wrap(err, "create sites")
	}
	return row, s.audit(ctx, "org.site.write", "sites", row.ID, nil, nil, row)
}

// UpdateSite changes one row of sites.
func (s *Service) UpdateSite(ctx context.Context, p db.UpdateSiteParams) (db.Site, error) {
	before, err := s.GetSite(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateSite(ctx, p)
	if err != nil {
		return before, wrap(err, "update sites")
	}
	if err := s.voidSignatures(ctx, "sites", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.site.write", "sites", row.ID, &row.ID, before, row)
}

// ListDepartment returns a page of departments.
func (s *Service) ListDepartment(ctx context.Context, p db.ListDepartmentParams) ([]db.Department, error) {
	return wrap2(s.q(ctx).ListDepartment(ctx, p))("list departments")
}

// GetDepartment returns one row of departments.
func (s *Service) GetDepartment(ctx context.Context, id uuid.UUID) (db.Department, error) {
	row, err := s.q(ctx).GetDepartment(ctx, id)
	if err != nil {
		return row, notFound(err, "departments")
	}
	return row, nil
}

// CreateDepartment adds one row to departments.
func (s *Service) CreateDepartment(ctx context.Context, p db.CreateDepartmentParams) (db.Department, error) {
	var zero db.Department
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDepartment(ctx, p)
	if err != nil {
		return zero, wrap(err, "create departments")
	}
	return row, s.audit(ctx, "org.department.write", "departments", row.ID, nil, nil, row)
}

// UpdateDepartment changes one row of departments.
func (s *Service) UpdateDepartment(ctx context.Context, p db.UpdateDepartmentParams) (db.Department, error) {
	before, err := s.GetDepartment(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDepartment(ctx, p)
	if err != nil {
		return before, wrap(err, "update departments")
	}
	if err := s.voidSignatures(ctx, "departments", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.department.write", "departments", row.ID, nil, before, row)
}

// ListCostCentre returns a page of cost centres.
func (s *Service) ListCostCentre(ctx context.Context, p db.ListCostCentreParams) ([]db.CostCentre, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListCostCentre(ctx, p))("list cost centres")
}

// GetCostCentre returns one row of cost centres.
func (s *Service) GetCostCentre(ctx context.Context, id uuid.UUID) (db.CostCentre, error) {
	row, err := s.q(ctx).GetCostCentre(ctx, id)
	if err != nil {
		return row, notFound(err, "cost centres")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateCostCentre adds one row to cost centres.
func (s *Service) CreateCostCentre(ctx context.Context, p db.CreateCostCentreParams) (db.CostCentre, error) {
	var zero db.CostCentre
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCostCentre(ctx, p)
	if err != nil {
		return zero, wrap(err, "create cost centres")
	}
	return row, s.audit(ctx, "org.cost_centre.write", "cost_centres", row.ID, p.SiteID, nil, row)
}

// UpdateCostCentre changes one row of cost centres.
func (s *Service) UpdateCostCentre(ctx context.Context, p db.UpdateCostCentreParams) (db.CostCentre, error) {
	before, err := s.GetCostCentre(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCostCentre(ctx, p)
	if err != nil {
		return before, wrap(err, "update cost centres")
	}
	if err := s.voidSignatures(ctx, "cost_centres", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.cost_centre.write", "cost_centres", row.ID, row.SiteID, before, row)
}

// ListProductionLine returns a page of production lines.
func (s *Service) ListProductionLine(ctx context.Context, p db.ListProductionLineParams) ([]db.ProductionLine, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListProductionLine(ctx, p))("list production lines")
}

// GetProductionLine returns one row of production lines.
func (s *Service) GetProductionLine(ctx context.Context, id uuid.UUID) (db.ProductionLine, error) {
	row, err := s.q(ctx).GetProductionLine(ctx, id)
	if err != nil {
		return row, notFound(err, "production lines")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateProductionLine adds one row to production lines.
func (s *Service) CreateProductionLine(ctx context.Context, p db.CreateProductionLineParams) (db.ProductionLine, error) {
	var zero db.ProductionLine
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateProductionLine(ctx, p)
	if err != nil {
		return zero, wrap(err, "create production lines")
	}
	return row, s.audit(ctx, "org.production_line.write", "production_lines", row.ID, &p.SiteID, nil, row)
}

// UpdateProductionLine changes one row of production lines.
func (s *Service) UpdateProductionLine(ctx context.Context, p db.UpdateProductionLineParams) (db.ProductionLine, error) {
	before, err := s.GetProductionLine(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateProductionLine(ctx, p)
	if err != nil {
		return before, wrap(err, "update production lines")
	}
	if err := s.voidSignatures(ctx, "production_lines", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.production_line.write", "production_lines", row.ID, &row.SiteID, before, row)
}

// ListTill returns a page of tills.
func (s *Service) ListTill(ctx context.Context, p db.ListTillParams) ([]db.Till, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListTill(ctx, p))("list tills")
}

// GetTill returns one row of tills.
func (s *Service) GetTill(ctx context.Context, id uuid.UUID) (db.Till, error) {
	row, err := s.q(ctx).GetTill(ctx, id)
	if err != nil {
		return row, notFound(err, "tills")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateTill adds one row to tills.
func (s *Service) CreateTill(ctx context.Context, p db.CreateTillParams) (db.Till, error) {
	var zero db.Till
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateTill(ctx, p)
	if err != nil {
		return zero, wrap(err, "create tills")
	}
	return row, s.audit(ctx, "org.till.write", "tills", row.ID, &p.SiteID, nil, row)
}

// UpdateTill changes one row of tills.
func (s *Service) UpdateTill(ctx context.Context, p db.UpdateTillParams) (db.Till, error) {
	before, err := s.GetTill(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateTill(ctx, p)
	if err != nil {
		return before, wrap(err, "update tills")
	}
	if err := s.voidSignatures(ctx, "tills", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.till.write", "tills", row.ID, &row.SiteID, before, row)
}

// ListSafe returns a page of safes.
func (s *Service) ListSafe(ctx context.Context, p db.ListSafeParams) ([]db.Safe, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListSafe(ctx, p))("list safes")
}

// GetSafe returns one row of safes.
func (s *Service) GetSafe(ctx context.Context, id uuid.UUID) (db.Safe, error) {
	row, err := s.q(ctx).GetSafe(ctx, id)
	if err != nil {
		return row, notFound(err, "safes")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateSafe adds one row to safes.
func (s *Service) CreateSafe(ctx context.Context, p db.CreateSafeParams) (db.Safe, error) {
	var zero db.Safe
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateSafe(ctx, p)
	if err != nil {
		return zero, wrap(err, "create safes")
	}
	return row, s.audit(ctx, "org.safe.write", "safes", row.ID, &p.SiteID, nil, row)
}

// UpdateSafe changes one row of safes.
func (s *Service) UpdateSafe(ctx context.Context, p db.UpdateSafeParams) (db.Safe, error) {
	before, err := s.GetSafe(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateSafe(ctx, p)
	if err != nil {
		return before, wrap(err, "update safes")
	}
	if err := s.voidSignatures(ctx, "safes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.safe.write", "safes", row.ID, &row.SiteID, before, row)
}

// ListStorageLocation returns a page of storage locations.
func (s *Service) ListStorageLocation(ctx context.Context, p db.ListStorageLocationParams) ([]db.StorageLocation, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListStorageLocation(ctx, p))("list storage locations")
}

// GetStorageLocation returns one row of storage locations.
func (s *Service) GetStorageLocation(ctx context.Context, id uuid.UUID) (db.StorageLocation, error) {
	row, err := s.q(ctx).GetStorageLocation(ctx, id)
	if err != nil {
		return row, notFound(err, "storage locations")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateStorageLocation adds one row to storage locations.
func (s *Service) CreateStorageLocation(ctx context.Context, p db.CreateStorageLocationParams) (db.StorageLocation, error) {
	var zero db.StorageLocation
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateStorageLocation(ctx, p)
	if err != nil {
		return zero, wrap(err, "create storage locations")
	}
	return row, s.audit(ctx, "org.storage_location.write", "storage_locations", row.ID, &p.SiteID, nil, row)
}

// UpdateStorageLocation changes one row of storage locations.
func (s *Service) UpdateStorageLocation(ctx context.Context, p db.UpdateStorageLocationParams) (db.StorageLocation, error) {
	before, err := s.GetStorageLocation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateStorageLocation(ctx, p)
	if err != nil {
		return before, wrap(err, "update storage locations")
	}
	if err := s.voidSignatures(ctx, "storage_locations", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.storage_location.write", "storage_locations", row.ID, &row.SiteID, before, row)
}

// ListCurrency returns a page of currencies.
func (s *Service) ListCurrency(ctx context.Context, p db.ListCurrencyParams) ([]db.Currency, error) {
	return wrap2(s.q(ctx).ListCurrency(ctx, p))("list currencies")
}

// GetCurrency returns one row of currencies.
func (s *Service) GetCurrency(ctx context.Context, id uuid.UUID) (db.Currency, error) {
	row, err := s.q(ctx).GetCurrency(ctx, id)
	if err != nil {
		return row, notFound(err, "currencies")
	}
	return row, nil
}

// CreateCurrency adds one row to currencies.
func (s *Service) CreateCurrency(ctx context.Context, p db.CreateCurrencyParams) (db.Currency, error) {
	var zero db.Currency
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCurrency(ctx, p)
	if err != nil {
		return zero, wrap(err, "create currencies")
	}
	return row, s.audit(ctx, "org.currency.write", "currencies", row.ID, nil, nil, row)
}

// UpdateCurrency changes one row of currencies.
func (s *Service) UpdateCurrency(ctx context.Context, p db.UpdateCurrencyParams) (db.Currency, error) {
	before, err := s.GetCurrency(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCurrency(ctx, p)
	if err != nil {
		return before, wrap(err, "update currencies")
	}
	if err := s.voidSignatures(ctx, "currencies", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.currency.write", "currencies", row.ID, nil, before, row)
}

// ListFxRate returns a page of fx rates.
func (s *Service) ListFxRate(ctx context.Context, p db.ListFxRateParams) ([]db.FxRate, error) {
	return wrap2(s.q(ctx).ListFxRate(ctx, p))("list fx rates")
}

// GetFxRate returns one row of fx rates.
func (s *Service) GetFxRate(ctx context.Context, id uuid.UUID) (db.FxRate, error) {
	row, err := s.q(ctx).GetFxRate(ctx, id)
	if err != nil {
		return row, notFound(err, "fx rates")
	}
	return row, nil
}

// CreateFxRate adds one row to fx rates.
func (s *Service) CreateFxRate(ctx context.Context, p db.CreateFxRateParams) (db.FxRate, error) {
	var zero db.FxRate
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateFxRate(ctx, p)
	if err != nil {
		return zero, wrap(err, "create fx rates")
	}
	return row, s.audit(ctx, "org.fx_rate.write", "fx_rates", row.ID, nil, nil, row)
}

// ListPublicHoliday returns a page of public holidays.
func (s *Service) ListPublicHoliday(ctx context.Context, p db.ListPublicHolidayParams) ([]db.PublicHoliday, error) {
	return wrap2(s.q(ctx).ListPublicHoliday(ctx, p))("list public holidays")
}

// GetPublicHoliday returns one row of public holidays.
func (s *Service) GetPublicHoliday(ctx context.Context, id uuid.UUID) (db.PublicHoliday, error) {
	row, err := s.q(ctx).GetPublicHoliday(ctx, id)
	if err != nil {
		return row, notFound(err, "public holidays")
	}
	return row, nil
}

// CreatePublicHoliday adds one row to public holidays.
func (s *Service) CreatePublicHoliday(ctx context.Context, p db.CreatePublicHolidayParams) (db.PublicHoliday, error) {
	var zero db.PublicHoliday
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePublicHoliday(ctx, p)
	if err != nil {
		return zero, wrap(err, "create public holidays")
	}
	return row, s.audit(ctx, "org.public_holiday.write", "public_holidays", row.ID, nil, nil, row)
}

// UpdatePublicHoliday changes one row of public holidays.
func (s *Service) UpdatePublicHoliday(ctx context.Context, p db.UpdatePublicHolidayParams) (db.PublicHoliday, error) {
	before, err := s.GetPublicHoliday(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePublicHoliday(ctx, p)
	if err != nil {
		return before, wrap(err, "update public holidays")
	}
	if err := s.voidSignatures(ctx, "public_holidays", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "org.public_holiday.write", "public_holidays", row.ID, nil, before, row)
}

// DeletePublicHoliday removes one row of public holidays.
func (s *Service) DeletePublicHoliday(ctx context.Context, id uuid.UUID) error {
	before, err := s.GetPublicHoliday(ctx, id)
	if err != nil {
		return err
	}
	if err := s.q(ctx).DeletePublicHoliday(ctx, id); err != nil {
		return wrap(err, "delete public holidays")
	}
	return s.audit(ctx, "public_holidays.delete", "public_holidays", id, nil, before, nil)
}

// ListEmployee returns a page of employees.
func (s *Service) ListEmployee(ctx context.Context, p db.ListEmployeeParams) ([]db.Employee, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListEmployee(ctx, p))("list employees")
}

// GetEmployee returns one row of employees.
func (s *Service) GetEmployee(ctx context.Context, id uuid.UUID) (db.Employee, error) {
	row, err := s.q(ctx).GetEmployee(ctx, id)
	if err != nil {
		return row, notFound(err, "employees")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateEmployee adds one row to employees.
func (s *Service) CreateEmployee(ctx context.Context, p db.CreateEmployeeParams) (db.Employee, error) {
	var zero db.Employee
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "EMP", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateEmployee(ctx, p)
	if err != nil {
		return zero, wrap(err, "create employees")
	}
	return row, s.audit(ctx, "iam.employee.write", "employees", row.ID, &p.SiteID, nil, row)
}

// UpdateEmployee changes one row of employees.
func (s *Service) UpdateEmployee(ctx context.Context, p db.UpdateEmployeeParams) (db.Employee, error) {
	before, err := s.GetEmployee(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateEmployee(ctx, p)
	if err != nil {
		return before, wrap(err, "update employees")
	}
	if err := s.voidSignatures(ctx, "employees", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "iam.employee.write", "employees", row.ID, &row.SiteID, before, row)
}

// ListEmployeeStatusChange returns a page of employee status changes.
func (s *Service) ListEmployeeStatusChange(ctx context.Context, p db.ListEmployeeStatusChangeParams) ([]db.EmployeeStatusChange, error) {
	return wrap2(s.q(ctx).ListEmployeeStatusChange(ctx, p))("list employee status changes")
}

// GetEmployeeStatusChange returns one row of employee status changes.
func (s *Service) GetEmployeeStatusChange(ctx context.Context, id uuid.UUID) (db.EmployeeStatusChange, error) {
	row, err := s.q(ctx).GetEmployeeStatusChange(ctx, id)
	if err != nil {
		return row, notFound(err, "employee status changes")
	}
	return row, nil
}

// ListEmployeeTransfer returns a page of employee transfers.
func (s *Service) ListEmployeeTransfer(ctx context.Context, p db.ListEmployeeTransferParams) ([]db.EmployeeTransfer, error) {
	return wrap2(s.q(ctx).ListEmployeeTransfer(ctx, p))("list employee transfers")
}

// GetEmployeeTransfer returns one row of employee transfers.
func (s *Service) GetEmployeeTransfer(ctx context.Context, id uuid.UUID) (db.EmployeeTransfer, error) {
	row, err := s.q(ctx).GetEmployeeTransfer(ctx, id)
	if err != nil {
		return row, notFound(err, "employee transfers")
	}
	return row, nil
}

// CreateEmployeeTransfer adds one row to employee transfers.
func (s *Service) CreateEmployeeTransfer(ctx context.Context, p db.CreateEmployeeTransferParams) (db.EmployeeTransfer, error) {
	var zero db.EmployeeTransfer
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateEmployeeTransfer(ctx, p)
	if err != nil {
		return zero, wrap(err, "create employee transfers")
	}
	return row, s.audit(ctx, "iam.employee_transfer.write", "employee_transfers", row.ID, nil, nil, row)
}

// SubmitEmployeeTransfer is the submit for approval action on employee transfers.
func (s *Service) SubmitEmployeeTransfer(ctx context.Context, p db.ActSubmitEmployeeTransferParams) (db.EmployeeTransfer, error) {
	before, err := s.GetEmployeeTransfer(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "iam.employee_transfer.submit", "employee_transfers", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitEmployeeTransfer(ctx, p)
	if err != nil {
		return before, stateError(err, "submit employee transfers")
	}
	if err := s.audit(ctx, "iam.employee_transfer.submit", "employee_transfers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.startApproval(ctx, "employee_transfers", row.ID, nil, row)
}

// ListEmployeeDocument returns a page of employee documents.
func (s *Service) ListEmployeeDocument(ctx context.Context, p db.ListEmployeeDocumentParams) ([]db.EmployeeDocument, error) {
	return wrap2(s.q(ctx).ListEmployeeDocument(ctx, p))("list employee documents")
}

// GetEmployeeDocument returns one row of employee documents.
func (s *Service) GetEmployeeDocument(ctx context.Context, id uuid.UUID) (db.EmployeeDocument, error) {
	row, err := s.q(ctx).GetEmployeeDocument(ctx, id)
	if err != nil {
		return row, notFound(err, "employee documents")
	}
	return row, nil
}

// CreateEmployeeDocument adds one row to employee documents.
func (s *Service) CreateEmployeeDocument(ctx context.Context, p db.CreateEmployeeDocumentParams) (db.EmployeeDocument, error) {
	var zero db.EmployeeDocument
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateEmployeeDocument(ctx, p)
	if err != nil {
		return zero, wrap(err, "create employee documents")
	}
	return row, s.audit(ctx, "iam.employee_document.write", "employee_documents", row.ID, nil, nil, row)
}

// DeleteEmployeeDocument removes one row of employee documents.
func (s *Service) DeleteEmployeeDocument(ctx context.Context, id uuid.UUID) error {
	before, err := s.GetEmployeeDocument(ctx, id)
	if err != nil {
		return err
	}
	if err := s.q(ctx).DeleteEmployeeDocument(ctx, id); err != nil {
		return wrap(err, "delete employee documents")
	}
	return s.audit(ctx, "employee_documents.delete", "employee_documents", id, nil, before, nil)
}

// ListEmergencyContact returns a page of emergency contacts.
func (s *Service) ListEmergencyContact(ctx context.Context, p db.ListEmergencyContactParams) ([]db.EmergencyContact, error) {
	return wrap2(s.q(ctx).ListEmergencyContact(ctx, p))("list emergency contacts")
}

// GetEmergencyContact returns one row of emergency contacts.
func (s *Service) GetEmergencyContact(ctx context.Context, id uuid.UUID) (db.EmergencyContact, error) {
	row, err := s.q(ctx).GetEmergencyContact(ctx, id)
	if err != nil {
		return row, notFound(err, "emergency contacts")
	}
	return row, nil
}

// CreateEmergencyContact adds one row to emergency contacts.
func (s *Service) CreateEmergencyContact(ctx context.Context, p db.CreateEmergencyContactParams) (db.EmergencyContact, error) {
	var zero db.EmergencyContact
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateEmergencyContact(ctx, p)
	if err != nil {
		return zero, wrap(err, "create emergency contacts")
	}
	return row, s.audit(ctx, "iam.emergency_contact.write", "emergency_contacts", row.ID, nil, nil, row)
}

// UpdateEmergencyContact changes one row of emergency contacts.
func (s *Service) UpdateEmergencyContact(ctx context.Context, p db.UpdateEmergencyContactParams) (db.EmergencyContact, error) {
	before, err := s.GetEmergencyContact(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateEmergencyContact(ctx, p)
	if err != nil {
		return before, wrap(err, "update emergency contacts")
	}
	if err := s.voidSignatures(ctx, "emergency_contacts", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "iam.emergency_contact.write", "emergency_contacts", row.ID, nil, before, row)
}

// DeleteEmergencyContact removes one row of emergency contacts.
func (s *Service) DeleteEmergencyContact(ctx context.Context, id uuid.UUID) error {
	before, err := s.GetEmergencyContact(ctx, id)
	if err != nil {
		return err
	}
	if err := s.q(ctx).DeleteEmergencyContact(ctx, id); err != nil {
		return wrap(err, "delete emergency contacts")
	}
	return s.audit(ctx, "emergency_contacts.delete", "emergency_contacts", id, nil, before, nil)
}

// ListUser returns a page of users.
func (s *Service) ListUser(ctx context.Context, p db.ListUserParams) ([]db.User, error) {
	return wrap2(s.q(ctx).ListUser(ctx, p))("list users")
}

// GetUser returns one row of users.
func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (db.User, error) {
	row, err := s.q(ctx).GetUser(ctx, id)
	if err != nil {
		return row, notFound(err, "users")
	}
	return row, nil
}

// UpdateUser changes one row of users.
func (s *Service) UpdateUser(ctx context.Context, p db.UpdateUserParams) (db.User, error) {
	before, err := s.GetUser(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateUser(ctx, p)
	if err != nil {
		return before, wrap(err, "update users")
	}
	if err := s.voidSignatures(ctx, "users", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "iam.user.write", "users", row.ID, nil, before, row)
}

// ListRole returns a page of roles.
func (s *Service) ListRole(ctx context.Context, p db.ListRoleParams) ([]db.Role, error) {
	return wrap2(s.q(ctx).ListRole(ctx, p))("list roles")
}

// GetRole returns one row of roles.
func (s *Service) GetRole(ctx context.Context, id uuid.UUID) (db.Role, error) {
	row, err := s.q(ctx).GetRole(ctx, id)
	if err != nil {
		return row, notFound(err, "roles")
	}
	return row, nil
}

// CreateRole adds one row to roles.
func (s *Service) CreateRole(ctx context.Context, p db.CreateRoleParams) (db.Role, error) {
	var zero db.Role
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRole(ctx, p)
	if err != nil {
		return zero, wrap(err, "create roles")
	}
	return row, s.audit(ctx, "iam.role.write", "roles", row.ID, nil, nil, row)
}

// UpdateRole changes one row of roles.
func (s *Service) UpdateRole(ctx context.Context, p db.UpdateRoleParams) (db.Role, error) {
	before, err := s.GetRole(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRole(ctx, p)
	if err != nil {
		return before, wrap(err, "update roles")
	}
	if err := s.voidSignatures(ctx, "roles", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "iam.role.write", "roles", row.ID, nil, before, row)
}

// ListRoleAssignment returns a page of role assignments.
func (s *Service) ListRoleAssignment(ctx context.Context, p db.ListRoleAssignmentParams) ([]db.RoleAssignment, error) {
	return wrap2(s.q(ctx).ListRoleAssignment(ctx, p))("list role assignments")
}

// GetRoleAssignment returns one row of role assignments.
func (s *Service) GetRoleAssignment(ctx context.Context, id uuid.UUID) (db.RoleAssignment, error) {
	row, err := s.q(ctx).GetRoleAssignment(ctx, id)
	if err != nil {
		return row, notFound(err, "role assignments")
	}
	return row, nil
}

// UpdateRoleAssignment changes one row of role assignments.
func (s *Service) UpdateRoleAssignment(ctx context.Context, p db.UpdateRoleAssignmentParams) (db.RoleAssignment, error) {
	before, err := s.GetRoleAssignment(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRoleAssignment(ctx, p)
	if err != nil {
		return before, wrap(err, "update role assignments")
	}
	if err := s.voidSignatures(ctx, "role_assignments", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "iam.role_assignment.write", "role_assignments", row.ID, nil, before, row)
}

// DeleteRoleAssignment removes one row of role assignments.
func (s *Service) DeleteRoleAssignment(ctx context.Context, id uuid.UUID) error {
	before, err := s.GetRoleAssignment(ctx, id)
	if err != nil {
		return err
	}
	if err := s.q(ctx).DeleteRoleAssignment(ctx, id); err != nil {
		return wrap(err, "delete role assignments")
	}
	return s.audit(ctx, "role_assignments.delete", "role_assignments", id, nil, before, nil)
}

// ListDelegation returns a page of delegations.
func (s *Service) ListDelegation(ctx context.Context, p db.ListDelegationParams) ([]db.Delegation, error) {
	return wrap2(s.q(ctx).ListDelegation(ctx, p))("list delegations")
}

// GetDelegation returns one row of delegations.
func (s *Service) GetDelegation(ctx context.Context, id uuid.UUID) (db.Delegation, error) {
	row, err := s.q(ctx).GetDelegation(ctx, id)
	if err != nil {
		return row, notFound(err, "delegations")
	}
	return row, nil
}

// CreateDelegation adds one row to delegations.
func (s *Service) CreateDelegation(ctx context.Context, p db.CreateDelegationParams) (db.Delegation, error) {
	var zero db.Delegation
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDelegation(ctx, p)
	if err != nil {
		return zero, wrap(err, "create delegations")
	}
	return row, s.audit(ctx, "iam.delegation.write", "delegations", row.ID, nil, nil, row)
}

// UpdateDelegation changes one row of delegations.
func (s *Service) UpdateDelegation(ctx context.Context, p db.UpdateDelegationParams) (db.Delegation, error) {
	before, err := s.GetDelegation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDelegation(ctx, p)
	if err != nil {
		return before, wrap(err, "update delegations")
	}
	if err := s.voidSignatures(ctx, "delegations", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "iam.delegation.write", "delegations", row.ID, nil, before, row)
}

// SubmitDelegation is the submit for approval action on delegations.
func (s *Service) SubmitDelegation(ctx context.Context, p db.ActSubmitDelegationParams) (db.Delegation, error) {
	before, err := s.GetDelegation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "iam.delegation.submit", "delegations", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitDelegation(ctx, p)
	if err != nil {
		return before, stateError(err, "submit delegations")
	}
	if err := s.audit(ctx, "iam.delegation.submit", "delegations", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.startApproval(ctx, "delegations", row.ID, nil, row)
}

// RevokeDelegation is the end the delegation early action on delegations.
func (s *Service) RevokeDelegation(ctx context.Context, p db.ActRevokeDelegationParams) (db.Delegation, error) {
	before, err := s.GetDelegation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "iam.delegation.revoke", "delegations", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRevokeDelegation(ctx, p)
	if err != nil {
		return before, stateError(err, "revoke delegations")
	}
	if err := s.audit(ctx, "iam.delegation.revoke", "delegations", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListSodRule returns a page of sod rules.
func (s *Service) ListSodRule(ctx context.Context, p db.ListSodRuleParams) ([]db.SodRule, error) {
	return wrap2(s.q(ctx).ListSodRule(ctx, p))("list sod rules")
}

// GetSodRule returns one row of sod rules.
func (s *Service) GetSodRule(ctx context.Context, id uuid.UUID) (db.SodRule, error) {
	row, err := s.q(ctx).GetSodRule(ctx, id)
	if err != nil {
		return row, notFound(err, "sod rules")
	}
	return row, nil
}

// CreateSodRule adds one row to sod rules.
func (s *Service) CreateSodRule(ctx context.Context, p db.CreateSodRuleParams) (db.SodRule, error) {
	var zero db.SodRule
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateSodRule(ctx, p)
	if err != nil {
		return zero, wrap(err, "create sod rules")
	}
	return row, s.audit(ctx, "iam.sod_rule.write", "sod_rules", row.ID, nil, nil, row)
}

// UpdateSodRule changes one row of sod rules.
func (s *Service) UpdateSodRule(ctx context.Context, p db.UpdateSodRuleParams) (db.SodRule, error) {
	before, err := s.GetSodRule(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateSodRule(ctx, p)
	if err != nil {
		return before, wrap(err, "update sod rules")
	}
	if err := s.voidSignatures(ctx, "sod_rules", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "iam.sod_rule.write", "sod_rules", row.ID, nil, before, row)
}

// ListDevice returns a page of devices.
func (s *Service) ListDevice(ctx context.Context, p db.ListDeviceParams) ([]db.Device, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListDevice(ctx, p))("list devices")
}

// GetDevice returns one row of devices.
func (s *Service) GetDevice(ctx context.Context, id uuid.UUID) (db.Device, error) {
	row, err := s.q(ctx).GetDevice(ctx, id)
	if err != nil {
		return row, notFound(err, "devices")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateDevice adds one row to devices.
func (s *Service) CreateDevice(ctx context.Context, p db.CreateDeviceParams) (db.Device, error) {
	var zero db.Device
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDevice(ctx, p)
	if err != nil {
		return zero, wrap(err, "create devices")
	}
	return row, s.audit(ctx, "iam.device.write", "devices", row.ID, p.SiteID, nil, row)
}

// UpdateDevice changes one row of devices.
func (s *Service) UpdateDevice(ctx context.Context, p db.UpdateDeviceParams) (db.Device, error) {
	before, err := s.GetDevice(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateDevice(ctx, p)
	if err != nil {
		return before, wrap(err, "update devices")
	}
	if err := s.voidSignatures(ctx, "devices", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "iam.device.write", "devices", row.ID, row.SiteID, before, row)
}

// RevokeDevice is the revoke the device and its offline data action on devices.
func (s *Service) RevokeDevice(ctx context.Context, p db.ActRevokeDeviceParams) (db.Device, error) {
	before, err := s.GetDevice(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "iam.device.revoke", "devices", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRevokeDevice(ctx, p)
	if err != nil {
		return before, stateError(err, "revoke devices")
	}
	if err := s.audit(ctx, "iam.device.revoke", "devices", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListAuditEvent returns a page of audit events.
func (s *Service) ListAuditEvent(ctx context.Context, p db.ListAuditEventParams) ([]db.AuditEvent, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListAuditEvent(ctx, p))("list audit events")
}

// GetAuditEvent returns one row of audit events.
func (s *Service) GetAuditEvent(ctx context.Context, id uuid.UUID) (db.AuditEvent, error) {
	row, err := s.q(ctx).GetAuditEvent(ctx, id)
	if err != nil {
		return row, notFound(err, "audit events")
	}
	return row, inScope(ctx, row.SiteID)
}
