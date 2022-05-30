package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const dateFormat = "2006-01-02"

type JSONDate struct {
	time.Time
}

// interface Marshaler
func (d *JSONDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(d.Time).Format(dateFormat))
	return []byte(stamp), nil
}

// interface Unmarshaler
func (d *JSONDate) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
    if s == "null" {
       d.Time = time.Time{}
       return
    }
    d.Time, err = time.Parse(dateFormat, s)
    return
}

func main() {
	r := gin.New()

	r.GET("/pets", listPetsHandler)
	r.POST("/pets", createPetHandler)
	r.DELETE("/pets/:id", deletePetHandler)

	r.Run()
}

type Pet struct {
	ID   string    `json:"id"`
	Name string `json:"name"`

	// to be replaced with id
	OwnerName string   `json:"ownerName"`
	Kind      string   `json:"kind"`
	Breed     string   `json:"breed"`
	BirthDate JSONDate `json:"birthDate"`
}

var pets = []Pet{
	{
		ID:        "1",
		Name:      "Barsique",
		OwnerName: "Ivan",
		Kind:      "Cat",
		Breed:     "Domestic Shorthair",
		BirthDate: JSONDate{time.Date(2022, 2, 24, 0, 0, 0, 0, time.UTC)},
	},
	{
		ID:        "2",
		Name:      "Manya",
		OwnerName: "Vasya",
		Kind:      "Cat",
		Breed:     "Maine Coon",
		BirthDate: JSONDate{time.Date(2018, 3, 12, 0, 0, 0, 0, time.UTC)},
	},
	{
		ID:        "3",
		Name:      "Bob",
		OwnerName: "Alex",
		Kind:      "Dog",
		Breed:     "Labrador Retriever",
		BirthDate: JSONDate{time.Date(2016, 12, 10, 0, 0, 0, 0, time.UTC)},
	},
}

func listPetsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, pets)
}

func createPetHandler(c *gin.Context) {
	var pet Pet

	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pets = append(pets, pet)

	c.JSON(http.StatusCreated, pet)
}

func deletePetHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range pets {
		if a.ID == id {
			pets = append(pets[:i], pets[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
