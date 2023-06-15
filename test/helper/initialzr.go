package testing

import (
	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/models"
	"github.com/maulerrr/banana/pkg/utils"
	"github.com/maulerrr/banana/services/auth"
	"github.com/maulerrr/banana/services/comment"
	"github.com/maulerrr/banana/services/post"
	"time"
)

var (
	mockDBURL     = "postgres://postgres:1111@localhost:5432/banana_test?sslmode=disable"
	authSvcURL    = ":50051"
	postSvcURL    = ":50052"
	commentSvcURL = ":50053"
)

func InitAuthMock() *auth.Handler {
	dbHandler := db.InitDB(mockDBURL)
	authSvc, _ := auth.InitAuthService(authSvcURL, dbHandler)
	authHandler := auth.NewHandler(authSvc)

	return authHandler
}

func InitPostsMock() *post.Handler {
	dbHandler := db.InitDB(mockDBURL)
	Svc, _ := post.InitPostService(postSvcURL, dbHandler)
	Handler := post.NewHandler(Svc)

	return Handler
}

func InitCommentMock() *comment.Handler {
	dbHandler := db.InitDB(mockDBURL)
	Svc, _ := comment.InitCommentService(commentSvcURL, dbHandler)
	Handler := comment.NewHandler(Svc)

	return Handler
}

func ResetDB() {
	dbHandler := db.GetDBHandler()
	dbHandler.DB.Exec("TRUNCATE TABLE users, posts, comments RESTART IDENTITY CASCADE")
}

func InitTestInstances(h db.DBHandler, withUser, withPost, withComment bool) {
	if withUser {
		h.DB.Create(&models.User{
			UserID:    1,
			Username:  "testuser",
			Email:     "test@example.com",
			Password:  utils.HashPassword([]byte("password")),
			CreatedAt: time.Now(),
		})
	}
	if withPost {
		h.DB.Create(
			&models.Post{
				PostID: 1,
				UserID: 1,
				Header: "Test Post 1",
				Body:   "This is test post 1",
			},
		)
	}
	if withComment {
		h.DB.Create(
			&models.Comment{
				CommentID: 1,
				UserID:    1,
				PostID:    1,
				Text:      "Test comment",
			},
		)
	}
}
