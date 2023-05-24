package grpc

import (
	"bicycle/bicycle_go_order_service/config"
	"bicycle/bicycle_go_order_service/genproto/order_service"
	"bicycle/bicycle_go_order_service/grpc/client"
	"bicycle/bicycle_go_order_service/grpc/service"
	"bicycle/bicycle_go_order_service/pkg/logger"
	"bicycle/bicycle_go_order_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	order_service.RegisterOrderServiceServer(grpcServer, service.NewOrderService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
