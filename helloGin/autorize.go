package helloGin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523433"},
}

func authorized() {

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// 地址：/admin/secrets
	// 访问地址后会有一个弹窗，要求输入帐号和密码，输入内容是 authorized 中设置的账号密码
	authorized.GET("/secrets", func(context *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := context.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": "no secret :("})
		}
	})

}
