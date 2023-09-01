package db

import (
	"github.com/zhosyaaa/foodDeliverySystems-payment/internal/config"
	"github.com/zhosyaaa/foodDeliverySystems-payment/internal/models"
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

	db.AutoMigrate(&models.GenerateInvoiceRequest{},
		&models.GetTransactionHistoryRequest{},
		&models.Transaction{},
		&models.Response{},
		&models.ProcessPaymentRequest{},
		&models.GenerateInvoiceRequest{},
		&models.GetTransactionHistoryResponse{},
	)

	return &DB{
		DB: db,
	}, nil
}
