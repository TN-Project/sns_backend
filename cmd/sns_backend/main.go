package main

import (
	"sns_backend/pkg/db/create"
	"sns_backend/pkg/engine"

	"github.com/gin-gonic/gin"
)

func init() {
	create.CreateDefaultTable()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r = engine.Engine(r)

	r.Run()
}
