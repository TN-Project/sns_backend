package group

import (
	"sns_backend/pkg/db/read"
	"strconv"

	"github.com/gin-gonic/gin"
)

func groupPictureListGet(c *gin.Context) {
	group_id, err := strconv.Atoi(c.Param("group_id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}
	pictures, err := read.GetPicture(group_id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "success get group pictures",
		"pictures": pictures,
	})
}
