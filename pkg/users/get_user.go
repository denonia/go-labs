package users

import (
	"net/http"

	"github.com/denonia/go-labs/pkg"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

type GetUserDto struct {
	Name        string      `json:"name"`
	PhoneNumber string      `json:"phone_number"`
	Pets        []GetPetDto `json:"pets"`
}

type GetPetDto struct {
	Name      string       `json:"name"`
	Kind      string       `json:"kind"`
	Breed     string       `json:"breed"`
	BirthDate pkg.JSONDate `json:"birthDate"`
}

func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var petRecords []models.Pet
	h.DB.Where("owner_id = ?", user.ID).Find(&petRecords)

	var petDtos []GetPetDto
	for i := 0; i < len(petRecords); i++ {
		petDtos = append(petDtos, GetPetDto{
			Name:      petRecords[i].Name,
			Kind:      petRecords[i].Kind,
			Breed:     petRecords[i].Breed,
			BirthDate: pkg.JSONDate{petRecords[i].BirthDate},
		})
	}

	dto := GetUserDto{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Pets:        petDtos,
	}

	c.JSON(http.StatusOK, &dto)
}
