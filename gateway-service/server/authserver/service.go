package server

import (
	"context"
	"strings"

	"gateway-service/errors"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated/auth"
	"google.golang.org/grpc"
)

type GatewayServiceInterface interface {
	Signup(ctx context.Context, req *SignupRequest) (*SignupResponse, error)
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenResponse, error)
	Logout(ctx context.Context, refreshToken string) (*LogoutResponse, error)
}

type GatewayService struct {
	authClient pb.AuthServiceClient
}

func NewGatewayService(authConn *grpc.ClientConn) *GatewayService {
	return &GatewayService{
		authClient: pb.NewAuthServiceClient(authConn),
	}
}

var _ GatewayServiceInterface = (*GatewayService)(nil)

func (s *GatewayService) Signup(ctx context.Context, req *SignupRequest) (*SignupResponse, error) {
	if req == nil {
		return nil, errors.Internal.New("signup request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest.New(err.Error())
	}

	protoResp, err := s.authClient.Signup(ctx, mapSignupRequestToProto(req))
	if err != nil {
		return nil, err
	}

	if protoResp == nil {
		return nil, errors.Internal.New("received nil response from auth service for signup")
	}

	if protoResp.Error != nil {
		appErr := errors.ProtoErrorToAppError(protoResp.Error)
		return nil, appErr
	}

	return &SignupResponse{
		Message: protoResp.Message,
	}, nil
}

func (s *GatewayService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	if req == nil {
		return nil, errors.Internal.New("login request cannot be nil")
	}

	if err := req.Validate(); err != nil {
		return nil, errors.BadRequest.New(err.Error())
	}

	protoResp, err := s.authClient.Login(ctx, mapLoginRequestToProto(req))
	if err != nil {
		return nil, err
	}

	if protoResp == nil {
		return nil, errors.Internal.New("received nil response from auth service")
	}

	if protoResp.Error != nil {
		appErr := errors.ProtoErrorToAppError(protoResp.Error)
		return nil, appErr
	}

	return &LoginResponse{
		AccessToken:  protoResp.AccessToken,
		RefreshToken: protoResp.RefreshToken,
	}, nil
}

func (s *GatewayService) RefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenResponse, error) {
	if strings.TrimSpace(refreshToken) == "" {
		return nil, errors.BadRequest.New("refresh token cannot be empty")
	}

	protoResp, err := s.authClient.RefreshToken(ctx, &pb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return nil, err
	}

	if protoResp == nil {
		return nil, errors.Internal.New("received nil response from auth service")
	}

	if protoResp.Error != nil {
		appErr := errors.ProtoErrorToAppError(protoResp.Error)
		return nil, appErr
	}

	return &RefreshTokenResponse{
		AccessToken:  protoResp.AccessToken,
		RefreshToken: protoResp.RefreshToken,
	}, nil
}

func (s *GatewayService) Logout(ctx context.Context, refreshToken string) (*LogoutResponse, error) {
	if strings.TrimSpace(refreshToken) == "" {
		return nil, errors.BadRequest.New("refresh token cannot be empty to logout")
	}

	protoResp, err := s.authClient.Logout(ctx, &pb.LogoutRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return nil, err
	}

	if protoResp == nil {
		return nil, errors.Internal.New("received nil response from auth service")
	}

	if protoResp.Error != nil {
		appErr := errors.ProtoErrorToAppError(protoResp.Error)
		return nil, appErr
	}

	return &LogoutResponse{
		Message: protoResp.Message,
	}, nil
} //
