package we

type (
	codeToOpenid struct {
		SessionKey string `json:"session_key"` // 	会话密钥
		Unionid    string `json:"unionid"`     // 	用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见 UnionID 机制说明。
		Errmsg     string `json:"errmsg"`      // 	错误信息
		Openid     string `json:"openid"`      // 	用户唯一标识
		Errcode    int32  `json:"errcode"`     // 	错误码
	}
	accountToken struct {
		AccessToken string `json:"access_token"` // 	获取到的凭证
		ExpiresIn   int32  `json:"expires_in"`   // 	凭证有效时间，单位：秒。目前是7200秒之内的值。
		Errcode     int32  `json:"errcode"`      // 	错误码
		Errmsg      string `json:"errmsg"`       // 	错误信息
	}
)
