package interfaces

import "github.com/zhosyaaa/foodDeliverySystems-order/internal/models"

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	CancelOrder(order *models.Order) error
	GetUserOrders(userId uint64) ([]models.Order, error)
	GetOrderDetails(orderId uint64) (*models.Order, error)
}
