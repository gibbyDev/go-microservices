package server

import (
	"context"

	pb "go-microservices/proto/user"

	"go-microservices/services/user-service/internal/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *repository.Repository
}

func NewUserServer(repo *repository.Repository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// TODO: Implement create user logic
	user := &pb.User{Id: "1", Username: req.Username, Email: req.Email, Bio: req.Bio, AvatarUrl: req.AvatarUrl, CreatedAt: 0, UpdatedAt: 0}
	return &pb.CreateUserResponse{User: user}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// TODO: Implement get user logic
	user := &pb.User{Id: req.Id, Username: "user", Email: "user@example.com", Bio: "bio", AvatarUrl: "", CreatedAt: 0, UpdatedAt: 0}
	return &pb.GetUserResponse{User: user}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	// TODO: Implement update user logic
	user := &pb.User{Id: req.Id, Username: req.Username, Email: req.Email, Bio: req.Bio, AvatarUrl: req.AvatarUrl, CreatedAt: 0, UpdatedAt: 0}
	return &pb.UpdateUserResponse{User: user}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	// TODO: Implement delete user logic
	return &emptypb.Empty{}, nil
}

func (s *UserServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	// TODO: Implement list users logic
	users := []*pb.User{}
	return &pb.ListUsersResponse{Users: users}, nil
}
