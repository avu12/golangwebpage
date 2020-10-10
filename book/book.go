package book

import (
	"net/http"

	"github.com/avu12/golangwebpage/database"
	"github.com/gin-gonic/gin"
)

type Book struct {
	Title  string
	Author string
}

func UploadBookHandler(c *gin.Context) {
	var B Book
	B.Title = c.PostForm("title")
	B.Author = c.PostForm("author")
	err := database.InsertBook(B.Title, B.Author)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "showbooks.html", nil)
}
