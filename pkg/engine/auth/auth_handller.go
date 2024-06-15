package auth

import (
	"github.com/gin-gonic/gin"
)

func SignupPost() gin.HandlerFunc {
	return signupPost
}
