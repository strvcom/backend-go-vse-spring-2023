-- name: CreateCustomerAddress :exec
INSERT INTO customer_address (
    customer_id,
    location_name,
    address,
    created_at,
    updated_at
) VALUES
    ($1, $2, $3, $4, $5);
