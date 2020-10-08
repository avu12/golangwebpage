package upload

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	filename := header.Filename
	if err != nil {
		log.Println(err)
	}
	out, err := os.Create("static/upload/" + filename)
	if err != nil {
		log.Println(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println(err)
	}
}
