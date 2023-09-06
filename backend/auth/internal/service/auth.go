package service

import (
	"context"
	"github.com/zhosyaaa/foodDeliverySystems-auth/internal/protos/pb"
	repositories "github.com/zhosyaaa/foodDeliverySystems-auth/internal/repositories"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	repo repositories.AuthService
}

func (s *Service) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.Response, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	//email := req.GetEmail()
	//userRole := req.GetUserRole()
	//location := req.GetLocation()
	err := s.repo.RegisterUser(username, password)
	if err != nil {
		return &pb.Response{
			Status: 500,
			Error:  "Error registering the user",
		}, err
	}
	return &pb.Response{
		Status: 200,
		Error:  "",
	}, nil
}

func (s *Service) AuthenticateUser(ctx context.Context, req *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	isAuthenticated, err := s.repo.AuthenticateUser(username, password)
	if err != nil {
		return nil, err
	}

	response := &pb.AuthenticateUserResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
	}

	token := "need to work on security"
	if isAuthenticated {
		response.Token = token
	}

	return response, nil
}
