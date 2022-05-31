package pets

import (
	"github.com/denonia/go-labs/pkg/auth"
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
	routes.GET("/", h.GetPets)
	routes.GET("/:id", h.GetPet)

	protected := routes.Group("/", auth.AuthorizationMiddleware)
	protected.POST("/", h.AddPet)
	protected.PUT("/:id", h.UpdatePet)
	protected.DELETE("/:id", h.DeletePet)
}