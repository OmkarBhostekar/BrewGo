package api

import (
	"context"
	"database/sql"
	"encoding/json"

	db "github.com/omkarbhostekar/brewgo/order/db/sqlc"
	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/omkarbhostekar/brewgo/rabbitmq"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

func (server *CounterServer) PlaceOrder(ctx context.Context, req *gen.PlaceOrderRequest) (res *gen.PlaceOrderResponse, err error) {
	err = server.authorizeAdmin(ctx)
	if err != nil {
		return nil, err
	}
	txOrderParams := db.CreateOrderParams{
		UserID:        req.UserId,
		TotalAmount:   decimal.NewFromInt(0),
		PaymentMethod: req.PaymentMethod,
	}

	txOrderItems := make([]db.CreateOrderItemParams, 0, len(req.OrderItems))
	for _, item := range req.OrderItems {
		txOrderItems = append(txOrderItems, db.CreateOrderItemParams{
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
			Notes: sql.NullString{
				String: item.Notes,
				Valid:  item.Notes != "",
			},
		})
	}

	txParams := db.PlaceOrderTxParams{
		OrderParams: txOrderParams,
		OrderItems:  txOrderItems,
	}

	// Start a new transaction
	order, err := server.store.PlaceOrderTx(ctx, txParams)

	if err != nil {
		log.Error().Err(err).Msg("Failed to place order")
	}

	// Publish order items received event
	log.Info().Msg("Publishing order items received event")
	for _, item := range order.OrderItems {
		eventData := rabbitmq.NewOrderItemEventData{
			OrderId: item.CounterOrderID,
			OrderItemId: item.ID,
			ProductId: item.ProductID,
			Quantity: item.Quantity,
			Notes: item.Notes.String,
		}
		jsonData, err := json.Marshal(eventData)
		if err != nil {
			log.Error().Err(err).Msg("Failed to marshal order item event data")
		}
		server.rmq.Publish(rabbitmq.NewOrderItemReceived, string(jsonData))
	}


	return &gen.PlaceOrderResponse{
		OrderId: order.Order.ID,
		OrderStatus: order.Order.OrderStatus,
		PaymentMethod: order.Order.PaymentMethod,
		OrderItems: convertPlaceOrderItems(order.OrderItems),
	}, nil
}
