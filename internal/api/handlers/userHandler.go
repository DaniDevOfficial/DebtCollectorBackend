package handlers

import (
	"dept-collector/internal/domain/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	RegisterUserRoutes(router, db)
}

func RegisterAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/login", func(c *gin.Context) {
		user.Login(c, db)
	})
}
