package api

import (
	db "github.com/omkarbhostekar/brewgo/order/db/sqlc"
	"github.com/omkarbhostekar/brewgo/proto/gen"
)

func convertPlaceOrderItems(items []db.CounterOrderItem) []*gen.PlaceOrderItem {
	var orderItems []*gen.PlaceOrderItem
	for _, item := range items {
		orderItems = append(orderItems, &gen.PlaceOrderItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
			Notes:     item.Notes.String,
		})
	}
	return orderItems
}