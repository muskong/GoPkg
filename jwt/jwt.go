package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/muskong/GoService/pkg/idworker"
	"github.com/muskong/GoService/pkg/rsa"
)

type (
	Algorithm struct {
		/**
		 * jti(JWT ID)
		 * 签发jwt时给予当前token的唯一ID，通常用于一次性消费的token。
		 */
		Jti int64 `json:"jti"`

		/**
		 * iss(Issuer) jwt的颁发者，其值应为大小写敏感的字符串或Uri。
		**/
		Iss string `json:"iss"`
		/**
		 * aud(Audience)
		 * jwt的适用对象，其值应为大小写敏感的字符串或Uri。一般可以为特定的App、服务或模块。
		 * 比如我们颁发了一个jwt给一个叫”JsonWebToken”的app使用，sub可以是这个app的包签名或者标识。
		 * 服务器端的安全策略在签发时和验证时，aud必须是一致的。
		**/
		Aud string `json:"aud"`

		/**
		 * sub(Subject)
		 * jwt 的所有者，可以是用户ID、唯一标识。
		**/
		Sub interface{} `json:"sub"`

		/**
		 * iat(Issued At)
		 * jwt的签发时间。同exp一样，需为可以解析成时间的数字类型。
		 */
		Iat int64 `json:"iat"` /** 发布时间 **/
		/**
		 * exp(Expiration Time)
		 * jwt的过期时间，必须是可以解析为时间/时间戳的数字类型。服务器端在验证当前时间大于过期时间时，应当验证不予通过。
		 */
		Exp time.Time `json:"exp"`
		/**
		 * nbf(Not Before)
		 * 表示jwt在这个时间后启用。同exp一样，需为可以解析成时间的数字类型。
		 * 在此之前不可用, 表示 JWT Token 在这个时间之前是无效的
		 */
		Nbf time.Time `json:"nbf"`
	}

	_jwt struct {
		tokenName string
		algorithm *Algorithm
		rsa       *rsa.Rsa
	}
	JwtConfig struct {
		TokenName string
		Server    int64
		Exp       int64
		Iss       string
		Aud       string
		Pub       string
		Pri       string
	}
)

var (
	Jwt     *_jwt
	jwtOnce sync.Once
)

func JwtInit(jcfg *JwtConfig) {
	jwtOnce.Do(jcfg.initJwt)
}

func (jcfg *JwtConfig) initJwt() {
	expTime := time.Duration(jcfg.Exp) * time.Hour
	Jwt = &_jwt{
		tokenName: jcfg.TokenName,
		algorithm: &Algorithm{
			Iss: jcfg.Iss,
			Aud: jcfg.Aud,
			Nbf: time.Now(),
			Iat: time.Now().Unix(),
			Exp: time.Now().Add(expTime),
			Jti: idworker.IdInt64(jcfg.Server),
		},
		rsa: rsa.NewRsa(jcfg.Pub, jcfg.Pri),
	}
}

func (j *_jwt) GetTokenName() string {
	return j.tokenName
}

func (j *_jwt) GenerateToken(data interface{}) string {
	j.algorithm.Sub = data

	token, err := j.encode()
	if err != nil {
		return ""
	}

	return token
}

func (j *_jwt) decodeToken(token string, data *Algorithm) (err error) {
	dd, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		err = errors.New("invalid signature")
		return
	}

	fmt.Printf("j.rsa %+v", j.rsa)

	de, err := j.rsa.RsaDecrypt(dd)
	if err != nil {
		err = errors.New("invalid signature")
		return
	}

	return json.Unmarshal(de, &data)
}

func (j *_jwt) ValidateToken(token string) (data *Algorithm, err error) {
	var a Algorithm
	err = j.decodeToken(token, &a)
	if err != nil {
		err = errors.New("decode error")
		return
	}

	if err = j.validateExp(); err != nil {
		err = errors.New("failed to validate exp" + err.Error())
		return
	}

	if err = j.validateNbf(); err != nil {
		err = errors.New("failed to validate nbf" + err.Error())
	}

	data = &a
	return
}

func (j *_jwt) encode() (token string, err error) {
	orig, err := json.Marshal(j.algorithm)
	if err != nil {
		err = errors.New("unable to marshal payload" + err.Error())
		return
	}

	byteToken, err := j.rsa.RsaEncrypt(orig)
	if err != nil {
		err = errors.New("rsa encrypt encode " + err.Error())
		return
	}

	token = base64.RawURLEncoding.EncodeToString(byteToken)

	return
}

func (j *_jwt) validateExp() error {
	if j.algorithm.Exp.Before(time.Now()) {
		return errors.New("token has expired")
	}

	return nil
}

func (j *_jwt) validateNbf() error {
	if j.algorithm.Nbf.After(time.Now()) {
		return errors.New("token is invalid")
	}

	return nil
}
