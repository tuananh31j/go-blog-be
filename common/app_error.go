package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"_"`
	Message    string `json:"message"`
	Log        string `json:"log,omitempty"`
}

func NewErrorResponse(root error, msg, log string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
	}
}

func NewFullErrorResponse(stt int, root error, msg, log string) *AppError {
	return &AppError{
		StatusCode: stt,
		RootErr:    root,
		Message:    msg,
		Log:        log,
	}
}

func NewUnauthorized(root error, msg, log string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Log:        log,
	}
}

func NewCustomError(root error, msg, log string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error())
	}

	return NewErrorResponse(errors.New(msg), msg, msg)
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func ErrInternal(err error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError, // 500
		RootErr:    err,
		Message:    "Internal server error",
		Log:        err.Error(),
	}
}
