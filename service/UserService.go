package service

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"location_service_v1/ls_v2/models"
	"location_service_v1/ls_v2/repository"
)

// UsersService is a
type UsersService struct {
	usersRepository *repository.UsersRepository
}

// NewUsersService is a
func NewUsersService(repository *repository.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: repository,
	}
}

// List is a
func (s *UsersService) List(c buffalo.Context) (*models.Users, *pop.Query, error) {

	users, q, err := s.usersRepository.List(c)
	if err != nil {
		return nil, nil, err
	}

	return users, q, err
}

// Show gets the data for one Company. This function is mapped to
// the path GET /users/{user_id}
func (s *UsersService) Show(c buffalo.Context) (*models.User, error) {
	user, err := s.usersRepository.Show(c)
	if err != nil {
		return nil, err
	}
	return user, err
}

// New renders the form for creating a new Company.
// This function is mapped to the path GET /users/new
func (s *UsersService) New(c buffalo.Context) *models.User {
	return s.usersRepository.New(c)
}

// Create adds a Company to the DB. This function is mapped to the
// path POST /users
func (s *UsersService) Create(c buffalo.Context) (*validate.Errors, *models.User, error) {
	create, user, err := s.usersRepository.Create(c)
	if err != nil {
		return nil, nil, err
	}
	return create, user, nil
}

// Edit renders a edit form for a Company. This function is
// mapped to the path GET /users/{user_id}/edit
func (s *UsersService) Edit(c buffalo.Context) (*models.User, error) {
	user, err := s.usersRepository.Edit(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update changes a Company in the DB. This function is mapped to
// the path PUT /users/{user_id}
func (s *UsersService) Update(c buffalo.Context) (*validate.Errors, *models.User, error) {
	update, user, err := s.usersRepository.Update(c)
	if err != nil {
		return nil, nil, err
	}
	return update, user, nil
}

// Destroy deletes a Company from the DB. This function is mapped
// to the path DELETE /users/{user_id}
func (s *UsersService) Destroy(c buffalo.Context) (*models.User, error) {
	user, err := s.usersRepository.Destroy(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}
