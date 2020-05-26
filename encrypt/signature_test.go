package encrypt

import (
	"fmt"
	"testing"
)

const (
	msg1       = "hello"
	msg2       = "world"
	token1     = "ricnsmart"
	token2     = "ricn"
	nonce1     = "5xp7vzm3fN"
	nonce2     = "rp/BxZWCeOic"
	signature1 = "d/ncDg jrcW259dHOZByyw=="
	signature2 = "vikQgk0OCjj6RIiIIbLtxA=="
)

func TestNewSignature(t *testing.T) {
	type args struct {
		msg    string
		token  string
		nonces []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{
			msg:    msg1,
			token:  token1,
			nonces: []string{nonce1},
		}},
		{"test2", args{
			msg:    msg2,
			token:  token2,
			nonces: []string{nonce2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signature := NewSignature(tt.args.msg, tt.args.token, tt.args.nonces...)
			fmt.Println(signature)
		})
	}
}

func TestValidateSignature(t *testing.T) {
	type args struct {
		msg       string
		token     string
		nonce     string
		signature string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1", args{
				msg:       msg1,
				token:     token1,
				nonce:     nonce1,
				signature: signature1,
			}, true},
		{
			"test1", args{
				msg:       msg2,
				token:     token2,
				nonce:     nonce2,
				signature: signature2,
			}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateSignature(tt.args.msg, tt.args.token, tt.args.nonce, tt.args.signature); got != tt.want {
				t.Errorf("ValidateSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
