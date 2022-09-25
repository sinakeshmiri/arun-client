package arunclient

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Add(src string, name string, url string) error {
	encoded := base64.StdEncoding.EncodeToString([]byte(src))
	jsonBody := []byte(fmt.Sprintf(`{"name": "%s" , "src": "%s"}`, name, encoded))
	bodyReader := bytes.NewReader(jsonBody)
	url = url + "add/"
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)

	if err != nil {
		return err
	}
	client := http.Client{
		Timeout: 300 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		var a []byte
		req.Response.Body.Read(a)
		return errors.New(string(a))
	}
	return nil

}
