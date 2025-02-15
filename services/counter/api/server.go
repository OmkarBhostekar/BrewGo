package api

import (
	"github.com/omkarbhostekar/brewgo/proto/gen"
	db "github.com/omkarbhostekar/brewgo/services/counter/db/sqlc"
	"github.com/omkarbhostekar/brewgo/services/counter/util"
)

type CounterServer struct {
	gen.UnimplementedUserServiceServer
	store           db.Store
}

// Creates new gRPC instance
func NewServer(config util.Config, store db.Store) (*CounterServer, error) {
	server := &CounterServer{
		store: store,
	}
	return server, nil
}