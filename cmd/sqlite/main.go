package main

import (
	"log"

	"github.com/denonia/go-labs/pkg/auth"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/denonia/go-labs/pkg/pets"
	"github.com/denonia/go-labs/pkg/sessions"
	"github.com/denonia/go-labs/pkg/users"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatalln("Error loading .env file")
	}
	
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.User{}, &models.Session{}, &models.Pet{})

	r := gin.Default()

	auth.RegisterRoutes(r, db)
	pets.RegisterRoutes(r, db)
	sessions.RegisterRoutes(r, db)
	users.RegisterRoutes(r, db)

	r.Run()
}