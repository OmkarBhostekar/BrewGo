package api

import (
	"context"
	"fmt"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	db "github.com/omkarbhostekar/brewgo/services/counter/db/sqlc"
	"github.com/omkarbhostekar/brewgo/services/counter/util"
	"google.golang.org/grpc/metadata"
)

type CounterServer struct {
	gen.UnimplementedCounterServiceServer
	store           db.Store
}

// Creates new gRPC instance
func NewServer(config util.Config, store db.Store) (*CounterServer, error) {
	server := &CounterServer{
		store: store,
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