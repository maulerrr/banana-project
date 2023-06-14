package comment

import (
	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/models"
	"gorm.io/gorm"
)

func GetAllCommentsByPostID(postID int, handler *db.DBHandler) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := handler.DB.Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func DeleteComment(commentID int, handler *db.DBHandler) error {
	comment := &models.Comment{CommentID: commentID}
	err := handler.DB.Delete(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateComment(comment *models.Comment, handler *db.DBHandler) (*models.Comment, error) {
	err := handler.DB.Create(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func GetCommentByID(commentID int, handler *db.DBHandler) (models.Comment, error) {
	comment := models.Comment{}
	result := handler.DB.First(&comment, "comment_id = ?", commentID)

	if result.Error != nil {
		return comment, result.Error
	}

	if result.RowsAffected == 0 {
		return comment, gorm.ErrRecordNotFound
	}
	return comment, nil
}
