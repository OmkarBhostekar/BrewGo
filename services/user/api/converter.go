package api

import (
	"github.com/omkarbhostekar/brewgo/proto/gen"
	db "github.com/omkarbhostekar/brewgo/services/user/db/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *gen.User {
	return &gen.User{
		Name:              user.Name,
		PhoneNumber:       user.PhoneNumber,
		Email:             user.Email,
		CreatedAt:         timestamppb.New(user.CreatedAt),
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
	}
}
