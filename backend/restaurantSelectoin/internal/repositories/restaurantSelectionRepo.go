package repositories

import (
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type RestaurantSelectionService struct {
	DB *gorm.DB
}

func NewRestaurantSelectionService(DB *gorm.DB) interfaces.RestaurantSelectionRepository {
	return &RestaurantSelectionService{DB: DB}
}

func (r RestaurantSelectionService) GetRestaurants() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	if err := r.DB.Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r RestaurantSelectionService) GetRestaurantDetails(restId uint64) (*models.Restaurant, error) {
	restaurant := &models.Restaurant{}
	if err := r.DB.Preload("Locations").Preload("Menu.Dishes.Categories").First(restaurant, restId).Error; err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (r RestaurantSelectionService) AddNewRestaurant(rest *models.Restaurant) error {
	if err := r.DB.Create(rest).Error; err != nil {
		return err
	}
	return nil
}

func (r RestaurantSelectionService) UpdateRestaurantInfo(rest *models.Restaurant) error {
	if err := r.DB.Save(rest).Error; err != nil {
		return err
	}
	return nil
}

func (r RestaurantSelectionService) SearchRestaurants(query string) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	if query != "" {
		if err := r.DB.Where("name LIKE ?", "%"+query+"%").Find(&restaurants).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.DB.Find(&restaurants).Error; err != nil {
			return nil, err
		}
	}
	return restaurants, nil
}

func (r RestaurantSelectionService) DeleteRestaurant(rest *models.Restaurant) error {
	if err := r.DB.Delete(rest).Error; err != nil {
		return err
	}
	return nil
}
