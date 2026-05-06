package errors

import "google.golang.org/grpc/codes"

var BadRequest = ErrorBuilder{
	code: codes.InvalidArgument,
}
