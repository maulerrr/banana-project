package routes

import (
	"github.com/gin-gonic/gin"
	middlewares "github.com/maulerrr/banana/pkg/middleware"
	"github.com/maulerrr/banana/services/auth"
	"github.com/maulerrr/banana/services/comment"
	"github.com/maulerrr/banana/services/post"
)

func RegisterAuthRoutes(r *gin.Engine, handler *auth.Handler) {
	router := r.Group("/auth")

	router.POST("/login", handler.LoginHandler)
	router.POST("/signup", handler.SignUpHandler)
}

func RegisterPostRoutes(r *gin.Engine, handler *post.Handler) {
	router := r.Group("/post")

	router.Use(middlewares.AuthMiddleware())

	router.GET("/", handler.GetAllPostsHandler)
	router.POST("/", handler.CreatePostHandler)
	router.GET("/:id", handler.GetPostHandler)
	router.DELETE("/:id", middlewares.PostDeletionMiddleware(), handler.DeletePostHandler)

}

func RegisterCommentRoutes(r *gin.Engine, handler *comment.Handler) {
	router := r.Group("comment")

	router.Use(middlewares.AuthMiddleware())

	router.GET("/", handler.GetAllCommentsHandler)
	router.POST("/", handler.CreateCommentHandler)
	router.DELETE("/:id", middlewares.CommentDeletionMiddleware(), handler.DeleteCommentHandler)
}
