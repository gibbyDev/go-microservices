// -----------------------------------------------------------------------------
// File: main.go
//
// This is the entry point for the Auth Service microservice. It initializes the
// gRPC server, listens on a configured port, and registers the AuthService server
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
// - Serves as the central orchestrator for the Auth Service.
// - Exposes authentication operations to other services via gRPC.
// - Manages server startup, listening, and graceful error handling.
// -----------------------------------------------------------------------------
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/gibbyDev/go-microservices/proto/auth"
	"github.com/gibbyDev/go-microservices/services/auth-service/internal/server"
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
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server.NewAuthServer())
	log.Printf("Auth Service listening on %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
