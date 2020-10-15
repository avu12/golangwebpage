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
	var Bslice []types.Book
	B.Title = c.PostForm("title")
	B.Author = c.PostForm("author")
	err := database.InsertBook(B.Title, B.Author)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	Bslice, err = ShowAllBooks()
	log.Println(Bslice)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "bookstemplate.html", Bslice)

}

func ShowAllBooks() ([]types.Book, error) {
	err, Bsclice := database.SelectAllBooks()
	if err != nil {
		return nil, err
	}
	return Bsclice, nil
}
func BookRecommenderHandler(c *gin.Context) {
	uname, err := c.Cookie("username")
	if err != nil {
		c.HTML(http.StatusOK, "books.html", nil)
		return
	}
	c.HTML(http.StatusOK, "books.html", uname)
}
