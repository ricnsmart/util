package token

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func GeneratorSignature(accessKey, method, res, version string, et int) (string, error) {
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
