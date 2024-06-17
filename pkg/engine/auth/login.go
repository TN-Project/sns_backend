package auth

import (
	"github.com/gin-gonic/gin"
)


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginPost(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	// データベースからユーザ情報を取得
	
	// パスワードを検証

	// セッション情報を設定

	// セッションを設定(cookieにセット)

	c.JSON(200, gin.H{
		"error": nil,
		"message": "login success",
	})
}