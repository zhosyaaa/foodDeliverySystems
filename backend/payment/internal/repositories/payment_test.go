package repositories

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhosyaaa/foodDeliverySystems-payment/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "postgresql://postgres:1079@localhost:6543/payment"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	return db
}

func clearTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM transactions")
}

func TestPaymentService_ProcessPayment(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	paymentService := NewPaymentService(db)

	request := &models.ProcessPaymentRequest{
		UserID: 1,
		Amount: 100.0,
	}

	transaction, err := paymentService.ProcessPayment(request)
	assert.NoError(t, err)
	assert.NotNil(t, transaction)

	var fetchedTransaction models.Transaction
	result := db.First(&fetchedTransaction, transaction.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, transaction.ID, fetchedTransaction.ID)
}

func TestPaymentService_GenerateInvoice(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	paymentService := NewPaymentService(db)

	request := &models.GenerateInvoiceRequest{
		OrderID: strconv.Itoa(1),
	}

	err := paymentService.GenerateInvoice(request)
	assert.NoError(t, err)

	// Проверяем, что счет был создан (в данной версии GenerateInvoiceRequest и Invoice не используются)
}

func TestPaymentService_GetTransactionHistory(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	paymentService := NewPaymentService(db)
	userID := 1
	for i := 1; i <= 3; i++ {
		_ = &models.Transaction{
			UserID: strconv.Itoa(userID),
			Amount: float64(i * 100),
			Type:   "Payment",
			Status: "Processed",
		}
		_, err := paymentService.ProcessPayment(&models.ProcessPaymentRequest{
			UserID: uint(userID),
			Amount: float64(i * 100),
		})
		assert.NoError(t, err)
	}

	request := &models.GetTransactionHistoryRequest{
		CustomerID: strconv.Itoa(int(uint(userID))),
	}
	transactionHistory, err := paymentService.GetTransactionHistory(request)
	assert.NoError(t, err)
	assert.NotNil(t, transactionHistory)
	assert.Len(t, transactionHistory.Transactions, 3)
}
