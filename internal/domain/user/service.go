package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNewUser(c *gin.Context, db *gorm.DB) {

}

func Login(c *gin.Context, db *gorm.DB) {
	var loginRequest LoginRequest

	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
