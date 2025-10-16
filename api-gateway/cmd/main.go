// -----------------------------------------------------------------------------
// File: main.go
//
// This is the entry point for the API Gateway. It initializes the Fiber web
// server, sets up the gRPC client connection to the AuthService, creates handler
// instances, registers REST routes, and starts the HTTP server. The main function
// demonstrates how to bridge RESTful HTTP requests to gRPC microservices.
//
// Syntax:
// - Uses Go's main package and function for application entry.
// - Uses environment variables for configuration.
// - Initializes Fiber and gRPC client connections.
// - Registers routes and starts the server.
//
// Purpose:
// - Serves as the central orchestrator for the API Gateway.
// - Connects frontend REST requests to backend gRPC services.
// - Manages server lifecycle and error handling.
// -----------------------------------------------------------------------------
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gibbyDev/go-microservices/api-gateway/internal/handlers"
	"github.com/gibbyDev/go-microservices/api-gateway/internal/routes"

	"github.com/gibbyDev/go-microservices/api-gateway/internal/clients"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting API Gateway...")
	app := fiber.New()

	// Connect to AuthService gRPC
	authServiceAddr := os.Getenv("AUTH_SERVICE_GRPC")
	if authServiceAddr == "" {
		authServiceAddr = "localhost:50051"
	}
	conn, err := grpc.Dial(authServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to AuthService: %v", err)
	}
	defer conn.Close()

	authClient := clients.NewAuthClient(conn)
	authHandler := handlers.NewAuthHandler(authClient)

	routes.RegisterAuthRoutes(app, authHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
