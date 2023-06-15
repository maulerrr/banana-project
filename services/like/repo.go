package like

import (
	"errors"
	"fmt"

	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/models"
	"gorm.io/gorm"
)

type Response struct {
	Message string `json:"message"`
	Liked   bool   `json:"liked"`
}

func GetLike(query *models.Like, h *db.DBHandler) (*Response, error) {
	like := &models.Like{}
	err := h.DB.First(like, query).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Response{Message: "not liked", Liked: false}, nil
		}
		return nil, fmt.Errorf("failed to get like: %v", err)
	}
	return &Response{Message: "liked", Liked: true}, nil
}

func GetAllLikes(h *db.DBHandler) ([]models.Like, error) {
	var likes []models.Like
	err := h.DB.Find(&likes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get likes: %v", err)
	}
	return likes, nil
}

func GetLikesCountOnPost(postID int, h *db.DBHandler) (int, error) {
	var count int64
	err := h.DB.Model(&models.Like{}).Where("post_id = ?", postID).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("failed to get likes count: %v", err)
	}
	return int(count), nil
}

func Like(query *models.Like, h *db.DBHandler) (*Response, error) {
	like := &models.Like{}

	err := h.DB.First(like, query).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Like doesn't exist, create it
			err := h.DB.Create(query).Error
			if err != nil {
				return nil, fmt.Errorf("failed to create like: %v", err)
			}

			return &Response{
				Message: "liked",
				Liked:   true,
			}, nil
		}
		return nil, fmt.Errorf("failed to get like: %v", err)
	}

	// Like exists, delete it
	err = h.DB.Delete(like).Error
	if err != nil {
		return nil, fmt.Errorf("failed to delete like: %v", err)
	}

	return &Response{Message: "not liked", Liked: false}, nil
}
