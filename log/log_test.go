package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestInfo(t *testing.T) {
	Init("test.log")
	type args struct {
		msg    string
		fields []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		{"test", args{
			msg:    "test",
			fields: []zap.Field{zap.String("service", "GS524N")},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.msg, tt.args.fields...)
		})
	}
}
