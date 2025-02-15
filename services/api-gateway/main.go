package main

import (
	"context"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/omkarbhostekar/brewgo/api-gateway/util"
	"github.com/omkarbhostekar/brewgo/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func runGateway(config util.Config) error {
	ctx := context.Background()
	mux := runtime.NewServeMux(runtime.WithMetadata(ForwardMetadata))
	opts := []grpc.DialOption{grpc.WithInsecure()} //In production remove grpc.WithInsecure()

	// Register Handlers
	err := gen.RegisterUserServiceHandlerFromEndpoint(ctx, mux, config.UserServiceEndPoint, opts)
	if err != nil {
		return err
	}
	err = gen.RegisterCounterServiceHandlerFromEndpoint(ctx, mux, config.CounterServiceEndPoint, opts)
	if err != nil {
		return err
	}

	protectedMux := util.AuthenticationMiddleware(config.UserServiceEndPoint, mux)

	log.Info().Msg("Starting HTTP/REST gateway on :3000")
	return http.ListenAndServe(":3000", protectedMux)
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	runGateway(config)
}

// ForwardMetadata copies metadata from HTTP request to gRPC metadata
func ForwardMetadata(ctx context.Context, r *http.Request) metadata.MD {
	md := metadata.New(map[string]string{
		"X-Role": r.Header.Get("X-Role"), 
	})
	return md
}
 