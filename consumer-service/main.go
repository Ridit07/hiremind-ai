package main

import (
	"log"
	"net"
	"os"

	"consumer-service/config"
	"consumer-service/db"
	"consumer-service/services/authService"
	authServer "consumer-service/transport_grpc/auth"

	"consumer-service/redisclient"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated/auth"
	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"
)

func main() {

	config.LoadConfig()

	err := db.InitDB(
		config.AppConfig.DBReadURL,
		config.AppConfig.DBWriteURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // move to env later
	})
	redisWrapper := redisclient.NewClient(rdb)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set")
	}

	authSvc := authService.NewService(redisWrapper, jwtSecret)

	// ✅ Setup gRPC server WITH interceptors
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			authService.LoggingInterceptor(),
			authService.RateLimitInterceptor(rdb),
			authService.AuthInterceptor(authSvc),
		),
	)

	// ✅ Inject service into transport
	pb.RegisterAuthServiceServer(
		grpcServer,
		&authServer.AuthServer{
			Service: authSvc,
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
