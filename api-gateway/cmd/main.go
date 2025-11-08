package main

import (
	"fmt"
	"log"
	"os"

	"go-microservices/api-gateway/internal/handlers"
	"go-microservices/api-gateway/internal/routes"

	"go-microservices/api-gateway/internal/clients"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting API Gateway...")
	app := fiber.New()

	// Connect to AuthService gRPC
	authServiceAddr := os.Getenv("AUTH_SERVICE_GRPC")

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
	log.Printf("API Gateway listening on :%s", port)

	log.Fatal(app.Listen(":" + port))
}
