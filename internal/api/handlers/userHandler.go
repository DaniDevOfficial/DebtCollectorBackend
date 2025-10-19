package handlers

import (
	"dept-collector/internal/domain/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	RegisterAuthRoutes(router, db)
}

func RegisterAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/login", func(c *gin.Context) {
		user.Login(c, db)
	})

	router.POST("/signup", func(c *gin.Context) {
		user.SignUp(c, db)
	})

	router.GET("/checkAuth", func(c *gin.Context) {
		user.CheckAuth(c, db)
	})
}
