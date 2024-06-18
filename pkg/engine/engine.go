package engine

import (
	"sns_backend/pkg/engine/auth"
	"sns_backend/pkg/engine/group"
	"sns_backend/pkg/engine/picture"

	"github.com/gin-gonic/gin"
)

func Engine(r *gin.Engine) *gin.Engine {
	r.Use(gin.Logger())
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", auth.SignupPost())
		authGroup.POST("/login", auth.LoginPost())
	}
	groupGroup := r.Group("/group")
	{
		groupGroup.POST("/add", group.AddGroupPost())
	}
	pictureGroup := r.Group("/picture")
	{
		pictureGroup.POST("/upload", picture.UploadPost())
		pictureGroup.GET("/get", picture.GetPictureGet())
	}

	return r
}
