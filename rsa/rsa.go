package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type Rsa struct {
	publicKey  []byte
	privateKey []byte
}

func NewRsa(pub, pri string) *Rsa {
	return &Rsa{
		publicKey:  []byte(pub),
		privateKey: []byte(pri),
	}
}

// 私钥生成
//openssl genrsa -out rsa_private_key.pem 2048
// 公钥: 根据私钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem

// 加密
func (r *Rsa) RsaEncrypt(originData []byte) ([]byte, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode(r.publicKey)
	if block == nil {
		return nil, errors.New("公钥错误")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 类型段言
	pub := pubInterface.(*rsa.PublicKey)

	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, originData)
}

func (r *Rsa) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	// 解密
	block, _ := pem.Decode(r.privateKey)
	if block == nil {
		return nil, errors.New("公钥错误")
	}

	// 解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
