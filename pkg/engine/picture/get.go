package picture

import (
	"github.com/gin-gonic/gin"
)

func getPictureGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get picture",
	})
}
