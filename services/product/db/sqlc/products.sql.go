// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package product

import (
	"context"
	"database/sql"

	"github.com/shopspring/decimal"
)

const addProduct = `-- name: AddProduct :one
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
) RETURNING id, name, description, price, category, is_available, est_preparation_time, item_type, created_at, updated_at
`

type AddProductParams struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
	Category    string          `json:"category"`
	IsAvailable bool            `json:"is_available"`
	ItemType    string          `json:"item_type"`
}

func (q *Queries) AddProduct(ctx context.Context, arg AddProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, addProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Category,
		arg.IsAvailable,
		arg.ItemType,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Category,
		&i.IsAvailable,
		&i.EstPreparationTime,
		&i.ItemType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :one
DELETE FROM products WHERE id = $1 RETURNING id, name, description, price, category, is_available, est_preparation_time, item_type, created_at, updated_at
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, deleteProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Category,
		&i.IsAvailable,
		&i.EstPreparationTime,
		&i.ItemType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProductById = `-- name: GetProductById :one
SELECT id, name, description, price, category, is_available, est_preparation_time, item_type, created_at, updated_at FROM products WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProductById(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductById, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Category,
		&i.IsAvailable,
		&i.EstPreparationTime,
		&i.ItemType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const searchProductsByCategory = `-- name: SearchProductsByCategory :many
SELECT id, name, description, price, category, is_available, est_preparation_time, item_type, created_at, updated_at FROM products WHERE category = $1 LIMIT $3 OFFSET $2
`

type SearchProductsByCategoryParams struct {
	Category   string `json:"category"`
	PageNumber int32  `json:"page_number"`
	PageSize   int32  `json:"page_size"`
}

func (q *Queries) SearchProductsByCategory(ctx context.Context, arg SearchProductsByCategoryParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, searchProductsByCategory, arg.Category, arg.PageNumber, arg.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Category,
			&i.IsAvailable,
			&i.EstPreparationTime,
			&i.ItemType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchProductsByName = `-- name: SearchProductsByName :many
SELECT id, name, description, price, category, is_available, est_preparation_time, item_type, created_at, updated_at FROM products WHERE name ILIKE '%' || $1 || '%' LIMIT $3 OFFSET $2
`

type SearchProductsByNameParams struct {
	Query      sql.NullString `json:"query"`
	PageNumber int32          `json:"page_number"`
	PageSize   int32          `json:"page_size"`
}

func (q *Queries) SearchProductsByName(ctx context.Context, arg SearchProductsByNameParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, searchProductsByName, arg.Query, arg.PageNumber, arg.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Category,
			&i.IsAvailable,
			&i.EstPreparationTime,
			&i.ItemType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET
    name = COALESCE($2, name),
    description = COALESCE($3, description),
    price = COALESCE($4, price),
    category = COALESCE($5, category),
    is_available = COALESCE($6, is_available),
    item_type = COALESCE($7, item_type)
WHERE id = $1
RETURNING id, name, description, price, category, is_available, est_preparation_time, item_type, created_at, updated_at
`

type UpdateProductParams struct {
	ID          int32           `json:"id"`
	Name        sql.NullString  `json:"name"`
	Description sql.NullString  `json:"description"`
	Price       decimal.Decimal `json:"price"`
	Category    sql.NullString  `json:"category"`
	IsAvailable sql.NullBool    `json:"is_available"`
	ItemType    sql.NullString  `json:"item_type"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Category,
		arg.IsAvailable,
		arg.ItemType,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Category,
		&i.IsAvailable,
		&i.EstPreparationTime,
		&i.ItemType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
