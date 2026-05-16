package model

import (
	errorv2 "consumer-service/errors"
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID       string     `gorm:"column:user_id;type:uuid;default:gen_random_uuid();primaryKey"`
	Email        string     `gorm:"column:email; uniqueIndex:idx_users_email;not null"`
	PasswordHash string     `gorm:"column:password_hash"`
	UserType     UserType   `gorm:"column:user_type"`
	PhoneNumber  string     `gorm:"column:phone_number"`
	UserStatus   UserStatus `gorm:"column:user_status"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at"`
}

type GetUser struct {
	UserID      []string
	Email       []string
	UserType    UserType
	PhoneNumber string
	UserStatus  UserStatus
}

type UserType string

const (
	UserTypeCandidate UserType = "candidate"
	UserTypeHR        UserType = "hr"
)

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusNotActive UserStatus = "not_active"
)

func (User) TableName() string {
	return "users"
}

func CreateUser(ctx context.Context, db *gorm.DB, user *User) error {

	err := db.WithContext(ctx).Create(user).Error

	if err != nil {
		return errorv2.Database.Wrap(err, "failed to create user")
	}

	return nil
}

func GetUserDetails(ctx context.Context, db *gorm.DB, getUsers GetUser) ([]User, error) {

	var user []User

	query := db.WithContext(ctx).Model(&User{})

	validQuery := false

	if len(getUsers.UserID) > 0 {

		query = query.Where("user_id IN ?", getUsers.UserID)

		validQuery = true
	}

	if len(getUsers.Email) > 0 {

		query = query.Where("email IN ?", getUsers.Email)

		validQuery = true
	}

	if strings.TrimSpace(getUsers.PhoneNumber) != "" {
		query = query.Where("phone_number = ?", getUsers.PhoneNumber)
	}

	if strings.TrimSpace(string(getUsers.UserType)) != "" {
		query = query.Where("user_type = ?", getUsers.UserType)
	}

	if !validQuery {
		return nil, errorv2.Database.New(
			"at least one indexed filter is required to get users",
		)
	}

	err := query.Find(&user).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errorv2.Database.Wrap(
			err,
			"failed to get user",
		)
	}

	return user, nil
}
