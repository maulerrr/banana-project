package post

import (
	"errors"
	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/models"
	"gorm.io/gorm"
)

func CreatePost(post *models.Post, handler *db.DBHandler) (*models.Post, error) {
	err := handler.DB.Create(post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func GetAllPosts(handler *db.DBHandler) ([]*models.Post, error) {
	var posts []*models.Post

	err := handler.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func GetPostByID(postID int, handler *db.DBHandler) (*models.Post, error) {
	post := &models.Post{}

	err := handler.DB.First(post, postID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("post not found")
	}

	if err != nil {
		return nil, err
	}

	return post, nil
}

func DeletePost(postID int, handler *db.DBHandler) error {
	err := handler.DB.Delete(&models.Post{}, postID).Error
	if err != nil {
		return err
	}

	return nil
}
