package main

import (
	"sns_backend/pkg/engine"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r = engine.Engine(r)

	r.Run()
}
