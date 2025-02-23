package api

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	db "github.com/omkarbhostekar/brewgo/services/product/db/sqlc"
	"github.com/shopspring/decimal"
)

func (server *ProductServer) UpdateProduct(ctx context.Context, req *gen.UpdateProductRequest) (res *gen.Product, err error) {
	err = server.authorizeAdmin(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseInt(req.GetProductId(), 10, 32)
    if err != nil {
        return nil, err
    }
	product, err := server.store.GetProductById(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	price := product.Price
	if req.Price != nil {
		price = decimal.NewFromFloat32(req.GetPrice())
	}

	product, err = server.store.UpdateProduct(ctx, db.UpdateProductParams{
		ID: 		int32(id),
		Name:        sql.NullString{String: req.GetName(), Valid: req.Name != nil},
		Description: sql.NullString{String: req.GetDescription(), Valid: req.Description != nil},
		Category:    sql.NullString{String: req.GetCategory(), Valid: req.Category != nil},
		Price:       price,
		ItemType:    sql.NullString{String: req.GetItemType(), Valid: req.ItemType != nil},
		IsAvailable: sql.NullBool{Bool: req.GetIsAvailable(), Valid: req.IsAvailable != nil},
	})
	if err != nil {
		return nil, err
	}

	return convertProductModelToProto(product), nil
}
