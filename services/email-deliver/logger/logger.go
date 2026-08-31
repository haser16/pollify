package services_email_logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(config Config) (*Logger, error) {
	zapLvl := zap.NewAtomicLevel()

	if err := zapLvl.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("unmarshal log level: %w", err)
	}

	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeTime =
		zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapConfig.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		zapLvl,
	)

	return &Logger{
		Logger: zap.New(core, zap.AddCaller()),
	}, nil
}

func (l *Logger) Close() {
	_ = l.Logger.Sync()
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
	}
}
