package interfaces

import "github.com/zhosyaaa/foodDeliverySystems-auth/internal/models"

type AuthRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserById(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}
