package auth

import (
	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Userid   int    `json:"userid"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func signupPost(c *gin.Context) {
	var req SignupRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	// データベースに登録

	c.JSON(200, gin.H{
		"message": "signup",
	})
}
