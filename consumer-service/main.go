package main

import (
	"log"
	"net"

	"consumer-service/config"
	"consumer-service/db"
	authServer "consumer-service/transport_grpc/auth"

	pb "github.com/Ridit07/hiremind-proto-contracts/generated/auth"

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
		&authServer.AuthServer{},
	)

	log.Println(
		"consumer-service running on :50051",
	)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal(err)
	}
}
