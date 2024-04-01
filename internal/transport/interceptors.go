package transport

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"ms-template-golang/internal/logger"
	"time"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	var (
		log = logger.Logger()
	)
	t0 := time.Now()
	resp, err = handler(ctx, req)
	t1 := time.Now()

	d := t1.Sub(t0)
	log.Info("unary request",
		slog.String("method", info.FullMethod),
		slog.String("duration", fmt.Sprintf("%v", d)),
	)
	return resp, err
}

func StreamLoggingInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	var (
		log = logger.Logger()
		err error
	)
	log.Info("stream started",
		slog.String("method", info.FullMethod),
	)
	t0 := time.Now()
	err = handler(srv, ss)
	t1 := time.Now()
	d := t1.Sub(t0)
	log.Info("stream ended",
		slog.String("method", info.FullMethod),
		slog.String("duration", fmt.Sprintf("%v", d)),
	)
	return err
}
