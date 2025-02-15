package api

import (
	"context"
	gen "github.com/omkarbhostekar/brewgo/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *UserServer) ValidateToken(ctx context.Context,req *gen.ValidateTokenRequest) (res *gen.ValidateTokenResponse, err error) {
	authPayload, err := server.tokenMaker.VerifyToken(req.AccessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated %s", err)
	}

	return &gen.ValidateTokenResponse{
		UserId: authPayload.UserID,
		Role: authPayload.Role,
	}, nil
}