package sessions

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

	routes := r.Group("/sessions")
	routes.GET("/", h.GetSessions)
	routes.POST("/", h.AddSession, auth.AuthorizationMiddleware)
}