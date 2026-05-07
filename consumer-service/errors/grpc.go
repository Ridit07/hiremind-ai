package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

type ErrorBuilder struct {
	code codes.Code
}

func (b ErrorBuilder) New(
	message string,
) error {

	return &AppError{
		Code:    b.code,
		Message: message,
	}
}

func (b ErrorBuilder) Newf(
	format string,
	args ...any,
) error {

	return &AppError{
		Code: b.code,
		Message: fmt.Sprintf(
			format,
			args...,
		),
	}
}

func (b ErrorBuilder) Wrap(
	err error,
	message string,
) error {

	return &AppError{
		Code:    b.code,
		Message: message,
		Err:     err,
	}
}

func (b ErrorBuilder) Wrapf(
	err error,
	format string,
	args ...any,
) error {

	return &AppError{
		Code: b.code,
		Message: fmt.Sprintf(
			format,
			args...,
		),
		Err: err,
	}
}
