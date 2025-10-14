package logging

import (
	"context"
	"fmt"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (level LogLevel) String() string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	default:
		panic(fmt.Sprintf("unexpected logging.LogLevel: %#v", level))
	}
}

type ILogger interface {
	DEBUG(ctx context.Context, msg string, args ...interface{})
	INFO(ctx context.Context, msg string, args ...interface{})
	WARN(ctx context.Context, msg string, args ...interface{})
	ERROR(ctx context.Context, msg string, args ...interface{})
	FATAL(ctx context.Context, msg string, args ...interface{})
	Log(level LogLevel, ctx context.Context, msg string, args ...interface{})

	With(args ...any)
	WithGroup(name string, args ...any)
	WithContext(ctx context.Context)
	Clear()
}

type LoggerOptions struct {
	MinLevel    LogLevel
	JsonLogging bool
	LoggingRate int
}

// type LogEntry struct {
// 	TimeStamp time.Time `json:"time_stamp"`
// 	Level     string    `json:"level"`
// 	Message   string    `json:"message"`
// 	Request   struct {
// 		Method    string    `json:"method"`
// 		Path      string    `json:"ath"`
// 		Query     string    `json:"query"`
// 		UserAgent string    `json:"user_agent"`
// 		RequestId uuid.UUID `json:"request_id"`
// 		IpAdress  string    `json:"ip_address"`
// 	} `json:"request"`
// 	UserId uuid.UUID `json:"user_id"`

// }
