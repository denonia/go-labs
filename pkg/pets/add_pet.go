package pets

import (
	"net/http"

	"github.com/denonia/go-labs/pkg"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

type AddPetDto struct {
	Name      string       `json:"name"`
	OwnerID   uint         `json:"ownerId"`
	Kind      string       `json:"kind"`
	Breed     string       `json:"breed"`
	BirthDate pkg.JSONDate `json:"birthDate"`
}

func (h handler) AddPet(c *gin.Context) {
	body := AddPetDto{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pet := models.Pet{
		Name:      body.Name,
		OwnerID:   body.OwnerID,
		Kind:      body.Kind,
		Breed:     body.Breed,
		BirthDate: body.BirthDate.Time,
	}

	if result := h.DB.Create(&pet); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &pet)
}
