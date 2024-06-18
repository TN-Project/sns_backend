package engine

import (
	"sns_backend/pkg/engine/auth"
	"sns_backend/pkg/engine/create_group"
	"github.com/gin-gonic/gin"
)

func Engine(r *gin.Engine) *gin.Engine {
	r.Use(gin.Logger())
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", auth.SignupPost())
		authGroup.POST("/login", auth.LoginPost())
	}
	create_groupGroup := r.Group("/create_group")
	{
		create_groupGroup.POST("/exist_user",create_group.ExistUser())
	}
	return r
}
