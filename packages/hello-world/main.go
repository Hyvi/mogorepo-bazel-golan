package main

import (
	"bazel-go/packages/hello-world/ext"
	"bazel-go/packages/hello-world/router"
	"fmt"
)

func main() {
	fmt.Printf("hello, %d", ext.Counting(1, 2))
	// fmt.Println("111111")
	router.GetEngine().Run(":8080")
}
