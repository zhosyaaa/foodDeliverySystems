package repository

import (
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/api/models"
	interfaces "github/zhosyaaa/foodDeliverySystems-auth-service/pkg/api/repository/interface"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (u UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.DB.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepositoryImpl) CreateUser(user *models.User) error {
	return u.DB.Create(user).Error
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{DB: db}
}
