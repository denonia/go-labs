package sessions

import (
	"errors"
	"net/http"
	"time"

	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

type AddSessionDto struct {
	PetID     uint		`json:"petId"`
	SitterID  uint		`json:"sitterId"`
	StartDate time.Time	`json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

func (h handler) AddSession(c *gin.Context) {
	body := AddSessionDto{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if body.StartDate.After(body.EndDate) {
		c.AbortWithError(http.StatusBadRequest, errors.New("end date can't be before start date"))
		return
	}

	session := models.Session{
		PetID:     body.PetID,
		SitterID:  body.SitterID,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
	}

	if result := h.DB.Create(&session); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &session)
}
