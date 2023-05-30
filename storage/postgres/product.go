package postgres

import (
	"bicycle/bicycle_go_order_service/genproto/order_service"
	"bicycle/bicycle_go_order_service/pkg/helper"
	"bicycle/bicycle_go_order_service/storage"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) storage.ProductRepoI {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Create(ctx context.Context, req *order_service.CreateProduct) (resp *order_service.PrimaryKeyProduct, err error) {
	query := `
		INSERT INTO product
			(name,price)
		VALUES(
			$1,
			$2
		)
	`
	uuid, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	_, err = p.db.Exec(ctx, query,
		uuid.String(),
		req.Name,
		req.Price,
	)
	if err != nil {
		return resp, err
	}

	resp = &order_service.PrimaryKeyProduct{
		Id: uuid.String(),
	}

	return resp, nil
}

func (p *productRepo) GetById(ctx context.Context, req *order_service.PrimaryKeyProduct) (resp *order_service.Product, err error) {
	resp = &order_service.Product{}

	query := `
		SELECT 
			id,
			name,
			price
		FROM product
		WHERE id = $1
	`

	err = p.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.Name,
		&resp.Price,
	)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (p *productRepo) GetList(ctx context.Context, req *order_service.GetAllProductRequest) (resp *order_service.GetAllProductResponse, err error) {
	resp = &order_service.GetAllProductResponse{}

	var (
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query := `
		SELECT 
			id,
			name,
			price
		FROM product
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product order_service.Product

		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		resp.Product = append(resp.Product, &product)
	}

	return resp, nil
}

func (p *productRepo) Update(ctx context.Context, req *order_service.UpdateProductRequest) error {
	query := `
		UPDATE product SET
			id = :id,
			name = :name,
			price = :price
		WHERE id = :id
	`

	params := map[string]interface{}{
		"id":    req.Product.Id,
		"name":  req.Product.Name,
		"price": req.Product.Price,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	_, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		return nil
	}

	return nil
}

func (p *productRepo) Delete(ctx context.Context, req *order_service.PrimaryKeyProduct) error {
	query := `
		DELETE FROM product WHERE id = $1
	`

	_, err := p.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
