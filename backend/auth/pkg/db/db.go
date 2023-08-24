package db

import (
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/config"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func InitDB(cfg *config.Config) (*DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgreDsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return &DB{
		DB: db,
	}, nil
}
