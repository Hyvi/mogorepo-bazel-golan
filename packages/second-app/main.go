package main

import (
	"fmt"

	"github.com/Hyvi/mogorepo-bazel-golang/packages/shared/handlers/health"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("I am second APP %s", "Server")

	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/health", health.Health("Second APP"))
	r.Run(":8080")

}
