package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

type AppError struct {
	Code    codes.Code
	Message string
	Err     error
}

func (e *AppError) Error() string {

	if e.Err != nil {
		return fmt.Sprintf(
			"%s: %v",
			e.Message,
			e.Err,
		)
	}

	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}
