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
	CreateInterviewDraft(ctx context.Context, req *CreateInterviewDraftRequest) (*CreateInterviewDraftResponse, error)
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

func (s *InterviewService) CreateInterviewDraft(ctx context.Context, req *CreateInterviewDraftRequest) (*CreateInterviewDraftResponse, error) {
	if req == nil {
		return nil, errors.BadRequest.New("request cannot be empty")
	}

	protoRole := mapRoleToProto(req.Role)
	if protoRole == interviewpb.Role_ROLE_UNSPECIFIED {
		return nil, errors.BadRequest.New("invalid or missing role")
	}

	protoLevel := mapLevelToProto(req.Level)
	if protoLevel == interviewpb.Level_LEVEL_UNSPECIFIED {
		return nil, errors.BadRequest.New("invalid or missing level")
	}

	ctx = addAuthTokenToContext(ctx)

	protoResp, err := s.interviewClient.CreateInterviewDraft(ctx, &interviewpb.CreateInterviewDraftRequest{
		Role:    protoRole,
		Company: strings.TrimSpace(req.Company),
		Level:   protoLevel,
	})
	if err != nil {
		return nil, err
	}

	if protoResp == nil {
		return nil, errors.Internal.New("received nil response from interview service")
	}

	if protoResp.Error != nil {
		return nil, errors.ProtoErrorToAppError(protoResp.Error)
	}

	resp := &CreateInterviewDraftResponse{}
	if protoResp.ExpiresAt != nil {
		resp.ExpiresAt = protoResp.ExpiresAt.AsTime()
	}

	return resp, nil
}
