// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: CreateProduct.sql

package eshop

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO product (
    id,
    name,
    description,
    price,
    created_at,
    updated_at
) VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING id, name, description, price, created_at, updated_at
`

type CreateProductParams struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
