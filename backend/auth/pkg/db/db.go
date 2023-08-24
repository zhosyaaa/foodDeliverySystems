package db

import (
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/api/models"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/config"
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
