package interviewService

import (
	"consumer-service/common"
	errorv2 "consumer-service/errors"
	"consumer-service/model"
	"consumer-service/services/authService"
	"context"
	"strings"
)

func getValidateGetInterviewsRequest(ctx context.Context, req GetInterviewsRequest) (GetInterviewsRequest, error) {

	req.UserID = strings.TrimSpace(req.UserID)

	if authenticatedUserID, ok := authService.GetAuthenticatedUserID(ctx); ok {
		if req.UserID != "" && req.UserID != authenticatedUserID {
			return GetInterviewsRequest{}, errorv2.BadRequest.New("user_id does not match authenticated user")
		}
		req.UserID = authenticatedUserID
	}

	if err := common.ValidateUUID(req.UserID); err != nil {
		return GetInterviewsRequest{}, errorv2.BadRequest.Wrap(err, "invalid user_id")
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
