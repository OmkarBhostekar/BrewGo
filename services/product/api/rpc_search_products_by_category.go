package api

import (
	"context"

	db "github.com/omkarbhostekar/brewgo/services/product/db/sqlc"
	"github.com/rs/zerolog/log"

	"github.com/omkarbhostekar/brewgo/proto/gen"
)

func (server *ProductServer) SearchProductsByCategory(ctx context.Context, req *gen.SearchProductsByCategoryRequest) (
	*gen.SearchProductsByNameResponse, 
	error,
) {
	params := db.SearchProductsByCategoryParams{
		Category: req.CategoryName,
		PageNumber: req.Page,
		PageSize: req.PageSize,
	}
	list, err := server.store.SearchProductsByCategory(ctx, params)
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
	log.Info().Msgf("Response: %v", rsp)
	return rsp, nil
}