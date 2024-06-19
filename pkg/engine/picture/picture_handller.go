package picture

import (
	"github.com/gin-gonic/gin"
)

func UploadPost() gin.HandlerFunc {
	return uploadPost
}

func GetPictureGet() gin.HandlerFunc {
	return getPictureGet
}
