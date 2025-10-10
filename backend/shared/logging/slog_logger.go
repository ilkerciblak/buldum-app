package logging

import (
	"context"
)

type Slogger struct {
	Logger
}

func (l *Slogger) SetLevel(level LogLevel) {
	l.Level = level
}

func (l *Slogger) WithField(fieldKey string, fieldValue interface{}) {

}

func (l *Slogger) WithFields(fields map[string]interface{}) {

}

func (l *Slogger) WithContext(ctx context.Context) {

}
func (l *Slogger) Log(level LogLevel, message string) {

}
func (l *Slogger) LogSelf(message string) {

}

func A() ILogger {
	return &Slogger{}
}
