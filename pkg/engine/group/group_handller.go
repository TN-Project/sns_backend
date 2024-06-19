package group

import (
	"github.com/gin-gonic/gin"
)

func AddGroupPost() gin.HandlerFunc {
	return addGroupPost
}

func GroupPictureListGet() gin.HandlerFunc {
	return groupPictureListGet
}
