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

type InterviewType string

const (
	InterviewTypeTechnical    InterviewType = "Technical"
	InterviewTypeBehavioral   InterviewType = "Behavioral"
	InterviewTypeSystemDesign InterviewType = "System Design"
)

type Language string

const (
	LanguagePython     Language = "Python"
	LanguageJava       Language = "Java"
	LanguageCPP        Language = "C++"
	LanguageJavaScript Language = "JavaScript"
)

type DraftStep int32

const (
	DraftStepUnspecified DraftStep = 0
	DraftStepBasic       DraftStep = 1 // role, company, level, type, skills
	DraftStepCoding      DraftStep = 2 // coding_enabled, num_questions, difficulty, topics, language
	DraftStepSchedule    DraftStep = 3 // interview_datetime + candidate email/password/phone
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
	Step DraftStep

	// basic
	Role    Role
	Company string
	Level   Level
	Type    InterviewType
	Skills  []string

	// coding
	CodingEnabled bool
	NumQuestions  int32
	Difficulty    []string
	Topics        []string
	Language      Language

	// schedule + candidate
	InterviewDatetime    *time.Time
	CandidateEmail       string
	CandidatePassword    string
	CandidatePhoneNumber string
}

type CreateInterviewDraftResponse struct {
	ExpiresAt time.Time
}

type InterviewDraft struct {
	HrID string `json:"hr_id"`

	// basic
	Role    Role          `json:"role,omitempty"`
	Company string        `json:"company,omitempty"`
	Level   Level         `json:"level,omitempty"`
	Type    InterviewType `json:"type,omitempty"`
	Skills  []string      `json:"skills,omitempty"`

	// coding
	CodingEnabled bool     `json:"coding_enabled"`
	NumQuestions  int32    `json:"num_questions,omitempty"`
	Difficulty    []string `json:"difficulty,omitempty"`
	Topics        []string `json:"topics,omitempty"`
	Language      Language `json:"language,omitempty"`

	// schedule + candidate
	InterviewDatetime    *time.Time `json:"interview_datetime,omitempty"`
	CandidateEmail       string     `json:"candidate_email,omitempty"`
	CandidatePassword    string     `json:"candidate_password,omitempty"`
	CandidatePhoneNumber string     `json:"candidate_phone_number,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
