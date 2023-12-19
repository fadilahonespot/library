package logres

import (
	"context"
	"time"
)

const (
	LoggerContext = "logger-context"
	RequestTime   = "request-time"
	RequestIp     = "request-ip"
	ErrorMessage  = "error-message"
)

func SetCtxLogger(ctx context.Context, dataLog Context) context.Context {
	ctx = context.WithValue(ctx, LoggerContext, dataLog)
	ctx = context.WithValue(ctx, RequestTime, time.Now())
	return ctx
}

func GetCtxLogger(ctx context.Context) Context {
	s, ok := ctx.Value(LoggerContext).(Context)
	if !ok {
		return Context{}
	}

	return s
}

func GetRequestTimeFromContext(ctx context.Context) time.Time {
	s, ok := ctx.Value(RequestTime).(time.Time)
	if !ok {
		return time.Now()
	}

	return s
}

func SetErrorMessage(ctx context.Context, errMessage string) context.Context {
	ctx = context.WithValue(ctx, ErrorMessage, errMessage)
	return ctx
}

func GetErrorMessageFromContext(ctx context.Context) string {
	s, ok := ctx.Value(ErrorMessage).(string)
	if !ok {
		return ""
	}

	return s
}
