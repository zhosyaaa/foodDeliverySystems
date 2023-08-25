package repository

import (
	"gorm.io/gorm"
)

type RestaurantsService struct {
	DB *gorm.DB
}

//
//func NewRestaurantsService(DB *gorm.DB) interfaces.Repository {
//	return &RestaurantsService{DB: DB}
//}
