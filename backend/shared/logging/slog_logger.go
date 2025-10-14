package logging

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type Slogger struct {
	*slog.Logger
	LoggerOptions
}

func NewSlogger(opts LoggerOptions) *Slogger {
	var handler slog.Handler
	var handlerOptions slog.HandlerOptions = slog.HandlerOptions{
		Level: slog.Level((opts.MinLevel - 1) * 4),
	}

	if opts.JsonLogging {

		handler = slog.NewJSONHandler(
			os.Stdout,
			&handlerOptions,
		)
	} else {
		handler = slog.NewTextHandler(
			os.Stdout,
			&handlerOptions,
		)
	}

	logger := slog.New(handler)

	return &Slogger{
		Logger:        logger,
		LoggerOptions: opts,
	}
}

func (l *Slogger) Clear() {
	l.Logger = NewSlogger(l.LoggerOptions).Logger
}

func (l *Slogger) DEBUG(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.DebugContext(ctx, msg, args...)
}
func (l *Slogger) INFO(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.InfoContext(ctx, msg, args...)
}
func (l *Slogger) WARN(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.WarnContext(ctx, msg, args...)
}
func (l *Slogger) ERROR(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.ErrorContext(ctx, msg, args...)
}
func (l *Slogger) FATAL(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}

func (l *Slogger) Log(level LogLevel, ctx context.Context, msg string, args ...interface{}) {

	if level >= l.MinLevel {
		switch level {
		case DEBUG:
			l.DEBUG(ctx, msg, args...)
		case ERROR:
			l.ERROR(ctx, msg, args...)
		case FATAL:
			l.FATAL(ctx, msg, args...)
		case INFO:
			l.INFO(ctx, msg, args...)
		case WARN:
			l.WARN(ctx, msg, args...)
		default:
			panic(fmt.Sprintf("unexpected logging.LogLevel: %#v", level))
		}
	}
}

func (l *Slogger) With(args ...any) {
	l.Logger = l.Logger.With(args...)
}

func (l *Slogger) WithGroup(name string, args ...interface{}) {

	groupedAttributes := slog.Group(name, args...)
	l.With(groupedAttributes.Key, groupedAttributes.Value)
}

func (l *Slogger) WithContext(ctx context.Context) {
	if val := ctx.Value("request_id"); val != nil {
		l.With("request_id", val)
	}

	if val := ctx.Value("user_id"); val != nil {
		l.With("user_id", val)
	}

	if val := ctx.Value("device_info"); val != nil {
		l.With("user_id", val)
	}
}

func A() ILogger {
	return &Slogger{}
}
