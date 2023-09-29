package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/protos/pb"
	repositories "github.com/zhosyaaa/foodDeliverySystems-auth/internal/repositories"
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	repo repositories.AuthService
}

func (s *Service) RegisterUser(context context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("error hashing password: " + err.Error()) // Исправленная строка
	}
	user := models.User{
		Username: req.Username,
		Password: password,
		Email:    req.Email,
		UserRole: req.UserRole,
		Location: models.Location{
			Latitude:  float64(req.Location.Latitude),
			Longitude: float64(req.Location.Longitude),
		},
	}
	_, err = s.repo.CreateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
	}
	accessToken, refreshToken, err := utils.CreateOAuth2Tokens(strconv.Itoa(int(user.ID)), user.Email, user.UserRole)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create JWT token: %v", err)
	}
	return &pb.RegisterUserResponse{Response: &pb.Response{Status: 200, Error: ""}, AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *Service) Login(context context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	email := req.Email
	password := req.Password
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.Errorf("error getting user", err)
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.Errorf("incorrect password", err)
	}
	accessToken, refreshToken, err := utils.CreateOAuth2Tokens(strconv.Itoa(int(user.ID)), user.Email, user.UserRole)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create JWT token: %v", err)
	}
	return &pb.LoginResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) AuthenticateUser(context context.Context, req *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateUser not implemented")
}
