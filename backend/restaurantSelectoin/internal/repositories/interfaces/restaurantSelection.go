package interfaces

import "github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/models"

type RestaurantSelectionRepository interface {
	GetRestaurants() ([]models.Restaurant, error)
	GetRestaurantDetails(restId uint64) (*models.Restaurant, error)
	AddNewRestaurant(rest *models.Restaurant) error
	UpdateRestaurantInfo(rest *models.Restaurant) error
	SearchRestaurants(query string) ([]models.Restaurant, error)
	DeleteRestaurant(rest *models.Restaurant) error
}
