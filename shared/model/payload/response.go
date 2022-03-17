package payload

import (
	"your/path/project/shared/model/apperror"
)

type Response struct {
	Success      bool        `json:"success"`
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
	TraceID      string      `json:"traceId"`
}

func NewSuccessResponse(data interface{}, traceID string) interface{} {
	var res Response
	res.Success = true
	res.Data = data
	res.TraceID = traceID
	return res
}

func NewErrorResponse(err error, traceID string) interface{} {
	var res Response
	res.Success = false

	et, ok := err.(apperror.ErrorType)
	if !ok {
		res.ErrorCode = "UNDEFINED"
		res.ErrorMessage = err.Error()
		return res
	}

	res.ErrorCode = et.Code()
	res.ErrorMessage = et.Error()
	res.TraceID = traceID
	return res
}
