package main

import (
	"database/sql"
	"net"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/omkarbhostekar/brewgo/proto/gen"
	"github.com/omkarbhostekar/brewgo/services/user/api"
	db "github.com/omkarbhostekar/brewgo/services/user/db/sqlc"
	"github.com/omkarbhostekar/brewgo/services/user/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func runGrpcServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}
	grpcServer := grpc.NewServer()
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

// func runGatewayServer(config util.Config, store db.Store) {
// 	server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal().Msg("cannot create server")
// 	}
// 	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
// 		MarshalOptions: protojson.MarshalOptions{
// 			UseProtoNames: true,
// 		},
// 		UnmarshalOptions: protojson.UnmarshalOptions{
// 			DiscardUnknown: true,
// 		},
// 	})
// 	grpcMux := runtime.NewServeMux(jsonOption)
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	err = gen.RegisterUserServiceHandlerServer(ctx, grpcMux, server)

// 	if err != nil {
// 		log.Fatal().Msg("cannot register gateway server")
// 	}

// 	mux := http.NewServeMux()
// 	mux.Handle("/", grpcMux)

// 	listener, err := net.Listen("tcp", "0.0.0.0:3002")
// 	if err != nil {
// 		log.Fatal().Msg("cannot create listener")
// 	}
// 	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
// 	err = http.Serve(listener, mux)
// 	if err != nil {
// 		log.Fatal().Msg("cannot start HTTP gateway server")
// 	}

// }