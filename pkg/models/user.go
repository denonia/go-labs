package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Pets 	 []Pet 	   `gorm:"foreignKey:OwnerID"`
	Sessions []Session `gorm:"foreignKey:SitterID"`

	Name 		string `gorm:"not_null"`
	PhoneNumber string
}