package repositories

import (
	"github.com/zhosyaaa/foodDeliverySystems-payment/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-payment/internal/repositories/interfaces"
	"gorm.io/gorm"
	"strconv"
)

type PaymentService struct {
	DB *gorm.DB
}

func NewPaymentService(DB *gorm.DB) interfaces.PaymentRepository {
	return &PaymentService{DB: DB}
}

func (p PaymentService) ProcessPayment(request *models.ProcessPaymentRequest) (*models.Transaction, error) {
	transaction := &models.Transaction{
		UserID: strconv.Itoa(int(request.UserID)),
		Amount: request.Amount,
		Type:   "Payment",
		Status: "Processed",
	}

	if err := p.DB.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (p PaymentService) GenerateInvoice(request *models.GenerateInvoiceRequest) error {
	//invoice := &models.Invoice{
	//	OrderID: request.OrderID,
	//}
	//
	//if err := p.DB.Create(invoice).Error; err != nil {
	//	return err
	//}
	//
	return nil
}

func (p PaymentService) GetTransactionHistory(request *models.GetTransactionHistoryRequest) (*models.GetTransactionHistoryResponse, error) {
	var transactions []models.Transaction
	if err := p.DB.Where("user_id = ?", request.CustomerID).Find(&transactions).Error; err != nil {
		return nil, err
	}

	response := &models.GetTransactionHistoryResponse{
		Transactions: transactions,
	}

	return response, nil
}
