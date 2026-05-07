package auth

import (
	"context"

	"consumer-service/services/authService"
	service "consumer-service/services/authService"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated/auth"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	Service *authService.Service
}

func (s *AuthServer) Signup(
	ctx context.Context,
	req *pb.SignupRequest,
) (*pb.SignupResponse, error) {

	err := service.Signup(ctx, mapSignupRequestToService(req))

	if err != nil {
		return &pb.SignupResponse{
			Error: ToProtoError(err),
		}, nil
	}

	return &pb.SignupResponse{
		Message: "signup successful",
	}, nil
}
