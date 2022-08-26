package resterrors

import (
	"errors"
	"net/http"
	"time"
)

type RestErr struct {
	TimeStamp time.Time `json:"time_stamp"`
	Message   string    `json:"message"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusBadRequest,
		Error:     "bad_request",
	}
}
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusInternalServerError,
		Error:     "internal_server_error",
	}
}
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusNotFound,
		Error:     "not_found",
	}
}
