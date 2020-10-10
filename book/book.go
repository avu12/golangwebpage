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
	err = ShowAllBooks()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "showbooks.html", nil)
}

func ShowAllBooks() error {
	err, Bsclice := database.SelectAllBooks()
	if err != nil {
		return err
	}
	log.Println("no problem until here")
	log.Println(Bsclice)
	return nil
}
