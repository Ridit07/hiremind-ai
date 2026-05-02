package grpc

import (
	"context"

	service "consumer-service/services/authService"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthServer) Signup(
	ctx context.Context,
	req *pb.SignupRequest,
) (*pb.SignupResponse, error) {

	err := service.Signup(
		ctx,
		req.Email,
		req.Password,
		req.UserType,
		req.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &pb.SignupResponse{
		Message: "signup successful",
	}, nil
}
