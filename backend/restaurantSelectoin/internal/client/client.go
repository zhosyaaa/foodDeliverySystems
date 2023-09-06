package client

import (
	"context"
	"github.com/zhosyaaa/foodDeliverySystems-restaurantSelection/internal/protos/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RestaurantsClient struct {
	Client pb.RestaurantsServiceClient
}

func NewRestaurantsClient(addr string) (*RestaurantsClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewRestaurantsServiceClient(conn)

	return &RestaurantsClient{
		Client: client,
	}, nil
}

func (c *RestaurantsClient) GetMenu(ctx context.Context, MenuID uint64) (*pb.GetMenuResponse, error) {
	request := &pb.GetMenuRequest{
		RestID: MenuID,
	}

	response, err := c.Client.GetMenu(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
