package service

import (
	"context"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/protos/pb"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/repositories"
)

type Service struct {
	pb.UnimplementedRestaurantSelectionServiceServer
	repo repositories.RestaurantSelectionService
}

func (s *Service) GetRestaurants(ctx context.Context, req *pb.GetRestaurantsRequest) (*pb.GetRestaurantsResponse, error) {

}

//func (s *Service) GetRestaurantDetails(ctx context.Context,req *pb.GetRestaurantDetailsRequest) (*pb.RestaurantDetails, error) {
//}
//func (s *Service) AddNewRestaurant(ctx context.Context,req *pb.AddNewRestaurantRequest) (*pb.Response, error) {
//}
//func (s *Service) UpdateRestaurantInfo(ctx context.Context,req *pb.UpdateRestaurantInfoRequest) (*pb.Response, error) {
//}
//func (s *Service) SearchRestaurants(ctx context.Context,req *pb.SearchRestaurantsRequest) (*pb.SearchRestaurantsResponse, error) {
//}
//func (s *Service) DeleteRestaurant(ctx context.Context,req *pb.DeleteRestaurantRequest) (*pb.Response, error) {
//}
