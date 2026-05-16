package interviewService

import "time"

type InterviewStatus string

const (
	InterviewStatusOngoing     InterviewStatus = "ongoing"
	InterviewStatusCompleted   InterviewStatus = "completed"
	InterviewStatusCancelled   InterviewStatus = "cancelled"
	InterviewStatusScheduled   InterviewStatus = "scheduled"
	InterviewStatusNoShow      InterviewStatus = "no_show"
	InterviewStatusRescheduled InterviewStatus = "rescheduled"
)

type GetInterviewsRequest struct {
	UserID string
}

type GetInterviewsResponse struct {
	Interviews []Interview
}

type Interview struct {
	InterviewID         string
	HrID                string
	HrName              string
	CandidateID         string
	CandidateName       string
	InterviewDatetime   time.Time
	Status              InterviewStatus
	InterviewReportPath string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
