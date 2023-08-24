package _interface

import "github/zhosyaaa/foodDeliverySystems-auth-service/pkg/api/models"

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}
