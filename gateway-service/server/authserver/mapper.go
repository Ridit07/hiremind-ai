package server

import (
	pb "github.com/Ridit07/hiremind-proto-contracts/generated/auth"
)

func mapSignupRequestToProto(req *SignupRequest) *pb.SignupRequest {
	if req == nil {
		return &pb.SignupRequest{}
	}

	userType := pb.UserType_USER_TYPE_UNSPECIFIED
	switch req.UserType {
	case string(UserTypeCandidate):
		userType = pb.UserType_USER_TYPE_CANDIDATE
	case string(UserTypeHR):
		userType = pb.UserType_USER_TYPE_HR
	}

	return &pb.SignupRequest{
		Email:       req.Email,
		Password:    req.Password,
		UserType:    userType,
		PhoneNumber: req.PhoneNumber,
	}
}

func mapLoginRequestToProto(req *LoginRequest) *pb.LoginRequest {
	if req == nil {
		return &pb.LoginRequest{}
	}

	return &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}
