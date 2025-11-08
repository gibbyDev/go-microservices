package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "go-microservices/proto/auth"

	"go-microservices/services/auth-service/internal/database"
	"go-microservices/services/auth-service/internal/repository"
	"go-microservices/services/auth-service/internal/server"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Auth Service...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db, err := database.Init()
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	repo := repository.NewRepository(db)
	srv := server.NewAuthServer(repo)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, srv)
	log.Printf("Auth Service listening on %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
