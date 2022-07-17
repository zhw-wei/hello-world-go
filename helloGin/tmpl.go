package helloGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func tmpl() {
	router.GET("/posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts",
		})
	})
	router.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users",
		})
	})
}
