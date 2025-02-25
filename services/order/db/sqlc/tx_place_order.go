package orders

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"
)

// PlaceOrderTx places an order and inserts related order items in a transaction
func (store *SQLStore) PlaceOrderTx(ctx context.Context, arg PlaceOrderTxParams) (PlaceOrderTxResult, error) {
	var res PlaceOrderTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Create order
		res.Order, err = q.CreateOrder(ctx, arg.OrderParams)
		if err != nil {
			return wrapError("failed to create order", err)
		}

		// Insert order items
		res.OrderItems = make([]CounterOrderItem, 0, len(arg.OrderItems))

		for _, item := range arg.OrderItems {
			orderItem, err := q.CreateOrderItem(ctx, CreateOrderItemParams{
				CounterOrderID: res.Order.ID,
				ProductID:      item.ProductID,
				Quantity:       item.Quantity,
				Notes:          item.Notes,
			})
			if err != nil {
				return wrapError("failed to create order item", err)
			}
			res.OrderItems = append(res.OrderItems, orderItem)
		}

		// Get total amount by order id
		totalAmount, err := q.GetTotalAmountByOrderId(ctx, res.Order.ID)
		if err != nil {
			return wrapError("failed to get total amount by order id", err)
		}

		// Update order total amount
		res.Order, err = q.UpdateOrder(ctx, UpdateOrderParams{
			ID:          res.Order.ID,
			TotalAmount: decimal.NewFromInt(totalAmount),
		})

		if err != nil {
			return wrapError("failed to update order total amount", err)
		}

		return nil
	})

	return res, err
}

// wrapError helps provide better error messages
func wrapError(message string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", message, err)
	}
	return nil
}

// Input parameters for transaction
type PlaceOrderTxParams struct {
	OrderParams CreateOrderParams
	OrderItems  []CreateOrderItemParams
}

// Result struct for transaction
type PlaceOrderTxResult struct {
	Order      CounterOrder
	OrderItems []CounterOrderItem
}
