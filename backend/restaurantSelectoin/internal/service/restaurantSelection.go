package service

import (
	"context"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/client"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/models"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/protos/pb"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/repositories"
)

type Service struct {
	pb.UnimplementedRestaurantSelectionServiceServer
	repo        repositories.RestaurantSelectionService
	restService client.RestaurantsClient
}

func (s *Service) GetRestaurants(ctx context.Context, req *pb.GetRestaurantsRequest) (*pb.GetRestaurantsResponse, error) {
	restaurants, err := s.repo.GetRestaurants()
	if err != nil {
		return nil, err
	}

	restaurantPointers := make([]*pb.Restaurant, len(restaurants))
	for i, r := range restaurants {
		location := r.Locations[0]
		menuResponse, err := s.restService.GetMenu(ctx, uint64(r.ID))
		if err != nil {
			return nil, err
		}
		restaurant := &pb.Restaurant{
			Id:          uint64(r.ID),
			Name:        r.Name,
			Description: r.Description,
			Rating:      r.Rating,
			Location: []*pb.Location{
				{
					City:       location.City,
					PostalCode: location.PostalCode,
					Address:    location.Address,
					Country:    location.Country,
				},
			},
			Menu: &pb.Menu{
				Dishes: menuResponse.Menu.Dishes,
			},
		}
		restaurantPointers[i] = restaurant
	}

	response := &pb.GetRestaurantsResponse{
		Restaurants: restaurantPointers,
	}

	return response, nil
}

func (s *Service) GetRestaurantDetails(ctx context.Context, req *pb.GetRestaurantDetailsRequest) (*pb.RestaurantDetailsResponse, error) {
	restaurantID := req.RestaurantId
	restaurantDetails, err := s.repo.GetRestaurantDetails(restaurantID)
	if err != nil {
		return &pb.RestaurantDetailsResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "error when getting the restaurant",
			},
			Restaurant: nil,
		}, err
	}

	rest := &pb.Restaurant{
		Id:          uint64(restaurantDetails.ID),
		Name:        restaurantDetails.Name,
		Description: restaurantDetails.Description,
		Rating:      restaurantDetails.Rating,
	}

	if len(restaurantDetails.Locations) > 0 {
		location := restaurantDetails.Locations[0]
		rest.Location = []*pb.Location{
			{
				City:       location.City,
				PostalCode: location.PostalCode,
				Address:    location.Address,
				Country:    location.Country,
			},
		}
	}

	menuResponse, err := s.restService.GetMenu(ctx, uint64(restaurantDetails.ID))
	if err != nil {
		return &pb.RestaurantDetailsResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "error when getting the menu",
			},
			Restaurant: nil,
		}, err
	}
	rest.Menu = &pb.Menu{
		Dishes: menuResponse.Menu.Dishes,
	}

	response := &pb.RestaurantDetailsResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
		Restaurant: rest,
	}

	return response, nil
}

func (s *Service) AddNewRestaurant(ctx context.Context, req *pb.AddNewRestaurantRequest) (*pb.AddNewRestaurantResponse, error) {
	newRestaurant := req.GetRestaurant()

	internalRestaurant := pbRestaurantToModel(newRestaurant)
	err := s.repo.AddNewRestaurant(internalRestaurant)
	if err != nil {
		return &pb.AddNewRestaurantResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "Error adding the new restaurant",
			},
		}, err
	}
	return &pb.AddNewRestaurantResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
	}, nil
}

func (s *Service) UpdateRestaurantInfo(ctx context.Context, req *pb.UpdateRestaurantInfoRequest) (*pb.UpdateRestaurantInfoResponse, error) {
	rest := req.Restaurant

	restaurantModel := &models.Restaurant{
		Name:        rest.Name,
		Description: rest.Description,
		Rating:      rest.Rating,
		Locations:   make([]models.Location, len(rest.Location)),
	}
	for i, location := range rest.Location {
		restaurantModel.Locations[i] = models.Location{
			City:       location.City,
			PostalCode: location.PostalCode,
			Address:    location.Address,
			Country:    location.Country,
		}
	}

	err := s.repo.UpdateRestaurantInfo(restaurantModel)
	if err != nil {
		return &pb.UpdateRestaurantInfoResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "Error updating restaurant information",
			},
		}, err
	}

	return &pb.UpdateRestaurantInfoResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
	}, nil
}

func (s *Service) SearchRestaurants(ctx context.Context, req *pb.SearchRestaurantsRequest) (*pb.SearchRestaurantsResponse, error) {
	searchCriteria := req.SearchTerm

	matchedRestaurants, err := s.repo.SearchRestaurants(searchCriteria)
	if err != nil {
		return &pb.SearchRestaurantsResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "Error searching for restaurants",
			},
		}, err
	}

	var restaurantPointers []*pb.Restaurant

	for _, restaurant := range matchedRestaurants {
		pbRestaurant := &pb.Restaurant{
			Id:          uint64(restaurant.ID),
			Name:        restaurant.Name,
			Description: restaurant.Description,
			Rating:      restaurant.Rating,
			Location:    make([]*pb.Location, len(restaurant.Locations)),
		}

		for i, location := range restaurant.Locations {
			pbRestaurant.Location[i] = &pb.Location{
				City:       location.City,
				PostalCode: location.PostalCode,
				Address:    location.Address,
				Country:    location.Country,
			}
		}

		restaurantPointers = append(restaurantPointers, pbRestaurant)
	}

	response := &pb.SearchRestaurantsResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
		Restaurants: restaurantPointers,
	}

	return response, nil
}

func (s *Service) DeleteRestaurant(ctx context.Context, req *pb.DeleteRestaurantRequest) (*pb.DeleteRestaurantResponse, error) {
	restaurantID := req.RestaurantId
	rest, err := s.repo.GetRestaurantDetails(restaurantID)
	if err != nil {
		return &pb.DeleteRestaurantResponse{
			Response: &pb.Response{
				Status: 404,
				Error:  "Restaurant not found",
			},
		}, err
	}
	err = s.repo.DeleteRestaurant(rest)
	if err != nil {
		return &pb.DeleteRestaurantResponse{
			Response: &pb.Response{
				Status: 500,
				Error:  "Error deleting the restaurant",
			},
		}, err
	}

	return &pb.DeleteRestaurantResponse{
		Response: &pb.Response{
			Status: 200,
			Error:  "",
		},
	}, nil
}

func pbRestaurantToModel(pbRestaurant *pb.Restaurant) *models.Restaurant {
	return &models.Restaurant{
		Name:        pbRestaurant.Name,
		Description: pbRestaurant.Description,
		Rating:      float32(pbRestaurant.Rating),
		Locations:   nil,
		Menu:        models.Menu{},
	}
}
