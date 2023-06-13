package db

import (
	"log"

	"github.com/maulerrr/banana/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

type DBHandler struct {
	DB *gorm.DB
}

func InitDB(url string) DBHandler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalf("error occured while opening db conn: %s", err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		//Todo: add migration models
	)

	dbConn = db

	return DBHandler{db}
}

func GetDBHandler() *DBHandler {
	if dbConn == nil {
		log.Fatal("database connection is not initialized")
	}

	return &DBHandler{DB: dbConn}
}
