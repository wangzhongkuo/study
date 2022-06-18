package study

import (
	"context"
	slog "git.shiyou.kingsoft.com/go/log"
	"github.com/natefinch/lumberjack"
	plog "github.com/phuslu/log"
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
	// middleware 里给gin的context设置如下contex（同时设置了x-seayoo-request-id header）
	bc := context.Background()
	ctx := context.WithValue(bc, "keys", map[string]any{
		wlog.TidKey: "123",
		wlog.RidKey: "abc",
	})
	plog.Info()
	// cong gin context中获取，并创建go context
	wlog.Info(ctx).Str("app_id", "1064").Msg("login success")
	slog.Infox(context.Background(), "aaa")
}
