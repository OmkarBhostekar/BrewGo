package api

import (
	"context"
	"database/sql"

	db "github.com/omkarbhostekar/brewgo/services/counter/db/sqlc"
	"github.com/omkarbhostekar/brewgo/proto/gen"
)

func (server *CounterServer) SearchProductsByName(ctx context.Context, req *gen.SearchProductsByNameRequest) (
	*gen.SearchProductsByNameResponse, 
	error,
) {
	params := db.SearchProductsByNameParams{
		Query: sql.NullString{String: req.GetName(), Valid: true},
		PageNumber: req.Page,
		PageSize: req.PageSize,
	}
	list, err := server.store.SearchProductsByName(ctx, params)
	if err != nil {
		return nil, err
	}
	products := make([]*gen.Product, len(list))
	for i, product := range list {
		products[i] = convertProductModelToProto(product)
	}
	rsp := &gen.SearchProductsByNameResponse{
		Products: products,
	}
	return rsp, nil
}