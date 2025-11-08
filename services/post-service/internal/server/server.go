package server

import (
	"context"
	pb "go-microservices/proto/post"
	"go-microservices/services/post-service/internal/repository"

	"google.golang.org/protobuf/types/known/emptypb"
)

type PostServer struct {
	pb.UnimplementedPostServiceServer
	repo *repository.Repository
}

func NewPostServer(repo *repository.Repository) *PostServer {
	return &PostServer{repo: repo}
}

func (s *PostServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	// TODO: Implement create post logic
	post := &pb.Post{Id: "1", AuthorId: req.AuthorId, Title: req.Title, Content: req.Content, CreatedAt: 0, UpdatedAt: 0}
	return &pb.CreatePostResponse{Post: post}, nil
}

func (s *PostServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	// TODO: Implement get post logic
	post := &pb.Post{Id: req.Id, AuthorId: "author", Title: "title", Content: "content", CreatedAt: 0, UpdatedAt: 0}
	return &pb.GetPostResponse{Post: post}, nil
}

func (s *PostServer) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	// TODO: Implement update post logic
	post := &pb.Post{Id: req.Id, AuthorId: "author", Title: req.Title, Content: req.Content, CreatedAt: 0, UpdatedAt: 0}
	return &pb.UpdatePostResponse{Post: post}, nil
}

func (s *PostServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	// TODO: Implement delete post logic
	return &emptypb.Empty{}, nil
}

func (s *PostServer) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	// TODO: Implement list posts logic
	posts := []*pb.Post{}
	return &pb.ListPostsResponse{Posts: posts}, nil
}
