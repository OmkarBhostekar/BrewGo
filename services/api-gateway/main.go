package main

import (
	"context"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/omkarbhostekar/brewgo/api-gateway/util"
	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func runGateway(config util.Config) error {
	ctx := context.Background()
	mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gen.RegisterUserServiceHandlerFromEndpoint(ctx, mux, config.UserServiceEndPoint, opts)
    if err != nil {
        return err
    }
	err = gen.RegisterProductServiceHandlerFromEndpoint(ctx, mux, config.ProductServiceEndPoint, opts)
	if err != nil {
		return err
	}
    log.Info().Msg("Starting HTTP/REST gateway on :3000")
    return http.ListenAndServe(":3000", mux)
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