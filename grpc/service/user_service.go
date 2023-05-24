package service

import (
	"bicycle/bicycle_go_order_service/config"
	"bicycle/bicycle_go_order_service/genproto/user_service"
	"bicycle/bicycle_go_order_service/grpc/client"
	"bicycle/bicycle_go_order_service/pkg/logger"
	"bicycle/bicycle_go_order_service/storage"
)

type userService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StoragI
	services client.ServiceManagerI
	user_service.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StoragI, svcs client.ServiceManagerI) *userService {
	return &userService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}
