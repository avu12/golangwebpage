package services

import (
	"log"
	"math"
	"time"

	"github.com/avu12/golangwebpage/database"
	"github.com/avu12/golangwebpage/webpagego/internal/domain/weather"
	"github.com/avu12/golangwebpage/webpagego/internal/providers/openweather"
)

type weatherService struct{}

type weatherServiceInterface interface {
	GetWeather(request weather.WeatherRequest) (*weather.WeatherResponse, error)
}

var (
	WeatherService weatherServiceInterface
)

func init() {
	WeatherService = &weatherService{}
}

func (w *weatherService) GetWeather(request weather.WeatherRequest) (*weather.WeatherResponse, error) {

	response, err := openweather.GetWeather(request)
	if err != nil {
		return nil, err
	}
	log.Println(response)

	result := weather.WeatherResponse{
		//we want result in  Celsius
		Temperature: math.Floor(response.Main.Temp - 273.05),
		Code:        response.Code,
		Description: response.W[0].Description,
	}
	if result.Code != 200 {
		return &result, nil
	}
	err = database.InsertTempDateCityNameQuery("WEATHERTABLE", int(result.Temperature), time.Now(), request.CityName, request.PersonName)
	if err != nil {
		return nil, err
	}
	citynumber, number, err := database.CityRateQuery(request.CityName)
	if err != nil {
		return nil, err
	}
	result.Ratio.CityNumber = citynumber
	result.Ratio.Number = number

	return &result, nil
}
