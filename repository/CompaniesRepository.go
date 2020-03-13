package repository

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"location_service_v1/ls_v2/models"
	"net/http"
)

// CompaniesRepository is a
type CompaniesRepository struct {
}

// NewCompaniesRepository is a
func NewCompaniesRepository() *CompaniesRepository {
	return &CompaniesRepository{}
}

// List gets all Companies. This function is mapped to the path
// GET /companies
func (p *CompaniesRepository) List(c buffalo.Context) (*models.Companies, *pop.Query, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, nil, fmt.Errorf("no transaction found")
	}

	companies := &models.Companies{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Companies from the DB
	if err := q.All(companies); err != nil {
		return nil, nil, err
	}

	return companies, q, nil
}

// Show gets the data for one Point. This function is mapped to
// the path GET /companies/{company_id}
func (p *CompaniesRepository) Show(c buffalo.Context) (*models.Company, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty Company
	company := &models.Company{}

	// To find the Company the parameter company_id is used.
	if err := tx.Find(company, c.Param("company_id")); err != nil {
		return nil, c.Error(http.StatusNotFound, err)
	}

	return company, nil
}

// New renders the form for creating a new Company.
// This function is mapped to the path GET /points/new
func (p *CompaniesRepository) New(c buffalo.Context) *models.Company {
	return &models.Company{}
}

// Create adds a Company to the DB. This function is mapped to the
// path POST /companies
func (p *CompaniesRepository) Create(c buffalo.Context) (*validate.Errors, *models.Company, error) {
	// Allocate an empty Company
	company := &models.Company{}

	// Bind point to the html form elements
	if err := c.Bind(company); err != nil {
		return nil, nil, err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, nil, fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	created, err := tx.ValidateAndCreate(company)
	if err != nil {
		return nil, nil, err
	}

	return created, company, nil
}

// Edit renders a edit form for a Company. This function is
// mapped to the path GET /companies/{company_id}/edit
func (p *CompaniesRepository) Edit(c buffalo.Context) (*models.Company, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty Company
	company := &models.Company{}

	if err := tx.Find(company, c.Param("company_id")); err != nil {
		return nil, c.Error(http.StatusNotFound, err)
	}
	return company, nil
}

// Update changes a Company in the DB. This function is mapped to
// the path PUT /companies/{company_id}
func (p *CompaniesRepository) Update(c buffalo.Context) (*validate.Errors, *models.Company, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty Company
	company := &models.Company{}

	if err := tx.Find(company, c.Param("company_id")); err != nil {
		return nil, nil, c.Error(http.StatusNotFound, err)
	}

	// Bind Company to the html form elements
	if err := c.Bind(company); err != nil {
		return nil, nil, err
	}

	updated, err := tx.ValidateAndUpdate(company)
	if err != nil {
		return nil, nil, err
	}

	return updated, company, nil
}

// Destroy deletes a Company from the DB. This function is mapped
// to the path DELETE /companies/{company_id}
func (p *CompaniesRepository) Destroy(c buffalo.Context) (*models.Company, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty Company
	company := &models.Company{}

	// To find the Point the parameter company_id is used.
	if err := tx.Find(company, c.Param("company_id")); err != nil {
		return nil, c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(company); err != nil {
		return nil, err
	}

	return company, nil
}
