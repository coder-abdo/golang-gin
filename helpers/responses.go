package helpers

import (
	"strings"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Message: message,
		Status:  status,
		Errors:  nil,
		Data:    data,
	}
	return res
}
func BuildErrorResponse(message string, err string, data interface{}) Response {
	errs := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  errs,
		Data:    data,
	}
	return res
}
