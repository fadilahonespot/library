package response

import (
	"encoding/json"
	"math"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ItemPages struct {
	Page      int64 `json:"page"`
	Limit     int64 `json:"limit"`
	TotalData int64 `json:"totalData"`
	TotalPage int64 `json:"totalPage"`
}

type responsePagination struct {
	Response
	Pagination ItemPages `json:"pagination"`
}

func ResponseSuccess(data interface{}) Response {
	return Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}

func ResponseProcees(data interface{}) Response {
	return Response{
		Code:    http.StatusProcessing,
		Message: "Request has been successfully processed",
		Data:    data,
	}
}

func ResponseError(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    struct{}{},
	}
}

func HandleSuccessWithPagination(totalItems float64, limit, page int, data interface{}) responsePagination {
	var totalPage float64 = 1
	if limit != 0 && page != 0 {
		res := totalItems / float64(limit)
		totalPage = math.Ceil(res)
	}

	var values = []string{}
	tx, _ := json.Marshal(data)
	if string(tx) == "null" {
		data = values
	}

	resp := responsePagination{
		Response: Response{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    data,
		},
		Pagination: ItemPages{
			TotalData: int64(totalItems),
			TotalPage: int64(totalPage),
			Page:      int64(page),
			Limit:     int64(limit),
		},
	}
	return resp
}
