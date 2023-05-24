package service

import (
	"bicycle/bicycle_go_order_service/config"
	"bicycle/bicycle_go_order_service/genproto/order_service"
	"bicycle/bicycle_go_order_service/grpc/client"
	"bicycle/bicycle_go_order_service/pkg/logger"
	"bicycle/bicycle_go_order_service/storage"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type orderService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	order_service.UnimplementedOrderServiceServer
}

func NewOrderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *orderService {
	return &orderService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (o orderService) Create(ctx context.Context, req *order_service.CreateOrder) (resp *order_service.Order, err error) {
	pKey, err := o.strg.Order().Create(context.Background(), req)
	if err != nil {
		o.log.Error("!!!CreateBook!!!", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = o.strg.Order().GetById(ctx, pKey)

	return
}

func (o orderService) GetById(ctx context.Context, req *order_service.PrimaryKey) (resp *order_service.Order, err error) {
	resp, err = o.strg.Order().GetById(context.Background(), req)
	if err != nil {
		o.log.Error("!!!GetOrderById!!!", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

// func (o orderService) Create(){
// 	o.services.UserService().Create(context.Background(),&user_service.CreateUserRequest{})
// }

// func (o orderService) CreateOrder() {
// 	o.services.UserService().GetAll(context.Background(),&user_service.GetAllUserRequest{})
// 	fmt.Println()
// }
