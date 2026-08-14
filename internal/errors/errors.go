package app_errors

import (
	"errors"

	"google.golang.org/grpc/status"
)

type AppError struct {
	Code    Code
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err == nil {
		return e.Message
	}

	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) GetMessage() string {
	return e.Message
}

func (e *AppError) GetCodeName() string {
	return e.Code.Name
}

func NewError(code Code, message string, err error) error {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func FromError(err error) AppError {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return *appErr
	}

	return AppError{
		Code:    InternalServerErr,
		Message: "internal server error",
		Err:     err,
	}
}

func ToGRPCError(err error) error {
	appErr := FromError(err)

	return status.Error(
		appErr.Code.GRPCCode,
		appErr.Message,
	)
}
