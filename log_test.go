package study

import (
	"context"
	"github.com/natefinch/lumberjack"
	"github.com/wangzhong/golang/study/wlog"
	"testing"
)

func TestRollingLog(t *testing.T) {
	logger := lumberjack.Logger{
		Filename: "./test.log",
		MaxSize:  10,
		MaxAge:   1,
		Compress: false,
	}
	logger.Write([]byte("test log"))
}

func TestPhuslu(t *testing.T) {
	wlog.InitWlog()
	bc := context.Background()
	ctx := context.WithValue(bc, "keys", map[string]any{
		"x-client-trace-id":   "123",
		"x-seayoo-request-id": "abc",
	})
	wlog.Info(ctx).Str("app_id", "1064").Msg("login success")
}
