package api

import (
	"context"

	gen "github.com/omkarbhostekar/brewgo/proto/gen"
	db "github.com/omkarbhostekar/brewgo/services/user/db/sqlc"
	"github.com/omkarbhostekar/brewgo/services/user/util"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *UserServer) CreateUser(ctx context.Context,req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password %s", err)
	}
	arg := db.CreateUserParams{
		Name: req.GetName(),
		Email: req.GetEmail(),
		Password: hashedPassword,
		PhoneNumber: req.GetPhoneNumber(),
	}
	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user %s", err)
	}
	log.Info().Msgf("User created with ID: %d", user.ID)
	return &gen.CreateUserResponse{
		User: convertUser(user),
	}, nil
}