package arunclient

import (
	"net/http"
	"time"
)

func Get(name string, url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url+"get/"+"?fname="+name, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: 300 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
