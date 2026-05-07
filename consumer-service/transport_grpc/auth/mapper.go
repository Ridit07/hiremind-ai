package auth

import (
	service "consumer-service/services/authService"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated/auth"
)

func mapSignupRequestToService(req *pb.SignupRequest) service.SignupRequest {
	if req == nil {
		return service.SignupRequest{}
	}

	return service.SignupRequest{
		Email:       req.Email,
		Password:    req.Password,
		UserType:    mapUserType(req.UserType),
		PhoneNumber: req.PhoneNumber,
	}
}

func mapUserType(protoUserType pb.UserType) service.UserType {
	switch protoUserType {
	case pb.UserType_USER_TYPE_CANDIDATE:
		return service.UserTypeCandidate
	case pb.UserType_USER_TYPE_HR:
		return service.UserTypeHR
	default:
		return ""
	}
}
