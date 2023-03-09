package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (l ZapLogger) Debug(s string) {
	l.logger.Sugar().Debug(s)
}

func (l ZapLogger) Info(s string) {
	l.logger.Sugar().Info(s)
}

func (l ZapLogger) Warn(s string) {
	l.logger.Sugar().Warn(s)
}

func (l ZapLogger) Error(s string) {
	l.logger.Sugar().Error()
}

func (l ZapLogger) Panic(s string) {
	l.logger.Sugar().Panic(s)
}

func (l ZapLogger) Fatal(s string) {
	l.logger.Sugar().Fatal(s)
}

func NewZapLogger(m *string) *ZapLogger {
	mode := ""
	if m == nil {
		mode = "dev"
	} else {
		mode = *m
	}

	var logger *zap.Logger
	if mode == "dev" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	return &ZapLogger{logger: logger}
}
