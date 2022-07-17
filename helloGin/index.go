package helloGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index() {
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "hello world go web",
		})
	})
}
