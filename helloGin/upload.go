package helloGin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func upload() {

	router.MaxMultipartMemory = 8 << 20 // 8m
	router.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		log.Println(file.Filename)

		dst := "./logs/" + file.Filename
		context.SaveUploadedFile(file, dst)

		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}
