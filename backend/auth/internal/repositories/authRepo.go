package repositories

import (
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(DB *gorm.DB) interfaces.AuthRepository {
	return &AuthService{DB: DB}
}

func (a AuthService) RegisterUser(username string, password string) error {

	user := &models.User{
		Username: username,
		Password: password,
	}
	if err := a.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (a AuthService) AuthenticateUser(username string, password string) (bool, error) {
	var user models.User
	if err := a.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	if user.Password == password {
		return true, nil
	}

	return false, nil
}
