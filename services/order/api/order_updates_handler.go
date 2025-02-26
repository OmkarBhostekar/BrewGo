package api

import (
	"context"
	"database/sql"
	"encoding/json"

	db "github.com/omkarbhostekar/brewgo/order/db/sqlc"
	"github.com/omkarbhostekar/brewgo/order/util"
	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/omkarbhostekar/brewgo/rabbitmq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func HandleOrderItemUpdate(body []byte, config util.Config, store db.Store, sendNofication func(data string)) error {
	data := rabbitmq.OrderItemStatusEventData{}
	err := json.Unmarshal(body, &data)

	if err != nil {
		log.Error().Err(err).Msg("cannot unmarshal data")
		return err
	}

	_, err = store.UpdateOrderItemStatus(context.Background(), db.UpdateOrderItemStatusParams{
		ID:         data.OrderItemId,
		ItemStatus: data.Status,
	})

	if err != nil {
		return err
	}

	orderItems, err := store.GetOrderDetailById(context.Background(), data.OrderId)
	if err != nil {
		return err
	}

	readyItems := countReadyItems(orderItems)
	userId := orderItems[0].UserID

	if readyItems == len(orderItems) {
		// all items are ready, update order status
		// send comms order is ready
		_, err = store.UpdateOrder(context.Background(), db.UpdateOrderParams{
			ID: data.OrderId,
			OrderStatus: sql.NullString{
				String: "ready",
				Valid:  true,
			},
		})
		if err != nil {
			return err
		}
		data := getNotificationEventData(userId, config, "get ready for yumminess! your order is ready to serve")
		if data != nil {
			sendNofication(*data)
		}
	} else if readyItems == 1 {
		// first item is ready
		// send comms order is being prepared
		data := getNotificationEventData(userId, config, "yay! your order is being prepared")
		if data != nil {
			sendNofication(*data)
		}
	}

	return nil
}

func getNotificationEventData(userId int32, config util.Config, message string) *string {
	// dial user service to get user details, email, phone number
	result, err := getUserFromUserService(config.UserServiceEndPoint, userId)
	if err != nil {
		log.Error().Err(err).Msg("cannot get user")
		return nil
	}
	
	email := result.User.GetEmail()

	// will send email notification for now
	data := rabbitmq.NotificationEventData{
		Type:    "email",
		EmailData: rabbitmq.EmailNotification{
			To:      email,
			Subject: "Update on your order",
			Body:    message,
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("cannot marshal notification")
		return nil
	}
	jsonStr := string(jsonData)
	return &jsonStr
}

func countReadyItems(items []db.GetOrderDetailByIdRow) int {
	count := 0
	for _, item := range items {
		if item.ItemStatus.String == "ready" {
			count++
		}
	}
	return count
}

func getUserFromUserService(userServiceEndPoint string, userId int32) (*gen.GetUserByPhoneNumberResponse, error) {
	// Connect to UserService
	conn, err := grpc.Dial(userServiceEndPoint, grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to UserService")
		return nil, err
	}
	defer conn.Close()

	client := gen.NewUserServiceClient(conn)

	// Attach "X-Role=admin" in gRPC metadata
	md := metadata.Pairs("X-Role", "admin")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Call the gRPC method with metadata
	result, err := client.GetUserById(ctx, &gen.GetUserByIdRequest{UserId: userId})
	if err != nil {
		log.Error().Err(err).Msg("cannot get user")
		return nil, err
	}

	return result, nil
}
