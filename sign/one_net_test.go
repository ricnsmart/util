package sign

import (
	"fmt"
	"testing"
)

const (
	accessKey = "KuF3NT/jUBJ62LNBB/A8XZA9CqS3Cu79B/ABmfA1UCw="
	method    = "sha1"
	version   = "2018-10-31"
	res       = "products/123123"
	et        = 1537255523
	sign      = "lsaPSiiGvEFFjXu5WU7a6IkScqE="

	msg1       = "hello"
	msg2       = "world"
	token1     = "ricnsmart"
	token2     = "ricn"
	nonce1     = "5xp7vzm3fN"
	nonce2     = "rp/BxZWCeOic"
	signature1 = "d/ncDg jrcW259dHOZByyw=="
	signature2 = "vikQgk0OCjj6RIiIIbLtxA=="
)

func TestNewToken(t *testing.T) {
	type args struct {
		accessKey string
		method    string
		res       string
		version   string
		et        int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"test1", args{
			accessKey: accessKey,
			method:    method,
			res:       res,
			version:   version,
			et:        et,
		}, sign, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSignatureWithTimestamp(tt.args.accessKey, tt.args.method, tt.args.res, tt.args.version, tt.args.et)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTokenWithTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewSignatureWithTimestamp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSignature(t *testing.T) {
	type args struct {
		msg   string
		token string
		nonce string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{
			msg:   msg1,
			token: token1,
			nonce: nonce1,
		}},
		{"test2", args{
			msg:   msg2,
			token: token2,
			nonce: nonce2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signature := NewSignatureWithNonce(tt.args.msg, tt.args.token, tt.args.nonce)
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
			if got := ValidateSignatureWithNonce(tt.args.msg, tt.args.token, tt.args.nonce, tt.args.signature); got != tt.want {
				t.Errorf("ValidateSignatureWithNonce() = %v, want %v", got, tt.want)
			}
		})
	}
}
