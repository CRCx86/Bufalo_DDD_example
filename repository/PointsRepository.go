package repository

import (
	"encoding/json"
	"fmt"
	"github.com/gobuffalo/validate"
	"io/ioutil"
	"location_service_v1/ls_v2/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// PointsRepository is a
type PointsRepository struct {
}

// NewPointsRepository is a
func NewPointsRepository() *PointsRepository {
	return &PointsRepository{}
}

// List gets all Points. This function is mapped to the path
// GET /points
func (p *PointsRepository) List(c buffalo.Context) (*models.Points, *pop.Query, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, nil, fmt.Errorf("no transaction found")
	}

	points := &models.Points{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Points from the DB
	if err := q.All(points); err != nil {
		return nil, nil, err
	}

	return points, q, nil
}

// Show gets the data for one Point. This function is mapped to
// the path GET /points/{point_id}
func (p *PointsRepository) Show(c buffalo.Context) (*models.Point, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty Point
	point := &models.Point{}

	// To find the Point the parameter point_id is used.
	if err := tx.Find(point, c.Param("point_id")); err != nil {
		return nil, c.Error(http.StatusNotFound, err)
	}

	return point, nil
}

// New renders the form for creating a new Point.
// This function is mapped to the path GET /points/new
func (p *PointsRepository) New(c buffalo.Context) *models.Point {
	return &models.Point{}
}

// Create adds a Point to the DB. This function is mapped to the
// path POST /points
func (p *PointsRepository) Create(c buffalo.Context) (*validate.Errors, *models.Point, error) {
	// Allocate an empty Point
	point := &models.Point{}

	// Bind point to the html form elements
	if err := c.Bind(point); err != nil {
		return nil, nil, err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, nil, fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	created, err := tx.ValidateAndCreate(point)
	if err != nil {
		return nil, nil, err
	}

	return created, point, nil
}

// Edit renders a edit form for a Point. This function is
// mapped to the path GET /points/{point_id}/edit
func (p *PointsRepository) Edit(c buffalo.Context) (*models.Point, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty Point
	point := &models.Point{}

	if err := tx.Find(point, c.Param("point_id")); err != nil {
		return nil, c.Error(http.StatusNotFound, err)
	}
	return point, nil
}

// Update changes a Point in the DB. This function is mapped to
// the path PUT /points/{point_id}
func (p *PointsRepository) Update(c buffalo.Context) (*validate.Errors, *models.Point, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty Point
	point := &models.Point{}

	if err := tx.Find(point, c.Param("point_id")); err != nil {
		return nil, nil, c.Error(http.StatusNotFound, err)
	}

	// Bind Point to the html form elements
	if err := c.Bind(point); err != nil {
		return nil, nil, err
	}

	updated, err := tx.ValidateAndUpdate(point)
	if err != nil {
		return nil, nil, err
	}

	return updated, point, nil
}

// Destroy deletes a Point from the DB. This function is mapped
// to the path DELETE /points/{point_id}
func (p *PointsRepository) Destroy(c buffalo.Context) (*models.Point, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty Point
	point := &models.Point{}

	// To find the Point the parameter point_id is used.
	if err := tx.Find(point, c.Param("point_id")); err != nil {
		return nil, c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(point); err != nil {
		return nil, err
	}

	return point, nil
}

func (p *PointsRepository) PickPointsList(c buffalo.Context) (*validate.Errors, error) {

	resp, err := http.Get("http://e-solution.pickpoint.ru/api/postamatlist")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var points models.Points
	err = json.Unmarshal(body, &points)
	if err != nil {
		return nil, err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	created, _ := tx.ValidateAndCreate(points)
	if created != nil {
		return nil, err
	}

	return created, nil

}
