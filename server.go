package main

import (
	"blog/apis/version1_0"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/blog", "./resource/htmls")

	v1 := router.Group("/v1.0")
	{
		v1.POST("/blog/user/login", version1_0.Login)
	}
	router.Run(":8887")
}
