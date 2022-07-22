package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type appLogs struct {
	log *zap.Logger
}

func NewAppLogs() AppLog {
	var log *zap.Logger

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return appLogs{log: log}
}

func (l appLogs) Info(msg string, fields ...zap.Field) {
	l.log.Info(msg, fields...)
}

func (l appLogs) Debug(msg string, fields ...zap.Field) {
	l.log.Debug(msg)
}

func (l appLogs) Warning(msg string, fields ...zap.Field) {
	l.log.Warn(msg, fields...)
}

func (l appLogs) Error(msg interface{}, fields ...zap.Field) {
	switch v := msg.(type) {
	case error:
		l.log.Error(v.Error(), fields...)
	case string:
		l.log.Error(v, fields...)
	}
}
