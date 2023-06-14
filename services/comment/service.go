package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/models"
	"github.com/maulerrr/banana/services/comment/pb"
	"google.golang.org/grpc"
	"net/http"
)

type Service struct {
	pb.UnimplementedCommentServiceServer
	Handler    db.DBHandler
	Connection *grpc.ClientConn
	Client     pb.CommentServiceClient
}

func InitCommentService(svcURL string, handler db.DBHandler) (*Service, error) {
	conn, err := grpc.Dial(svcURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	service := &Service{
		Handler:    handler,
		Connection: conn,
		Client:     pb.NewCommentServiceClient(conn),
	}

	return service, nil
}

func (s *Service) GetAllComments(c *gin.Context, req *pb.GetAllCommentRequest) (*pb.GetAllCommentResponse, error) {
	comments, err := GetAllCommentsByPostID(int(req.PostId), &s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, fmt.Errorf("failed to get all comments: %v", err)
	}

	pbComments := make([]*pb.Comment, len(comments))
	for i, comment := range comments {
		pbComments[i] = &pb.Comment{
			CommentId: int32(comment.CommentID),
			PostId:    int32(comment.PostID),
			UserId:    int32(comment.UserID),
			Text:      comment.Text,
		}
	}

	response := &pb.GetAllCommentResponse{
		Comments: pbComments,
	}

	return response, nil
}

func (s *Service) DeleteComment(c *gin.Context, req *pb.DeleteCommentRequest) (*pb.Empty, error) {
	err := DeleteComment(int(req.CommentId), &s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, fmt.Errorf("failed to delete comment: %v", err)
	}

	return &pb.Empty{}, nil
}

func (s *Service) CreateComment(c *gin.Context, req *pb.CreateCommentRequest) (*pb.Comment, error) {
	comment := &models.Comment{
		PostID: int(req.PostId),
		UserID: int(req.UserId),
		Text:   req.Text,
	}

	createdComment, err := CreateComment(comment, &s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, fmt.Errorf("failed to create comment: %v", err)
	}

	pbComment := &pb.Comment{
		CommentId: int32(createdComment.CommentID),
		PostId:    int32(createdComment.PostID),
		UserId:    int32(createdComment.UserID),
		Text:      createdComment.Text,
	}

	return pbComment, nil
}
