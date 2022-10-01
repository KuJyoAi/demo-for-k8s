package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	eng := gin.Default()
	eng.GET("/get", func(ctx *gin.Context) {
		fmt.Println("GET SUCCESS")
	})

	eng.POST("/post", func(ctx *gin.Context) {
		fmt.Println("POST SUCCESS")
	})

	eng.Run("44440")
}
