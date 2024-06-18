package engine

import (
	"sns_backend/pkg/engine/auth"
	"sns_backend/pkg/engine/group"

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

	return r
}
