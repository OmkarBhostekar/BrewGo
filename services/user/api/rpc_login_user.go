package api

import (
	"context"
	"strings"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/omkarbhostekar/brewgo/services/user/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *UserServer) LoginUser(ctx context.Context, req *gen.LoginUserRequest) (*gen.LoginUserResponse, error) {
	
	user, err := server.store.GetUserByEmail(ctx, strings.ToLower(req.Email))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no user found %s", err)
	}
	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password %s", err)
	}
	access, accessPayload, err := server.tokenMaker.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}
	refresh, refreshPayload, err := server.tokenMaker.CreateToken(user.ID, server.config.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}
	
	rsp := &gen.LoginUserResponse{
		AccessToken:           access,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refresh,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
		User:                  convertUser(user),
	}
	return rsp, nil
}