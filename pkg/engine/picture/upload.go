package picture

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sns_backend/pkg/common/model"
	"sns_backend/pkg/db/read"
	"sns_backend/pkg/session"

	"github.com/gin-gonic/gin"
)

func uploadPost(c *gin.Context) {
	data := session.Default(c, "session", &model.Session{}).Get(c)
	if data == nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
			"error":   "session not found",
		})
		return
	}

	// ユーザがグループに所属しているか確認
	group_name := c.PostForm("group_name")
	users, err := read.GetGroupsUser(group_name)
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

	// グループIDを取得
	group, err := read.GetGroup(group_name)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
			"error":   "failed to get group",
		})
		return
	}

	// グループ用のディレクトリを作成
	os.MkdirAll(filepath.Join("upload", fmt.Sprintf("%d", group.Group_id)), 0755)

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "file not found",
			"error":   "file not found",
		})
		return
	}
	defer file.Close()

	savepath := filepath.Join("upload", fmt.Sprintf("%d", group.Group_id), header.Filename)
	slog.Info("savepath: " + savepath)
	err = c.SaveUploadedFile(header, savepath)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "upload failed",
			"error":   "failed to save file",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "upload",
		"error":   nil,
	})
}
