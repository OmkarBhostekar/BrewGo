package api

import (
	"context"
	"database/sql"

	gen "github.com/omkarbhostekar/brewgo/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *UserServer) GetUserById(ctx context.Context, req *gen.GetUserByIdRequest) (res *gen.GetUserByPhoneNumberResponse, err error) {
	role, err := server.getRoleFromHeader(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated %s", err)
	}
	if role != "admin" && role != "service" {
		return nil, status.Errorf(codes.PermissionDenied, "you're not authorized to complete this action")
	}
	
	user, err := server.store.GetUserById(ctx, req.GetUserId())

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get user by user id %s", err)
	}

	return &gen.GetUserByPhoneNumberResponse{
		User: convertUser(user),
	}, nil
}