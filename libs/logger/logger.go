package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var (
	ZeroLog *zerolog.Logger
	logOnce sync.Once
)

func NewLogger() *zerolog.Logger {
	logOnce.Do(func() {
		consoleWriter := zerolog.ConsoleWriter{
			Out: os.Stderr,

			NoColor: false,

			TimeFormat: time.RFC3339,
			FormatLevel: func(v interface{}) string {
				return fmt.Sprintf("[%v]", v)
			},
			FormatCaller: func(v interface{}) string {
				return fmt.Sprintf("(%v)", v)
			},
			FormatMessage: func(v interface{}) string {
				return fmt.Sprintf(">>>>>%v", v)
			},
			FormatFieldName: func(v interface{}) string {
				return fmt.Sprintf("{%v}", v)
			},
			FormatFieldValue: func(v interface{}) string {
				return fmt.Sprintf("??%v??", v)
			},
			FormatErrFieldName: func(v interface{}) string {
				return fmt.Sprintf("[[%v]]", v)
			},
			FormatErrFieldValue: func(v interface{}) string {
				return fmt.Sprintf("!!%v!!", v)
			},
		}

		logger := zerolog.New(consoleWriter).Level(zerolog.DebugLevel).With().Caller().Timestamp().Logger()
		ZeroLog = &logger
	})

	return ZeroLog
}
