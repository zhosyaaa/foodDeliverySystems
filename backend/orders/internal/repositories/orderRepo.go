package repositories

import (
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type OrderService struct {
	DB *gorm.DB
}

func newOrderService(DB *gorm.DB) interfaces.OrderRepository {
	return &OrderService{DB: DB}
}

func (o OrderService) CreateOrder(order *models.Order) error {
	if err := o.DB.Create(order).Error; err != nil {
		return err
	}

	return nil
}

func (o OrderService) CancelOrder(order *models.Order) error {
	if err := o.DB.Model(order).Update("Status", "Canceled").Error; err != nil {
		return err
	}
	return nil
}

func (o OrderService) GetUserOrders(userId uint64) ([]models.Order, error) {
	var orders []models.Order
	if err := o.DB.Where("customer_id = ?", userId).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (o OrderService) GetOrderDetails(orderId uint64) (*models.Order, error) {
	order := &models.Order{}
	if err := o.DB.Preload("Items").First(order, orderId).Error; err != nil {
		return nil, err
	}

	return order, nil
}
