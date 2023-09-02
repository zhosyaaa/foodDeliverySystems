package repositories

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "postgresql://postgres:1079@localhost:8008/restSelection"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	return db
}

func clearTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM restaurants")
}

func TestRestaurantSelectionService_GetRestaurants(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	restaurantService := NewRestaurantSelectionService(db)

	for i := 1; i <= 3; i++ {
		var (
			restaurant = &models.Restaurant{
				Name:        "test",
				Description: "tesst",
				Locations:   nil,
				Rating:      5,
			}
		)
		err := restaurantService.AddNewRestaurant(restaurant)
		assert.NoError(t, err)
	}

	restaurants, err := restaurantService.GetRestaurants()
	assert.NoError(t, err)
	assert.Len(t, restaurants, 3)
}

func TestRestaurantSelectionService_GetRestaurantDetails(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	restaurantService := NewRestaurantSelectionService(db)

	location := models.Location{
		Restaurants: nil,
		City:        "as",
		Country:     "sfas",
		PostalCode:  "132",
	}
	locations := make([]models.Location, 1)
	locations = append(locations, location)
	restaurant := &models.Restaurant{
		Name:        "test",
		Description: "tesst",
		Locations:   locations,
		Rating:      5,
	}
	err := restaurantService.AddNewRestaurant(restaurant)
	assert.NoError(t, err)

	restaurantDetails, err := restaurantService.GetRestaurantDetails(uint64(restaurant.ID))
	assert.NoError(t, err)
	assert.NotNil(t, restaurantDetails)
}

func TestRestaurantSelectionService_AddNewRestaurant(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	restaurantService := NewRestaurantSelectionService(db)

	location := models.Location{
		Restaurants: nil,
		City:        "as",
		Country:     "sfas",
		PostalCode:  "132",
	}
	locations := make([]models.Location, 1)
	locations = append(locations, location)
	restaurant := &models.Restaurant{
		Name:        "test",
		Description: "tesst",
		Locations:   locations,
		Rating:      5,
	}

	err := restaurantService.AddNewRestaurant(restaurant)
	assert.NoError(t, err)

	var fetchedRestaurant models.Restaurant
	result := db.First(&fetchedRestaurant, restaurant.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, restaurant.ID, fetchedRestaurant.ID)
}

func TestRestaurantSelectionService_UpdateRestaurantInfo(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	restaurantService := NewRestaurantSelectionService(db)

	location := models.Location{
		Restaurants: nil,
		City:        "as",
		Country:     "sfas",
		PostalCode:  "132",
	}
	locations := make([]models.Location, 1)
	locations = append(locations, location)
	restaurant := &models.Restaurant{
		Name:        "test",
		Description: "tesst",
		Locations:   locations,
		Rating:      5,
	}
	err := restaurantService.AddNewRestaurant(restaurant)
	assert.NoError(t, err)

	updatedRestaurant := &models.Restaurant{
		Name: "Updated Restaurant Name",
	}
	err = restaurantService.UpdateRestaurantInfo(updatedRestaurant)
	assert.NoError(t, err)

	var fetchedRestaurant models.Restaurant
	result := db.First(&fetchedRestaurant, restaurant.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, updatedRestaurant.Name, fetchedRestaurant.Name)
}

func TestRestaurantSelectionService_SearchRestaurants(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	restaurantService := NewRestaurantSelectionService(db)

	location := models.Location{
		Restaurants: nil,
		City:        "as",
		Country:     "sfas",
		PostalCode:  "132",
	}
	locations := make([]models.Location, 1)
	locations = append(locations, location)
	for i := 1; i <= 3; i++ {
		restaurant := &models.Restaurant{
			Name:        "Restaurant" + strconv.Itoa(i),
			Description: "tesst",
			Locations:   locations,
			Rating:      5,
		}
		err := restaurantService.AddNewRestaurant(restaurant)
		assert.NoError(t, err)
	}

	query := "Restaurant2"
	restaurants, err := restaurantService.SearchRestaurants(query)
	assert.NoError(t, err)
	assert.Len(t, restaurants, 1)
	assert.Equal(t, query, restaurants[0].Name)
}

func TestRestaurantSelectionService_DeleteRestaurant(t *testing.T) {
	db := setupTestDB(t)
	clearTestDB(db)

	restaurantService := NewRestaurantSelectionService(db)

	location := models.Location{
		Restaurants: nil,
		City:        "as",
		Country:     "sfas",
		PostalCode:  "132",
	}
	locations := make([]models.Location, 1)
	locations = append(locations, location)
	restaurant := &models.Restaurant{
		Name:        "test",
		Description: "tesst",
		Locations:   locations,
		Rating:      5,
	}
	err := restaurantService.AddNewRestaurant(restaurant)
	assert.NoError(t, err)

	err = restaurantService.DeleteRestaurant(restaurant)
	assert.NoError(t, err)

	var deletedRestaurant models.Restaurant
	result := db.First(&deletedRestaurant, restaurant.ID)
	assert.Error(t, result.Error)
}
