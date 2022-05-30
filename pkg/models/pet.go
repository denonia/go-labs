package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Sessions  []Session  `gorm:"foreignKey:PetID"`
	OwnerID   uint       `gorm:"not_null"`

	Name 	  string 	`gorm:"not_null"`
	Kind      string    `gorm:"not_null"`
	Breed     string    `gorm:"not_null"`
	BirthDate time.Time `gorm:"not_null"`
}