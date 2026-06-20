package interviewService

import (
	"context"

	"consumer-service/db"
	"consumer-service/model"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
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
	//nhjubhjb
	return GetInterviewsResponse{
		Interviews: mapGetUserDetailsRespToGetInterviews(users, interviews),
	}, nil
}
