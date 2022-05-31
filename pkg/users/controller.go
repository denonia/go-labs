package users

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

	routes := r.Group("/users")
	routes.GET("/:id", h.GetUser)
	routes.POST("/", h.AddUser, auth.AuthorizationMiddleware)
}