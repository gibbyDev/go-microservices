// -----------------------------------------------------------------------------
// File: server.go
//
// This file implements the gRPC server for the AuthService microservice. It defines
// the AuthServer struct, which embeds the generated UnimplementedAuthServiceServer
// from the protobuf code. The server provides method stubs for Register, Login,
// ValidateToken, and GetUserInfo, which are called by the API Gateway via gRPC.
//
// Syntax:
// - Uses Go's struct and method syntax to implement gRPC service methods.
// - Embeds the protobuf-generated server interface for compatibility.
// - Uses context for request handling and cancellation.
//
// Purpose:
// - Serves as the backend for authentication operations.
// - Handles registration, login, token validation, and user info retrieval.
// - Provides a contract for future business logic and database integration.
// -----------------------------------------------------------------------------
package server

import (
	"context"
	pb "github.com/gibbyDev/go-microservices/proto/auth"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	// Add dependencies: repository, config, etc.
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (s *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// TODO: Implement registration logic
	return &pb.RegisterResponse{UserId: "123", Message: "registered"}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// TODO: Implement login logic
	return &pb.LoginResponse{AccessToken: "token", RefreshToken: "refresh", UserId: "123", Message: "logged in"}, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	// TODO: Implement token validation
	return &pb.ValidateTokenResponse{Valid: true, UserId: "123", Message: "valid"}, nil
}

func (s *AuthServer) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	// TODO: Implement user info retrieval
	return &pb.GetUserInfoResponse{UserId: "123", Username: "user", Email: "user@example.com", Roles: []string{"user"}}, nil
}