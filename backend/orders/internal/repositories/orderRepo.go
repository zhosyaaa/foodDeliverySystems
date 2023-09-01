package repositories

import (
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type orderService struct {
	DB *gorm.DB
}

func NewOrderService(DB *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{DB: DB}
}
