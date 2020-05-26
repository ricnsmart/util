package rand

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{10}, 10},
		{"test2", args{12}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := RandomString(tt.args.len)
			fmt.Println(str)
		})
	}
}
