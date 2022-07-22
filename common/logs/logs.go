package logs

import "go.uber.org/zap"

type AppLog interface {
	Info(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Warning(msg string, fields ...zap.Field)
	Error(msg interface{}, fields ...zap.Field)
}
