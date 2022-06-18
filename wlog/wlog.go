package wlog

import (
	"context"
	"fmt"
	plog "github.com/phuslu/log"
	"io"
	"strings"
)

var defaultWlog wlog

type wlog struct {
	plog.Logger
}

func InitWlog() {
	writer := plog.MultiEntryWriter{
		&plog.ConsoleWriter{
			ColorOutput: true,
			//QuoteString:    true,
			EndWithMessage: true,
			Formatter: func(w io.Writer, a *plog.FormatterArgs) (n int, err error) {
				tid := a.Get("x-client-trace-id")
				rid := a.Get("x-seayoo-request-id")
				n, _ = fmt.Fprintf(w, "%s %s %s [%s,%s] >", a.Time, strings.ToUpper(a.Level), a.Caller, tid, rid)
				for _, kv := range a.KeyValues {
					if kv.Key != "x-client-trace-id" && kv.Key != "x-seayoo-request-id" {
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

func Info(ctx context.Context) *plog.Entry {
	return defaultWlog.Info().Context(logContext(ctx))
}

func logContext(ctx context.Context) plog.Context {
	pc := plog.NewContext(nil)
	keys, ok := ctx.Value("keys").(map[string]any)
	if !ok {
		return pc.Value()
	}
	tid, _ := keys["x-client-trace-id"].(string)
	rid, _ := keys["x-seayoo-request-id"].(string)
	pc.Str("x-client-trace-id", tid)
	pc.Str("x-seayoo-request-id", rid)
	return pc.Value()
}
