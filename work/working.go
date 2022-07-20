package work

import (
	"encoding/json"
	"encoding/xml"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/muskong/GoService/pkg/jwt"
	"github.com/muskong/GoService/pkg/utils"
	"github.com/muskong/GoService/pkg/zaplog"
)

type (
	working struct {
		Response  http.ResponseWriter
		Request   *http.Request
		Result    Result
		TokenName string
	}
	Result struct {
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}
)

var Context = &working{}

func WorkInit(tokenName string) {
	Context.TokenName = tokenName
}

func WorkNew(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			zaplog.Sugar.Error(err)
		}
	}()

	Context.Response = w
	Context.Request = r
	Context.Result = Result{}
	//解析url传递的参数
	//err = c.Request.ParseForm()

	return
}

func (c *working) CheckAuth() bool {
	token := c.Request.Header.Get(c.TokenName)

	if token == "" {
		c.SetMessage("token empty")
		return true
	}

	err := jwt.Jwt.ValidateToken(token)
	if err != nil {
		c.SetMessage("auth error")
	}

	return err != nil
}

func (c *working) GetUserId() interface{} {
	token := c.Request.Header.Get(c.TokenName)

	var data jwt.Algorithm
	err := jwt.Jwt.DecodeToken(token, &data)
	if err != nil {
		c.SetMessage("auth error")
		return 0
	}

	return data.Sub
}

func (c *working) Get(key string) string {
	return strings.TrimSpace(c.Request.FormValue(key))
}
func (c *working) GetInt64(key string) int64 {
	return utils.StrToInt64(c.Get(key))
}
func (c *working) GetInt(key string) int {
	return utils.StrToInt(c.Get(key))
}
func (c *working) GetJson(data interface{}) error {
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	return err
}

func (c *working) Var(key string) string {
	rv := mux.Vars(c.Request)
	return rv[key]
}

func (c *working) Post(key string) string {
	return strings.TrimSpace(c.Request.PostFormValue(key))
}

func (c *working) Header(key string) string {
	return strings.TrimSpace(c.Request.Header.Get(key))
}

func (c *working) GetIp() string {
	clientIP := c.Header("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = c.Header("XRemoteAddr")
	}
	if clientIP == "" {
		clientIP = strings.TrimSpace(c.Header("X-Real-Ip"))
	}
	if clientIP == "" {
		if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr)); err == nil {
			clientIP = ip
		}
	}
	return clientIP
}

func (c *working) SetCode(code int) {
	zaplog.Sugar.Error(code)
	//c.Result.Message = setting.LangCode(code)
}
func (c *working) SetMessage(msg string) {
	c.Result.Message = msg
}
func (c *working) SetData(data interface{}) {
	c.Result.Data = data
}

func (c *working) WriteJson() {
	var err error
	defer func() {
		if err != nil {
			zaplog.Sugar.Error(err)
		}
	}()
	c.Response.Header().Set("content-type", "text/json")
	c.Response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(c.Response).Encode(c.Result)
}

func (c *working) Write(out interface{}) {
	var err error
	defer func() {
		if err != nil {
			zaplog.Sugar.Error(err)
		}
	}()

	c.Response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(c.Response).Encode(out)
}

func (c *working) WriteXml() {
	var err error
	defer func() {
		if err != nil {
			zaplog.Sugar.Error(err)
		}
	}()
	c.Response.WriteHeader(http.StatusOK)
	err = xml.NewEncoder(c.Response).Encode(c.Result.Data)
}
