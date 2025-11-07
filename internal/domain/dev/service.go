package dev

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func ValidUUID(c *gin.Context) {
	var req ValidUUIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "❌ Invalid JSON structure",
		})
		return
	}

	_, err := uuid.Parse(req.UUID)
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"uuid":    req.UUID,
			"valid":   false,
			"message": "❌ Invalid UUID format",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid":    req.UUID,
		"valid":   true,
		"message": "✅ UUID is valid",
	})
}

type ValidUUIDRequest struct {
	UUID string `json:"uuid"`
}
