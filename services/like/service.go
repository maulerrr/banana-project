package like

import (
	"encoding/json"
	"fmt"
	"github.com/maulerrr/banana/pkg/models"

	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/rabbitmq"
)

type Service struct {
	DBHandler *db.DBHandler
	Producer  *rabbitmq.Producer
}

func NewLikeService(dbHandler *db.DBHandler, producer *rabbitmq.Producer) *Service {
	return &Service{
		DBHandler: dbHandler,
		Producer:  producer,
	}
}

type LikeMessage struct {
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}

func (s *Service) GetLike(query models.Like) (*Response, error) {
	return GetLike(&query, s.DBHandler)
}

func (s *Service) GetAllLikes() ([]models.Like, error) {
	return GetAllLikes(s.DBHandler)
}

func (s *Service) GetLikesCountOnPost(postID int) (int, error) {
	return GetLikesCountOnPost(postID, s.DBHandler)
}

func (s *Service) Like(userID int, postID int) (*Response, error) {
	like := &models.Like{
		UserID: userID,
		PostID: postID,
	}
	resp, err := Like(like, s.DBHandler)
	if err != nil {
		return nil, err
	}

	messageJSON, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal like message: %v", err)
	}

	// Publish the message to RabbitMQ
	err = s.Producer.Publish(messageJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to publish like message: %v", err)
	}

	return resp, nil
}
