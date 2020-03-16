package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
	"time"
)

// Point is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Point struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	PointId        int       `json:"point_id" db:"point_id"`
	Address        string    `json:"Address" db:"address"`
	CityName       string    `json:"CitiName" db:"citi_name"`
	OutDescription string    `json:"OutDescription" db:"out_description"`
	OwnerId        int       `json:"OwnerId" db:"owner_id"`
	OwnerName      string    `json:"OwnerName" db:"owner_name"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// PointDTO
type PointDTO struct {
	ID             int    `json:"Id"`
	Name           string `json:"Name"`
	Address        string `json:"Address"`
	CityName       string `json:"CitiName"`
	OutDescription string `json:"OutDescription"`
	OwnerId        int    `json:"OwnerId"`
	OwnerName      string `json:"OwnerName"`
}

// String is not required by pop and may be deleted
func (p Point) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Points
type Points []Point

// Points DTO
type PointsDTO []PointDTO

// String is not required by pop and may be deleted
func (p Points) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Point) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Point) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Point) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
