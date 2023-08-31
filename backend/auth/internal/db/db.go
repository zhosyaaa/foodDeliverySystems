package db

import (
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/config"
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/models"
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

	db.AutoMigrate(&models.User{}, &models.Location{})
	return &DB{
		DB: db,
	}, nil
}
