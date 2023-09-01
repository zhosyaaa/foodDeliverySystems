package repositories

import (
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type orderService struct {
	DB *gorm.DB
}

func newOrderService(DB *gorm.DB) interfaces.OrderRepository {
	return &orderService{DB: DB}
}

func (o orderService) CreateOrder(order *models.Order) error {
	if err := o.DB.Create(order).Error; err != nil {
		return err
	}

	return nil
}

func (o orderService) CancelOrder(order *models.Order) error {
	if err := o.DB.Model(order).Update("Status", "Canceled").Error; err != nil {
		return err
	}
	return nil
}

func (o orderService) GetUserOrders(userId uint64) ([]models.Order, error) {
	var orders []models.Order
	if err := o.DB.Where("customer_id = ?", userId).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (o orderService) GetOrderDetails(orderId uint64) (*models.Order, error) {
	order := &models.Order{}
	if err := o.DB.Preload("Items").First(order, orderId).Error; err != nil {
		return nil, err
	}

	return order, nil
}
