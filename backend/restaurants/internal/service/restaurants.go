package service

import (
	"context"
	"github.com/zhosyaaa/foodDeliverySystems-restaurants/internal/models"
	pb "github.com/zhosyaaa/foodDeliverySystems-restaurants/internal/protos/pb"
	"github.com/zhosyaaa/foodDeliverySystems-restaurants/internal/repositories"
)

type Service struct {
	pb.UnimplementedRestaurantManagementServiceServer
	repo repositories.RestaurantsService
}

func (s *Service) AddDish(ctx context.Context, req *pb.CreateMenuItemRequest) (*models.Dish, error) {
	dish := &models.Dish{
		RestaurantID: req.Dish.RestaurantId,
		Name:         req.Dish.Name,
		Description:  req.Dish.Description,
		Availability: req.Dish.Availability,
		Price:        req.Dish.Price,
		Images:       req.Dish.Images,
		Ingredients:  req.Dish.Ingredients,
		Categories:   ConvertPbCategoriesToModels(req.Dish.Categories),
	}
	if err := s.repo.AddDish(dish); err != nil {
		return nil, err
	}
	return dish, nil

}

func (s *Service) UpdateDish(ctx context.Context, req *pb.UpdateDishRequest) (*pb.UpdateDishResponse, error) {
	dish, err := s.repo.GetDishDetails(req.Dish.ID)
	if err != nil {
		return &pb.UpdateDishResponse{
			Response: &pb.Response{
				Status: 404,
				Error:  "Dish not found",
			},
		}, err
	}
	if req.Dish.Name != "" {
		dish.Name = req.Dish.Name
	}
	if req.Dish.Description != "" {
		dish.Description = req.Dish.Description
	}
	if len(req.Dish.Images) != 0 {
		dish.Images = req.Dish.Images
	}
	if req.Dish.Price != 0 {
		dish.Price = req.Dish.Price
	}
	if len(req.Dish.Categories) != 0 {
		dish.Categories = ConvertPbCategoriesToModels(req.Dish.Categories)
	}
	if req.Dish.RestaurantId != 0 {
		dish.RestaurantID = req.Dish.RestaurantId
	}
	if req.Dish.Availability != dish.Availability {
		dish.Availability = req.Dish.Availability
	}
	if len(req.Dish.Ingredients) != 0 {
		dish.Ingredients = req.Dish.Ingredients
	}
	if err := s.repo.UpdateDish(dish); err != nil {
		return &pb.UpdateDishResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "Dish updated fail",
			},
		}, err
	}
	return &pb.UpdateDishResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "Dish updated successfully",
		},
	}, nil
}

func (s *Service) DeleteDish(ctx context.Context, req *pb.DeleteDishRequest) (*pb.DeleteDishResponse, error) {
	dish, err := s.repo.GetDishDetails(req.Dish.ID)
	if err != nil {
		return &pb.DeleteDishResponse{
			Response: &pb.Response{
				Status: 404,
				Error:  "Dish not found",
			},
		}, err
	}

	if err := s.repo.DeleteDish(dish); err != nil {
		return &pb.DeleteDishResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "Failed to delete dish",
			},
		}, err
	}

	return &pb.DeleteDishResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "Dish deleted successfully",
		},
	}, nil
}

//
//func (s *Service) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (*pb.Menu, error) {
//
//}
//func (s *Service) GetDishDetails(context.Context, *GetDishDetailsRequest) (*Dish, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetDishDetails not implemented")
//}
//func (s *Service) UpdateDishIngredients(context.Context, *UpdateDishIngredientsRequest) (*UpdateDishResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method UpdateDishIngredients not implemented")
//}
//func (s *Service) GetDishCategories(context.Context, *GetDishDetailsRequest) (*GetDishCategoriesResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetDishCategories not implemented")
//}
//func (s *Service) ToggleDishAvailability(context.Context, *ToggleDishAvailabilityRequest) (*ToggleDishAvailabilityResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method ToggleDishAvailability not implemented")
//}
//func (s *Service) UploadDishImages(context.Context, *UploadDishImagesRequest) (*UploadDishImagesResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method UploadDishImages not implemented")
//}
//func (s *Service) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
//}
//func (s *Service) UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
//}

func ConvertPbCategoriesToModels(pbCategories []*pb.Category) []*models.Category {
	var modelsCategories []*models.Category
	for _, pbCategory := range pbCategories {
		modelsCategory := &models.Category{
			ID:          pbCategory.Id,
			Name:        pbCategory.Name,
			Description: pbCategory.Description,
		}
		modelsCategories = append(modelsCategories, modelsCategory)
	}
	return modelsCategories
}
