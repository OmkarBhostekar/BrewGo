-- name: AddProduct :one
INSERT INTO products(
    name,
    description,
    price,
    category,
    is_available,
    item_type
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: SearchProductsByName :many
SELECT * FROM products WHERE name ILIKE '%' || @query || '%' LIMIT @page_size OFFSET @page_number;

-- name: SearchProductsByCategory :many
SELECT * FROM products WHERE category = @category LIMIT @page_size OFFSET @page_number;

-- name: GetProductById :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: UpdateProduct :one
UPDATE products
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    price = COALESCE(sqlc.narg(price), price),
    category = COALESCE(sqlc.narg(category), category),
    is_available = COALESCE(sqlc.narg(is_available), is_available),
    item_type = COALESCE(sqlc.narg(item_type), item_type)
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :one
DELETE FROM products WHERE id = $1 RETURNING *;