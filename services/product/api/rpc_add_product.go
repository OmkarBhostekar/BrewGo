package api

import (
	"context"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	db "github.com/omkarbhostekar/brewgo/services/product/db/sqlc"
	"github.com/shopspring/decimal"
)

func (server *ProductServer) AddProduct(ctx context.Context, req *gen.AddProductRequest) (res *gen.Product, err error) {
	err = validateAddProductRequest(req)
	if err != nil {
		return nil, err
	}
	
	err = server.authorizeAdmin(ctx)
	if err != nil {
		return nil, err
	}

	product, err := server.store.AddProduct(ctx, db.AddProductParams{
		Name:  req.Name,
		Description: req.Description,
		Category: req.Category,
		Price: decimal.NewFromFloat32(req.Price),
		ItemType: req.ItemType,
		IsAvailable: true,
	})
	if err != nil {
		return nil, err
	}

	return convertProductModelToProto(product), nil
}