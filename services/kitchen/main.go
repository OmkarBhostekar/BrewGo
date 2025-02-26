package main

import (
	"os"

	"github.com/omkarbhostekar/brewgo/rabbitmq"
	"github.com/omkarbhostekar/brewgo/services/kitchen/workers"
	"github.com/omkarbhostekar/brewgo/services/kitchen/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	rmq, err := rabbitmq.NewRabbitMQ(rabbitmq.EventsExchange, config.RabbitMQAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to rabbitmq")
	}
	log.Info().Msg("connected to rabbitmq")

	listenNewOrderItemReceived(rmq, config)

	select {}
}

func listenNewOrderItemReceived(rmq *rabbitmq.RabbitMQ, config util.Config) {
	err := rmq.Consume(rabbitmq.QueueNewOrders, rabbitmq.NewOrderItemReceived, func(msg amqp.Delivery) {
		handlers.NewOrderItemWorker(msg, config, func(data string) {
			publishOrderItemStatus(rmq, data)
		})
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot consume message")
	}
}

func publishOrderItemStatus(rmq *rabbitmq.RabbitMQ, data string) {
	err := rmq.Publish(rabbitmq.OrderItemStatusUpdated, data)
	if err != nil {
		log.Error().Err(err).Msg("cannot publish message")
	}
}