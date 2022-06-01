package users

import (
	"net/http"

	"github.com/denonia/go-labs/pkg"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(c *gin.Context) {
    var users []models.User

    if result := h.DB.Find(&users); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	var userDtos []GetUserDto

	for i := 0; i < len(users); i++ {
		var petDtos []GetPetDto

		for i := 0; i < len(users[i].Pets); i++ {
			petDtos = append(petDtos, GetPetDto{
				Name:      users[i].Pets[i].Name,
				Kind:      users[i].Pets[i].Kind,
				Breed:     users[i].Pets[i].Breed,
				BirthDate: pkg.JSONDate{users[i].Pets[i].BirthDate},
			})
		}

		userDtos = append(userDtos, GetUserDto{
			Name: users[i].Name,
			PhoneNumber: users[i].PhoneNumber,
			Pets: petDtos,
		})
	}

    c.JSON(http.StatusOK, &userDtos)
}