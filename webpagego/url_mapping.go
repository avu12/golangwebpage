package webpagego

import (
	"net/http"

	"github.com/avu12/golangwebpage/mail"
	"github.com/avu12/golangwebpage/upload"
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

	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})
	router.POST("/uploadfile", upload.UploadHandler)
}
