-- name: CreateOrder :one
INSERT INTO counter_orders(
    user_id,
    total_amount,
    payment_method
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;