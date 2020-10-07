package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/avu12/golangwebpage/webpagego/internal/domain/nameday"
	"github.com/avu12/golangwebpage/webpagego/internal/restclient"
)

type namedayService struct{}

type namedayServiceInterface interface {
	GetNameday(request *http.Request) (*nameday.NamedayResponse, error)
}

var (
	NamedayService namedayServiceInterface
)

func init() {
	NamedayService = &namedayService{}
}

func (n *namedayService) GetNameday(request *http.Request) (*nameday.NamedayResponse, error) {
	response, err := restclient.Get("https://api.abalin.net/today")
	if err != nil {
		return nil, err
	}
	result := nameday.NamedayResponse{}
	bytes, err := ioutil.ReadAll(response.Body)
	json.Unmarshal(bytes, &result)
	return &result, err
}
