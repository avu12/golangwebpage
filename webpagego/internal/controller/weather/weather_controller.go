package weather

import (
	"log"
	"net/http"
	"os"

	"github.com/avu12/golangwebpage/login"
	"github.com/avu12/golangwebpage/webpagego/internal/domain/weather"
	"github.com/avu12/golangwebpage/webpagego/internal/services"
	"github.com/gin-gonic/gin"
)

func GetWeatherNow(c *gin.Context) {
	var request weather.WeatherRequest
	request.CityName = c.PostForm("cityname")
	request.PersonName = c.PostForm("personname")
	request.Token = os.Getenv("OPEN_WEATHER")

	result, err := services.WeatherService.GetWeather(request)
	if err != nil {
		log.Println("Error happened during openweather information acquiring!")
		c.HTML(http.StatusOK, "error.html", nil)
		return
	}
	uname, err := login.GetUsername(c)
	datas := map[string]interface{}{
		"city":        request.CityName,
		"personname":  request.PersonName,
		"result":      result.Temperature,
		"code":        result.Code,
		"citynumber":  result.Ratio.CityNumber,
		"number":      result.Ratio.Number,
		"description": result.Description,
		"uname":       uname,
	}
	if err != nil {
		datas["uname"] = nil
	}

	c.HTML(http.StatusOK, "result.html", datas)
}
