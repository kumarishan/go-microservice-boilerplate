package logger

import (
	"context"
	"errors"
	"strings"

	"github.com/kumarishan/go-microservice-boilerplate/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	base *zap.SugaredLogger
}

func NewLogger(config *config.Config) (*Logger, error) {
	var cfg zap.Config

	switch strings.ToLower(config.Environment) {
	case "dev", "development":
		cfg = zap.NewDevelopmentConfig()
	case "prod", "production":
		cfg = zap.NewProductionConfig()
	default:
		return nil, errors.New("logger environment not supported.")
	}

	cfg.Level = zap.NewAtomicLevelAt(getLevel(config.Logger.LogLevel))
	cfg.OutputPaths = []string{config.Logger.FileName}
	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	s := log.Sugar()
	return &Logger{base: s}, nil
}

func getLevel(level string) zapcore.Level {
	var zapLevel zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		zapLevel = zap.DebugLevel

	case "info":
		zapLevel = zap.InfoLevel

	case "warn", "warning":
		zapLevel = zap.WarnLevel

	case "error":
		zapLevel = zap.ErrorLevel
	}
	return zapLevel
}

func (s *Logger) with(ctx context.Context) *Logger {
	return s
}

func (s *Logger) With(args ...interface{}) *Logger {
	return &Logger{base: s.base.With(args...)}
}

// Debug uses fmt.Sprint to construct and log a message.
func (s *Logger) Debug(ctx context.Context, args ...interface{}) {
	s.with(ctx).base.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (s *Logger) Info(ctx context.Context, args ...interface{}) {
	s.with(ctx).base.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (s *Logger) Warn(ctx context.Context, args ...interface{}) {
	s.with(ctx).base.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (s *Logger) Error(ctx context.Context, args ...interface{}) {
	s.with(ctx).base.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *Logger) DPanic(ctx context.Context, args ...interface{}) {
	s.with(ctx).base.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (s *Logger) Panic(ctx context.Context, args ...interface{}) {
	s.base.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (s *Logger) Fatal(ctx context.Context, args ...interface{}) {
	s.base.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (s *Logger) Debugf(ctx context.Context, template string, args ...interface{}) {
	s.base.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (s *Logger) Infof(ctx context.Context, template string, args ...interface{}) {
	s.base.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (s *Logger) Warnf(ctx context.Context, template string, args ...interface{}) {
	s.base.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (s *Logger) Errorf(ctx context.Context, template string, args ...interface{}) {
	s.base.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *Logger) DPanicf(ctx context.Context, template string, args ...interface{}) {
	s.base.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (s *Logger) Panicf(ctx context.Context, template string, args ...interface{}) {
	s.base.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (s *Logger) Fatalf(ctx context.Context, template string, args ...interface{}) {
	s.base.Fatalf(template, args)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues...).Debug(msg)
func (s *Logger) Debugw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.base.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *Logger) Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.base.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *Logger) Warnw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.base.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *Logger) Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.base.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (s *Logger) DPanicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.base.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (s *Logger) Panicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.base.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (s *Logger) Fatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.base.Fatalw(msg, keysAndValues...)
}

// Sync flushes any buffered log entries.
func (s *Logger) Sync() error {
	return s.base.Sync()
}
