package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/omkarbhostekar/brewgo/services/user/api"
	db "github.com/omkarbhostekar/brewgo/services/user/db/sqlc"
	"github.com/omkarbhostekar/brewgo/services/user/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	store := db.NewStore(conn)

	log.Info().Msg("connected to db")

	// go runGatewayServer(config, store)
	runGrpcServer(config, store)
}

func metadataInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	userID := metadata.ValueFromIncomingContext(ctx, "X-User-Id")
	role := metadata.ValueFromIncomingContext(ctx, "X-Role")

	log.Info().Msgf("gRPC Server - Received Metadata: X-User-Id=%s, X-Role=%s", userID, role)

	// Proceed with the request
	return handler(ctx, req)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(metadataInterceptor))
	gen.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:3001")
	if err != nil {
		log.Fatal().Msg("cannot create listener:")
	}
	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Msg("cannot start grpc server")
	}

}

func runGatewayServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = gen.RegisterUserServiceHandlerServer(ctx, grpcMux, server)

	if err != nil {
		log.Fatal().Msg("cannot register gateway server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", "0.0.0.0:3002")
	if err != nil {
		log.Fatal().Msg("cannot create listener")
	}
	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal().Msg("cannot start HTTP gateway server")
	}

}