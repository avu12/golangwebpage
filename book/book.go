package book

import (
	"log"
	"net/http"

	"github.com/avu12/golangwebpage/database"
	"github.com/avu12/golangwebpage/types"
	"github.com/gin-gonic/gin"
)

func UploadBookHandler(c *gin.Context) {
	var B types.Book
	B.Title = c.PostForm("title")
	B.Author = c.PostForm("author")
	err := database.InsertBook(B.Title, B.Author)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	err, Bsclice := database.SelectAllBooks()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	log.Println(Bsclice)
	c.HTML(http.StatusOK, "showbooks.html", nil)
}
