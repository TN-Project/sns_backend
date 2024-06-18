package group

import (
	"sns_backend/pkg/db/create"
	"sns_backend/pkg/db/read"

	"github.com/gin-gonic/gin"
)

type AddGroupRequest struct {
	GroupName string   `json:"groupname"`
	Usernames []string `json:"usernames"`
}

func addGroupPost(c *gin.Context) {
	var req AddGroupRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	// ユーザ情報(id)を取得
	userids := []int{}
	for _, username := range req.Usernames {
		userdata, err := read.GetUser(username)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "internal server error",
				"error":   err.Error(),
			})
			return
		}
		userids = append(userids, userdata.User_id)
	}

	// グループを作成
	group_id, err := create.CreateGroup(req.GroupName)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	// グループにユーザを追加
	if err := create.AddUserToGroup(userids, group_id); err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success add group",
		"error":   nil,
	})
}
