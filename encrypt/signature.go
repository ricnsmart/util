package encrypt

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/ricnsmart/util/rand"
	"strings"
)

// 生成签名
// AES对称加密
func NewSignature(msg, token string, nonces ...string) string {
	m := md5.New()
	var nonce string
	if len(nonces) != 0 {
		nonce = nonces[0]
	} else {
		nonce = rand.RandomString(10)
	}
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
