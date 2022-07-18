package helloGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pureJson() {
	router.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"html": "<b>hello, world</b>",
		})
	})

	router.GET("/pureJson", func(context *gin.Context) {
		context.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello, world</b>",
		})
	})
}
