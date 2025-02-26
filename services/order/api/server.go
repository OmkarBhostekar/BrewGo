package api

import (
	"context"
	"fmt"

	db "github.com/omkarbhostekar/brewgo/order/db/sqlc"
	"github.com/omkarbhostekar/brewgo/order/util"
	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/omkarbhostekar/brewgo/rabbitmq"
	"google.golang.org/grpc/metadata"
)

type CounterServer struct {
	gen.UnimplementedOrderServiceServer
	store           db.Store
	rmq 		   *rabbitmq.RabbitMQ
}

// Creates new gRPC instance
func NewServer(config util.Config, store db.Store, rmq *rabbitmq.RabbitMQ) (*CounterServer, error) {
	server := &CounterServer{
		store: store,
		rmq: rmq,
	}
	return server, nil
}

func (server *CounterServer) authorizeAdmin(Ctx context.Context) error {
	mtdt, ok := metadata.FromIncomingContext(Ctx)
	if !ok {
		return fmt.Errorf("metadata is not provided")
	}
	values := mtdt.Get("X-Role")
	if len(values) == 0 {
		return fmt.Errorf("role is not provided")
	}
	role := values[0]

	if role != "admin" {
		return fmt.Errorf("you're not allowed to perform this action")
	}
	return nil
}