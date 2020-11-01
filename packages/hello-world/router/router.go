package router

import (
	"bazel-go/packages/shared/handlers/health"

	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/health", health.Health("Hello World"))
	return r
}
