package api

import (
	"github.com/omkarbhostekar/brewgo/proto/gen"
	db "github.com/omkarbhostekar/brewgo/services/user/db/sqlc"
	"github.com/omkarbhostekar/brewgo/services/user/util"
)

type UserServer struct {
	gen.UnimplementedUserServiceServer
	config          util.Config
	store           db.Store
}

// Creates new gRPC instance
func NewServer(config util.Config, store db.Store) (*UserServer, error) {
	server := &UserServer{
		store: store, 
		config: config,
	}
	return server, nil
}