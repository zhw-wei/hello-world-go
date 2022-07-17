package helloGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonp() {
	router.GET("/jsonp", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		context.JSONP(http.StatusOK, data)
	})
}
