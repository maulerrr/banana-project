package main

import (
	"github.com/gin-contrib/cors"
	"github.com/maulerrr/banana/pkg/rabbitmq"
	"github.com/maulerrr/banana/services/like"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maulerrr/banana/pkg/config"
	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/routes"
	"github.com/maulerrr/banana/services/auth"
	"github.com/maulerrr/banana/services/comment"
	"github.com/maulerrr/banana/services/post"
)

func main() {
	dbHandler := db.InitDB(config.DBUrl)
	producer, err := rabbitmq.NewProducer(config.AMPQUrl, "likes")
	consumer, err := rabbitmq.NewConsumer(config.AMPQUrl, "likes")
	if err != nil {
		log.Fatal("rabbit error:", err)
		return
	}

	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	authService, err := auth.InitAuthService(config.AuthSvcUrl, dbHandler)
	if err != nil {
		log.Fatal("error initializing auth service: ", err)
		return
	}
	authHandlers := auth.NewHandler(authService)
	routes.RegisterAuthRoutes(app, authHandlers)

	postService, err := post.InitPostService(config.PostSvcUrl, dbHandler)
	if err != nil {
		log.Fatal("error initializing post service: ", err)
		return
	}
	postHandlers := post.NewHandler(postService)
	routes.RegisterPostRoutes(app, postHandlers)

	commentService, err := comment.InitCommentService(config.CommentSvcUrl, dbHandler)
	if err != nil {
		log.Fatal("error initializing comment service: ", err)
		return
	}
	commentHandlers := comment.NewHandler(commentService)
	routes.RegisterCommentRoutes(app, commentHandlers)

	likeService := like.NewLikeService(&dbHandler, producer)
	likeHandler := like.NewLikeHandler(likeService, consumer)
	routes.RegisterLikeRoutes(app, likeHandler)

	err = app.Run(config.Port)
	if err != nil {
		log.Fatal("error running server: ", err)
	}
}
