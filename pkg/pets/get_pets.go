package pets

import (
	"net/http"

	"github.com/denonia/go-labs/pkg"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

type GetPetDto struct {
	Name 	  string 	`json:"name"`
	OwnerName string 	`json:"ownerName"`
	Kind      string    `json:"kind"`
	Breed     string    `json:"breed"`
	BirthDate pkg.JSONDate `json:"birthDate"`
}

func (h handler) GetPets(c *gin.Context) {
    var pets []models.Pet

    if result := h.DB.Find(&pets); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	var petDtos []GetPetDto
	for i := 0; i < len(pets); i++ {
		petDtos = append(petDtos, GetPetDto{
			Name: pets[i].Name,
			OwnerName: pets[i].OwnerName,
			Kind: pets[i].Kind,
			Breed: pets[i].Breed,
			BirthDate: pkg.JSONDate{pets[i].BirthDate},
		})
	}

    c.JSON(http.StatusOK, &petDtos)
}