// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: CreateCustomerAddress.sql

package eshop

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createCustomerAddress = `-- name: CreateCustomerAddress :exec
INSERT INTO customer_address (
    customer_id,
    location_name,
    address,
    created_at,
    updated_at
) VALUES
    ($1, $2, $3, $4, $5)
`

type CreateCustomerAddressParams struct {
	CustomerID   uuid.UUID
	LocationName string
	Address      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (q *Queries) CreateCustomerAddress(ctx context.Context, arg CreateCustomerAddressParams) error {
	_, err := q.db.ExecContext(ctx, createCustomerAddress,
		arg.CustomerID,
		arg.LocationName,
		arg.Address,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}
