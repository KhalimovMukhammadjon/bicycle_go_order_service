package postgres

import (
	"bicycle/bicycle_go_order_service/genproto/order_service"
	"bicycle/bicycle_go_order_service/storage"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) storage.OrderRepoI {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) Create(ctx context.Context, req *order_service.CreateOrder) (resp *order_service.PrimaryKey, err error) {
	query :=
		`
		INSERT INTO orders
		(
			id,
			userID,
			productID,
			totalSum
		)
		VALUES(
			$1,
			$2,
			$3,
			$4
		)
	`
	uuid, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	_, err = o.db.Exec(ctx, query,
		uuid.String(),
		req.ProdutcID,
		req.UserID,
		req.TotalSum,
	)
	if err != nil {
		return resp, err
	}

	resp = &order_service.PrimaryKey{
		Id: uuid.String(),
	}

	return resp, nil
}

func (o *orderRepo) GetById(ctx context.Context, req *order_service.PrimaryKey) (resp *order_service.Order, err error) {
	query :=
		`
		SELECT
			id,
			userID,
			productID,
			totalSum
		FROM orders
		WHERE id = $1
	`
	err = o.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.ProdutcID,
		&resp.UserID,
		&resp.TotalSum,
	)
	if err != nil {
		return resp, err
	}

	return
}
