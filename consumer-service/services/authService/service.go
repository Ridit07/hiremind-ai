package authService

import (
	"context"
	"time"

	"consumer-service/db"
	"consumer-service/errors"
	"consumer-service/model"
)

func Signup(ctx context.Context, req SignupRequest) error {

	err := validateSignUpRequest(req)

	if err != nil {
		return err
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

	now := time.Now()

	err = model.CreateUser(
		ctx,
		db.WriteConnection(),
		&model.User{
			Email:        req.Email,
			PasswordHash: hashedPassword,
			UserType:     model.UserType(req.UserType),
			PhoneNumber:  req.PhoneNumber,
			UserStatus:   model.UserStatusActive,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
