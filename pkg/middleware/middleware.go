package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/banana/pkg/config"
	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/utils"
	"github.com/maulerrr/banana/services/comment"
	"github.com/maulerrr/banana/services/post"
	"strconv"

	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			utils.SendMessageWithStatus(context, "authorize!", 401)
			context.Abort()
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			utils.SendMessageWithStatus(context, "incorrect authorization header", 400)
			context.Abort()
			return
		}

		token := headerParts[1]
		claims, err := utils.ParseToken(token, config.JWTSecretKey)
		if err != nil {
			utils.SendMessageWithStatus(context, err.Error(), 401)
			context.Abort()
			return
		}

		context.Set("claims", claims)
		context.Next()
	}
}

func PostDeletionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			utils.SendMessageWithStatus(context, "Authorize!", 401)
			context.Abort()
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			utils.SendMessageWithStatus(context, "Incorrect authorization header", 401)
			context.Abort()
			return
		}

		token := headerParts[1]
		claims, err := utils.ParseToken(token, config.JWTSecretKey)
		if err != nil {
			utils.SendMessageWithStatus(context, err.Error(), 401)
			context.Abort()
			return
		}

		postID, err := strconv.Atoi(context.Param("post_id"))
		if err != nil {
			utils.SendMessageWithStatus(context, err.Error(), 401)
			context.Abort()
			return
		}

		dbConn := db.GetDBHandler()
		which, _ := post.GetPostByID(postID, dbConn)

		if claims.ID != which.UserID {
			utils.SendMessageWithStatus(context, "You are not allowed to delete someone's post!", 403)
			context.Abort()
			return
		}

		context.Next()
	}
}

func CommentDeletionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			utils.SendMessageWithStatus(context, "Authorize!", 401)
			context.Abort()
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			utils.SendMessageWithStatus(context, "Incorrect authorization header", 401)
			context.Abort()
			return
		}

		token := headerParts[1]
		claims, err := utils.ParseToken(token, config.JWTSecretKey)
		if err != nil {
			utils.SendMessageWithStatus(context, err.Error(), 401)
			context.Abort()
			return
		}

		commentID, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			utils.SendMessageWithStatus(context, err.Error(), 401)
			context.Abort()
			return
		}

		dbConn := db.GetDBHandler()
		which, _ := comment.GetCommentByID(commentID, dbConn)

		if claims.ID != which.UserID {
			utils.SendMessageWithStatus(context, "You are not allowed to delete someone's comment!", 403)
			context.Abort()
			return
		}

		context.Next()
	}
}
