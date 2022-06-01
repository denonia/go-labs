package pets

import (
	"net/http"

	"github.com/denonia/go-labs/pkg"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

type GetPetDto struct {
	Name 	  string 	   `json:"name"`
	Owner     GetUserDto   `json:"ownerId"`
	Kind      string       `json:"kind"`
	Breed     string       `json:"breed"`
	BirthDate pkg.JSONDate `json:"birthDate"`
}

type GetUserDto struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func (h handler) GetPet(c *gin.Context) {
    id := c.Param("id")

    var pet models.Pet

    if result := h.DB.First(&pet, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	var owner models.User

	if result := h.DB.First(&owner, pet.OwnerID); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
	}

	userDto := GetUserDto{
		Name: owner.Name,
		PhoneNumber: owner.PhoneNumber,
	}

	dto := GetPetDto{
		Name: pet.Name,
		Owner: userDto,
		Kind: pet.Kind,
		Breed: pet.Breed,
		BirthDate: pkg.JSONDate{pet.BirthDate},
	}

    c.JSON(http.StatusOK, &dto)
}