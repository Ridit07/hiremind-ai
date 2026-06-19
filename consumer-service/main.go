package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"

	"consumer-service/config"
	"consumer-service/db"
	"consumer-service/services/authService"
	"consumer-service/services/interviewService"
	authServer "consumer-service/transport_grpc/auth"
	interviewServer "consumer-service/transport_grpc/interview"

	"consumer-service/redisclient"

	authpb "github.com/Ridit07/hiremind-proto-contracts/generated/auth"
	interviewpb "github.com/Ridit07/hiremind-proto-contracts/generated/interview"
	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	err = db.InitDB(
		cfg.DBReadURL,
		cfg.DBWriteURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})

	redisWrapper := redisclient.NewClient(rdb)

	ctx := context.Background()

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("REDIS PING FAILED: %v", err)
	}

	log.Printf("REDIS CONNECTED: %s", pong)

	jwtSecret := cfg.JWTSecret
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}

	authSvc := authService.NewService(redisWrapper, jwtSecret)
	interviewSvc := interviewService.NewService()

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			authService.LoggingInterceptor(),
			authService.RateLimitInterceptor(rdb),
			authService.AuthInterceptor(authSvc),
		),
	)

	authpb.RegisterAuthServiceServer(
		grpcServer,
		&authServer.AuthServer{
			Service: authSvc,
		},
	)

	interviewpb.RegisterInterviewServiceServer(
		grpcServer,
		&interviewServer.InterviewServer{
			Service: interviewSvc,
		},
	)

	listener, err := net.Listen(
		"tcp",
		":50051",
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		"consumer-service running on :50051",
	)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal(err)
	}
}
