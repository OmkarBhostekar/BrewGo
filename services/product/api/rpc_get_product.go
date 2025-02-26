package api

import (
	"context"

	"github.com/omkarbhostekar/brewgo/proto/gen"
)

func (server *ProductServer) GetProduct(ctx context.Context, req *gen.GetProductRequest) (res *gen.Product, err error) {
	product, err := server.store.GetProductById(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	return convertProductModelToProto(product), nil
}