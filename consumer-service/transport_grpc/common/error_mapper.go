package auth

import (
	"consumer-service/errors"

	commonpb "github.com/Ridit07/hiremind-proto-contracts/generated/common"
	"google.golang.org/grpc/codes"
)

func ToProtoError(err error) *commonpb.Error {

	if err == nil {
		return nil
	}

	appErr, ok := err.(*errors.AppError)

	// fallback for unknown errors
	if !ok {
		return &commonpb.Error{
			Code:    500,
			Type:    commonpb.ErrorType_SYSTEM,
			Message: err.Error(),
		}
	}

	return &commonpb.Error{
		Code:    mapCode(appErr.Code),
		Type:    mapType(appErr.Code),
		Message: appErr.Message,
	}
}

func mapCode(code codes.Code) int32 {

	switch code {

	case codes.InvalidArgument:
		return 400

	case codes.NotFound:
		return 404

	case codes.Unauthenticated:
		return 401

	case codes.PermissionDenied:
		return 403

	case codes.AlreadyExists:
		return 409

	default:
		return 500
	}
}

func mapType(code codes.Code) commonpb.ErrorType {

	switch code {

	case codes.InvalidArgument:
		return commonpb.ErrorType_BAD_REQUEST

	case codes.NotFound:
		return commonpb.ErrorType_NOT_FOUND

	case codes.Unauthenticated:
		return commonpb.ErrorType_UNAUTHORIZED

	case codes.PermissionDenied:
		return commonpb.ErrorType_FORBIDDEN

	case codes.AlreadyExists:
		return commonpb.ErrorType_CONFLICT

	case codes.Internal:
		return commonpb.ErrorType_DATABASE

	default:
		return commonpb.ErrorType_SYSTEM
	}
}
