package helloGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func someJson() {
	router.GET("/someJson", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "golang",
			"tag":  "<br/>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})
}
