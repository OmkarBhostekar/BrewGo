package main

import (
	"os"

	"github.com/omkarbhostekar/brewgo/notification/util"
	"github.com/omkarbhostekar/brewgo/notification/workers"
	"github.com/omkarbhostekar/brewgo/rabbitmq"
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
	err := rmq.Consume(rabbitmq.QueueNotifications, rabbitmq.SendNotification, func(msg amqp.Delivery) {
		workers.NotificationWorker(msg, config)
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot consume message")
	}
}