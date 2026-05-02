package model

import (
	"context"

	"consumer-service/db"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Email        string
	PasswordHash string
	UserType     string
	PhoneNumber  string
}

func CreateUser(
	ctx context.Context,
	tx pgx.Tx,
	user User,
) error {

	query := `
	INSERT INTO users (
		email,
		password_hash,
		user_type,
		phone_number
	)
	VALUES ($1, $2, $3, $4)
	`

	_, err := tx.Exec(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.UserType,
		user.PhoneNumber,
	)

	return err
}

func GetUserByEmail(
	ctx context.Context,
	email string,
) (*User, error) {

	query := `
	SELECT
		email,
		password_hash,
		user_type,
		phone_number
	FROM users
	WHERE email = $1
	`

	row := db.ReadConnection().QueryRow(
		ctx,
		query,
		email,
	)

	var user User

	err := row.Scan(
		&user.Email,
		&user.PasswordHash,
		&user.UserType,
		&user.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
