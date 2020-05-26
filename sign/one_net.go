package sign

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
)

func NewSignatureWithTimestamp(accessKey, method, res, version string, et int) (string, error) {
	// 对access_key进行decode
	key, err := base64.StdEncoding.DecodeString(accessKey)
	if err != nil {
		return "", err
	}
	var StringForSignature = fmt.Sprintf("%v\n%v\n%v\n%v", et, method, res, version)
	// 计算sign = base64(hmac_<method>(base64decode(accessKey), utf-8(StringForSignature)))
	h := hmac.New(sha1.New, key)
	_, err = h.Write([]byte(StringForSignature))
	if err != nil {
		return "", err
	}
	// 拼装token 对value部分需要经过URL编码
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

// 生成签名
// AES对称加密
func NewSignatureWithNonce(msg, token, nonce string) string {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf(`%v%v%v`, token, nonce, msg)))
	str := base64.StdEncoding.EncodeToString(m.Sum(nil))
	// + 号替换为 空格
	return strings.Replace(str, "+", " ", -1)
}

// 验证签名
func ValidateSignatureWithNonce(msg, token, nonce, signature string) bool {
	expect := NewSignatureWithNonce(msg, token, nonce)
	return expect == signature
}
