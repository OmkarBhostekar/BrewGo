package api

import (
	"context"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/rs/zerolog/log"
)

func (server *CounterServer) PlaceOrder(ctx context.Context, req *gen.PlaceOrderRequest) (res *gen.OrderResponse, err error) {
	
	log.Info().Msgf("Place Order Request: %v", req)
	return &gen.OrderResponse{}, nil
}