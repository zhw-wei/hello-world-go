package helloGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func params() {
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello %s", name)
	})
}
