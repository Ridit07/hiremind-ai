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

func mapCreateInterviewDraftRequestToService(req *pb.CreateInterviewDraftRequest) interviewService.CreateInterviewDraftRequest {

	if req == nil {
		return interviewService.CreateInterviewDraftRequest{}
	}

	out := interviewService.CreateInterviewDraftRequest{
		Step: interviewService.DraftStep(req.Step),

		Role:    mapProtoRoleToService(req.Role),
		Company: req.Company,
		Level:   mapProtoLevelToService(req.Level),
		Type:    mapProtoTypeToService(req.Type),
		Skills:  req.Skills,

		CodingEnabled: req.CodingEnabled,
		NumQuestions:  req.NumQuestions,
		Difficulty:    req.Difficulty,
		Topics:        req.Topics,
		Language:      mapProtoLanguageToService(req.Language),

		CandidateEmail:       req.CandidateEmail,
		CandidatePassword:    req.CandidatePassword,
		CandidatePhoneNumber: req.CandidatePhoneNumber,
	}

	if req.InterviewDatetime != nil {
		t := req.InterviewDatetime.AsTime()
		out.InterviewDatetime = &t
	}

	return out
}

func mapInterviewDraftToProto(draft *interviewService.InterviewDraft) *pb.CreateInterviewDraftRequest {

	if draft == nil {
		return nil
	}

	out := &pb.CreateInterviewDraftRequest{
		Role:    mapServiceRoleToProto(draft.Role),
		Company: draft.Company,
		Level:   mapServiceLevelToProto(draft.Level),
		Type:    mapServiceTypeToProto(draft.Type),
		Skills:  draft.Skills,

		CodingEnabled: draft.CodingEnabled,
		NumQuestions:  draft.NumQuestions,
		Difficulty:    draft.Difficulty,
		Topics:        draft.Topics,
		Language:      mapServiceLanguageToProto(draft.Language),

		CandidateEmail:       draft.CandidateEmail,
		CandidatePassword:    draft.CandidatePassword,
		CandidatePhoneNumber: draft.CandidatePhoneNumber,
	}

	if draft.InterviewDatetime != nil {
		out.InterviewDatetime = timestamppb.New(*draft.InterviewDatetime)
	}

	return out
}

func mapServiceRoleToProto(role interviewService.Role) pb.Role {
	switch role {
	case interviewService.RoleBackendEngineer:
		return pb.Role_ROLE_BACKEND_ENGINEER
	case interviewService.RoleFrontendEngineer:
		return pb.Role_ROLE_FRONTEND_ENGINEER
	case interviewService.RoleFullStackEngineer:
		return pb.Role_ROLE_FULL_STACK_ENGINEER
	case interviewService.RoleDataEngineer:
		return pb.Role_ROLE_DATA_ENGINEER
	default:
		return pb.Role_ROLE_UNSPECIFIED
	}
}

func mapServiceLevelToProto(level interviewService.Level) pb.Level {
	switch level {
	case interviewService.LevelJunior:
		return pb.Level_LEVEL_JUNIOR
	case interviewService.LevelMid:
		return pb.Level_LEVEL_MID
	case interviewService.LevelSenior:
		return pb.Level_LEVEL_SENIOR
	default:
		return pb.Level_LEVEL_UNSPECIFIED
	}
}

func mapServiceTypeToProto(t interviewService.InterviewType) pb.InterviewType {
	switch t {
	case interviewService.InterviewTypeTechnical:
		return pb.InterviewType_INTERVIEW_TYPE_TECHNICAL
	case interviewService.InterviewTypeBehavioral:
		return pb.InterviewType_INTERVIEW_TYPE_BEHAVIORAL
	case interviewService.InterviewTypeSystemDesign:
		return pb.InterviewType_INTERVIEW_TYPE_SYSTEM_DESIGN
	default:
		return pb.InterviewType_INTERVIEW_TYPE_UNSPECIFIED
	}
}

func mapServiceLanguageToProto(l interviewService.Language) pb.Language {
	switch l {
	case interviewService.LanguagePython:
		return pb.Language_LANGUAGE_PYTHON
	case interviewService.LanguageJava:
		return pb.Language_LANGUAGE_JAVA
	case interviewService.LanguageCPP:
		return pb.Language_LANGUAGE_CPP
	case interviewService.LanguageJavaScript:
		return pb.Language_LANGUAGE_JAVASCRIPT
	default:
		return pb.Language_LANGUAGE_UNSPECIFIED
	}
}

func mapProtoTypeToService(t pb.InterviewType) interviewService.InterviewType {
	switch t {
	case pb.InterviewType_INTERVIEW_TYPE_TECHNICAL:
		return interviewService.InterviewTypeTechnical
	case pb.InterviewType_INTERVIEW_TYPE_BEHAVIORAL:
		return interviewService.InterviewTypeBehavioral
	case pb.InterviewType_INTERVIEW_TYPE_SYSTEM_DESIGN:
		return interviewService.InterviewTypeSystemDesign
	default:
		return ""
	}
}

func mapProtoLanguageToService(l pb.Language) interviewService.Language {
	switch l {
	case pb.Language_LANGUAGE_PYTHON:
		return interviewService.LanguagePython
	case pb.Language_LANGUAGE_JAVA:
		return interviewService.LanguageJava
	case pb.Language_LANGUAGE_CPP:
		return interviewService.LanguageCPP
	case pb.Language_LANGUAGE_JAVASCRIPT:
		return interviewService.LanguageJavaScript
	default:
		return ""
	}
}

func mapProtoRoleToService(role pb.Role) interviewService.Role {
	switch role {
	case pb.Role_ROLE_BACKEND_ENGINEER:
		return interviewService.RoleBackendEngineer
	case pb.Role_ROLE_FRONTEND_ENGINEER:
		return interviewService.RoleFrontendEngineer
	case pb.Role_ROLE_FULL_STACK_ENGINEER:
		return interviewService.RoleFullStackEngineer
	case pb.Role_ROLE_DATA_ENGINEER:
		return interviewService.RoleDataEngineer
	default:
		return ""
	}
}

func mapProtoLevelToService(level pb.Level) interviewService.Level {
	switch level {
	case pb.Level_LEVEL_JUNIOR:
		return interviewService.LevelJunior
	case pb.Level_LEVEL_MID:
		return interviewService.LevelMid
	case pb.Level_LEVEL_SENIOR:
		return interviewService.LevelSenior
	default:
		return ""
	}
}
