package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	With(fields ...Field) Logger
}

type Field struct {
	Key   string
	Value interface{}
}

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(production bool) *ZapLogger {
	var config zap.Config
	if production {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	logger, _ := config.Build()
	return &ZapLogger{logger: logger}
}

func (l *ZapLogger) Info(msg string, fields ...Field) {
	zapFields := toZapFields(fields)
	l.logger.Info(msg, zapFields...)
}

func (l *ZapLogger) Error(msg string, fields ...Field) {
	zapFields := toZapFields(fields)
	l.logger.Error(msg, zapFields...)
}

func (l *ZapLogger) With(fields ...Field) Logger {
	return &ZapLogger{
		logger: l.logger.With(toZapFields(fields)...),
	}
}

func toZapFields(fields []Field) []zapcore.Field {
	zapFields := make([]zapcore.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}

// HTTP中间件日志记录器
func HttpLogger(logger Logger) http.MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(ctx http.Context) error {
			start := time.Now()

			err := next(ctx)

			fields := []Field{
				{"method", ctx.Request().Method},
				{"path", ctx.Request().URL.Path},
				{"status", ctx.Writer().Status()},
				{"duration", time.Since(start)},
			}

			if err != nil {
				fields = append(fields, Field{"error", err.Error()})
				logger.Error("HTTP request error", fields...)
			} else {
				logger.Info("HTTP request", fields...)
			}

			return err
		}
	}
}
