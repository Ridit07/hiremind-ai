package errors

import "google.golang.org/grpc/codes"

var Database = ErrorBuilder{
	code: codes.Internal,
}
