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
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
	}))

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", auth.SignupPost())
		authGroup.POST("/login", auth.LoginPost())
		authGroup.GET("/logout",auth.LogoutGet())
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
