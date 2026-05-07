package errors

import "google.golang.org/grpc/codes"

var Internal = ErrorBuilder{
	code: codes.Internal,
}
