package repositories

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "postgresql://postgres:1079@localhost:8000/orders"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	return db
}

func clearTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM orders")
}

func TestOrderService_CreateOrder(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	orderService := newOrderService(db)

	order := &models.Order{
		ID:           1,
		CustomerID: strconv.Itoa(1),
		RestaurantID: 2,
	}

	err := orderService.CreateOrder(order)
	assert.NoError(t, err)

	var fetchedOrder models.Order
	result := db.First(&fetchedOrder, order.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, order.ID, fetchedOrder.ID)
}

func TestOrderService_CancelOrder(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	orderService := newOrderService(db)

	order := &models.Order{
		ID:           1,
		CustomerID: strconv.Itoa(1),
		RestaurantID: 2,
	}

	err := orderService.CreateOrder(order)
	assert.NoError(t, err)

	err = orderService.CancelOrder(order)
	assert.NoError(t, err)

	var updatedOrder models.Order
	result := db.First(&updatedOrder, order.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Canceled", updatedOrder.Status)
}

func TestOrderService_GetUserOrders(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	orderService := newOrderService(db)

	userID := uint64(1)
	for i := 1; i <= 3; i++ {
		order := &models.Order{
			CustomerID: strconv.FormatUint(userID, 10),
			ID:           1,
			RestaurantID: 2,
		}
		err := orderService.CreateOrder(order)
		assert.NoError(t, err)
	}

	orders, err := orderService.GetUserOrders(userID)
	assert.NoError(t, err)
	assert.Len(t, orders, 3)
}

func TestOrderService_GetOrderDetails(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	orderService := newOrderService(db)


	order := &models.Order{
		ID:           1,
		CustomerID: strconv.Itoa(1),
		RestaurantID: 2,
	}
	err := orderService.CreateOrder(order)
	assert.NoError(t, err)

	orderDetails, err := orderService.GetOrderDetails(uint64(order.ID))
	assert.NoError(t, err)
	assert.NotNil(t, orderDetails)
}
