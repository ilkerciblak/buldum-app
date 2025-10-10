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
	SetLevel(level LogLevel)
	WithContext(ctx context.Context)
	WithField(fieldKey string, fieldValue interface{})
	WithFields(fields map[string]interface{})
	Log(level LogLevel, message string)
	LogSelf(message string)
}

type LogEntry struct {
	Level        string                 `json:"level,omitempty"`
	Service      string                 `json:"service,omitempty"`
	Status       string                 `json:"status,omitempty"`
	RequestID    string                 `json:"request_id,omitempty"`
	UserId       string                 `json:"user_id,omitempty"`
	DeviceInfo   string                 `json:"device_info,omitempty"`
	AttemptNum   string                 `json:"attempt_num,omitempty"`
	Timestamp    string                 `json:"timestamp,omitempty"`
	Request      map[string]interface{} `json:"request,omitempty"`
	Response     map[string]interface{} `json:"response,omitempty"`
	EllapsedTime string                 `json:"ellapsed_time,omitempty"`
}

type Logger struct {
	Service string
	Level   LogLevel
	LogEntry
}
