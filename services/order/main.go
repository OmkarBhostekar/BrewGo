package main

import (
	"database/sql"
	"net"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/omkarbhostekar/brewgo/order/api"
	db "github.com/omkarbhostekar/brewgo/order/db/sqlc"
	"github.com/omkarbhostekar/brewgo/order/util"
	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/omkarbhostekar/brewgo/rabbitmq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	store := db.NewStore(conn)
	log.Info().Msg("connected to db")

	rmq, err := rabbitmq.NewRabbitMQ(rabbitmq.EventsExchange, config.RabbitMQAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to rabbitmq")
	}
	log.Info().Msg("connected to rabbitmq")

	go listenOrderItemUpdates(rmq, config, store)
	runGrpcServer(config, store, rmq)
}

func runGrpcServer(config util.Config, store db.Store, rmq *rabbitmq.RabbitMQ) {
	server, err := api.NewServer(config, store, rmq)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}
	grpcServer := grpc.NewServer()
	gen.RegisterOrderServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:3003")
	if err != nil {
		log.Fatal().Msg("cannot create listener:")
	}
	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Msg("cannot start grpc server")
	}
	defer rmq.Close()
}

func listenOrderItemUpdates(rmq *rabbitmq.RabbitMQ, config util.Config, store db.Store) {
	err := rmq.Consume(rabbitmq.QueueOrderItemStatus, rabbitmq.OrderItemStatusUpdated, func(msg amqp.Delivery) {
		err := api.HandleOrderItemUpdate(msg.Body, config, store, func(data string) {
			publishNotificationEvent(rmq, data)
		})
		if err != nil {
			log.Error().Err(err).Msg("cannot handle order item update")
		}
		msg.Ack(false)
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot consume message")
	}
	select {}
}

func publishNotificationEvent(rmq *rabbitmq.RabbitMQ, data string) {
	err := rmq.Publish(rabbitmq.SendNotification, data)
	if err != nil {
		log.Error().Err(err).Msg("cannot publish message")
	}
}