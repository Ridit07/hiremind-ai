package authService

import (
	"context"
	"time"

	"consumer-service/db"
	"consumer-service/errors"
	"consumer-service/model"

	"github.com/google/uuid"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type Service struct {
	redis     RedisClient
	jwtSecret string
}

func NewService(redis RedisClient, jwtSecret string) *Service {
	return &Service{
		redis:     redis,
		jwtSecret: jwtSecret,
	}
}

func Signup(ctx context.Context, req SignupRequest) error {

	err := validateSignUpRequest(req)

	if err != nil {
		return err
	}

	existingUser, err := model.GetUserDetails(ctx, db.ReadConnection(), model.GetUser{
		Email: []string{req.Email},
	})

	if err != nil {
		return err
	}

	if len(existingUser) != 0 {
		return errors.BadRequest.New("user already existsss")
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

func (s *Service) Login(ctx context.Context, req LoginRequest) (resp LoginResponse, err error) {

	user, err := model.GetUserDetails(ctx, db.ReadConnection(), model.GetUser{
		Email: []string{req.Email},
	})

	if err != nil {
		return LoginResponse{}, err
	}

	if len(user) == 0 {
		return LoginResponse{}, errors.BadRequest.New("invalid credentials")
	}

	err = ComparePassword(user[0].PasswordHash, req.Password)
	if err != nil {
		return LoginResponse{}, errors.BadRequest.New("invalid credentials")
	}

	accessToken, err := s.GenerateJWT(user[0].UserID)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken := uuid.NewString()

	err = s.redis.Set(ctx, refreshToken, user[0].UserID, time.Hour*24)

	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (LoginResponse, error) {

	userID, err := s.redis.Get(ctx, refreshToken)
	if err != nil {
		return LoginResponse{}, errors.BadRequest.New("invalid refresh token")
	}

	err = s.redis.Delete(ctx, refreshToken)
	if err != nil {
		return LoginResponse{}, err
	}

	newRefreshToken := uuid.NewString()

	err = s.redis.Set(ctx, newRefreshToken, userID, 24*time.Hour)
	if err != nil {
		return LoginResponse{}, err
	}

	accessToken, err := s.GenerateJWT(userID)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *Service) Logout(ctx context.Context, refreshToken string) error {

	err := s.redis.Delete(ctx, refreshToken)
	if err != nil {
		return errors.Internal.Wrap(err, "failed to logout")
	}

	return nil
}
