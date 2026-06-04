package interviewserver

import (
	"context"
	"strings"

	"gateway-service/errors"

	interviewpb "github.com/Ridit07/hiremind-proto-contracts/generated/interview"
	"google.golang.org/grpc"
)

type InterviewServiceInterface interface {
	GetInterviews(ctx context.Context, userID string) (*GetInterviewsResponse, error)
}

type InterviewService struct {
	interviewClient interviewpb.InterviewServiceClient
}

func NewInterviewService(interviewConn *grpc.ClientConn) *InterviewService {
	return &InterviewService{
		interviewClient: interviewpb.NewInterviewServiceClient(interviewConn),
	}
}

var _ InterviewServiceInterface = (*InterviewService)(nil)

func (s *InterviewService) GetInterviews(ctx context.Context, userID string) (*GetInterviewsResponse, error) {
	if strings.TrimSpace(userID) == "" {
		return nil, errors.BadRequest.New("user_id cannot be empty")
	}

	ctx = addAuthTokenToContext(ctx)

	protoResp, err := s.interviewClient.GetInterviews(ctx, &interviewpb.GetInterviewsRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	if protoResp == nil {
		return nil, errors.Internal.New("received nil response from interview service")
	}

	if protoResp.Error != nil {
		appErr := errors.ProtoErrorToAppError(protoResp.Error)
		return nil, appErr
	}

	return &GetInterviewsResponse{
		Interviews: mapInterviewResponse(protoResp),
	}, nil
}
