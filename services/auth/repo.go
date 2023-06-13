package auth

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/maulerrr/banana/pkg/db"
	"github.com/maulerrr/banana/pkg/models"
	"github.com/maulerrr/banana/pkg/utils"
	"gorm.io/gorm"
	"time"
)

func GetUserByEmail(email string, handler db.DBHandler) (*models.User, error) {
	user := &models.User{}
	query := models.User{Email: email}
	err := handler.DB.First(user, &query).Error

	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func CreateUser(registration *models.User, handler *db.DBHandler) (*models.User, error) {
	if len(registration.Password) < 4 {
		return nil, errors.New("minimum password length is 4")
	}

	password := utils.HashPassword([]byte(registration.Password))
	err := checkmail.ValidateFormat(registration.Email)
	if err != nil {
		return nil, errors.New("invalid email address")
	}

	user := &models.User{
		Password:  password,
		Email:     registration.Email,
		Username:  registration.Username,
		CreatedAt: time.Now(),
	}

	err = handler.DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
