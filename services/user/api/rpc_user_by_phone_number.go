package api

import (
	"context"
	"database/sql"

	gen "github.com/omkarbhostekar/brewgo/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *UserServer) GetUserByPhoneNumber(ctx context.Context, req *gen.GetUserByPhoneNumberRequest) (res *gen.GetUserByPhoneNumberResponse, err error) {
	user, err := server.store.GetUserByPhoneNumber(ctx, req.GetPhoneNumber())

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get user by phone number %s", err)
	}

	return &gen.GetUserByPhoneNumberResponse{
		User: convertUser(user),
	}, nil
}