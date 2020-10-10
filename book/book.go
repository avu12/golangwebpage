package book

import (
	"log"
	"net/http"
	"text/template"

	"github.com/avu12/golangwebpage/database"
	"github.com/avu12/golangwebpage/types"
	"github.com/gin-gonic/gin"
)

func UploadBookHandler(c *gin.Context) {
	var B types.Book
	var Bsclice []types.Book
	B.Title = c.PostForm("title")
	B.Author = c.PostForm("author")
	err := database.InsertBook(B.Title, B.Author)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	Bsclice, err = ShowAllBooks()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	t, err := template.ParseFiles("./templates/bookstemplate.html")
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	_ = Bsclice
	c.HTML(http.StatusOK, t.Name(), 42)
	//c.Writer.Header().Set("Content-Type", "text/html")
	/*err = t.Execute(c.Writer, Bsclice)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}*/

	//c.HTML(http.StatusOK, "showbooks.html", nil)
}

func ShowAllBooks() ([]types.Book, error) {
	err, Bsclice := database.SelectAllBooks()
	if err != nil {
		return nil, err
	}
	return Bsclice, nil
}
