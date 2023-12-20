package logres

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/spf13/cast"
	"golang.org/x/exp/slog"
)

type Logres interface {
	Info(ctx context.Context, title string, message ...interface{})
	Error(ctx context.Context, title string, message ...interface{})
	TDR(ctx context.Context, request []byte, response []byte)
}

type defaultLogger struct {
	syslog *slog.Logger
	tdrlog *slog.Logger
}

func SetLogger(config LogresConfig) Logres {
	fmt.Println("Start Logger.....")

	var loggingLevel = new(slog.LevelVar)
	jsonHandlerSys := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: loggingLevel})
	jsonHandlerTdr := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: loggingLevel})

	if config.LogsWrite {
		_, err := os.Stat(config.FolderPath)
		if os.IsNotExist(err) {
			err := os.MkdirAll(config.FolderPath, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		loggerSys := getLoggerConfig(config)
		loggerSys.Filename = config.FolderPath + "/sys.log"

		loggerTdr := getLoggerConfig(config)
		loggerTdr.Filename = config.FolderPath + "/tdr.log"

		jsonHandlerSys = slog.NewJSONHandler(loggerSys, &slog.HandlerOptions{Level: loggingLevel})
		jsonHandlerTdr = slog.NewJSONHandler(loggerTdr, &slog.HandlerOptions{Level: loggingLevel})
	}

	syslogger := slog.New(jsonHandlerSys)
	slog.SetDefault(syslogger)
	loggingLevel.Set(slog.LevelDebug)

	tdrlogger := slog.New(jsonHandlerTdr)
	slog.SetDefault(tdrlogger)

	return &defaultLogger{
		syslog: syslogger,
		tdrlog: tdrlogger,
	}
}

func (s *defaultLogger) Info(ctx context.Context, title string, message ...interface{}) {
	atr := formatMessage(message...)
	contexLog := GetCtxLogger(ctx)
	s.syslog.InfoContext(ctx, title, slog.Any("SYS", contexLog), slog.Any("atribute", atr))
}

func (s *defaultLogger) Error(ctx context.Context, title string, message ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[Error] %s:%d", filename, line)
	message = append(message, msg)

	atr := formatMessage(message...)
	contexLog := GetCtxLogger(ctx)
	s.syslog.ErrorContext(ctx, title, slog.Any("SYS", contexLog), slog.Any("atribute", atr))
}

func (s *defaultLogger) TDR(ctx context.Context, request []byte, response []byte) {
	rt := time.Since(GetRequestTimeFromContext(ctx)).Nanoseconds() / 1000000
	errMsg := GetErrorMessageFromContext(ctx)
	contexLog := GetCtxLogger(ctx)

	var resp interface{}
	var req interface{}
	json.Unmarshal(request, &req)
	json.Unmarshal(response, &resp)

	s.syslog.Info("[RESPONSE]", slog.Any("SYS", contexLog), slog.Any("resp", resp))

	tdrLog := LogTdrModel{
		RequestId: contexLog.ThreadID,
		Path:      contexLog.ReqURI,
		Method:    contexLog.ReqMethod,
		Port:      contexLog.ServicePort,
		RespTime:  rt,
		Request:   req,
		Response:  resp,
		Error:     errMsg,
	}

	if header, ok := contexLog.Header.(http.Header); ok {
		var headerMap = map[string]string{}
		for key, values := range header {
			for _, value := range values {
				headerMap[key] = value
			}
		}
		tdrLog.Header = headerMap
	}

	responseMap := make(map[string]interface{})
	if err := json.Unmarshal(response, &responseMap); err != nil {
		s.tdrlog.InfoCtx(ctx, "TDR", slog.Any("TDR", tdrLog))
		return
	}

	if responseMap["code"] != nil {
		tdrLog.ResponseCode = cast.ToString(responseMap["code"])
	}

	s.tdrlog.InfoCtx(ctx, "TDR", slog.Any("TDR", tdrLog))
}
