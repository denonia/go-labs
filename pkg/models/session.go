package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	PetID    uint `gorm:"not_null"`
	SitterID uint `gorm:"not_null"`

	StartDate time.Time `gorm:"not_null"`
	EndDate	  time.Time `gorm:"not_null"`
}