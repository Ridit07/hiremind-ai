package authService

import (
	"context"
	"strings"
	"time"

	"consumer-service/db"
	"consumer-service/errors"
	"consumer-service/model"
)

func Signup(ctx context.Context, req SignupRequest) error {

	if strings.TrimSpace(req.Email) == "" {
		return errors.BadRequest.New("email is required")
	}

	if req.UserType == "" {
		return errors.BadRequest.New("user_type is required")
	}

	existingUser, err := model.GetUserDetails(ctx, db.ReadConnection(), model.User{
		Email: req.Email,
	})

	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.BadRequest.New("user already exists")
	}

	hashedPassword, err := HashPassword(req.Password)

	if err != nil {
		return err
	}

	err = model.CreateUser(
		ctx,
		db.WriteConnection(),
		&model.User{
			Email:        req.Email,
			PasswordHash: hashedPassword,
			UserType:     model.UserType(req.UserType),
			PhoneNumber:  req.PhoneNumber,
			UserStatus:   model.UserStatusActive,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
