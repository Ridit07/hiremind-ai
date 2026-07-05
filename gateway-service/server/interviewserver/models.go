package interviewserver

import "time"

const (
	draftStepBasic    int32 = 1 // role, company, level, type, skills
	draftStepCoding   int32 = 2 // coding_enabled, num_questions, difficulty, topics, language
	draftStepSchedule int32 = 3 // interview_datetime + candidate email/password/phone
)

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
	Step int32 `json:"step"` // 1=basic, 2=coding, 3=schedule

	// basic
	Role    string   `json:"role"`
	Company string   `json:"company"`
	Level   string   `json:"level"`
	Type    string   `json:"type"`
	Skills  []string `json:"skills"`

	// coding
	CodingEnabled bool     `json:"coding_enabled"`
	NumQuestions  int32    `json:"num_questions"`
	Difficulty    []string `json:"difficulty"`
	Topics        []string `json:"topics"`
	Language      string   `json:"language"`

	// schedule + candidate
	InterviewDatetime    *time.Time `json:"interview_datetime,omitempty"`
	CandidateEmail       string     `json:"candidate_email"`
	CandidatePassword    string     `json:"candidate_password"`
	CandidatePhoneNumber string     `json:"candidate_phone_number"`
}

type CreateInterviewDraftResponse struct {
	ExpiresAt time.Time `json:"expires_at"`
}

type GetInterviewDraftResponse struct {
	Found     bool                `json:"found"`
	Draft     *InterviewDraftData `json:"draft,omitempty"`
	UpdatedAt *time.Time          `json:"updated_at,omitempty"`
}

type InterviewDraftData struct {
	Role    string   `json:"role"`
	Company string   `json:"company"`
	Level   string   `json:"level"`
	Type    string   `json:"type"`
	Skills  []string `json:"skills"`

	CodingEnabled bool     `json:"coding_enabled"`
	NumQuestions  int32    `json:"num_questions"`
	Difficulty    []string `json:"difficulty"`
	Topics        []string `json:"topics"`
	Language      string   `json:"language"`

	InterviewDatetime    *time.Time `json:"interview_datetime,omitempty"`
	CandidateEmail       string     `json:"candidate_email"`
	CandidatePhoneNumber string     `json:"candidate_phone_number"`
}
