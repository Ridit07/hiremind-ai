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

func mapTypeToProto(t string) interviewpb.InterviewType {
	switch normalizeEnumKey(t) {
	case "technical", "interview type technical":
		return interviewpb.InterviewType_INTERVIEW_TYPE_TECHNICAL
	case "behavioral", "interview type behavioral":
		return interviewpb.InterviewType_INTERVIEW_TYPE_BEHAVIORAL
	case "system design", "interview type system design":
		return interviewpb.InterviewType_INTERVIEW_TYPE_SYSTEM_DESIGN
	default:
		return interviewpb.InterviewType_INTERVIEW_TYPE_UNSPECIFIED
	}
}

func mapLanguageToProto(l string) interviewpb.Language {
	switch normalizeEnumKey(l) {
	case "python", "language python":
		return interviewpb.Language_LANGUAGE_PYTHON
	case "java", "language java":
		return interviewpb.Language_LANGUAGE_JAVA
	case "c++", "cpp", "language cpp":
		return interviewpb.Language_LANGUAGE_CPP
	case "javascript", "language javascript":
		return interviewpb.Language_LANGUAGE_JAVASCRIPT
	default:
		return interviewpb.Language_LANGUAGE_UNSPECIFIED
	}
}

func mapDraftToResponse(protoResp *interviewpb.GetInterviewDraftResponse) *InterviewDraftData {
	if protoResp == nil || protoResp.Draft == nil {
		return nil
	}

	d := protoResp.Draft
	draft := &InterviewDraftData{
		Role:    mapRoleToString(d.Role),
		Company: d.Company,
		Level:   mapLevelToString(d.Level),
		Type:    mapTypeToString(d.Type),
		Skills:  d.Skills,

		CodingEnabled: d.CodingEnabled,
		NumQuestions:  d.NumQuestions,
		Difficulty:    d.Difficulty,
		Topics:        d.Topics,
		Language:      mapLanguageToString(d.Language),

		CandidateEmail:       d.CandidateEmail,
		CandidatePhoneNumber: d.CandidatePhoneNumber,
	}

	if d.InterviewDatetime != nil {
		t := d.InterviewDatetime.AsTime()
		draft.InterviewDatetime = &t
	}

	return draft
}

func mapRoleToString(role interviewpb.Role) string {
	switch role {
	case interviewpb.Role_ROLE_BACKEND_ENGINEER:
		return "Backend Engineer"
	case interviewpb.Role_ROLE_FRONTEND_ENGINEER:
		return "Frontend Engineer"
	case interviewpb.Role_ROLE_FULL_STACK_ENGINEER:
		return "Full Stack Engineer"
	case interviewpb.Role_ROLE_DATA_ENGINEER:
		return "Data Engineer"
	default:
		return ""
	}
}

func mapLevelToString(level interviewpb.Level) string {
	switch level {
	case interviewpb.Level_LEVEL_JUNIOR:
		return "Junior"
	case interviewpb.Level_LEVEL_MID:
		return "Mid"
	case interviewpb.Level_LEVEL_SENIOR:
		return "Senior"
	default:
		return ""
	}
}

func mapTypeToString(t interviewpb.InterviewType) string {
	switch t {
	case interviewpb.InterviewType_INTERVIEW_TYPE_TECHNICAL:
		return "Technical"
	case interviewpb.InterviewType_INTERVIEW_TYPE_BEHAVIORAL:
		return "Behavioral"
	case interviewpb.InterviewType_INTERVIEW_TYPE_SYSTEM_DESIGN:
		return "System Design"
	default:
		return ""
	}
}

func mapLanguageToString(l interviewpb.Language) string {
	switch l {
	case interviewpb.Language_LANGUAGE_PYTHON:
		return "Python"
	case interviewpb.Language_LANGUAGE_JAVA:
		return "Java"
	case interviewpb.Language_LANGUAGE_CPP:
		return "C++"
	case interviewpb.Language_LANGUAGE_JAVASCRIPT:
		return "JavaScript"
	default:
		return ""
	}
}
