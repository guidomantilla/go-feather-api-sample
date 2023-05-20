package security

import (
	"net/http"
)

type Exception struct {
	Code    int      `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

func BadRequestException(message string, err ...error) *Exception {

	var errors []string
	if len(err) >= 0 {
		for _, err2 := range err {
			errors = append(errors, err2.Error())
		}
	}

	return &Exception{
		Message: message,
		Errors:  errors,
		Code:    http.StatusBadRequest,
	}
}

func UnauthorizedException(message string, err ...error) *Exception {

	var errors []string
	if len(err) >= 0 {
		for _, err2 := range err {
			errors = append(errors, err2.Error())
		}
	}

	return &Exception{
		Message: message,
		Errors:  errors,
		Code:    http.StatusUnauthorized,
	}
}

func NotFoundException(message string, err ...error) *Exception {

	var errors []string
	if len(err) >= 0 {
		for _, err2 := range err {
			errors = append(errors, err2.Error())
		}
	}

	return &Exception{
		Message: message,
		Errors:  errors,
		Code:    http.StatusNotFound,
	}
}

func InternalServerErrorException(message string, err ...error) *Exception {

	var errors []string
	if len(err) >= 0 {
		for _, err2 := range err {
			errors = append(errors, err2.Error())
		}
	}

	return &Exception{
		Message: message,
		Errors:  errors,
		Code:    http.StatusInternalServerError,
	}
}
