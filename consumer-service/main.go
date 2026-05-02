package main

import (
	"log"
	"net"
	"os"

	"consumer-service/db"
	grpcserver "consumer-service/transport/grpc"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated"

	"google.golang.org/grpc"
)

func main() {

	readDBURL := os.Getenv(
		"READ_DATABASE_URL",
	)

	writeDBURL := os.Getenv(
		"WRITE_DATABASE_URL",
	)

	err := db.InitDB(
		readDBURL,
		writeDBURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen(
		"tcp",
		":50051",
	)

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(
		grpcServer,
		&grpcserver.AuthServer{},
	)

	log.Println(
		"consumer-service running on :50051",
	)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal(err)
	}
}
