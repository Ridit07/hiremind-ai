package errors

import (
	"fmt"
	"net/http"

	commonpb "github.com/Ridit07/hiremind-proto-contracts/generated/common"
)

type AppError struct {
	Code      int
	Message   string
	Err       error
	ErrorType string
	ProtoCode int32
	ProtoType string
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

type ErrorBuilder struct {
	httpCode  int
	errorType string
}

func (b ErrorBuilder) New(message string) error {
	return &AppError{
		Code:      b.httpCode,
		ErrorType: b.errorType,
		Message:   message,
	}
}

func (b ErrorBuilder) Newf(format string, args ...any) error {
	return &AppError{
		Code:      b.httpCode,
		ErrorType: b.errorType,
		Message:   fmt.Sprintf(format, args...),
	}
}

func (b ErrorBuilder) Wrap(err error, message string) error {
	return &AppError{
		Code:      b.httpCode,
		ErrorType: b.errorType,
		Message:   message,
		Err:       err,
	}
}

func (b ErrorBuilder) Wrapf(err error, format string, args ...any) error {
	return &AppError{
		Code:      b.httpCode,
		ErrorType: b.errorType,
		Message:   fmt.Sprintf(format, args...),
		Err:       err,
	}
}

var (
	BadRequest = ErrorBuilder{
		httpCode:  http.StatusBadRequest,
		errorType: "BAD_REQUEST",
	}
	Unauthorized = ErrorBuilder{
		httpCode:  http.StatusUnauthorized,
		errorType: "UNAUTHORIZED",
	}
	Forbidden = ErrorBuilder{
		httpCode:  http.StatusForbidden,
		errorType: "FORBIDDEN",
	}
	NotFound = ErrorBuilder{
		httpCode:  http.StatusNotFound,
		errorType: "NOT_FOUND",
	}
	Conflict = ErrorBuilder{
		httpCode:  http.StatusConflict,
		errorType: "CONFLICT",
	}
	Internal = ErrorBuilder{
		httpCode:  http.StatusInternalServerError,
		errorType: "INTERNAL_SERVER",
	}
	Database = ErrorBuilder{
		httpCode:  http.StatusInternalServerError,
		errorType: "DATABASE",
	}
)

func ProtoErrorToAppError(protoErr *commonpb.Error) *AppError {
	if protoErr == nil {
		return &AppError{
			Code:      http.StatusInternalServerError,
			ErrorType: "INTERNAL_SERVER",
			Message:   "unknown error",
		}
	}

	httpCode := protoErrorTypeToHTTPCode(protoErr.Type)
	errorType := protoErrorTypeToString(protoErr.Type)

	return &AppError{
		Code:      httpCode,
		ErrorType: errorType,
		Message:   protoErr.Message,
		ProtoCode: protoErr.Code,
		ProtoType: errorType,
	}
}

func protoErrorTypeToHTTPCode(errType commonpb.ErrorType) int {
	switch errType {
	case commonpb.ErrorType_BAD_REQUEST:
		return http.StatusBadRequest
	case commonpb.ErrorType_UNAUTHORIZED:
		return http.StatusUnauthorized
	case commonpb.ErrorType_FORBIDDEN:
		return http.StatusForbidden
	case commonpb.ErrorType_NOT_FOUND:
		return http.StatusNotFound
	case commonpb.ErrorType_CONFLICT:
		return http.StatusConflict
	case commonpb.ErrorType_DATABASE:
		return http.StatusInternalServerError
	case commonpb.ErrorType_SYSTEM:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func protoErrorTypeToString(errType commonpb.ErrorType) string {
	switch errType {
	case commonpb.ErrorType_BAD_REQUEST:
		return "BAD_REQUEST"
	case commonpb.ErrorType_UNAUTHORIZED:
		return "UNAUTHORIZED"
	case commonpb.ErrorType_FORBIDDEN:
		return "FORBIDDEN"
	case commonpb.ErrorType_NOT_FOUND:
		return "NOT_FOUND"
	case commonpb.ErrorType_CONFLICT:
		return "CONFLICT"
	case commonpb.ErrorType_DATABASE:
		return "DATABASE"
	case commonpb.ErrorType_SYSTEM:
		return "SYSTEM"
	default:
		return "UNKNOWN"
	}
}

func IsAppError(err error) (*AppError, bool) {
	appErr, ok := err.(*AppError)
	return appErr, ok
}
