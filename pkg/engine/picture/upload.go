package picture

import (
	"github.com/gin-gonic/gin"
)

func uploadPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "upload",
	})
}
