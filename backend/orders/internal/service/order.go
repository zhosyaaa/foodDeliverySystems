package service

import (
	"context"
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/client"
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/protos/pb"
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/repositories"
)

type Service struct {
	pb.UnimplementedOrderServiceServer
	repo        repositories.OrderService
	DishService *client.DishServiceClient
}

func (s *Service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	restaurantID := req.RestaurantId
	customerID := req.CustomerId
	orderItems := req.Items

	var totalPrice float64
	for _, item := range orderItems {
		tempDish, _ := s.DishService.GetDishById(ctx, item.DishID)
		totalPrice += float64(item.Quantity) * tempDish.Dish.Price
	}

	order := models.Order{
		RestaurantID: restaurantID,
		CustomerID:   customerID,
		Items:        convertPBOrderItemsToModel(orderItems),
		TotalPrice:   totalPrice,
		Status:       "pending",
	}

	if err := s.repo.CreateOrder(&order); err != nil {
		return &pb.CreateOrderResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "error when creating an order",
			}, Order: nil,
		}, err
	}

	return &pb.CreateOrderResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		}, Order: convertModelOrderToPB(order),
	}, nil
}

func (s *Service) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.Response, error) {
	orderID := req.OrderId

	order, err := s.repo.GetOrderDetails(orderID)
	if err != nil {
		return &pb.Response{
			Status: 404,
			Error:  "Order not found",
		}, nil
	}
	if order.Status != "pending" {
		return &pb.Response{
			Status: 400,
			Error:  "Order cannot be canceled",
		}, nil
	}

	order.Status = "canceled"
	if err := s.repo.CancelOrder(order); err != nil {
		return &pb.Response{
			Status: 500,
			Error:  "Failed to cancel the order",
		}, err
	}

	return &pb.Response{
		Status: 200,
		Error:  "Order canceled successfully",
	}, nil
}

func (s *Service) GetUserOrders(ctx context.Context, req *pb.GetUserOrdersRequest) (*pb.GetUserOrdersResponse, error) {
	userID := req.CustomerId
	userOrders, err := s.repo.GetUserOrders(userID)
	if err != nil {
		return nil, err
	}
	var pbOrders []*pb.Order
	for _, userOrder := range userOrders {
		pbOrder := convertModelOrderToPB(userOrder)
		pbOrders = append(pbOrders, pbOrder)
	}

	userOrdersResponse := &pb.UserOrders{
		Orders: pbOrders,
	}

	response := &pb.GetUserOrdersResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		}, UserOrders: userOrdersResponse,
	}

	return response, nil
}

func (s *Service) GetOrderDetails(ctx context.Context, req *pb.GetOrderDetailsRequest) (*pb.GetOrderDetailsResponse, error) {
	orderID := req.OrderId
	orderDetails, err := s.repo.GetOrderDetails(orderID)
	if err != nil {
		return &pb.GetOrderDetailsResponse{
			Response: &pb.Response{
				Status: 404,
				Error:  "Order not found",
			},
			Order: nil,
		}, nil
	}

	pbOrderDetails := convertModelOrderToPB(*orderDetails)

	response := &pb.GetOrderDetailsResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
		Order: pbOrderDetails,
	}

	return response, nil
}

func convertPBOrderItemsToModel(items []*pb.OrderItem) []models.OrderItem {
	var convertedItems []models.OrderItem
	for _, item := range items {
		convertedItem := models.OrderItem{
			DishID:   item.DishID,
			Quantity: item.Quantity,
		}
		convertedItems = append(convertedItems, convertedItem)
	}
	return convertedItems
}

func convertModelOrderToPB(order models.Order) *pb.Order {
	pbOrder := &pb.Order{
		Id:           uint64(order.ID),
		RestaurantId: order.RestaurantID,
		CustomerId:   order.CustomerID,
		TotalPrice:   order.TotalPrice,
		Status:       order.Status,
		Items:        convertModelOrderItemsToPB(order.Items),
	}

	return pbOrder
}
func convertModelOrderItemsToPB(items []models.OrderItem) []*pb.OrderItem {
	var pbItems []*pb.OrderItem

	for _, item := range items {
		pbItem := &pb.OrderItem{
			DishID:   item.DishID,
			Quantity: int32(item.Quantity),
		}
		pbItems = append(pbItems, pbItem)
	}

	return pbItems
}
