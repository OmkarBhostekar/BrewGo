package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/omkarbhostekar/brewgo/services/user/token"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
)

func (server *UserServer) authorizeUser(Ctx context.Context) (*token.Payload, error) {
	mtdt, ok := metadata.FromIncomingContext(Ctx)

	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}

	values := mtdt.Get("authorization")	
	if len(values) == 0 {
		return nil, fmt.Errorf("authorization token is not provided")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) != 2 || fields[0] != "Bearer" {
		return nil, fmt.Errorf("invalid authorization token format")
	}

	authType := strings.ToLower(fields[0])

	if authType != "bearer" {	
		return nil, fmt.Errorf("unsupported authorization type")
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)

	if err != nil {
		return nil, fmt.Errorf("access token is not valid")
	}

	return payload, nil
}

func (server *UserServer) getRoleFromHeader(Ctx context.Context) (string, error) {
	mtdt, ok := metadata.FromIncomingContext(Ctx)
	if !ok {
		return "", fmt.Errorf("metadata is not provided")
	}

	values := mtdt.Get("X-Role")
	serviceToken := mtdt.Get("X-Service-Token")
	log.Info().Msgf("service token: %v", serviceToken)
	if len(serviceToken) != 0 && serviceToken[0] == server.config.ServiceToken {
		return "service", nil
	}

	if len(values) == 0 {
		return "", fmt.Errorf("role is not provided")
	}

	return values[0], nil
}