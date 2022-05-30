package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name string `json:"name"`

	// to be replaced with id
	OwnerName string    `json:"ownerName"`
	Kind      string    `json:"kind"`
	Breed     string    `json:"breed"`
	BirthDate time.Time `json:"birthDate"`
}