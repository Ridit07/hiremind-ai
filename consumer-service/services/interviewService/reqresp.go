package interviewService

import "time"

type InterviewStatus string

const (
	InterviewStatusScheduled   InterviewStatus = "scheduled"
	InterviewStatusOngoing     InterviewStatus = "ongoing"
	InterviewStatusCompleted   InterviewStatus = "completed"
	InterviewStatusCancelled   InterviewStatus = "cancelled"
	InterviewStatusNoShow      InterviewStatus = "no_show"
	InterviewStatusRescheduled InterviewStatus = "rescheduled"
)

type Role string

const (
	RoleBackendEngineer   Role = "Backend Engineer"
	RoleFrontendEngineer  Role = "Frontend Engineer"
	RoleFullStackEngineer Role = "Full Stack Engineer"
	RoleDataEngineer      Role = "Data Engineer"
)

type Level string

const (
	LevelJunior Level = "Junior"
	LevelMid    Level = "Mid"
	LevelSenior Level = "Senior"
)

const draftTTL = 24 * time.Hour

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

type CreateInterviewDraftRequest struct {
	Role    Role
	Company string
	Level   Level
}

type CreateInterviewDraftResponse struct {
	ExpiresAt time.Time
}

type InterviewDraft struct {
	HrID      string    `json:"hr_id"`
	Role      Role      `json:"role,omitempty"`
	Company   string    `json:"company,omitempty"`
	Level     Level     `json:"level,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
