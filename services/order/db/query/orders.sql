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

-- name: UpdateOrder :one
UPDATE counter_orders
SET 
    payment_method = COALESCE(sqlc.narg(payment_method), payment_method),
    order_status = COALESCE(sqlc.narg(order_status), order_status),
    total_amount = COALESCE(sqlc.narg(total_amount), total_amount),
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: GetOrderDetailById :many
SELECT 
    o.id AS order_id,
    o.user_id,
    o.order_date,
    o.total_amount,
    o.payment_method,
    o.order_status,
    o.created_at AS order_created_at,
    o.updated_at AS order_updated_at,
    oi.id AS item_id,
    oi.counter_order_id,
    oi.product_id,
    oi.item_status,
    oi.quantity,
    oi.notes,
    oi.created_at AS item_created_at,
    oi.updated_at AS item_updated_at,
    p.name AS product_name,
    p.price AS product_price
FROM counter_orders o
LEFT JOIN counter_order_items oi ON o.id = oi.counter_order_id
LEFT JOIN products p ON oi.product_id = p.id
WHERE o.id = $1;

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
SELECT 
    coi.*,
    p.name as product_name,
    p.price as product_price
FROM counter_order_items coi
JOIN products p ON coi.product_id = p.id
WHERE coi.counter_order_id = $1;

-- name: GetTotalAmountByOrderId :one
SELECT SUM(quantity * price) as total_amount FROM counter_order_items coi
JOIN products p ON coi.product_id = p.id
WHERE coi.counter_order_id = $1;

-- name: AreAllOrderItemsReady :one
SELECT BOOL_AND(item_status = 'ready') 
FROM counter_order_items 
WHERE counter_order_id = $1;

-- name: UpdateOrderItemStatus :one
UPDATE counter_order_items
    SET item_status = $2
    WHERE id = $1
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM counter_order_items WHERE id = $1;