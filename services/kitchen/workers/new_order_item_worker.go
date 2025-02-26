package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/omkarbhostekar/brewgo/rabbitmq"
	"github.com/omkarbhostekar/brewgo/services/kitchen/util"
)

func NewOrderItemWorker(msg amqp.Delivery, config util.Config, updateStatus func(string)) error {
	data := rabbitmq.NewOrderItemEventData{}
	err := json.Unmarshal(msg.Body, &data)

	if err != nil {
		log.Error().Err(err).Msg("cannot unmarshal data")
		return err
	}

	// get product details
	conn, err := grpc.Dial(config.ProductServiceEndPoint, grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("connect to product service")
		return err
	}
	defer conn.Close()

	log.Info().Msgf("received order for product id: %d", data.ProductId)
	client := gen.NewProductServiceClient(conn)
	product, err := client.GetProduct(context.Background(), &gen.GetProductRequest{ProductId: data.ProductId})
	if err != nil {
		log.Error().Err(err).Msg("cannot get product")
		return err
	}
	// currently using seconds but should be in minutes
	estimatedTime := product.EstimatedPreparationTime

	log.Info().Msgf("preparing estimated time: %v", estimatedTime)
	preparing, err := getUpdateStatusEventData(data.OrderId, data.OrderItemId, "preparing")
	if err != nil {
		return err
	}
	updateStatus(preparing)

	// sleep for estimated time
	time.Sleep(time.Duration(estimatedTime) * time.Second)

	log.Info().Msg("order item is ready")
	ready, err := getUpdateStatusEventData(data.OrderId, data.OrderItemId, "ready")
	if err != nil {
		return err
	}
	updateStatus(ready)
	
	msg.Ack(false)
	return nil
}

func getUpdateStatusEventData(orderId int32, orderItemId int32, status string) (string, error) {
	data := rabbitmq.OrderItemStatusEventData{
		Status: status,
		OrderId: orderId,
		OrderItemId: orderItemId,
	}
	jsonData, err := json.Marshal(data)
	return string(jsonData), err
}