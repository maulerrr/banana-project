package test

import (
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/banana/pkg/models"
	"github.com/maulerrr/banana/pkg/utils"
	"github.com/maulerrr/banana/services/comment"
	"net/http"
	"testing"
	"time"

	helper "github.com/maulerrr/banana/test/helper"

	authpb "github.com/maulerrr/banana/services/auth/pb"
	postpb "github.com/maulerrr/banana/services/post/pb"
)

func TestDBReset(t *testing.T) {
	helper.InitAuthMock()
	helper.ResetDB()
}

func TestLoginHandler(t *testing.T) {
	handler := helper.InitAuthMock()

	helper.InitTestInstances(handler.Service.Handler, true, false, false)

	testcases := []helper.Testcase{
		{
			Name: "Valid credentials",
			Payload: &authpb.LoginRequest{
				Email:    "test@example.com",
				Password: "password",
			},
			ExpectedCode: http.StatusOK,
			ExpectedData: &authpb.AuthResponse{
				UserId:   1,
				Username: "testuser",
				Email:    "test@example.com",
				Token:    "<mocked_token>",
			},
		},
		{
			Name: "User not found",
			Payload: &authpb.LoginRequest{
				Email:    "nonexistent@example.com",
				Password: "password",
			},
			ExpectedCode: http.StatusNotFound,
			ExpectedData: gin.H{"error": "such user was not found"},
		},
		{
			Name: "Incorrect password",
			Payload: &authpb.LoginRequest{
				Email:    "test@example.com",
				Password: "wrong_password",
			},
			ExpectedCode: http.StatusUnauthorized,
			ExpectedData: gin.H{"error": "incorrect password"},
		},
		{
			Name:         "Empty payload",
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: gin.H{"error": "Invalid request body"},
		},
	}

	helper.TestRun(testcases, handler.LoginHandler, http.MethodPost, true, t)
	helper.ResetDB()
}

func TestSignUpHandler(t *testing.T) {
	handler := helper.InitAuthMock()

	handler.Service.Handler.DB.Create(&models.User{
		UserID:    2,
		Username:  "existinguser",
		Email:     "exintingtest@example.com",
		Password:  utils.HashPassword([]byte("password")),
		CreatedAt: time.Now(),
	})

	testcases := []helper.Testcase{
		{
			Name: "Valid registration",
			Payload: &authpb.RegisterRequest{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password",
			},
			ExpectedCode: http.StatusOK,
			ExpectedData: &authpb.AuthResponse{
				UserId:   1,
				Username: "testuser",
				Email:    "test@example.com",
				Token:    "<mocked_token>",
			},
		},
		{
			Name: "User already exists",
			Payload: &authpb.RegisterRequest{
				Username: "existinguser",
				Email:    "exintingtest@example.com",
				Password: "password",
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: gin.H{"error": "user already exists"},
		},
		{
			Name:         "Empty payload",
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: gin.H{"error": "Invalid request body"},
		},
	}

	helper.TestRun(testcases, handler.SignUpHandler, http.MethodPost, true, t)
	helper.ResetDB()
}

// test post handler
func TestGetAllPostsHandler(t *testing.T) {
	handler := helper.InitPostsMock()

	helper.InitTestInstances(handler.Service.Handler, true, true, false)

	testcases := []helper.Testcase{
		{
			Name:         "Valid Request",
			ExpectedCode: http.StatusOK,
			ExpectedData: []*postpb.Post{
				{
					PostId: 1,
					Header: "Test Post 1",
					Body:   "This is test post 1",
					UserId: 1,
				},
			},
		},
	}

	helper.TestRun(testcases, handler.GetAllPostsHandler, http.MethodGet, false, t)
	helper.ResetDB()
}

func TestGetPostHandler(t *testing.T) {
	handler := helper.InitPostsMock()

	helper.InitTestInstances(handler.Service.Handler, true, true, false)

	testcases := []helper.Testcase{
		{
			Name:         "Valid Request",
			Params:       gin.Params{gin.Param{Key: "id", Value: "1"}},
			ExpectedCode: http.StatusOK,
			ExpectedData: postpb.Post{
				PostId: 1,
				Header: "Test Post 1",
				Body:   "This is a test post 1",
				UserId: 1,
			},
		},
		{
			Name:         "Invalid Request",
			Params:       gin.Params{gin.Param{Key: "id", Value: "invalid"}},
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: map[string]interface{}{
				"error": "Invalid post ID",
			},
		},
	}

	helper.TestRun(testcases, handler.GetPostHandler, http.MethodGet, false, t)
	helper.ResetDB()
}

func TestCreatePostHandler(t *testing.T) {
	handler := helper.InitPostsMock()

	helper.InitTestInstances(handler.Service.Handler, true, false, false)

	testcases := []helper.Testcase{
		{
			Name: "Valid Request",
			Payload: &postpb.CreatePostRequest{
				Header: "Test Post",
				Body:   "This is a test post",
				UserId: 1,
			},
			ExpectedCode: http.StatusOK,
			ExpectedData: &postpb.Post{
				PostId: 1,
				Header: "Test Post",
				Body:   "This is a test post",
				UserId: 1,
			},
		},
		{
			Name:         "Invalid request body/Empty",
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: map[string]interface{}{
				"error": "Invalid request body",
			},
		},
	}

	helper.TestRun(testcases, handler.CreatePostHandler, http.MethodPost, true, t)
	helper.ResetDB()
}

func TestDeletePostHandler(t *testing.T) {
	handler := helper.InitPostsMock()

	helper.InitTestInstances(handler.Service.Handler, true, true, false)

	testcases := []helper.Testcase{
		{
			Name:         "Valid Request",
			Params:       gin.Params{gin.Param{Key: "id", Value: "1"}},
			ExpectedCode: http.StatusOK,
			ExpectedData: map[string]interface{}{},
		},
		{
			Name:         "Invalid Request",
			Params:       gin.Params{gin.Param{Key: "id", Value: "invalid"}},
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: map[string]interface{}{"error": "Invalid post ID"},
		},
	}

	helper.TestRun(testcases, handler.DeletePostHandler, http.MethodDelete, false, t)
	helper.ResetDB()
}

/ test comment handler
func TestGetAllCommentsHandler(t *testing.T) {
	handler := helper.InitCommentMock()

	helper.InitTestInstances(handler.Service.Handler, true, true, true)

	testcases := []helper.Testcase{
		{
			Name:         "Valid post ID",
			Params:       gin.Params{gin.Param{Key: "id", Value: "1"}},
			ExpectedCode: http.StatusOK,
			ExpectedData: []comment.ResponseToGetAllComments{
				{
					CommentID: 1,
					UserID:    1,
					PostID:    1,
					Username:  "testuser",
					Text:      "Test comment",
				},
			},
		},
		{
			Name:         "Invalid post ID",
			Params:       gin.Params{gin.Param{Key: "id", Value: "invalid"}},
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: gin.H{"error": "Invalid post ID"},
		},
	}

	helper.TestRun(testcases, handler.GetAllCommentsHandler, http.MethodGet, false, t)
	helper.ResetDB()
}

func TestCreateCommentHandler(t *testing.T) {
	handler := helper.InitCommentMock()

	helper.InitTestInstances(handler.Service.Handler, true, true, false)

	testcases := []helper.Testcase{
		{
			Name: "Valid request",
			Payload: commentpb.CreateCommentRequest{
				PostId: 1,
				UserId: 1,
				Text:   "New comment",
			},
			ExpectedCode: http.StatusOK,
			ExpectedData: gin.H{
				"comment_id": 1,
				"post_id":    1,
				"user_id":    1,
				"text":       "New comment",
			},
		},
		{
			Name:         "Invalid request",
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: gin.H{"error": "Invalid request body"},
		},
	}

	helper.TestRun(testcases, handler.CreateCommentHandler, http.MethodPost, true, t)
	helper.ResetDB()
}

func TestDeleteCommentHandler(t *testing.T) {
	handler := helper.InitCommentMock()

	helper.InitTestInstances(handler.Service.Handler, true, true, true)

	testcases := []helper.Testcase{
		{
			Name:         "Valid comment ID",
			Params:       gin.Params{gin.Param{Key: "id", Value: "1"}},
			ExpectedCode: http.StatusOK,
			ExpectedData: gin.H{},
		},
		{
			Name:         "Invalid comment ID",
			Params:       gin.Params{gin.Param{Key: "id", Value: "invalid"}},
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: gin.H{"error": "Invalid comment ID"},
		},
	}

	helper.TestRun(testcases, handler.DeleteCommentHandler, http.MethodDelete, false, t)
	helper.ResetDB()
}