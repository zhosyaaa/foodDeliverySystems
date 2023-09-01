package interfaces

import "github.com/zhosyaaa/foodDeliverySystems-payment/internal/models"

type PaymentRepository interface {
	ProcessPayment(request *models.ProcessPaymentRequest) (*models.Transaction, error)
	GenerateInvoice(request *models.GenerateInvoiceRequest) error
	GetTransactionHistory(request *models.GetTransactionHistoryRequest) (*models.GetTransactionHistoryResponse, error)
}
