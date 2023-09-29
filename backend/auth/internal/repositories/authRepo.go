package auth

import (
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (a AuthService) CreateUser(user models.User) (models.User, error) {
	if err := a.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (a AuthService) GetUserById(id int) (models.User, error) {
	var user models.User
	if err := a.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (a AuthService) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := a.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func NewAuthService(DB *gorm.DB) interfaces.AuthRepository {
	return &AuthService{DB: DB}
}
