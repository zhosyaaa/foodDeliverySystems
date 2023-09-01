package repositories

import (
	"github.com/zhosyaaa/foodDeliverySystems-restaurants/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-restaurants/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type RestaurantsService struct {
	DB *gorm.DB
}

func NewRestaurantsService(DB *gorm.DB) interfaces.RestaurantsRepository {
	return &RestaurantsService{DB: DB}
}

func (r RestaurantsService) AddDish(dish *models.Dish) error {
	if err := r.DB.Create(dish).Error; err != nil {
		return err
	}

	restaurant := models.Menu{}
	if err := r.DB.First(&restaurant, dish.RestaurantID).Error; err != nil {
		return err
	}

	restaurant.Dishes = append(restaurant.Dishes, *dish)
	if err := r.DB.Save(&restaurant).Error; err != nil {
		return err
	}
	return nil
}

func (r RestaurantsService) UpdateDish(dish *models.Dish) error {
	if err := r.DB.Save(dish).Error; err != nil {
		return err
	}
	restaurant := models.Menu{}
	if err := r.DB.First(&restaurant, dish.RestaurantID).Error; err != nil {
		return err
	}
	for i, d := range restaurant.Dishes {
		if d.ID == dish.ID {
			restaurant.Dishes[i] = *dish
			break
		}
	}

	if err := r.DB.Save(&restaurant).Error; err != nil {
		return err
	}

	return nil
}

func (r RestaurantsService) DeleteDish(dish *models.Dish) error {
	if err := r.DB.Delete(dish).Error; err != nil {
		return err
	}

	restaurant := models.Menu{}
	if err := r.DB.First(&restaurant, dish.RestaurantID).Error; err != nil {
		return err
	}

	var updatedDishes []models.Dish
	for _, d := range restaurant.Dishes {
		if d.ID != dish.ID {
			updatedDishes = append(updatedDishes, d)
		}
	}

	restaurant.Dishes = updatedDishes
	if err := r.DB.Save(&restaurant).Error; err != nil {
		return err
	}

	return nil
}

func (r RestaurantsService) GetMenu(restId uint64) (*models.Menu, error) {
	restaurant := &models.Menu{}
	if err := r.DB.Preload("Dishes").First(restaurant, restId).Error; err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (r RestaurantsService) GetDishDetails(dishId uint64) (*models.Dish, error) {
	dish := &models.Dish{}
	if err := r.DB.First(dish, dishId).Error; err != nil {
		return nil, err
	}
	return dish, nil
}

func (r RestaurantsService) UpdateDishIngredients(dishID uint64, ingredients []string) (*models.Dish, error) {
	dish := &models.Dish{}
	if err := r.DB.Preload("Ingredients").First(dish, dishID).Error; err != nil {
		return nil, err
	}

	dish.Ingredients = ingredients
	if err := r.UpdateDish(dish); err != nil {
		return nil, err
	}
	return dish, nil
}

func (r RestaurantsService) GetDishCategories(dishID uint64) ([]string, error) {
	dish := &models.Dish{}
	if err := r.DB.Preload("Categories").First(dish, dishID).Error; err != nil {
		return nil, err
	}

	var categories []string
	for _, category := range dish.Categories {
		categories = append(categories, category.Name)
	}

	return categories, nil
}

func (r RestaurantsService) ToggleDishAvailability(dishID uint64, count int64) (*models.Dish, error) {
	dish := &models.Dish{}
	if err := r.DB.First(dish, dishID).Error; err != nil {
		return nil, err
	}

	dish.Availability = uint64(count)

	if err := r.DB.Save(dish).Error; err != nil {
		return nil, err
	}

	return dish, nil
}

func (r RestaurantsService) UploadDishImages(dishID uint64, img []string) (*models.Dish, error) {
	dish := &models.Dish{}
	if err := r.DB.First(dish, dishID).Error; err != nil {
		return nil, err
	}

	dish.Images = img

	if err := r.DB.Save(dish).Error; err != nil {
		return nil, err
	}

	return dish, nil
}

func (r RestaurantsService) GetOrder(orderId uint64) (*models.Order, error) {
	order := &models.Order{}
	if err := r.DB.Preload("Items").First(order, orderId).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r RestaurantsService) UpdateOrderStatus(orderId uint64, status string) (*models.Order, error) {
	order := &models.Order{}
	if err := r.DB.First(order, orderId).Error; err != nil {
		return nil, err
	}

	order.Status = status

	if err := r.DB.Save(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
