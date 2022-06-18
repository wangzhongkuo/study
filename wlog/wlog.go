package wlog

import (
	"context"
	"fmt"
	plog "github.com/phuslu/log"
	"io"
	"strings"
)

var defaultWlog = wlog{plog.DefaultLogger}

type CTX context.Context

var TidKey = "x-client-trace-id"
var RidKey = "x-seayoo-request-id"

type wlog struct {
	plog.Logger
}

func InitWlog() {
	writer := plog.MultiEntryWriter{
		&plog.ConsoleWriter{
			Formatter: func(w io.Writer, a *plog.FormatterArgs) (n int, err error) {
				tid := a.Get(TidKey)
				rid := a.Get(RidKey)
				n, _ = fmt.Fprintf(w, "%s %s %s [%s,%s] >", a.Time, strings.ToUpper(a.Level), a.Caller, tid, rid)
				for _, kv := range a.KeyValues {
					if kv.Key != TidKey && kv.Key != RidKey {
						i, _ := fmt.Fprintf(w, " %s=%s", kv.Key, kv.Value)
						n += i
					}
				}
				i, err := fmt.Fprintf(w, " %s\n", a.Message)
				return n + i, err
			},
		},
		//&plog.ConsoleWriter{
		//	ColorOutput:    false,
		//	QuoteString:    true,
		//	EndWithMessage: true,
		//	Writer: &plog.FileWriter{
		//		Filename:     "./phuslu.log",
		//		TimeFormat:   "2006-01-02",
		//		EnsureFolder: true,
		//	},
		//},
	}
	logger := plog.Logger{
		Caller:     1,
		TimeFormat: "2006-01-02 15:04:05.000",
		Writer:     &writer,
	}
	defaultWlog = wlog{logger}
}

func Info(ctx ...context.Context) *plog.Entry {
	if len(ctx) == 0 {
		return defaultWlog.Info()
	}
	return defaultWlog.Info().Context(logContext(ctx[0]))
}

func logContext(ctx context.Context) plog.Context {
	pc := plog.NewContext(nil)
	keys, ok := ctx.Value("keys").(map[string]any)
	if !ok {
		return pc.Value()
	}
	tid, _ := keys[TidKey].(string)
	rid, _ := keys[RidKey].(string)
	pc.Str(TidKey, tid)
	pc.Str(RidKey, rid)
	return pc.Value()
}
