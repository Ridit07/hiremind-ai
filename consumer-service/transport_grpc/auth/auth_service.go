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

func (s *AuthServer) Login(
	ctx context.Context,
	req *pb.LoginRequest,
) (*pb.LoginResponse, error) {

	resp, err := s.Service.Login(ctx, mapLoginRequestToService(req))

	if err != nil {
		return &pb.LoginResponse{
			Error: ToProtoError(err),
		}, nil
	}

	return &pb.LoginResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}

func (s *AuthServer) RefreshToken(
	ctx context.Context,
	req *pb.RefreshTokenRequest,
) (*pb.RefreshTokenResponse, error) {

	resp, err := s.Service.RefreshToken(ctx, req.RefreshToken)

	if err != nil {
		return &pb.RefreshTokenResponse{
			Error: ToProtoError(err),
		}, nil
	}

	return &pb.RefreshTokenResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}

func (s *AuthServer) Logout(
	ctx context.Context,
	req *pb.LogoutRequest,
) (*pb.LogoutResponse, error) {

	err := s.Service.Logout(ctx, req.RefreshToken)

	if err != nil {
		return &pb.LogoutResponse{
			Error: ToProtoError(err),
		}, nil
	}

	return &pb.LogoutResponse{
		Message: "logged out successfully",
	}, nil
}
