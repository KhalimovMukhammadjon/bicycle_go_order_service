package storage

import (
	"bicycle/bicycle_go_order_service/genproto/order_service"
	"context"
)

type StorageI interface {
	CloseDB()
	Order() OrderRepoI
}

type OrderRepoI interface {
	Create(ctx context.Context, req *order_service.CreateOrder) (resp *order_service.PrimaryKey, err error)
	GetById(ctx context.Context, req *order_service.PrimaryKey) (resp *order_service.Order, err error)
	//GetList(ctx context.Context, req *order_service.GetAllOrderRequest) (resp *order_service.GetAllOrderResponse, err error)
	// Update(ctx context.Context, req *order_service.PrimaryKey) error
	// Delete(ctx context.Context, req *order_service.PrimaryKey) error
}

type ProductRepoI interface {
	Create(ctx context.Context, req *order_service.CreateProduct) (resp *order_service.PrimaryKeyProduct, err error)
	GetById(ctx context.Context, req *order_service.PrimaryKeyProduct) (resp *order_service.Product, err error)
	GetList(ctx context.Context, req *order_service.GetAllProductRequest) (resp *order_service.GetAllProductResponse, err error)
	// Update(ctx context.Context, req *order_service.PrimaryKeyProduct) error
	// Delete(ctx context.Context, req *order_service.PrimaryKeyProduct) error
}
