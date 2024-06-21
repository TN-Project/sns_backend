package auth

import (
	"sns_backend/pkg/common/encrypt"
	"sns_backend/pkg/common/model"
	"sns_backend/pkg/db/read"
	"sns_backend/pkg/session"

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
	user, err := read.GetUser(req.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	// パスワードを検証
	if err := encrypt.CompareHashAndPassword(user.Password, req.Password); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid password",
			"error":   err.Error(),
		})
		return
	}

	// セッション情報を設定
	session_data := model.Session{
		Userid:   user.User_id,
		Nickname: user.Nickname,
		Username: user.Username,
	}

	// セッションを設定(cookieにセット)
	session.Default(c, "session", &model.Session{}).Set(c, session_data)
	c.SetCookie("login", "true", 0, "/", "localhost", false, false)
	c.JSON(200, gin.H{
		"error":   nil,
		"message": "login success",
	})
}
