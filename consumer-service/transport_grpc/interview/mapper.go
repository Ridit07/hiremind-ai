package interview

import (
	"consumer-service/services/interviewService"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated/interview"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapGetInterviewsRequestToService(req *pb.GetInterviewsRequest) interviewService.GetInterviewsRequest {

	if req == nil {
		return interviewService.GetInterviewsRequest{}
	}

	return interviewService.GetInterviewsRequest{
		UserID: req.UserId,
	}
}

func mapInterviewToProto(interview *interviewService.Interview) *pb.Interview {

	if interview == nil {
		return nil
	}

	return &pb.Interview{
		InterviewId:         interview.InterviewID,
		HrId:                interview.HrID,
		HrName:              interview.HrName,
		CandidateId:         interview.CandidateID,
		CandidateName:       interview.CandidateName,
		InterviewDatetime:   timestamppb.New(interview.InterviewDatetime),
		Status:              mapStatusToProto(interview.Status),
		InterviewReportPath: interview.InterviewReportPath,
		CreatedAt:           timestamppb.New(interview.CreatedAt),
		UpdatedAt:           timestamppb.New(interview.UpdatedAt),
	}
}

func mapStatusToProto(status interviewService.InterviewStatus) pb.InterviewStatus {
	switch status {
	case interviewService.InterviewStatusScheduled:
		return pb.InterviewStatus_INTERVIEW_STATUS_SCHEDULED
	case interviewService.InterviewStatusOngoing:
		return pb.InterviewStatus_INTERVIEW_STATUS_ONGOING
	case interviewService.InterviewStatusCompleted:
		return pb.InterviewStatus_INTERVIEW_STATUS_COMPLETED
	case interviewService.InterviewStatusCancelled:
		return pb.InterviewStatus_INTERVIEW_STATUS_CANCELLED
	case interviewService.InterviewStatusNoShow:
		return pb.InterviewStatus_INTERVIEW_STATUS_NO_SHOW
	case interviewService.InterviewStatusRescheduled:
		return pb.InterviewStatus_INTERVIEW_STATUS_RESCHEDULED
	default:
		return pb.InterviewStatus_INTERVIEW_STATUS_UNSPECIFIED
	}
}
