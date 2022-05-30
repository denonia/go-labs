package sessions

import (
	"net/http"
	"time"

	"github.com/denonia/go-labs/pkg"
	"github.com/denonia/go-labs/pkg/models"
	"github.com/denonia/go-labs/pkg/pets"
	"github.com/gin-gonic/gin"
)

type GetSessionDto struct {
	Sitter    GetSitterDto	 `json:"sitter"`
	Pet 	  pets.GetPetDto `json:"pet"`
	StartDate time.Time 	 `json:"startDate"`
	EndDate	  time.Time 	 `json:"endDate"`
}

type GetSitterDto struct {
	Name 		string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func (h handler) GetSessions(c *gin.Context) {
    var sessions []models.Session

    if result := h.DB.Find(&sessions); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	var sessionDtos []GetSessionDto
	for i := 0; i < len(sessions); i++ {
		var pet models.Pet

		if result := h.DB.First(&pet, sessions[i].PetID); result.Error != nil {
			continue
		}

		var sitter models.User

		if result := h.DB.First(&sitter, sessions[i].SitterID); result.Error != nil {
			continue
		}

		petDto := pets.GetPetDto{
			Name: pet.Name,
			OwnerID: pet.OwnerID,
			Kind: pet.Kind,
			Breed: pet.Breed,
			BirthDate: pkg.JSONDate{pet.BirthDate},
		}

		sitterDto := GetSitterDto{
			Name: sitter.Name,
			PhoneNumber: sitter.PhoneNumber,
		}

		sessionDtos = append(sessionDtos, GetSessionDto{
			Pet: petDto,
			Sitter: sitterDto,
			StartDate: sessions[i].StartDate,
			EndDate: sessions[i].EndDate,
		})
	}

    c.JSON(http.StatusOK, &sessionDtos)
}