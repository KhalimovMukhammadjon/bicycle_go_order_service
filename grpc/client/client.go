package client

import (
	"bicycle/bicycle_go_order_service/config"
	"bicycle/bicycle_go_order_service/genproto/user_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	UserService() user_service.UserServiceClient
}

type grpcClients struct {
	userService user_service.UserServiceClient
}

func NewGrpcClient(cfg config.Config) (ServiceManagerI, error) {
	connOrderService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService: user_service.NewUserServiceClient(connOrderService),
	}, nil
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}
