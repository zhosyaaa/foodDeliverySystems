package db

import (
	"github.com/zhosyaaa/foodDeliverySystems-restaurants/internal/config"
	"github.com/zhosyaaa/foodDeliverySystems-restaurants/internal/models"
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

	db.AutoMigrate(&models.Category{}, &models.Menu{}, &models.Dish{}, &models.Order{}, &models.OrderItem{}, &models.Location{})

	return &DB{
		DB: db,
	}, nil
}
