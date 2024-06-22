package engine

import (
	"sns_backend/pkg/engine/auth"
	"sns_backend/pkg/engine/group"
	"sns_backend/pkg/engine/picture"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Engine(r *gin.Engine) *gin.Engine {
	r.Use(gin.Logger())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://y-f.natyosu.com", "http://y-f.natyosu.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Set-Cookie"},
		AllowCredentials: true,
	}))

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", auth.SignupPost())
		authGroup.POST("/login", auth.LoginPost())
	}
	groupGroup := r.Group("/group")
	{
		groupGroup.POST("/add", group.AddGroupPost())
		groupGroup.GET("/acquisition-affiliation-user", group.AcquisitionAffiliationUserGet())
		groupGroup.GET("/:group_id/pictures-list", group.GroupPictureListGet())
	}
	pictureGroup := r.Group("/picture")
	{
		pictureGroup.POST("/upload", picture.UploadPost())
		pictureGroup.GET("/get/:group_id/:id", picture.GetPictureGet())
	}

	return r
}
