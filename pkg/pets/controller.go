package pets

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler {
		DB: db,
	}

	routes := r.Group("/pets")
	routes.POST("/", h.AddPet)
	routes.GET("/", h.GetPets)
	routes.GET("/:id", h.GetPet)
	routes.PUT("/:id", h.UpdatePet)
	routes.DELETE("/:id", h.DeletePet)
}