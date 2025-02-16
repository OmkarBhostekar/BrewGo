package api

import (
	"context"
	"strconv"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *CounterServer) DeleteProduct(ctx context.Context, req *gen.DeleteProductRequest) (res *emptypb.Empty, err error) {
	err = server.authorizeAdmin(ctx)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseInt(req.GetProductId(), 10, 32)
    if err != nil {
        return nil, err
    }
	_, err = server.store.DeleteProduct(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}