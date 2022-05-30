package main

import (
	"fmt"
	"log"
	"os"

	"github.com/denonia/go-labs/pkg/db"
	"github.com/denonia/go-labs/pkg/pets"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatalln("Error loading .env file")
	}

	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbUrl := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	
	db := db.Init(dbUrl)

	r := gin.Default()

	pets.RegisterRoutes(r, db)

	r.Run()
}