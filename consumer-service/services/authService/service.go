package authService

import (
	"context"
	"errors"

	"consumer-service/db"
	"consumer-service/model"
)

func Signup(
	ctx context.Context,
	email string,
	password string,
	userType string,
	phone string,
) error {

	existingUser, _ := model.GetUserByEmail(
		ctx,
		email,
	)

	if existingUser != nil {
		return errors.New(
			"user already exists",
		)
	}

	hashedPassword, err := HashPassword(
		password,
	)

	if err != nil {
		return err
	}

	tx, err := db.WriteConnection().Begin(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	err = model.CreateUser(
		ctx,
		tx,
		model.User{
			Email:        email,
			PasswordHash: hashedPassword,
			UserType:     userType,
			PhoneNumber:  phone,
		},
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
