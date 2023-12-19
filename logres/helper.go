package logger

import (
	"fmt"

	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func ToField(key string, val interface{}) (field slog.Attr) {
	field = slog.Any(key, val)
	return
}

func getLoggerConfig(config LogresConfig) *lumberjack.Logger {
	loggerTdr := lumberjack.Logger{
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
		LocalTime:  config.LocalTime,
	}

	return &loggerTdr
}

func formatMessage(msg ...interface{}) []slog.Attr {
	var logField []slog.Attr
	for index, message := range msg {
		logField = append(logField, ToField(fmt.Sprintf("message_%v", index), message))
	}

	return logField
}
