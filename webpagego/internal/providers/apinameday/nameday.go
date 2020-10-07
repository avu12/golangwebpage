package apinameday

import (
	"encoding/json"
	"io/ioutil"

	"github.com/avu12/golangwebpage/webpagego/internal/domain/nameday"
	"github.com/avu12/golangwebpage/webpagego/internal/restclient"
)

func GetNameDay() (*nameday.NamedayResponse, error) {
	response, err := restclient.Get("https://api.abalin.net/today")
	if err != nil {
		return nil, err
	}
	var result nameday.NamedayResponse
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &result)
	return &result, nil
}
