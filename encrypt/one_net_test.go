package encrypt

import (
	"testing"
)

const (
	accessKey = "KuF3NT/jUBJ62LNBB/A8XZA9CqS3Cu79B/ABmfA1UCw="
	method    = "sha1"
	version   = "2018-10-31"
	res       = "products/123123"
	et        = 1537255523
	sign      = "lsaPSiiGvEFFjXu5WU7a6IkScqE="
	token     = "version=2018-10-31&res=products%2F123123&et=1537255523&method=sha1&sign=lsaPSiiGvEFFjXu5WU7a6IkScqE%3D"
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
			got, err := GeneratorSignature(tt.args.accessKey, tt.args.method, tt.args.res, tt.args.version, tt.args.et)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTokenWithTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeneratorSignature() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewToken1(t *testing.T) {
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
		}, token, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTokenWithTimestamp(tt.args.accessKey, tt.args.method, tt.args.res, tt.args.version, tt.args.et)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTokenWithTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewTokenWithTimestamp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
