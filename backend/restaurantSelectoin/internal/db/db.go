package db

import (
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/config"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/models"
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

	db.AutoMigrate(&models.Restaurant{}, &models.Location{}, &models.Menu{}, &models.Dish{}, &models.Category{})
	return &DB{
		DB: db,
	}, nil
}
