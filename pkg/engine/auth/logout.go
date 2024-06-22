package auth

import (
	"sns_backend/pkg/common/encrypt"
	"sns_backend/pkg/common/model"
	"sns_backend/pkg/db/read"
	"sns_backend/pkg/session"

	"github.com/gin-gonic/gin"
)


func logoutGet(c *gin.Context) {
	 session.Default(c, "session", &model.Session{}).Delete(c)
	 c.SetCookie("login",false,-1,"/","",false,false)
	 
}
