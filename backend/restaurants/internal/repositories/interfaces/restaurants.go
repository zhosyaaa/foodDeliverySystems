package interfaces

import "github.com/zhosyaaa/foodDeliverySystems-restaurants/internal/models"

type RestaurantsRepository interface {
	AddDish(dish *models.Dish) error
	UpdateDish(dish *models.Dish) error
	DeleteDish(dish *models.Dish) error
	GetMenu(restId uint64) (*models.Menu, error)
	GetDishDetails(dishId uint64) (*models.Dish, error)
	UpdateDishIngredients(dishID uint64, ingredients []string) (*models.Dish, error)
	GetDishCategories(dishID uint64) ([]string, error)
	ToggleDishAvailability(dishID uint64, count int64) (*models.Dish, error)
	UploadDishImages(dishID uint64, img []string) (*models.Dish, error)
	GetOrder(orderId uint64) (*models.Order, error)
	UpdateOrderStatus(orderId uint64, status string) (*models.Order, error)
}
