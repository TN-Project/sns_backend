package auth

import (
	"github.com/gin-gonic/gin"
)

func SignupPost() gin.HandlerFunc {
	return signupPost
}

func LoginPost() gin.HandlerFunc {
	return loginPost
}