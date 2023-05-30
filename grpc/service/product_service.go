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

type productService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	order_service.UnimplementedProductServiceServer
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *productService {
	return &productService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (o orderService) CreateProduct(ctx context.Context, req *order_service.CreateProduct) (resp *order_service.Product, err error) {
	o.log.Info("---CreateProduct--->", logger.Any("req", req))

	pKey, err := o.strg.Product().Create(context.Background(), req)
	if err != nil {
		o.log.Error("!!!CreateProduct!!!", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = o.strg.Product().GetById(ctx, pKey)

	return
}
