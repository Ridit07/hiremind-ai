package interviewserver

import (
	"context"
	"gateway-service/errors"
	"strings"

	interviewpb "github.com/Ridit07/hiremind-proto-contracts/generated/interview"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//  type authTokenContextKey struct{}
//  var authTokenKey = authTokenContextKey{}

func addAuthTokenToContext(ctx context.Context) context.Context {
	//if token, ok := ctx.Value(authTokenKey).(string); ok && token != "" {
	if token, ok := ctx.Value("token").(string); ok && token != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)
	}
	return ctx
}

func buildCreateDraftProtoRequest(req *CreateInterviewDraftRequest) (*interviewpb.CreateInterviewDraftRequest, error) {
	out := &interviewpb.CreateInterviewDraftRequest{
		Step: interviewpb.DraftStep(req.Step),
	}

	switch req.Step {
	case draftStepBasic:
		protoRole := mapRoleToProto(req.Role)
		if protoRole == interviewpb.Role_ROLE_UNSPECIFIED {
			return nil, errors.BadRequest.New("invalid or missing role")
		}
		protoLevel := mapLevelToProto(req.Level)
		if protoLevel == interviewpb.Level_LEVEL_UNSPECIFIED {
			return nil, errors.BadRequest.New("invalid or missing level")
		}
		protoType := mapTypeToProto(req.Type)
		if protoType == interviewpb.InterviewType_INTERVIEW_TYPE_UNSPECIFIED {
			return nil, errors.BadRequest.New("invalid or missing type")
		}

		out.Role = protoRole
		out.Level = protoLevel
		out.Type = protoType
		out.Company = strings.TrimSpace(req.Company)
		out.Skills = req.Skills

	case draftStepCoding:
		out.CodingEnabled = req.CodingEnabled
		if req.CodingEnabled {
			if req.NumQuestions <= 0 {
				return nil, errors.BadRequest.New("num_questions must be greater than 0 when coding is enabled")
			}
			protoLanguage := mapLanguageToProto(req.Language)
			if protoLanguage == interviewpb.Language_LANGUAGE_UNSPECIFIED {
				return nil, errors.BadRequest.New("invalid or missing language")
			}

			out.NumQuestions = req.NumQuestions
			out.Difficulty = req.Difficulty
			out.Topics = req.Topics
			out.Language = protoLanguage
		}

	case draftStepSchedule:
		if req.InterviewDatetime == nil || req.InterviewDatetime.IsZero() {
			return nil, errors.BadRequest.New("invalid or missing interview_datetime")
		}
		email := strings.TrimSpace(req.CandidateEmail)
		if email == "" {
			return nil, errors.BadRequest.New("missing candidate email")
		}
		if strings.TrimSpace(req.CandidatePassword) == "" {
			return nil, errors.BadRequest.New("missing candidate password")
		}

		out.InterviewDatetime = timestamppb.New(*req.InterviewDatetime)
		out.CandidateEmail = email
		out.CandidatePassword = req.CandidatePassword
		out.CandidatePhoneNumber = strings.TrimSpace(req.CandidatePhoneNumber)

	default:
		return nil, errors.BadRequest.New("invalid or missing step")
	}

	return out, nil
}
