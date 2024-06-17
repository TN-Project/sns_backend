package engine

import (
	"sns_backend/pkg/engine/auth"

	"github.com/gin-gonic/gin"
)

func Engine(r *gin.Engine) *gin.Engine {
	r.Use(gin.Logger())
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", auth.SignupPost())
		authGroup.POST("/login", auth.LoginPost())
	}

	return r
}
