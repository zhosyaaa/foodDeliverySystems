package client

import (
	"context"
	"github.com/zhosyaaa/foodDeliverySystems-order/internal/protos/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DishServiceClient struct {
	Client pb.DishServiceClient
}

func NewDishServiceClient(addr string) (*DishServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewDishServiceClient(conn)

	return &DishServiceClient{
		Client: client,
	}, nil
}

func (c *DishServiceClient) GetDishById(ctx context.Context, DishID uint64) (*pb.GetDishByIdResponse, error) {
	request := &pb.GetDishByIdRequest{
		DishID: DishID,
	}
	response, err := c.Client.GetDishById(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
