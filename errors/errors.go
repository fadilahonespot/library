package errors

import (
	"net/http"
	"strconv"
)

type ApplicationError struct {
	ErrorCode      int
	Message        string
	Data           interface{}
	OverideMessage bool
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetError(errorCode int, message string) error {
	return &ApplicationError{
		ErrorCode: errorCode,
		Message:   message,
	}
}

func SetErrorMessage(errorCode int, message string) error {
	return &ApplicationError{
		ErrorCode:      errorCode,
		Message:        message,
		OverideMessage: true,
	}
}

func SetErrorMessageWithData(errorCode int, message string, data interface{}) error {
	return &ApplicationError{
		ErrorCode:      errorCode,
		Message:        message,
		Data:           data,
		OverideMessage: true,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func (e *ApplicationError) Code() int {
	return e.ErrorCode
}

func GetErrorCode(err error) int {
	if err == nil {
		return 0
	}

	if se, ok := err.(interface {
		Code() int
	}); ok {
		return se.Code()
	}
	return 0
}

func ErrorHandle(err error) Response {
	code := GetErrorCode(err)
	if code == 0 {
		code = http.StatusInternalServerError
	}

	message := http.StatusText(code)
	var data interface{}
	if he, ok := err.(*ApplicationError); ok {
		if he.OverideMessage {
			message = he.Message
			if he.Data != nil {
				data = he.Data
			}
		}
	}

	return Response{
		Code:    strconv.Itoa(code),
		Message: message,
		Data:    data,
	}
}
