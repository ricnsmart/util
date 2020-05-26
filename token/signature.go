package token

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
)

// 生成签名
// AES对称加密
func NewSignature(msg, token, nonce string) string {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf(`%v%v%v`, token, nonce, msg)))
	str := base64.StdEncoding.EncodeToString(m.Sum(nil))
	// + 号替换为 空格
	return strings.Replace(str, "+", " ", -1)
}

// 验证签名
func ValidateSignature(msg, token, nonce, signature string) bool {
	expect := NewSignature(msg, token, nonce)
	return expect == signature
}

func NewTokenWithNonce(msg, token, nonce string) string {
	signature := NewSignature(msg, token, nonce)
	return fmt.Sprintf(`msg=%v&nonce=%v&signature=%v`, msg, nonce, signature)
}
