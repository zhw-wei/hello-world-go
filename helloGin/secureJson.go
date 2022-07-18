package helloGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func secureJson() {
	router.GET("/someJSON", func(context *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// 防止dns劫持，默认预置 while(1);到响应题
		// while(1);["lena","austin","foo"]
		context.SecureJSON(http.StatusOK, names)
	})
}
