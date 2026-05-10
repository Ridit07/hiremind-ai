package main

import (
	"fmt"
	"log"
	"net/http"

	"gateway-service/config"
	server "gateway-service/server/authserver"
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

	svc := server.NewGatewayService(authConn)

	router := transport.SetupRoutes(svc)

	httpAddr := fmt.Sprintf(":%s", cfg.HTTPPort)
	log.Printf("Starting HTTP server on %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, router); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
