package interviewService

import (
	"consumer-service/common"
	errorv2 "consumer-service/errors"
	"consumer-service/model"
	"consumer-service/services/authService"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func getValidateGetInterviewsRequest(ctx context.Context, req GetInterviewsRequest) (GetInterviewsRequest, error) {

	req.UserID = strings.TrimSpace(req.UserID)

	authenticatedUserID, ok := authService.GetAuthenticatedUserID(ctx)
	if !ok {
		return GetInterviewsRequest{}, errorv2.BadRequest.New("authenticated user_id not found in context")
	}

	if req.UserID != "" && req.UserID != authenticatedUserID {
		return GetInterviewsRequest{}, errorv2.BadRequest.New("user_id does not match authenticated user")
	}
	req.UserID = authenticatedUserID // user from request to be fixed

	if err := common.ValidateUUID(req.UserID); err != nil {
		return GetInterviewsRequest{}, errorv2.BadRequest.Wrap(err, "invalid user_id format")
	}

	return req, nil
}

func makeGetUserDetailsRequest(interviews []model.Interview) []string {

	if len(interviews) == 0 {
		return nil
	}

	userIDMap := make(map[string]bool)
	for idx := range interviews {
		userIDMap[interviews[idx].HrID] = true
		userIDMap[interviews[idx].CandidateID] = true
	}

	var userIDs []string
	for id := range userIDMap {
		userIDs = append(userIDs, id)
	}

	return userIDs
}

func mapGetUserDetailsRespToGetInterviews(users []model.User, interviews []model.Interview) []Interview {

	if len(users) == 0 || len(interviews) == 0 {
		return nil
	}

	userMap := make(map[string]model.User)
	for idx := range users {
		userMap[users[idx].UserID] = users[idx]
	}

	var enrichedInterviews []Interview
	for idx := range interviews {
		hrUser := userMap[interviews[idx].HrID]
		candUser := userMap[interviews[idx].CandidateID]

		enrichedInterviews = append(enrichedInterviews, Interview{
			InterviewID:         interviews[idx].InterviewID,
			HrID:                interviews[idx].HrID,
			HrName:              hrUser.Email,
			CandidateID:         interviews[idx].CandidateID,
			CandidateName:       candUser.Email,
			InterviewDatetime:   interviews[idx].InterviewDatetime,
			Status:              InterviewStatus(interviews[idx].Status),
			InterviewReportPath: interviews[idx].InterviewReportPath,
			CreatedAt:           interviews[idx].CreatedAt,
			UpdatedAt:           interviews[idx].UpdatedAt,
		})
	}
	return enrichedInterviews
}

func validateCreateInterviewDraftRequest(ctx context.Context, req *CreateInterviewDraftRequest) (string, error) {

	hrID, ok := authService.GetAuthenticatedUserID(ctx)
	if !ok {
		return "", errorv2.BadRequest.New("authenticated user_id not found in context")
	}

	if err := common.ValidateUUID(hrID); err != nil {
		return "", errorv2.BadRequest.Wrap(err, "invalid hr_id format")
	}

	req.Company = strings.TrimSpace(req.Company)

	if !req.Role.IsValid() {
		return "", errorv2.BadRequest.New("invalid or missing role")
	}

	if !req.Level.IsValid() {
		return "", errorv2.BadRequest.New("invalid or missing level")
	}

	return hrID, nil
}

func draftKey(hrID string) string {
	return "interview:draft:" + hrID
}

func (r Role) IsValid() bool {
	switch r {
	case RoleBackendEngineer, RoleFrontendEngineer, RoleFullStackEngineer, RoleDataEngineer:
		return true
	}
	return false
}

func (l Level) IsValid() bool {
	switch l {
	case LevelJunior, LevelMid, LevelSenior:
		return true
	}
	return false
}

func (s *Service) getDraftForCreateInterviewDraft(ctx context.Context, hrID string) (InterviewDraft, bool, error) {
	data, err := s.redis.Get(ctx, draftKey(hrID))
	if errors.Is(err, redis.Nil) {
		return InterviewDraft{}, false, nil
	}
	if err != nil {
		return InterviewDraft{}, false, errorv2.Internal.Wrap(err, "failed to get interview draft")
	}

	var draft InterviewDraft
	if err := json.Unmarshal([]byte(data), &draft); err != nil {
		return InterviewDraft{}, false, errorv2.Internal.Wrap(err, "failed to unmarshal interview draft")
	}

	return draft, true, nil
}

func (s *Service) saveDraftForCreateInterviewDraft(ctx context.Context, draft InterviewDraft, draftTTL time.Duration) error {
	data, err := json.Marshal(draft)
	if err != nil {
		return errorv2.Internal.Wrap(err, "failed to marshal interview draft")
	}

	if err := s.redis.Set(ctx, draftKey(draft.HrID), data, draftTTL); err != nil {
		return errorv2.Internal.Wrap(err, "failed to save interview draft")
	}

	return nil
}
