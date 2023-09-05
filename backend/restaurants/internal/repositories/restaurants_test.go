package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhosyaaa/foodDeliverySystems-restaurants/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TestRestaurantsService is a test suite for RestaurantsService methods.
func TestRestaurantsService(t *testing.T) {
	// Replace the following values with your PostgreSQL connection details.
	dsn := "postgresql://postgres:1079@localhost:5432/restaurants"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}

	service := NewRestaurantsService(db)

	// Test AddDish method.
	t.Run("Test AddDish", func(t *testing.T) {
		dish := &models.Dish{
			RestaurantID: 1,
			Name:         "Test Dish",
			Description:  "Test Description",
			Price:        9.99,
		}

		err := service.AddDish(dish)
		assert.NoError(t, err)
		// Verify that the dish has been added to the database.
		var fetchedDish models.Dish
		db.First(&fetchedDish, dish.ID)
		assert.Equal(t, dish.Name, fetchedDish.Name)
	})

	// Test UpdateDish method.
	t.Run("Test UpdateDish", func(t *testing.T) {
		dish := &models.Dish{
			RestaurantID: 1,
			Name:         "Update Test Dish",
			Description:  "Update Test Description",
			Price:        19.99,
		}

		err := service.UpdateDish(dish)
		assert.NoError(t, err)

		// Verify that the dish has been updated in the database.
		var updatedDish models.Dish
		db.First(&updatedDish, dish.ID)
		assert.Equal(t, dish.Name, updatedDish.Name)
	})

	// Test DeleteDish method.
	t.Run("Test DeleteDish", func(t *testing.T) {
		dish := &models.Dish{
			RestaurantID: 1,
			Name:         "Delete Test Dish",
			Description:  "Delete Test Description",
			Price:        29.99,
		}

		err := service.AddDish(dish)
		assert.NoError(t, err)

		err = service.DeleteDish(dish)
		assert.NoError(t, err)

		// Verify that the dish has been deleted from the database.
		var deletedDish models.Dish
		err = db.First(&deletedDish, dish.ID).Error
		assert.Error(t, err) // Dish should not exist in the database.
	})

	// Test GetMenu method.
	t.Run("Test GetMenu", func(t *testing.T) {
		restID := uint64(1)

		menu, err := service.GetMenu(restID)
		assert.NoError(t, err)
		assert.NotNil(t, menu)

		// You can add assertions to check menu contents if needed.
	})
	// Test GetDishDetails method.
	t.Run("Test GetDishDetails", func(t *testing.T) {
		// Assuming you have a dish ID from a previously added dish.
		dishID := uint64(1)

		dish, err := service.GetDishDetails(dishID)
		assert.NoError(t, err)
		assert.NotNil(t, dish)

		// You can add assertions to check dish details if needed.
	})

	// Test UpdateDishIngredients method.
	t.Run("Test UpdateDishIngredients", func(t *testing.T) {
		// Assuming you have a dish ID from a previously added dish.
		dishID := uint64(1)
		newIngredients := []string{"Ingredient1", "Ingredient2"}

		updatedDish, err := service.UpdateDishIngredients(dishID, newIngredients)
		assert.NoError(t, err)
		assert.NotNil(t, updatedDish)

		// Verify that the dish's ingredients have been updated.
		assert.Equal(t, newIngredients, updatedDish.Ingredients)
	})

	// Test GetDishCategories method.
	t.Run("Test GetDishCategories", func(t *testing.T) {
		// Assuming you have a dish ID from a previously added dish.
		dishID := uint64(1)

		categories, err := service.GetDishCategories(dishID)
		assert.NoError(t, err)
		assert.NotNil(t, categories)

		// You can add assertions to check categories if needed.
	})

	// Test ToggleDishAvailability method.
	t.Run("Test ToggleDishAvailability", func(t *testing.T) {
		// Assuming you have a dish ID from a previously added dish.
		dishID := uint64(1)
		newAvailability := int64(10)

		updatedDish, err := service.ToggleDishAvailability(dishID, newAvailability)
		assert.NoError(t, err)
		assert.NotNil(t, updatedDish)

		// Verify that the dish's availability has been updated.
		assert.Equal(t, uint64(newAvailability), updatedDish.Availability)
	})

	// Test UploadDishImages method.
	t.Run("Test UploadDishImages", func(t *testing.T) {
		// Assuming you have a dish ID from a previously added dish.
		dishID := uint64(1)
		newImages := []string{"image1.jpg", "image2.jpg"}

		updatedDish, err := service.UploadDishImages(dishID, newImages)
		assert.NoError(t, err)
		assert.NotNil(t, updatedDish)

		// Verify that the dish's images have been updated.
		assert.Equal(t, newImages, updatedDish.Images)
	})

	// Test GetOrder method.
	t.Run("Test GetOrder", func(t *testing.T) {
		// Assuming you have an order ID from a previously added order.
		orderID := uint64(1)

		order, err := service.GetOrder(orderID)
		assert.NoError(t, err)
		assert.NotNil(t, order)

		// You can add assertions to check order details if needed.
	})

	// Test UpdateOrderStatus method.
	t.Run("Test UpdateOrderStatus", func(t *testing.T) {
		// Assuming you have an order ID from a previously added order.
		orderID := uint64(1)
		newStatus := "Delivered"

		updatedOrder, err := service.UpdateOrderStatus(orderID, newStatus)
		assert.NoError(t, err)
		assert.NotNil(t, updatedOrder)

		// Verify that the order's status has been updated.
		assert.Equal(t, newStatus, updatedOrder.Status)
	})

}
