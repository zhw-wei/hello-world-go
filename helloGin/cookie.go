package helloGin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func cookie() {
	router.GET("/cookie", func(context *gin.Context) {
		cookie, err := context.Cookie("gin_cookie")

		if err != nil {
			cookie = "not cookie"
			context.SetCookie("gin_cookie", "test", 3600, "/", "127.0.0.1", false, true)
		}

		fmt.Printf("cookie value : %s \n\n", cookie)
	})
}
