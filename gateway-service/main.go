package main

import (
	"fmt"
	"log"
	"net/http"

	"gateway-service/config"
	authserver "gateway-service/server/authserver"
	interviewserver "gateway-service/server/interviewserver"
	transport "gateway-service/transport"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	authConn, err := grpc.Dial(
		cfg.AuthServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth service: %v", err)
	}
	defer authConn.Close()
	// why insecure is it because we are in a local environment and not using TLS? In production, we should use TLS for secure communication between services. For local development, we can skip TLS to simplify setup.
	interviewConn, err := grpc.Dial(
		cfg.InterviewServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to connect to interview service: %v", err)
	}
	defer interviewConn.Close()

	authSvc := authserver.NewGatewayService(authConn)
	interviewSvc := interviewserver.NewInterviewService(interviewConn)

	router := transport.SetupRoutes(authSvc, interviewSvc)

	httpAddr := fmt.Sprintf(":%s", cfg.HTTPPort)
	log.Printf("Starting HTTP server on %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, router); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
