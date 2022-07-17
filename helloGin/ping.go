package helloGin

import (
	"github.com/gin-gonic/gin"
)

func ping() {
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})
}
