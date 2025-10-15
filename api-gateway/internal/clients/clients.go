// -----------------------------------------------------------------------------
// File: clients.go
//
// This file defines the gRPC client for the AuthService microservice. It imports
// the generated protobuf code and wraps the gRPC client in a struct for easy use
// throughout the API Gateway. The AuthClient struct provides methods to call the
// Register, Login, ValidateToken, and GetUserInfo RPCs defined in the proto file.
//
// Syntax:
// - Uses Go's package system and imports for modularity.
// - Defines a struct to encapsulate the gRPC client.
// - Provides constructor and method wrappers for each RPC.
//
// Purpose:
// - Allows the API Gateway to communicate with the AuthService using gRPC.
// - Abstracts away connection details and provides a clean interface for handlers.
// -----------------------------------------------------------------------------
package clients

import (
	"context"
	"google.golang.org/grpc"
	pb "github.com/gibbyDev/go-microservices/proto/auth"
)

type AuthClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	return &AuthClient{
		client: pb.NewAuthServiceClient(conn),
	}
}

func (a *AuthClient) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return a.client.Register(ctx, req)
}

func (a *AuthClient) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return a.client.Login(ctx, req)
}

func (a *AuthClient) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	return a.client.ValidateToken(ctx, req)
}

func (a *AuthClient) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return a.client.GetUserInfo(ctx, req)
}
// -----------------------------------------------------------------------------
// The following code defines the PostClient for the PostService microservice.
// It follows the same pattern as AuthClient, providing methods for CRUD operations
// on posts via gRPC. This enables the API Gateway to communicate with the PostService
// for creating, retrieving, updating, deleting, and listing posts.
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// The following code defines the UserClient for the UserService microservice.
// It follows the same pattern as AuthClient and PostClient, providing methods for
// CRUD operations on users via gRPC. This enables the API Gateway to communicate
// with the UserService for creating, retrieving, updating, deleting, and listing users.
// -----------------------------------------------------------------------------

import (
	pbUser "github.com/gibbyDev/go-microservices/proto/user"
)

type UserClient struct {
	client pbUser.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: pbUser.NewUserServiceClient(conn),
	}
}

func (u *UserClient) CreateUser(ctx context.Context, req *pbUser.CreateUserRequest) (*pbUser.CreateUserResponse, error) {
	return u.client.CreateUser(ctx, req)
}

func (u *UserClient) GetUser(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.GetUserResponse, error) {
	return u.client.GetUser(ctx, req)
}

func (u *UserClient) UpdateUser(ctx context.Context, req *pbUser.UpdateUserRequest) (*pbUser.UpdateUserResponse, error) {
	return u.client.UpdateUser(ctx, req)
}

func (u *UserClient) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest) error {
	_, err := u.client.DeleteUser(ctx, req)
	return err
}

func (u *UserClient) ListUsers(ctx context.Context, req *pbUser.ListUsersRequest) (*pbUser.ListUsersResponse, error) {
	return u.client.ListUsers(ctx, req)
}

import (
	pbPost "github.com/gibbyDev/go-microservices/proto/post"
)

type PostClient struct {
	client pbPost.PostServiceClient
}

func NewPostClient(conn *grpc.ClientConn) *PostClient {
	return &PostClient{
		client: pbPost.NewPostServiceClient(conn),
	}
}

func (p *PostClient) CreatePost(ctx context.Context, req *pbPost.CreatePostRequest) (*pbPost.CreatePostResponse, error) {
	return p.client.CreatePost(ctx, req)
}

func (p *PostClient) GetPost(ctx context.Context, req *pbPost.GetPostRequest) (*pbPost.GetPostResponse, error) {
	return p.client.GetPost(ctx, req)
}

func (p *PostClient) UpdatePost(ctx context.Context, req *pbPost.UpdatePostRequest) (*pbPost.UpdatePostResponse, error) {
	return p.client.UpdatePost(ctx, req)
}

func (p *PostClient) DeletePost(ctx context.Context, req *pbPost.DeletePostRequest) error {
	_, err := p.client.DeletePost(ctx, req)
	return err
}

func (p *PostClient) ListPosts(ctx context.Context, req *pbPost.ListPostsRequest) (*pbPost.ListPostsResponse, error) {
	return p.client.ListPosts(ctx, req)
}