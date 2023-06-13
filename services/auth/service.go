package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/banana/pkg/config"
	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/models"
	"github.com/maulerrr/banana/pkg/utils"
	"github.com/maulerrr/banana/services/auth/pb"
	"google.golang.org/grpc"
	"net/http"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	Handler    db.DBHandler
	Connection *grpc.ClientConn
	Client     pb.AuthServiceClient
}

func InitAuthService(svcURL string, handler db.DBHandler) (*Service, error) {
	conn, err := grpc.Dial(svcURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	service := &Service{
		Handler:    handler,
		Connection: conn,
		Client:     pb.NewAuthServiceClient(conn),
	}

	return service, nil
}

type LoginReqBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpReqBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Service) Login(c *gin.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	b := LoginReqBody{}

	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil, errors.New("invalid request body")
	}

	found, err := GetUserByEmail(b.Email, s.Handler)
	if err != nil && found == nil {
		c.AbortWithError(http.StatusNotFound, err)
		return nil, errors.New("such user was not found")
	}

	if !utils.ComparePasswords(found.Password, b.Password) {
		c.AbortWithError(http.StatusUnauthorized, errors.New("incorrect password"))
		return nil, errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(*found, config.JWTSecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, errors.New("apologies for inconvenience, some error occurred")
	}

	return &pb.AuthResponse{
		UserId:   int64(found.UserID),
		Username: found.Username,
		Email:    found.Email,
		Token:    token,
	}, nil

}

func (s *Service) SignUp(c *gin.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	b := SignUpReqBody{}

	if err := c.BindJSON(&b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil, errors.New("invalid request body")
	}

	found, err := GetUserByEmail(b.Email, s.Handler)
	if err == nil && found != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil, errors.New("user already exists")
	}

	newUser := models.User{
		Username: b.Username,
		Email:    b.Email,
		Password: b.Password,
	}

	created, err := CreateUser(&newUser, &s.Handler)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, errors.New("could not create an account")
	}

	token, err := utils.GenerateToken(*found, config.JWTSecretKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return nil, errors.New("apologies for inconvenience, some error occurred")
	}

	return &pb.AuthResponse{
		UserId:   int64(created.UserID),
		Username: created.Username,
		Email:    created.Email,
		Token:    token,
	}, nil
}
