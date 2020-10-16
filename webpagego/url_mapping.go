package webpagego

import (
	"net/http"

	"github.com/avu12/golangwebpage/book"
	"github.com/avu12/golangwebpage/login"
	"github.com/avu12/golangwebpage/mail"
	"github.com/avu12/golangwebpage/webpagego/internal/controller/weather"
	"github.com/gin-gonic/gin"
)

func mapUrls() {

	router.GET("/", login.HomepageHandler)

	router.POST("/weather", weather.GetWeatherNow)

	router.POST("/emaildidreg", mail.MailHandler)

	router.GET("/loadweatherpage", func(c *gin.Context) {
		uname, err := login.GetUsername(c)
		if err != nil {
			c.HTML(http.StatusOK, "weather.html", nil)
			return
		}
		c.HTML(http.StatusOK, "weather.html", uname)
	})
	router.GET("/emailregpage", func(c *gin.Context) {
		uname, err := login.GetUsername(c)
		if err != nil {
			c.HTML(http.StatusOK, "emailreg.html", nil)
			return
		}
		c.HTML(http.StatusOK, "emailreg.html", uname)
	})
	router.GET("/emailregistered/:emailhash", mail.ConfirmRegistration)

	router.GET("/recommendabook", book.BookRecommenderHandler)
	router.GET("/loginpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.POST("/uploadbook", book.UploadBookHandler)
	router.POST("/login", login.LoginHandler)
	router.GET("/logout", login.LogoutHandler)
}
