package service

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"location_service_v1/ls_v2/models"
	"location_service_v1/ls_v2/repository"
)

// PointsService is a
type PointsService struct {
	pointsRepository *repository.PointsRepository
}

// NewPointsService is a
func NewPointsService(repository *repository.PointsRepository) *PointsService {
	return &PointsService{
		pointsRepository: repository,
	}
}

// List is a
func (s *PointsService) List(c buffalo.Context) (*models.Points, *pop.Query, error) {

	points, q, err := s.pointsRepository.List(c)
	if err != nil {
		return nil, nil, err
	}

	return points, q, err
}

// Show gets the data for one Point. This function is mapped to
// the path GET /points/{point_id}
func (s *PointsService) Show(c buffalo.Context) (*models.Point, error) {
	point, err := s.pointsRepository.Show(c)
	if err != nil {
		return nil, err
	}
	return point, err
}

// New renders the form for creating a new Point.
// This function is mapped to the path GET /points/new
func (s *PointsService) New(c buffalo.Context) *models.Point {
	return s.pointsRepository.New(c)
}

// Create adds a Point to the DB. This function is mapped to the
// path POST /points
func (s *PointsService) Create(c buffalo.Context) (*validate.Errors, *models.Point, error) {
	create, point, err := s.pointsRepository.Create(c)
	if err != nil {
		return nil, nil, err
	}
	return create, point, nil
}

// Edit renders a edit form for a Point. This function is
// mapped to the path GET /points/{point_id}/edit
func (s *PointsService) Edit(c buffalo.Context) (*models.Point, error) {
	point, err := s.pointsRepository.Edit(c)
	if err != nil {
		return nil, err
	}
	return point, nil
}

// Update changes a Point in the DB. This function is mapped to
// the path PUT /points/{point_id}
func (s *PointsService) Update(c buffalo.Context) (*validate.Errors, *models.Point, error) {
	update, point, err := s.pointsRepository.Update(c)
	if err != nil {
		return nil, nil, err
	}
	return update, point, nil
}

// Destroy deletes a Point from the DB. This function is mapped
// to the path DELETE /points/{point_id}
func (s *PointsService) Destroy(c buffalo.Context) (*models.Point, error) {
	point, err := s.pointsRepository.Destroy(c)
	if err != nil {
		return nil, err
	}
	return point, nil
}
