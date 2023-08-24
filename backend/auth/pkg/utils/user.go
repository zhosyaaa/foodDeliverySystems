package utils

import (
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/api/models"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/pb"
)

func UserPbToUserModels(user *pb.User) *models.User {
	return &models.User{
		UserId:   user.UserId,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UserModelsToUserPb(user *models.User) *pb.User {
	return &pb.User{
		UserId:   user.UserId,
		Email:    user.Email,
		Password: user.Password,
		Fullname: user.Fullname,
	}
}
