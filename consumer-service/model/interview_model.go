package model

import (
	errorv2 "consumer-service/errors"
	"context"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Interview struct {
	InterviewID         string          `gorm:"column:interview_id;type:uuid;default:gen_random_uuid();primaryKey"`
	HrID                string          `gorm:"column:hr_id;type:uuid;not null;index"`
	CandidateID         string          `gorm:"column:candidate_id;type:uuid;not null;index"`
	ParticipantID       string          `gorm:"-"`
	InterviewDatetime   time.Time       `gorm:"column:interview_datetime;not null"`
	Status              InterviewStatus `gorm:"column:status;default:'scheduled'"`
	InterviewReportPath string          `gorm:"column:interview_report_path"`
	CreatedAt           time.Time       `gorm:"column:created_at"`
	UpdatedAt           time.Time       `gorm:"column:updated_at"`
}

type InterviewStatus string

const (
	InterviewStatusScheduled   InterviewStatus = "scheduled"
	InterviewStatusOngoing     InterviewStatus = "ongoing"
	InterviewStatusCompleted   InterviewStatus = "completed"
	InterviewStatusCancelled   InterviewStatus = "cancelled"
	InterviewStatusNoShow      InterviewStatus = "no_show"
	InterviewStatusRescheduled InterviewStatus = "rescheduled"
)

func (Interview) TableName() string {
	return "interviews"
}

func CreateInterview(ctx context.Context, db *gorm.DB, interview *Interview) error {
	err := db.WithContext(ctx).Create(interview).Error

	if err != nil {
		return errorv2.Database.Wrap(err, "failed to create interview")
	}

	return nil
}

func GetInterview(ctx context.Context, db *gorm.DB, getInterview Interview) ([]Interview, error) {
	var interviews []Interview

	query := db.WithContext(ctx).Model(&Interview{})

	validQuery := false

	if strings.TrimSpace(getInterview.InterviewID) != "" {
		query = query.Where("interview_id = ?", getInterview.InterviewID)
		validQuery = true
	}

	getInterview.HrID = strings.TrimSpace(getInterview.HrID)
	getInterview.CandidateID = strings.TrimSpace(getInterview.CandidateID)
	getInterview.ParticipantID = strings.TrimSpace(getInterview.ParticipantID)

	if getInterview.ParticipantID != "" {
		query = query.Where(
			"hr_id = ? OR candidate_id = ?",
			getInterview.ParticipantID,
			getInterview.ParticipantID,
		)
		validQuery = true
	} else if getInterview.HrID != "" && getInterview.CandidateID != "" {
		query = query.Where("hr_id = ? AND candidate_id = ?", getInterview.HrID, getInterview.CandidateID)
		validQuery = true
	} else if getInterview.HrID != "" {
		query = query.Where("hr_id = ?", getInterview.HrID)
		validQuery = true
	} else if getInterview.CandidateID != "" {
		query = query.Where("candidate_id = ?", getInterview.CandidateID)
		validQuery = true
	}

	if string(getInterview.Status) != "" {
		query = query.Where("status = ?", getInterview.Status)
		validQuery = true
	}

	if !getInterview.InterviewDatetime.IsZero() {
		query = query.Where("interview_datetime = ?", getInterview.InterviewDatetime)
		validQuery = true
	}

	if !validQuery {
		return nil, errorv2.Database.New(
			"at least one indexed filter is required to get interviews",
		)
	}

	err := query.Order("interview_datetime DESC").Find(&interviews).Error

	if err != nil {
		return nil, errorv2.Database.Wrap(
			err,
			"failed to get interviews",
		)
	}

	return interviews, nil
}
