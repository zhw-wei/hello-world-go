package helloGin

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var router *gin.Engine

func CreateR() {
	router = gin.Default()
	router.Use(logger())

	router.LoadHTMLGlob("template/**/*")

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()
	// 日志颜色开启
	//gin.ForceConsoleColor()

	// 日志文件
	f, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 定义路由日志的格式
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	ping()
	someJson()
	index()
	tmpl()
	jsonp()
	login()
	formPost()
	pureJson()
	secureJson()
	xmlJson()
	upload()
	moreUpload()
	authorized()
	cookie()
	params()
	group()

	// 默认 http 配置
	//router.Run() //127.0.0.1:8080

	// 自定义 http 配置
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// Logger 自定义中间件
func logger() gin.HandlerFunc {

	return func(context *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		context.Set("example", "12345")

		// 请求前

		context.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的status
		status := context.Writer.Status()
		log.Println("logger status: ", status)
	}

}
