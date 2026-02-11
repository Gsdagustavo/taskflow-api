package util

import (
	"context"
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	zapLogger := InitZap()

	handler := NewZapHandler(zapLogger)
	slog.SetDefault(slog.New(handler))

	return zapLogger
}

func InitZap() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.ConsoleSeparator = " | "
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		DisableStacktrace: true,
		DisableCaller:     true,
	}

	return zap.Must(cfg.Build())
}

type ZapHandler struct {
	logger *zap.Logger
}

func NewZapHandler(l *zap.Logger) slog.Handler {
	return &ZapHandler{logger: l}
}

func (h *ZapHandler) Enabled(_ context.Context, level slog.Level) bool {
	switch level {
	case slog.LevelDebug:
		return h.logger.Core().Enabled(zap.DebugLevel)
	case slog.LevelInfo:
		return h.logger.Core().Enabled(zap.InfoLevel)
	case slog.LevelWarn:
		return h.logger.Core().Enabled(zap.WarnLevel)
	case slog.LevelError:
		return h.logger.Core().Enabled(zap.ErrorLevel)
	default:
		return h.logger.Core().Enabled(zap.InfoLevel)
	}
}

func (h *ZapHandler) Handle(_ context.Context, r slog.Record) error {
	fields := make([]zap.Field, 0, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields = append(fields, zap.Any(a.Key, a.Value.Any()))
		return true
	})

	switch r.Level {
	case slog.LevelDebug:
		h.logger.Debug(r.Message, fields...)
	case slog.LevelInfo:
		h.logger.Info(r.Message, fields...)
	case slog.LevelWarn:
		h.logger.Warn(r.Message, fields...)
	case slog.LevelError:
		h.logger.Error(r.Message, fields...)
	default:
		h.logger.Info(r.Message, fields...)
	}

	return nil
}

func (h *ZapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	fields := make([]zap.Field, 0, len(attrs))
	for _, a := range attrs {
		fields = append(fields, zap.Any(a.Key, a.Value.Any()))
	}
	return &ZapHandler{
		logger: h.logger.With(fields...),
	}
}

func (h *ZapHandler) WithGroup(name string) slog.Handler {
	return &ZapHandler{
		logger: h.logger.Named(name),
	}
}
