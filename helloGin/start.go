package helloGin

import "github.com/gin-gonic/gin"

var router *gin.Engine

func CreateR() {
	router = gin.Default()

	router.LoadHTMLGlob("template/**/*")

	ping()
	someJson()
	index()
	tmpl()
	jsonp()
	login()
	formPost()

	router.Run() //127.0.0.1:8080
}
