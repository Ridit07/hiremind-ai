package interviewService

import (
	"context"
	"log"
	"time"

	"consumer-service/db"
	"consumer-service/model"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type Service struct {
	redis RedisClient
}

func NewService(redis RedisClient) *Service {
	return &Service{redis: redis}
}

func (s *Service) GetInterviews(ctx context.Context, req GetInterviewsRequest) (resp GetInterviewsResponse, err error) {

	req, err = getValidateGetInterviewsRequest(ctx, req)
	if err != nil {
		return GetInterviewsResponse{}, err
	}

	interviews, err := model.GetInterview(ctx, db.ReadConnection(), model.Interview{
		ParticipantID: req.UserID,
	})

	if err != nil {
		return GetInterviewsResponse{}, err
	}

	userIDs := makeGetUserDetailsRequest(interviews)

	if len(userIDs) == 0 {
		return GetInterviewsResponse{}, nil
	}

	users, err := model.GetUserDetails(ctx, db.ReadConnection(), model.GetUser{
		UserID: userIDs,
	})
	if err != nil {
		return GetInterviewsResponse{}, err
	}

	return GetInterviewsResponse{
		Interviews: mapGetUserDetailsRespToGetInterviews(users, interviews),
	}, nil
}

func (s *Service) CreateInterviewDraft(ctx context.Context, req CreateInterviewDraftRequest) (CreateInterviewDraftResponse, error) {

	hrID, err := validateCreateInterviewDraftRequest(ctx, &req)
	if err != nil {
		return CreateInterviewDraftResponse{}, err
	}

	now := time.Now()

	draft, found, err := s.getDraftForCreateInterviewDraft(ctx, hrID)
	if err != nil {
		// to raise notice error
		log.Printf("Error retrieving draft for HR ID %s: %v", hrID, err)
	}

	if !found {
		draft = InterviewDraft{
			HrID:      hrID,
			CreatedAt: now,
		}
	}

	switch req.Step {
	case DraftStepBasic:
		draft.Role = req.Role
		draft.Company = req.Company
		draft.Level = req.Level
		draft.Type = req.Type
		draft.Skills = req.Skills

	case DraftStepCoding:
		draft.CodingEnabled = req.CodingEnabled
		if !req.CodingEnabled {
			draft.NumQuestions = 0
			draft.Difficulty = nil
			draft.Topics = nil
			draft.Language = ""
			break
		}

		draft.NumQuestions = req.NumQuestions
		draft.Difficulty = req.Difficulty
		draft.Topics = req.Topics
		draft.Language = req.Language

	case DraftStepSchedule:
		draft.InterviewDatetime = req.InterviewDatetime
		draft.CandidateEmail = req.CandidateEmail
		draft.CandidatePassword = req.CandidatePassword
		draft.CandidatePhoneNumber = req.CandidatePhoneNumber
	}

	draft.UpdatedAt = now

	if err := s.saveDraftForCreateInterviewDraft(ctx, draft, draftTTL); err != nil {
		return CreateInterviewDraftResponse{}, err
	}

	return CreateInterviewDraftResponse{
		ExpiresAt: now.Add(draftTTL),
	}, nil
}
