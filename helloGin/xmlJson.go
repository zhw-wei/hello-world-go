package helloGin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func xmlJson() {
	// gin.H 是 map[string]interface{} 的一种快捷模式
	router.GET("/json/someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	router.GET("/json/moreJSON", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		//注意 msg.Name 在json中变成了 user
		context.JSON(http.StatusOK, msg)
	})

	router.GET("/xml/someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	router.GET("/yaml/someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	router.GET("/proto/someProtoBuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"

		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 数据在响应中变为二进制数据
		context.ProtoBuf(http.StatusOK, data)
	})
}
