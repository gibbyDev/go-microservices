package main

import (
	"fmt"
	pb "go-microservices/proto/user"
	"go-microservices/services/user-service/internal/database"
	"go-microservices/services/user-service/internal/repository"
	"go-microservices/services/user-service/internal/server"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting User Service...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "50052"
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

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, server.NewUserServer(repo))
	log.Printf("User Service listening on %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
