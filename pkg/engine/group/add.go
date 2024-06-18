package group

import (
	"github.com/gin-gonic/gin"
)

func addGroupPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "add group",
		"error":   nil,
	})
}
