package plugin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type (
	PluginApi struct {
		ContentType string // json|form
		Data        url.Values
		Url         string // 接口域名或IP
		ApiID       string // 商户号
		ApiKey      string // 商户密钥
	}
	PluginInfo struct {
		Key   string
		Title string
		Value string
	}
)

func (p *PluginApi) post(path string, result *map[string]any) (err error) {
	u, err := url.ParseRequestURI(p.Url)
	u.Path = path

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(p.Data.Encode()))
	if err != nil {
		return
	}

	contentType := "application/x-www-form-urlencoded"
	if strings.ToUpper(p.ContentType) == "JSON" {
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
