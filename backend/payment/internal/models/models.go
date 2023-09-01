package models

import "gorm.io/gorm"

type ProcessPaymentRequest struct {
	gorm.Model
	UserID        uint    `json:"userID"`
	Amount        float64 `gorm:"column:amount"`
	PaymentMethod string  `gorm:"column:payment_method"`
}

type GenerateInvoiceRequest struct {
	gorm.Model
	OrderID string `gorm:"column:order_id"`
}

type GetTransactionHistoryRequest struct {
	gorm.Model
	CustomerID string `gorm:"column:customer_id"`
}

type Transaction struct {
	gorm.Model
	ID     string  `gorm:"column:id"`
	UserID string  `gorm:"column:user_id"`
	Amount float64 `gorm:"column:amount"`
	Type   string  `gorm:"column:type"`
	Status string  `gorm:"column:status"`
}

type GetTransactionHistoryResponse struct {
	gorm.Model
	Transactions []Transaction `gorm:"foreignKey:ResponseID"`
	Response     Response      `gorm:"embedded"`
}

type Response struct {
	gorm.Model
	Error  string `gorm:"column:error"`
	Status uint32 `gorm:"column:status"`
}
