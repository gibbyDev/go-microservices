package clients

import (
	"context"
	pb "go-microservices/proto/auth"
	pbPost "go-microservices/proto/post"
	pbUser "go-microservices/proto/user"

	"google.golang.org/grpc"
)

type AuthClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	return &AuthClient{
		client: pb.NewAuthServiceClient(conn),
	}
}

func (a *AuthClient) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return a.client.SignUp(ctx, req)
}

func (a *AuthClient) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	return a.client.SignIn(ctx, req)
}

func (a *AuthClient) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	return a.client.ValidateToken(ctx, req)
}

func (a *AuthClient) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	return a.client.GetUserInfo(ctx, req)
}

func (a *AuthClient) CreateTest(ctx context.Context, req *pb.CreateTestRequest) (*pb.CreateTestResponse, error) {
	return a.client.CreateTest(ctx, req)
}

func (a *AuthClient) ListTests(ctx context.Context, req *pb.ListTestsRequest) (*pb.ListTestsResponse, error) {
	return a.client.ListTests(ctx, req)
}

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
