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
	"fmt"
	"strconv"

	pb "go-microservices/proto/auth"
	"go-microservices/services/auth-service/internal/models"
	"go-microservices/services/auth-service/internal/repository"
	"go-microservices/services/auth-service/internal/utils"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	repo *repository.Repository
}

func NewAuthServer(repo *repository.Repository) *AuthServer {
	return &AuthServer{repo: repo}
}

func (s *AuthServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	auth := &models.Auth{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
		Role:     req.Role,
	}

	if err := s.repo.CreateAuth(auth); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	userID := fmt.Sprintf("%d", auth.ID)
	return &pb.SignUpResponse{UserId: userID, Message: "registered"}, nil
}

func (s *AuthServer) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	auth, err := s.repo.GetAuthByEmail(req.Email)
	if err != nil || auth == nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(req.Password)); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	accessToken, refreshToken, err := utils.GenerateJWT(*auth)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate tokens: %v", err)
	}

	userID := fmt.Sprintf("%d", auth.ID)
	return &pb.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserId:       userID,
		Message:      "logged in",
	}, nil
}

func (s *AuthServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	claims, err := utils.ValidateJWT(req.Token, false)
	if err != nil {
		return &pb.ValidateTokenResponse{Valid: false, UserId: "", Message: "invalid token"}, nil
	}

	var userID string
	if sub, ok := claims["sub"]; ok {
		switch v := sub.(type) {
		case string:
			userID = v
		case float64:
			userID = strconv.FormatFloat(v, 'f', 0, 64)
		default:
			userID = fmt.Sprintf("%v", v)
		}
	}

	if userID == "" {
		return &pb.ValidateTokenResponse{Valid: false, UserId: "", Message: "invalid token"}, nil
	}

	return &pb.ValidateTokenResponse{Valid: true, UserId: userID, Message: "valid"}, nil
}

func (s *AuthServer) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	if req.UserId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user id required")
	}
	u64, err := strconv.ParseUint(req.UserId, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id: %v", err)
	}
	user, err := s.repo.GetAuthByID(uint(u64))
	if err != nil || user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &pb.GetUserInfoResponse{
		UserId:   fmt.Sprintf("%d", user.ID),
		Username: user.Username,
		Email:    user.Email,
		Roles:    []string{user.Role},
	}, nil
}
