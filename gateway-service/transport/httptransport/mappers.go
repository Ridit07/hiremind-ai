package http

import (
	"gateway-service/server"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated/auth"
	commonpb "github.com/Ridit07/hiremind-proto-contracts/generated/common"
)

// ===== Request Mappers =====

// mapSignupRequestToProto maps signup request to proto
func mapSignupRequestToProto(req *server.SignupRequest) *pb.SignupRequest {
	if req == nil {
		return &pb.SignupRequest{}
	}

	return &pb.SignupRequest{
		Email:       req.Email,
		Password:    req.Password,
		UserType:    mapUserTypeToProto(req.UserType),
		PhoneNumber: req.PhoneNumber,
	}
}

// mapLoginRequestToProto maps login request to proto
func mapLoginRequestToProto(req *server.LoginRequest) *pb.LoginRequest {
	if req == nil {
		return &pb.LoginRequest{}
	}

	return &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

// mapUserTypeToProto maps user type string to proto enum
func mapUserTypeToProto(userType string) pb.UserType {
	switch userType {
	case "CANDIDATE":
		return pb.UserType_USER_TYPE_CANDIDATE
	case "HR":
		return pb.UserType_USER_TYPE_HR
	default:
		return pb.UserType_USER_TYPE_UNSPECIFIED
	}
}

// ===== Response Mappers =====

// mapProtoSignupResponseToService maps proto signup response to service response
func mapProtoSignupResponseToService(protoResp *pb.SignupResponse) *server.SignupResponse {
	if protoResp == nil {
		return &server.SignupResponse{}
	}

	resp := &server.SignupResponse{
		Message: protoResp.Message,
	}

	if protoResp.Error != nil {
		resp.Error = mapProtoErrorToService(protoResp.Error)
	}

	return resp
}

// mapProtoLoginResponseToService maps proto login response to service response
func mapProtoLoginResponseToService(protoResp *pb.LoginResponse) *server.LoginResponse {
	if protoResp == nil {
		return &server.LoginResponse{}
	}

	resp := &server.LoginResponse{
		AccessToken:  protoResp.AccessToken,
		RefreshToken: protoResp.RefreshToken,
	}

	if protoResp.Error != nil {
		resp.Error = mapProtoErrorToService(protoResp.Error)
	}

	return resp
}

// mapProtoRefreshTokenResponseToService maps proto refresh token response to service response
func mapProtoRefreshTokenResponseToService(protoResp *pb.RefreshTokenResponse) *server.RefreshTokenResponse {
	if protoResp == nil {
		return &server.RefreshTokenResponse{}
	}

	resp := &server.RefreshTokenResponse{
		AccessToken:  protoResp.AccessToken,
		RefreshToken: protoResp.RefreshToken,
	}

	if protoResp.Error != nil {
		resp.Error = mapProtoErrorToService(protoResp.Error)
	}

	return resp
}

// mapProtoLogoutResponseToService maps proto logout response to service response
func mapProtoLogoutResponseToService(protoResp *pb.LogoutResponse) *server.LogoutResponse {
	if protoResp == nil {
		return &server.LogoutResponse{}
	}

	resp := &server.LogoutResponse{
		Message: protoResp.Message,
	}

	if protoResp.Error != nil {
		resp.Error = mapProtoErrorToService(protoResp.Error)
	}

	return resp
}

// mapProtoErrorToService maps proto error to service error
func mapProtoErrorToService(protoErr *commonpb.Error) *server.Error {
	if protoErr == nil {
		return nil
	}

	return &server.Error{
		Code:    protoErr.Code,
		Type:    mapProtoErrorTypeToString(protoErr.Type),
		Message: protoErr.Message,
	}
}

// mapProtoErrorTypeToString maps proto error type enum to string
func mapProtoErrorTypeToString(errType commonpb.ErrorType) string {
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
