package main

import (
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
	router := gin.Default()

	authService, err := auth.InitAuthService(config.AuthSvcUrl, dbHandler)
	if err != nil {
		log.Fatal("error initializing auth service: ", err)
		return
	}
	authHandlers := auth.NewHandler(authService)
	routes.RegisterAuthRoutes(router, authHandlers)

	postService, err := post.InitPostService(config.PostSvcUrl, dbHandler)
	if err != nil {
		log.Fatal("error initializing post service: ", err)
		return
	}
	postHandlers := post.NewHandler(postService)
	routes.RegisterPostRoutes(router, postHandlers)

	commentService, err := comment.InitCommentService(config.CommentSvcUrl, dbHandler)
	if err != nil {
		log.Fatal("error initializing comment service: ", err)
		return
	}
	commentHandlers := comment.NewHandler(commentService)
	routes.RegisterCommentRoutes(router, commentHandlers)

	err = router.Run(config.Port)
	if err != nil {
		log.Fatal("error running server: ", err)
	}
}
