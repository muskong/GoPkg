package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func JsonPost(address, path string, result *map[string]any) error {
	return post("json", address, path, result)
}

func FormPost(address, path string, result *map[string]any) error {
	return post("form", address, path, result)
}

func post(methodType, address, path string, result *map[string]any) (err error) {
	u, err := url.ParseRequestURI(address)
	u.Path = path

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(a.Data.Encode()))
	if err != nil {
		return
	}

	contentType := "application/x-www-form-urlencoded"
	if strings.ToUpper(methodType) == "JSON" {
		contentType = "application/json"
	}
	req.Header.Add("Content-Type", contentType)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(rspBody, *result)

	return
}
