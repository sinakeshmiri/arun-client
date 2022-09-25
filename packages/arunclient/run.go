package arunclient

import (
	"net/http"
	"time"
)

func Run(name string, url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header["X-Func"] = []string{name}
	client := http.Client{
		Timeout: 300 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	return res,nil
}
