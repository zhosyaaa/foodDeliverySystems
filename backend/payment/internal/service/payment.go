package service

import (
	"context"
	"github.com/zhosyaaa/foodDeliverySystems-payment/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-payment/internal/protos/pb"
	"github.com/zhosyaaa/foodDeliverySystems-payment/internal/repositories"
)

type Service struct {
	pb.UnimplementedPaymentServiceServer
	repo repositories.PaymentService
}

func (s *Service) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.Response, error) {
	paymentInfo := models.ProcessPaymentRequest{
		UserID:        uint(req.UserID),
		PaymentMethod: req.PaymentMethod,
		Amount:        req.Amount,
	}

	_, err := s.repo.ProcessPayment(&paymentInfo)
	if err != nil {
		return &pb.Response{
			Status: 500,
			Error:  "Error processing the payment",
		}, err
	}

	return &pb.Response{
		Status: 200,
		Error:  "",
	}, nil
}

//	func (s *Service) GenerateInvoice(ctx context.Context, req *pb.GenerateInvoiceRequest) (*pb.Response, error) {
//		invoiceData := req.GetInvoiceData()
//
//		// Assuming you have a repository method to generate an invoice
//		err := s.repo.GenerateInvoice(invoiceData)
//		if err != nil {
//			return &pb.Response{
//				Status: 500,
//				Error:  "Error generating the invoice",
//			}, err
//		}
//
//		// Return a success response if the invoice was generated successfully
//		return &pb.Response{
//			Status: 200,
//			Error:  "",
//		}, nil
//	}
func (s *Service) GetTransactionHistory(ctx context.Context, req *pb.GetTransactionHistoryRequest) (*pb.GetTransactionHistoryResponse, error) {
	user := models.GetTransactionHistoryRequest{
		CustomerID: req.CustomerId,
	}
	responseFromRepo, err := s.repo.GetTransactionHistory(&user)
	if err != nil {
		return nil, err
	}
	transactions := make([]*pb.Transaction, len(responseFromRepo.Transactions))
	for i, transaction := range responseFromRepo.Transactions {
		transactions[i] = &pb.Transaction{
			Id:     transaction.ID,
			UserID: transaction.UserID,
			Amount: transaction.Amount,
			Type:   transaction.Type,
			Status: transaction.Status,
		}
	}

	response := &pb.GetTransactionHistoryResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
		Transactions: transactions,
	}

	return response, nil
}
