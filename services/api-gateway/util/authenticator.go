package util

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// authenticationMiddleware authenticates incoming HTTP requests.
func AuthenticationMiddleware(userServiceEndpoint string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/users/login" || r.URL.Path == "/v1/users/create" { // Adjust the path
			next.ServeHTTP(w, r)
			return
		}

		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Warn().Msg("Missing Authorization header")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")
		if accessToken == "" {
			log.Warn().Msg("Malformed Authorization header")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validate the token by calling the UserService
		rsp, err := validateTokenWithUserService(accessToken, userServiceEndpoint)
		if err != nil {
			log.Error().Err(err).Msg("Token validation failed") // Log the error
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := metadata.AppendToOutgoingContext(r.Context(), "X-Role", rsp.Role)
		r = r.WithContext(ctx)
		r.Header.Set("X-Role", rsp.Role)
		next.ServeHTTP(w, r)
	})
}

// validateTokenWithUserService calls the UserService to validate the token.
func validateTokenWithUserService(accessToken string, userServiceEndPoint string) (*gen.ValidateTokenResponse, error) {
	conn, err := grpc.Dial(userServiceEndPoint, grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to user service")
		return nil, fmt.Errorf("failed to connect to user service") //Return an error
	}
	defer conn.Close()

	client := gen.NewUserServiceClient(conn)
	rsp, err := client.ValidateToken(context.Background(), &gen.ValidateTokenRequest{AccessToken: accessToken})
	if err != nil {
		log.Error().Err(err).Msg("cannot validate token")
		return nil, fmt.Errorf("invalid token")
	}

	return rsp, nil
}