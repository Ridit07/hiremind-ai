package authService

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(svc *Service) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		// skip auth for login/signup
		if info.FullMethod == "/auth.AuthService/Login" ||
			info.FullMethod == "/auth.AuthService/Signup" ||
			info.FullMethod == "/auth.AuthService/RefreshToken" ||
			info.FullMethod == "/auth.AuthService/Logout" {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		authHeader := md["authorization"]
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing token")
		}

		token := authHeader[0]

		token = strings.TrimPrefix(token, "Bearer ")

		userID, err := svc.ValidateJWT(token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		ctx = context.WithValue(ctx, userIDKey, userID)

		return handler(ctx, req)
	}
}

func GetAuthenticatedUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	if !ok {
		return "", false
	}

	userID = strings.TrimSpace(userID)
	if userID == "" {
		return "", false
	}

	return userID, true
}

func RateLimitInterceptor(redis *redis.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (interface{}, error) {

		key := "rate:" + getClientIP(ctx)

		count, err := redis.Incr(ctx, key).Result()
		if err != nil {
			return nil, status.Error(codes.Internal, "rate limit error")
		}

		if count == 1 {
			redis.Expire(ctx, key, time.Minute)
		}

		if count > 100 {
			return nil, status.Error(codes.ResourceExhausted, "rate limit exceeded")
		}

		return handler(ctx, req)
	}
}

func LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (interface{}, error) {

		start := time.Now()

		resp, err := handler(ctx, req)

		log.Printf("method=%s duration=%s error=%v",
			info.FullMethod,
			time.Since(start),
			err,
		)

		return resp, err
	}
}

func getClientIP(ctx context.Context) string {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if ip := md.Get("x-forwarded-for"); len(ip) > 0 {
			return ip[0]
		}
		if ip := md.Get("x-real-ip"); len(ip) > 0 {
			return ip[0]
		}
	}

	// fallback: peer info (direct connection)
	if p, ok := peer.FromContext(ctx); ok {
		return p.Addr.String()
	}

	return "unknown"
}

func (s *Service) ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}
	exp, ok := claims["exp"].(float64)
	if !ok || int64(exp) < time.Now().Unix() {
		return "", errors.New("token expired")
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid user_id")
	}
	return userID, nil
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not set")
	}
	return []byte(secret)
}

func (s *Service) GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(s.jwtSecret))
}
