package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Point is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Point struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	PointID        int       `json:"point_id" db:"point_id"`
	Address        string    `json:"address" db:"address"`
	CityName       string    `json:"citiName" db:"citi_name"`
	OutDescription string    `json:"outDescription" db:"out_description"`
	OwnerID        int       `json:"ownerId" db:"owner_id"`
	OwnerName      string    `json:"ownerName" db:"owner_name"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	CompanyID      uuid.UUID `json:"company_id" db:"company_id"`
	Company        *Company  `json:"company,omitempty" belongs_to:"company"`
}

// PointDTO is a
type PointDTO struct {
	ID             int    `json:"Id"`
	Name           string `json:"Name"`
	Address        string `json:"Address"`
	CityName       string `json:"CitiName"`
	OutDescription string `json:"OutDescription"`
	OwnerID        int    `json:"OwnerId"`
	OwnerName      string `json:"OwnerName"`
	Company        string `json:"Company"`
}

// String is not required by pop and may be deleted
func (p Point) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Points is a
type Points []Point

// PointsDTO DTO is a
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
