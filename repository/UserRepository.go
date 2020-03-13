package repository

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"location_service_v1/ls_v2/models"
	"net/http"
)

// UsersRepository is a
type UsersRepository struct {
}

// NewUsersRepository is a
func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

// List gets all Users. This function is mapped to the path
// GET /users
func (p *UsersRepository) List(c buffalo.Context) (*models.Users, *pop.Query, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, nil, fmt.Errorf("no transaction found")
	}

	users := &models.Users{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Companies from the DB
	if err := q.All(users); err != nil {
		return nil, nil, err
	}

	return users, q, nil
}

// Show gets the data for one Point. This function is mapped to
// the path GET /users/{user_id}
func (p *UsersRepository) Show(c buffalo.Context) (*models.User, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return nil, c.Error(http.StatusNotFound, err)
	}

	return user, nil
}

// New renders the form for creating a new User.
// This function is mapped to the path GET /users/new
func (p *UsersRepository) New(c buffalo.Context) *models.User {
	return &models.User{}
}

// Create adds a User to the DB. This function is mapped to the
// path POST /users
func (p *UsersRepository) Create(c buffalo.Context) (*validate.Errors, *models.User, error) {
	// Allocate an empty User
	user := &models.User{}

	// Bind user to the html form elements
	if err := c.Bind(user); err != nil {
		return nil, nil, err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, nil, fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	created, err := tx.ValidateAndCreate(user)
	if err != nil {
		return nil, nil, err
	}

	return created, user, nil
}

// Edit renders a edit form for a User. This function is
// mapped to the path GET /users/{user_id}/edit
func (p *UsersRepository) Edit(c buffalo.Context) (*models.User, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return nil, c.Error(http.StatusNotFound, err)
	}
	return user, nil
}

// Update changes a User in the DB. This function is mapped to
// the path PUT /users/{user_id}
func (p *UsersRepository) Update(c buffalo.Context) (*validate.Errors, *models.User, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return nil, nil, c.Error(http.StatusNotFound, err)
	}

	// Bind User to the html form elements
	if err := c.Bind(user); err != nil {
		return nil, nil, err
	}

	updated, err := tx.ValidateAndUpdate(user)
	if err != nil {
		return nil, nil, err
	}

	return updated, user, nil
}

// Destroy deletes a User from the DB. This function is mapped
// to the path DELETE /users/{user_id}
func (p *UsersRepository) Destroy(c buffalo.Context) (*models.User, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return nil, fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the Point the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return nil, c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(user); err != nil {
		return nil, err
	}

	return user, nil
}
