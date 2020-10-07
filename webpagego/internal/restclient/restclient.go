package restclient

import (
	"net/http"
)

func Get(urlwithparams string) (*http.Response, error) {

	request, err := http.NewRequest(http.MethodGet, urlwithparams, nil)
	if err != nil {
		return nil, err
	}
	request.Header = http.Header{}

	client := http.Client{}
	return client.Do(request)
}
