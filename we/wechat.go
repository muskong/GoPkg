package we

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type wechat struct {
	AppID     string
	AppSecret string
}

var We = &wechat{}

func Init(id, secret string) {
	We.AppID = id
	We.AppSecret = secret
}

func (w *wechat) CodeToOpenid(code string) (openid string, err error) {
	val := url.Values{}
	val.Add("appid", w.AppID)
	val.Add("secret", w.AppSecret)
	val.Add("js_code", code)
	val.Add("grant_type", "authorization_code")

	rsp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?" + val.Encode())

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	var data codeToOpenid
	err = json.NewDecoder(rsp.Body).Decode(&data)
	if err != nil {
		return
	}

	switch data.Errcode {
	case 40029:
		err = errors.New("code 无效")

	case 45011:
		err = errors.New("API 调用太频繁")

	case 40226:
		err = errors.New("高风险等级用户，小程序登录拦截")

	case -1:
		err = errors.New("系统繁忙，此时请开发者稍候再试")
	default:
		openid = data.Openid
	}
	return
}

func (w *wechat) GetAccountToken(code string) (token string, expires int32, err error) {
	val := url.Values{}
	val.Add("appid", w.AppID)
	val.Add("secret", w.AppSecret)
	val.Add("grant_type", "client_credential")

	rsp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?" + val.Encode())

	if err != nil {
		return
	}
	defer rsp.Body.Close()

	var data accountToken
	err = json.NewDecoder(rsp.Body).Decode(&data)
	if err != nil {
		return
	}

	switch data.Errcode {
	case 40001:
		err = errors.New("AppSecret 错误或者 AppSecret 不属于这个小程序")

	case 40002:
		err = errors.New("请确保 grant_type 字段值为 client_credential")

	case 40013:
		err = errors.New("不合法的 AppID，请开发者检查 AppID 的正确性，避免异常字符，注意大小写")

	case -1:
		err = errors.New("系统繁忙，此时请开发者稍候再试")

	default:
		token = data.AccessToken
		expires = data.ExpiresIn
	}
	return

}
