package service

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"location_service_v1/ls_v2/models"
	"location_service_v1/ls_v2/repository"
)

// CompaniesService is a
type CompaniesService struct {
	companiesRepository *repository.CompaniesRepository
}

// NewCompaniesService is a
func NewCompaniesService(repository *repository.CompaniesRepository) *CompaniesService {
	return &CompaniesService{
		companiesRepository: repository,
	}
}

// List is a
func (s *CompaniesService) List(c buffalo.Context) (*models.Companies, *pop.Query, error) {

	companies, q, err := s.companiesRepository.List(c)
	if err != nil {
		return nil, nil, err
	}

	return companies, q, err
}

// Show gets the data for one Company. This function is mapped to
// the path GET /companies/{company_id}
func (s *CompaniesService) Show(c buffalo.Context) (*models.Company, error) {
	company, err := s.companiesRepository.Show(c)
	if err != nil {
		return nil, err
	}
	return company, err
}

// New renders the form for creating a new Company.
// This function is mapped to the path GET /companies/new
func (s *CompaniesService) New(c buffalo.Context) *models.Company {
	return s.companiesRepository.New(c)
}

// Create adds a Company to the DB. This function is mapped to the
// path POST /companies
func (s *CompaniesService) Create(c buffalo.Context) (*validate.Errors, *models.Company, error) {
	create, company, err := s.companiesRepository.Create(c)
	if err != nil {
		return nil, nil, err
	}
	return create, company, nil
}

// Edit renders a edit form for a Company. This function is
// mapped to the path GET /companies/{company_id}/edit
func (s *CompaniesService) Edit(c buffalo.Context) (*models.Company, error) {
	company, err := s.companiesRepository.Edit(c)
	if err != nil {
		return nil, err
	}
	return company, nil
}

// Update changes a Company in the DB. This function is mapped to
// the path PUT /companies/{company_id}
func (s *CompaniesService) Update(c buffalo.Context) (*validate.Errors, *models.Company, error) {
	update, company, err := s.companiesRepository.Update(c)
	if err != nil {
		return nil, nil, err
	}
	return update, company, nil
}

// Destroy deletes a Company from the DB. This function is mapped
// to the path DELETE /companies/{company_id}
func (s *CompaniesService) Destroy(c buffalo.Context) (*models.Company, error) {
	company, err := s.companiesRepository.Destroy(c)
	if err != nil {
		return nil, err
	}
	return company, nil
}
