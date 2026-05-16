package interviewserver

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func addAuthTokenToContext(ctx context.Context) context.Context {
	if token, ok := ctx.Value("token").(string); ok && token != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	}
	return ctx
}
