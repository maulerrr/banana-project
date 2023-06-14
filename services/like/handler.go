package like

import (
	"encoding/json"
	"fmt"
	"github.com/maulerrr/banana/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maulerrr/banana/pkg/rabbitmq"
)

type Handler struct {
	Service  *Service
	Consumer *rabbitmq.Consumer
}

func NewLikeHandler(service *Service, consumer *rabbitmq.Consumer) *Handler {
	return &Handler{
		Service:  service,
		Consumer: consumer,
	}
}

func (h *Handler) GetLikeHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	claims, _ := c.Get("claims")
	userClaims := claims.(*models.Claims)

	like, err := h.Service.GetLike(models.Like{UserID: userClaims.ID, PostID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get like"})
		return
	}

	c.JSON(http.StatusOK, like)
}

func (h *Handler) GetAllLikesHandler(c *gin.Context) {
	likes, err := h.Service.GetAllLikes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get likes"})
		return
	}

	c.JSON(http.StatusOK, likes)
}

func (h *Handler) GetLikesCountOnPostHandler(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	count, err := h.Service.GetLikesCountOnPost(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get likes count"})
		return
	}

	c.JSON(http.StatusOK, count)
}

func (h *Handler) LikeHandler(c *gin.Context) {
	var requestBody struct {
		UserID int `json:"user_id"`
		PostID int `json:"post_id"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	like, err := h.Service.Like(requestBody.UserID, requestBody.PostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like"})
		return
	}

	c.JSON(http.StatusOK, like)

	// Handle RabbitMQ message received asynchronously
	go h.handleRabbitMQMessages(c)
}

func (h *Handler) handleRabbitMQMessages(c *gin.Context) {
	err := h.Consumer.Consume(func(body []byte) bool {
		message := &LikeMessage{}
		err := json.Unmarshal(body, message)
		if err != nil {
			fmt.Printf("Failed to unmarshal RabbitMQ message: %v\n", err)
			return false
		}

		// Process the received message
		fmt.Printf("Received RabbitMQ message: %+v\n", message)

		c.JSON(http.StatusAccepted, gin.H{"Hello": "Rabbit"})
		//TODO: implement conversation between handler and service layer through rabbitMQ, priority = 3

		return true
	})

	if err != nil {
		fmt.Printf("Failed to consume RabbitMQ messages: %v\n", err)
	}
}
