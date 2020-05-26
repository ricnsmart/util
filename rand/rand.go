package rand

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func RandomString(len int) string {
	var container string
	b := bytes.NewBufferString(encodeStd)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(encodeStd[randomInt.Int64()])
	}
	return container
}
