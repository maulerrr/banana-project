package post

import (
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/models"
	"github.com/maulerrr/banana/services/post/pb"
	"google.golang.org/grpc"
	"net/http"
)

type Service struct {
	pb.UnimplementedPostServiceServer
	Handler    db.DBHandler
	Connection *grpc.ClientConn
	Client     pb.PostServiceClient
}

func InitPostService(svcURL string, handler db.DBHandler) (*Service, error) {
	conn, err := grpc.Dial(svcURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	service := &Service{
		Handler:    handler,
		Connection: conn,
		Client:     pb.NewPostServiceClient(conn),
	}

	return service, nil
}

func (s *Service) CreatePost(c *gin.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	post := &models.Post{
		Header: req.Header,
		Body:   req.Body,
		UserID: int(req.UserId),
	}

	createdPost, err := CreatePost(post, &s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, fmt.Errorf("failed to create post: %v", err)
	}

	pbPost := &pb.Post{
		PostId: int32(createdPost.PostID),
		Header: createdPost.Header,
		Body:   createdPost.Body,
		UserId: int32(createdPost.UserID),
	}

	return pbPost, nil
}

func (s *Service) GetAllPosts(c *gin.Context, req *pb.GetAllPostRequest) (*pb.GetAllPostResponse, error) {
	posts, err := GetAllPosts(&s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, fmt.Errorf("failed to get all posts: %v", err)
	}

	pbPosts := make([]*pb.Post, len(posts))
	for i, post := range posts {
		pbPosts[i] = &pb.Post{
			PostId: int32(post.PostID),
			Header: post.Header,
			Body:   post.Body,
			UserId: int32(post.UserID),
		}
	}

	response := &pb.GetAllPostResponse{
		Posts: pbPosts,
	}

	return response, nil
}

func (s *Service) GetPost(c *gin.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	post, err := GetPostByID(int(req.PostId), &s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, fmt.Errorf("failed to get post: %v", err)
	}

	pbPost := &pb.Post{
		PostId: int32(post.PostID),
		Header: post.Header,
		Body:   post.Body,
		UserId: int32(post.UserID),
	}

	return pbPost, nil
}

func (s *Service) DeletePost(c *gin.Context, req *pb.DeletePostRequest) (*pb.Empty, error) {
	err := DeletePost(int(req.PostId), &s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, fmt.Errorf("failed to delete post: %v", err)
	}

	return &pb.Empty{}, nil
}
