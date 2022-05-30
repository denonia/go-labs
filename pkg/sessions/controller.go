package sessions

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

	routes := r.Group("/sessions")
	routes.POST("/", h.AddSession)
	routes.GET("/", h.GetSessions)
	// routes.GET("/:id", h.GetSession)
	// routes.PUT("/:id", h.UpdatePet)
	// routes.DELETE("/:id", h.DeletePet)
}