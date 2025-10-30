// -----------------------------------------------------------------------------
// File: main.go
//
// This is the entry point for the Post Service microservice. It initializes the
// gRPC server, listens on a configured port, and registers the PostService server
// implementation. The main function demonstrates how to start a gRPC service in Go
// and connect it to the generated protobuf contract.
//
// Syntax:
// - Uses Go's main package and function for application entry.
// - Uses environment variables for configuration.
// - Initializes gRPC server and registers service implementation.
// - Handles server lifecycle and error logging.
//
// Purpose:
// - Serves as the central orchestrator for the Post Service.
// - Exposes post CRUD operations to other services via gRPC.
// - Manages server startup, listening, and graceful error handling.
// -----------------------------------------------------------------------------
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
	// initialize database
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
