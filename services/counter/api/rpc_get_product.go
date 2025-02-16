package api

import (
	"context"
	"strconv"

	"github.com/omkarbhostekar/brewgo/proto/gen"
)

func (server *CounterServer) GetProduct(ctx context.Context, req *gen.GetProductRequest) (res *gen.Product, err error) {
	id, err := strconv.ParseInt(req.GetProductId(), 10, 32)
    if err != nil {
        return nil, err
    }
	product, err := server.store.GetProductById(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return convertProductModelToProto(product), nil
}