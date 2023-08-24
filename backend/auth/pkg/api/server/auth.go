package server

import (
	"context"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/api/models"
	interfaces "github/zhosyaaa/foodDeliverySystems-auth-service/pkg/api/repository/interface"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/db"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/pb"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/utils"
	"net/http"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	DB          *db.DB
	Jwt         utils.JwtWrapper
	userService interfaces.UserRepository
}

var _ pb.AuthServiceServer = &Service{}

func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := s.userService.GetUserByEmail(req.User.Email)
	if err == nil {
		return &pb.RegisterResponse{
			Response: &pb.Response{
				Status: http.StatusConflict,
				Error:  "Email already exists",
			},
		}, nil
	}

	user.Email = req.User.Email
	hashed, err := utils.HashPassword(req.User.Password)
	if err != nil {
		return &pb.RegisterResponse{
			Response: &pb.Response{
				Error:  "Failed to prepare user data",
				Status: http.StatusInternalServerError,
			},
		}, nil
	}
	user.Password = hashed

	err = s.userService.CreateUser(user)
	if err != nil {
		return &pb.RegisterResponse{
			Response: &pb.Response{
				Error:  "Failed to prepare user data",
				Status: http.StatusInternalServerError,
			},
		}, nil
	}

	return &pb.RegisterResponse{
		Response: &pb.Response{
			Status: http.StatusCreated,
		},
		User: utils.UserModelsToUserPb(user),
	}, nil
}

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userService.GetUserByEmail(req.Email)
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Status: http.StatusNotFound,
				Error:  "User not found",
			},
		}, nil
	}

	if !utils.ComparePasswordAndHash(user.Password, req.Password) {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Status: http.StatusNotFound,
				Error:  "User not found",
			},
		}, nil
	}

	token, err := s.Jwt.GenerateToken(*user)
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				Error:  "Failed to prepare response",
				Status: http.StatusInternalServerError,
			},
		}, nil
	}

	return &pb.LoginResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
		Token: token,
	}, nil
}

func (s *Service) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	claims, err := s.Jwt.VerifyToken(req.Token)
	if err != nil {
		return &pb.AuthenticateResponse{
			Response: &pb.Response{
				Error:  err.Error(),
				Status: http.StatusBadRequest,
			},
		}, nil
	}

	var user *models.User
	user, err = s.userService.GetUserByEmail(claims.Email)
	if err != nil {
		return &pb.AuthenticateResponse{
			Response: &pb.Response{
				Error:  "User not found",
				Status: http.StatusNotFound,
			},
		}, nil
	}

	return &pb.AuthenticateResponse{
		Response: &pb.Response{
			Status: http.StatusOK,
		},
		UserId: user.UserId,
	}, nil
}
