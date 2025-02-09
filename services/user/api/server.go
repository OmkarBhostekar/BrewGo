package api

import (
	"fmt"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	db "github.com/omkarbhostekar/brewgo/services/user/db/sqlc"
	"github.com/omkarbhostekar/brewgo/services/user/token"
	"github.com/omkarbhostekar/brewgo/services/user/util"
)

type UserServer struct {
	gen.UnimplementedUserServiceServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
}

// Creates new gRPC instance
func NewServer(config util.Config, store db.Store) (*UserServer, error) {
	maker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &UserServer{
		store: store, 
		config: config,
		tokenMaker: maker,
	}
	return server, nil
}