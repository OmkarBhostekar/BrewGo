package api

import (
	"context"
	"fmt"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *CounterServer) GetOrderStatus(ctx context.Context,req *gen.GetOrderStatusRequest) (*gen.OrderResponse, error) {
	
	rows, err := server.store.GetOrderDetailById(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("order not found")
	}

	orderItems := make([]*gen.OrderItemResponse, 0, len(rows))
	for _, item := range rows {
		price, _ := item.ProductPrice.Float64()
		orderItems = append(orderItems, &gen.OrderItemResponse{
			ProductId: item.ProductID.Int32,
			ProductName: item.ProductName.String,
			Price: float32(price),
			Quantity: item.Quantity.Int32,
			Notes: item.Notes.String,
			ItemStatus: item.ItemStatus.String,
			UpdatedAt: timestamppb.New(item.ItemUpdatedAt.Time),
		})
	}
	totalAmount, ok := rows[0].TotalAmount.Float64()

	if !ok {
		return nil, fmt.Errorf("failed to convert total amount to float64")
	}

	return &gen.OrderResponse{
		OrderId: rows[0].OrderID,
		OrderStatus: rows[0].OrderStatus,
		OrderItems: orderItems,
		TotalAmount: float32(totalAmount),
	}, nil

}