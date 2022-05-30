package pets

import (
	"net/http"

	"github.com/denonia/go-labs/pkg"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetPet(c *gin.Context) {
    id := c.Param("id")

    var pet models.Pet

    if result := h.DB.First(&pet, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	dto := GetPetDto{
		Name: pet.Name,
		OwnerID: pet.OwnerID,
		Kind: pet.Kind,
		Breed: pet.Breed,
		BirthDate: pkg.JSONDate{pet.BirthDate},
	}

    c.JSON(http.StatusOK, &dto)
}