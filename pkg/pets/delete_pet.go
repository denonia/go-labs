package pets

import (
	"net/http"

	"github.com/denonia/go-labs/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeletePet(c *gin.Context) {
    id := c.Param("id")

    var pet models.Pet

    if result := h.DB.First(&pet, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    h.DB.Delete(&pet)

    c.Status(http.StatusOK)
}