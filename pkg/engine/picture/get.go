package picture

import (
	"path/filepath"
	"sns_backend/pkg/common/model"
	"sns_backend/pkg/db/read"
	"sns_backend/pkg/session"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPictureGet(c *gin.Context) {
	data := session.Default(c, "session", &model.Session{}).Get(c)
	if data == nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
			"error":   "session not found",
		})
		return
	}

	id := c.Param("id")
	strGroupID := c.Param("group_id")
	GroupID, err := strconv.Atoi(strGroupID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
			"error":   "invalid group_id",
		})
		return
	}

	group, err := read.GetGroupByID(GroupID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   "failed to get group",
		})
		return
	}

	// ユーザがグループに所属しているか確認
	users, err := read.GetGroupsUser(group.Group_name)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   "failed to get group users",
		})
		return
	}
	found := false
	for _, user := range users {
		if user.User_id == data.(*model.Session).Userid {
			found = true
			break
		}
	}
	if !found {
		c.JSON(403, gin.H{
			"message": "forbidden",
			"error":   "user is not in the group",
		})
		return
	}

	filepath := filepath.Join("upload", strGroupID, id)
	c.File(filepath)
}
