package main

import (
	"fmt"
	"log"
	"os"

	"github.com/denonia/go-labs/pkg/auth"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/denonia/go-labs/pkg/pets"
	"github.com/denonia/go-labs/pkg/sessions"
	"github.com/denonia/go-labs/pkg/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatalln("Error loading .env file")
	}

	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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