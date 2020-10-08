package upload

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("uploadedfile")
	log.Println(file, header, err)
	filename := header.Filename
	if err != nil {
		log.Println("Error in form", err)
	}
	newpath := filepath.Join(".", "uploadtest/")
	os.MkdirAll(newpath, os.ModePerm)
	out, err := os.Create(newpath + filename)
	log.Println(out)
	if err != nil {
		log.Println("Error in filecreation", err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	log.Println(out)
	if err != nil {
		log.Println("Error in copy data", err)
	}
	c.HTML(http.StatusOK, "uploadsuccesful.html", nil)
}
