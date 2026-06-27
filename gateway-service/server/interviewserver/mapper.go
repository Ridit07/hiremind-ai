package interviewserver

import (
	"strings"

	interviewpb "github.com/Ridit07/hiremind-proto-contracts/generated/interview"
)

func mapInterviewResponse(protoResp *interviewpb.GetInterviewsResponse) []InterviewData {
	if protoResp == nil || protoResp.Interviews == nil {
		return []InterviewData{}
	}

	interviews := make([]InterviewData, 0, len(protoResp.Interviews))
	for idx := range protoResp.Interviews {
		interviews = append(interviews, InterviewData{
			InterviewID:         protoResp.Interviews[idx].InterviewId,
			HrID:                protoResp.Interviews[idx].HrId,
			HrName:              protoResp.Interviews[idx].HrName,
			CandidateID:         protoResp.Interviews[idx].CandidateId,
			CandidateName:       protoResp.Interviews[idx].CandidateName,
			Status:              mapStatusToString(protoResp.Interviews[idx].Status),
			InterviewDatetime:   protoResp.Interviews[idx].InterviewDatetime.AsTime(),
			InterviewReportPath: protoResp.Interviews[idx].InterviewReportPath,
			CreatedAt:           protoResp.Interviews[idx].CreatedAt.AsTime(),
			UpdatedAt:           protoResp.Interviews[idx].UpdatedAt.AsTime(),
		})
	}

	return interviews
}

func mapStatusToString(status interviewpb.InterviewStatus) string {
	switch status {
	case interviewpb.InterviewStatus_INTERVIEW_STATUS_SCHEDULED:
		return "scheduled"
	case interviewpb.InterviewStatus_INTERVIEW_STATUS_ONGOING:
		return "ongoing"
	case interviewpb.InterviewStatus_INTERVIEW_STATUS_COMPLETED:
		return "completed"
	case interviewpb.InterviewStatus_INTERVIEW_STATUS_CANCELLED:
		return "cancelled"
	case interviewpb.InterviewStatus_INTERVIEW_STATUS_NO_SHOW:
		return "no_show"
	case interviewpb.InterviewStatus_INTERVIEW_STATUS_RESCHEDULED:
		return "rescheduled"
	default:
		return "unspecified"
	}
}

func normalizeEnumKey(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	return strings.Join(strings.Fields(s), " ")
}

func mapRoleToProto(role string) interviewpb.Role {
	switch normalizeEnumKey(role) {
	case "backend engineer", "role backend engineer":
		return interviewpb.Role_ROLE_BACKEND_ENGINEER
	case "frontend engineer", "role frontend engineer":
		return interviewpb.Role_ROLE_FRONTEND_ENGINEER
	case "full stack engineer", "fullstack engineer", "role full stack engineer":
		return interviewpb.Role_ROLE_FULL_STACK_ENGINEER
	case "data engineer", "role data engineer":
		return interviewpb.Role_ROLE_DATA_ENGINEER
	default:
		return interviewpb.Role_ROLE_UNSPECIFIED
	}
}

func mapLevelToProto(level string) interviewpb.Level {
	switch normalizeEnumKey(level) {
	case "junior", "level junior":
		return interviewpb.Level_LEVEL_JUNIOR
	case "mid", "level mid":
		return interviewpb.Level_LEVEL_MID
	case "senior", "level senior":
		return interviewpb.Level_LEVEL_SENIOR
	default:
		return interviewpb.Level_LEVEL_UNSPECIFIED
	}
}
