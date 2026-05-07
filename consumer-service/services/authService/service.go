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

func (s *Service) Login(ctx context.Context, req LoginRequest) (resp LoginResponse, err error) {

	user, err := model.GetUserDetails(ctx, db.ReadConnection(), model.User{
		Email: req.Email,
	})
	if err != nil {
		return LoginResponse{}, err
	}

	if user == nil {
		return LoginResponse{}, errors.BadRequest.New("invalid credentials")
	}

	err = ComparePassword(user.PasswordHash, req.Password)
	if err != nil {
		return LoginResponse{}, errors.BadRequest.New("invalid credentials")
	}

	accessToken, err := s.GenerateJWT(user.UserID)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken := uuid.NewString()

	err = s.redis.Set(ctx, refreshToken, user.UserID, time.Hour*24)

	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
