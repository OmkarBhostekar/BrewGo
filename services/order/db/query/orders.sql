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

-- name: UpdateOrderStatus :one
UPDATE counter_orders
    SET order_status = $2
    WHERE id = $1
RETURNING *;

-- name: GetOrderById :one
SELECT * FROM counter_orders WHERE id = $1;

-- name: DeleteOrder :exec
DELETE FROM counter_orders WHERE id = $1;

-- name: GetOrdersByUserId :many
SELECT * FROM counter_orders WHERE user_id = $1;

-- name: CreateOrderItem :one
INSERT INTO counter_order_items(
    counter_order_id,
    product_id,
    quantity,
    notes
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetOrderItemsByOrderId :many
SELECT * FROM counter_order_items WHERE counter_order_id = $1;

-- name: UpdateOrderItemStatus :one
UPDATE counter_order_items
    SET item_status = $2
    WHERE id = $1
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM counter_order_items WHERE id = $1;