package main

import (
	"fmt"
	pb "go-microservices/proto/post"
	"go-microservices/services/post-service/internal/database"
	"go-microservices/services/post-service/internal/repository"
	"go-microservices/services/post-service/internal/server"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Post Service...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "50053"
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
	pb.RegisterPostServiceServer(grpcServer, server.NewPostServer(repo))
	log.Printf("Post Service listening on %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
