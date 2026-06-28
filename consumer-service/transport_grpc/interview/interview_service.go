package interview

import (
	"context"

	"consumer-service/services/interviewService"

	errormapper "consumer-service/transport_grpc/common"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated/interview"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type InterviewServer struct {
	pb.UnimplementedInterviewServiceServer
	Service *interviewService.Service
}

func (s *InterviewServer) GetInterviews(
	ctx context.Context,
	req *pb.GetInterviewsRequest,
) (*pb.GetInterviewsResponse, error) {

	resp, err := s.Service.GetInterviews(ctx, mapGetInterviewsRequestToService(req))

	if err != nil {
		return &pb.GetInterviewsResponse{
			Error: errormapper.ToProtoError(err),
		}, nil
	}

	interviews := make([]*pb.Interview, len(resp.Interviews))
	for i, interview := range resp.Interviews {
		interviews[i] = mapInterviewToProto(&interview)
	}

	return &pb.GetInterviewsResponse{
		Interviews: interviews,
	}, nil
}

func (s *InterviewServer) CreateInterviewDraft(
	ctx context.Context,
	req *pb.CreateInterviewDraftRequest,
) (*pb.CreateInterviewDraftResponse, error) {

	resp, err := s.Service.CreateInterviewDraft(ctx, mapCreateInterviewDraftRequestToService(req))

	if err != nil {
		return &pb.CreateInterviewDraftResponse{
			Error: errormapper.ToProtoError(err),
		}, nil
	}

	return &pb.CreateInterviewDraftResponse{
		ExpiresAt: timestamppb.New(resp.ExpiresAt),
	}, nil
}
