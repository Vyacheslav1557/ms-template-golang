package app

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"ms-template-golang/internal/config"
	"ms-template-golang/internal/logger"
	"ms-template-golang/internal/services"
	"ms-template-golang/internal/transport"
	"net"
	"os"
)

type App struct {
	server *grpc.Server
}

func New() *App {
	var (
		server  *grpc.Server
		service services.Service
	)

	server = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				transport.LoggingInterceptor,
			),
		),
		grpc.StreamInterceptor(
			transport.StreamLoggingInterceptor,
		),
	)
	service = services.Service{}

	transport.Register(server, service)

	return &App{
		server: server,
	}
}

func (a *App) Run() {
	var (
		cfg      = config.Cfg()
		log      = logger.Logger()
		err      error
		listener net.Listener
		address  = fmt.Sprintf("%s:%d", cfg.HOST, cfg.PORT)
	)

	listener, err = net.Listen("tcp", address)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info(fmt.Sprintf("Listening on %s", address))
	if err = a.server.Serve(listener); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func (a *App) Stop() {
	var (
		log = logger.Logger()
	)
	log.Info("Gracefully shutting down")
	a.server.GracefulStop()
}
