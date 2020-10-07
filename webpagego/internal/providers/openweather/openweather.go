package openweather

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/avu12/golangwebpage/webpagego/internal/domain/openweather"
	"github.com/avu12/golangwebpage/webpagego/internal/domain/weather"
	"github.com/avu12/golangwebpage/webpagego/internal/restclient"
)

const (
	urlweather = "api.openweathermap.org/data/2.5/weather"
)

func GetWeather(request weather.WeatherRequest) (*openweather.WeatherResponse, error) {
	base, _ := url.Parse("https://api.openweathermap.org/data/2.5/weather")
	params := url.Values{}
	params.Add("q", request.CityName)
	params.Add("appid", request.Token)
	base.RawQuery = params.Encode()
	urlwithparams := base.String()

	response, err := restclient.Get(urlwithparams)
	if err != nil {
		return nil, err
	}

	var result openweather.WeatherResponse
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &result)
	if len(result.W) < 1 {
		emptydetail := openweather.Detail{
			Description: "Non defined description",
			Id:          -1,
		}
		result.W = append(result.W, emptydetail)
	}
	return &result, nil
}
