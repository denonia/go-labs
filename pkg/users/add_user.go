package users

import (
	"net/http"

	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

type AddUserDto struct {
	Name 		string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

func (h handler) AddUser(c *gin.Context) {
	body := AddUserDto{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := models.User{
		Name: body.Name,
		PhoneNumber: body.PhoneNumber,
	}

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}
