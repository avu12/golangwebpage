package webpagego

import (
	"net/http"

	"github.com/avu12/golangwebpage/book"
	"github.com/avu12/golangwebpage/mail"
	"github.com/avu12/golangwebpage/webpagego/internal/controller/weather"
	"github.com/gin-gonic/gin"
)

func mapUrls() {

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/weather", weather.GetWeatherNow)

	router.POST("/emaildidreg", mail.MailHandler)

	router.GET("/loadweatherpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "weather.html", nil)
	})
	router.GET("/emailregpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "emailreg.html", nil)
	})
	router.GET("/emailregistered/:emailhash", mail.ConfirmRegistration)

	router.GET("/recommendabook", func(c *gin.Context) {
		c.HTML(http.StatusOK, "books.html", nil)
	})
	router.GET("/loginpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.POST("/uploadbook", book.UploadBookHandler)
	router.POST("/login", login.LoginHandler)
}
