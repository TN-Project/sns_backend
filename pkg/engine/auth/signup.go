package auth

import (
	"sns_backend/pkg/common/encrypt"
	"sns_backend/pkg/common/model"
	"sns_backend/pkg/common/random"
	"sns_backend/pkg/db/create"
	"sns_backend/pkg/session"

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
	// ランダムなユーザIDを生成
	userid := random.GenerateRandomInt()
	// パスワードをハッシュ化
	hashed_password, err := encrypt.PasswordEncrypt(req.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	// ユーザ情報を登録
	if err := create.CreateUser(userid, req.Nickname, req.Username, hashed_password); err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	// セッション情報を設定
	session_data := model.Session{
		Userid:   userid,
		Nickname: req.Nickname,
		Username: req.Username,
	}

	// セッションを設定
	session.Default(c, "session", &model.Session{}).Set(c, session_data)

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "complete signup",
	})
}
