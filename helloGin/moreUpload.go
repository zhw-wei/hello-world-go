package helloGin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func moreUpload() {

	router.MaxMultipartMemory = 8 << 20

	router.POST("/moreUpload", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			dst := "./logs/" + file.Filename
			context.SaveUploadedFile(file, dst)
		}

		context.String(http.StatusOK, fmt.Sprintf("%d files upload!", len(files)))
	})
}
