package db

import (
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/config"
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/models"
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

	db.AutoMigrate(&models.Order{}, &models.OrderItem{}, &models.Location{})

	return &DB{
		DB: db,
	}, nil
}
