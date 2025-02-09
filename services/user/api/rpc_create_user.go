package api

import (
	"context"

	gen "github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/rs/zerolog/log"
)

func (server *UserServer) CreateUser(context.Context, *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	log.Info().Msg("CreateUser RPC called")
	return &gen.CreateUserResponse{}, nil
}