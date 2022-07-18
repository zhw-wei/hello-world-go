package helloGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由组
func group() {
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndPoint)
		v1.GET("/login2", loginEndPoint)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/login", loginEndPoint)
		v2.GET("/login2", loginEndPoint)
	}
}

func loginEndPoint(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
