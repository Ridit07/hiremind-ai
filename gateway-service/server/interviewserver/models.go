package interviewserver

import "time"

type GetInterviewsResponse struct {
	Interviews []InterviewData `json:"interviews"`
}

type InterviewData struct {
	InterviewID         string    `json:"interview_id"`
	HrID                string    `json:"hr_id"`
	HrName              string    `json:"hr_name"`
	CandidateID         string    `json:"candidate_id"`
	CandidateName       string    `json:"candidate_name"`
	Status              string    `json:"status"`
	InterviewDatetime   time.Time `json:"interview_datetime"`
	InterviewReportPath string    `json:"interview_report_path"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type CreateInterviewDraftRequest struct {
	Role    string `json:"role"`
	Company string `json:"company"`
	Level   string `json:"level"`
}

type CreateInterviewDraftResponse struct {
	ExpiresAt time.Time `json:"expires_at"`
}
