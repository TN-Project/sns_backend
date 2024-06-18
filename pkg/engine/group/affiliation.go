package group

import (
	"sns_backend/pkg/common/model"
	"sns_backend/pkg/db/read"
	"sns_backend/pkg/session"

	"github.com/gin-gonic/gin"
)

// ユーザが所属しているグループ一覧を取得
func acquisitionAffiliationUserGet(c *gin.Context) {
	data := session.Default(c, "session", &model.Session{}).Get(c)
	if data == nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
			"error":   "session not found",
		})
		return
	}

	groups, err := read.GetUsersGroup(data.(*model.Session).Username)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   "failed to get user's groups",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"groups":  groups,
	})
}
