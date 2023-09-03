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

func (s *Service) AddDish(ctx context.Context, req *pb.CreateMenuItemRequest) (*pb.Dish, error) {
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

	pbDish := &pb.Dish{
		ID:           dish.ID,
		RestaurantId: dish.RestaurantID,
		Name:         dish.Name,
		Description:  dish.Description,
		Availability: dish.Availability,
		Price:        dish.Price,
		Images:       dish.Images,
		Ingredients:  dish.Ingredients,
		Categories:   ConvertModelsCategoriesToPb(dish.Categories),
	}
	return pbDish, nil
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

func (s *Service) GetMenu(ctx context.Context, req *pb.GetMenuRequest) (*pb.Menu, error) {
	menu, err := s.repo.GetMenu(req.RestaurantId)
	if err != nil {
		return nil, err
	}

	rDish := make([]*pb.Dish, 0, len(menu.Dishes))
	for _, dish := range menu.Dishes {
		pbDish := &pb.Dish{
			ID:           dish.ID,
			RestaurantId: dish.RestaurantID,
			Name:         dish.Name,
			Description:  dish.Description,
			Availability: dish.Availability,
			Price:        dish.Price,
			Images:       dish.Images,
			Ingredients:  dish.Ingredients,
			Categories:   ConvertModelsCategoriesToPb(dish.Categories),
		}
		rDish = append(rDish, pbDish)
	}

	menuResponse := &pb.Menu{
		Dishes: rDish,
	}
	return menuResponse, nil
}

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

func (s *Service) GetDishDetails(ctx context.Context, req *pb.GetDishDetailsRequest) (*pb.Dish, error) {
	dishID := req.DishId

	dish, err := s.repo.GetDishDetails(dishID)
	if err != nil {
		return nil, err
	}
	pbDish := &pb.Dish{
		ID:           dish.ID,
		RestaurantId: dish.RestaurantID,
		Name:         dish.Name,
		Description:  dish.Description,
		Availability: dish.Availability,
		Price:        dish.Price,
		Images:       dish.Images,
		Ingredients:  dish.Ingredients,
		Categories:   ConvertModelsCategoriesToPb(dish.Categories),
	}
	return pbDish, nil
}

func (s *Service) UpdateDishIngredients(ctx context.Context, req *pb.UpdateDishIngredientsRequest) (*pb.UpdateDishResponse, error) {
	dishID := req.DishId
	newIngredients := req.Ingredients

	dish, err := s.repo.GetDishDetails(dishID)
	if err != nil {
		return &pb.UpdateDishResponse{
			Response: &pb.Response{
				Status: 404,
				Error:  "Dish not found",
			},
		}, err
	}
	dish.Ingredients = newIngredients

	if err := s.repo.UpdateDish(dish); err != nil {
		return &pb.UpdateDishResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "Failed to update dish",
			},
		}, err
	}
	return &pb.UpdateDishResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "Dish ingredients updated successfully",
		},
	}, nil
}

func (s *Service) GetDishCategories(ctx context.Context, req *pb.GetDishDetailsRequest) (*pb.GetDishCategoriesResponse, error) {
	dishID := req.DishId
	dish, err := s.repo.GetDishDetails(dishID)
	if err != nil {
		return nil, err
	}
	categories := ConvertModelsCategoriesToPb(dish.Categories)
	response := &pb.GetDishCategoriesResponse{
		Categories: categories,
	}

	return response, nil
}

func (s *Service) ToggleDishAvailability(ctx context.Context, req *pb.ToggleDishAvailabilityRequest) (*pb.ToggleDishAvailabilityResponse, error) {
	dishID := req.DishId
	dish, err := s.repo.GetDishDetails(dishID)
	if err != nil {
		return nil, err
	}
	dish.Availability = dish.Availability

	if err := s.repo.UpdateDish(dish); err != nil {
		return nil, err
	}

	response := &pb.ToggleDishAvailabilityResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
	}
	return response, nil
}

//func (s *Service) UploadDishImages(ctx context.Context, req *pb.UploadDishImagesRequest) (*pb.UploadDishImagesResponse, error) {
//
//}
//
//func (s *Service) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
//}
//func (s *Service) UpdateOrderStatus(context.Context, *UpdateOrderStatusRequest) (*UpdateOrderStatusResponse, error) {
//	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
//}

func ConvertModelsCategoriesToPb(categories []*models.Category) []*pb.Category {
	pbCategories := make([]*pb.Category, len(categories))
	for i, category := range categories {
		pbCategory := &pb.Category{
			Id:   category.ID,
			Name: category.Name,
			// Add other fields if needed.
		}
		pbCategories[i] = pbCategory
	}
	return pbCategories
}
