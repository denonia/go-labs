package pets

import (
	"net/http"

	"github.com/denonia/go-labs/pkg"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

type UpdatePetDto struct {
	Name      string       `json:"name"`
	OwnerID   uint         `json:"ownerId"`
	Kind      string       `json:"kind"`
	Breed     string       `json:"breed"`
	BirthDate pkg.JSONDate `json:"birthDate"`
}

func (h handler) UpdatePet(c *gin.Context) {
    id := c.Param("id")
    body := UpdatePetDto{}

    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var pet models.Pet

    if result := h.DB.First(&pet, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	pet.Name = body.Name
	pet.OwnerID = body.OwnerID
	pet.Kind = body.Kind
	pet.Breed = body.Breed
	pet.BirthDate = body.BirthDate.Time

    h.DB.Save(&pet)

    c.JSON(http.StatusOK, body)
}