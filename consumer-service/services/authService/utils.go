package authService

import (
	"consumer-service/errors"
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(hashedBytes), err
}

func ComparePassword(
	hashed string,
	plain string,
) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(plain),
	)
}

func validateSignUpRequest(req SignupRequest) error {

	if strings.TrimSpace(req.Email) == "" {
		return errors.BadRequest.New("email is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.BadRequest.New("password is required")
	}

	if req.UserType == "" {
		return errors.BadRequest.New("user_type is required")
	}

	if strings.TrimSpace(req.PhoneNumber) == "" {
		return errors.BadRequest.New("phone_number is required")
	}

	return nil
}

func (s *Service) getAccessAndRefreshToken(ctx context.Context, userID string) (getAccessAndRefreshTokenResp, error) {

	accessToken, err := s.GenerateJWT(userID)
	if err != nil {
		return getAccessAndRefreshTokenResp{}, err
	}

	refreshToken := uuid.NewString()

	err = s.redis.Set(ctx, refreshToken, userID, time.Hour*24)

	if err != nil {
		return getAccessAndRefreshTokenResp{}, err
	}

	return getAccessAndRefreshTokenResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
