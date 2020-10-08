package upload

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("uploadedfile")
	log.Println(file, header, err)
	filename := header.Filename
	if err != nil {
		log.Println("Error in form", err)
	}
	out, err := os.Create("uploadfiles/" + filename)
	if err != nil {
		log.Println("Error in filecreation", err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("Error in copy data", err)
	}
	c.HTML(http.StatusOK, "uploadsuccesful.html", nil)
}
