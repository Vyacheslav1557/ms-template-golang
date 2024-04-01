package transport

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"ms-template-golang/internal/logger"
	"ms-template-golang/pkg/go/gen"
)

type Service interface {
	Hello(name string) string
}

type server struct {
	gen.UnimplementedSampleServiceServer
	service Service
}

func Register(gRPCServer *grpc.Server, service Service) {
	gen.RegisterSampleServiceServer(gRPCServer, &server{service: service})
}

func (s *server) SayHello(ctx context.Context, in *gen.HelloRequest) (*gen.HelloResponse, error) {
	return &gen.HelloResponse{Greeting: s.service.Hello(in.GetName())}, nil
}

func (s *server) Count(stream gen.SampleService_CountServer) error {
	var (
		log = logger.Logger()
		err error
		req *gen.CountRequest
		i   = 0
	)
	for {
		err = stream.Context().Err()
		if err != nil {
			return err
		}
		req, err = stream.Recv()
		if err == io.EOF {
			break
		}
		log.Info("got stream message")
		err = stream.Send(&gen.CountResponse{Greeting: s.service.Hello(req.Name), Count: int64(i)})
		log.Info("sent stream message")
		if err != nil {
			log.Error(err.Error())
			return status.Errorf(codes.Unknown, "cannot receive stream request: %v", err)
		}
		i++
	}

	return nil
}
