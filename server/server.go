package server

import (
	"context"
	"fmt"
	"github.com/sha1n/grpc-echo-service/gen/echo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}
type StopServer func()

func (s *server) Echo(ctx context.Context, in *echo.Request) (*echo.Response, error) {
	return &echo.Response{Message: in.Message}, nil
}

func Start(address string) StopServer {
	log.Println(fmt.Sprintf("starting echo server on %s", address))

	retVal := make(chan StopServer)
	go func() {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		echo.RegisterEchoServiceServer(s, &server{})

		retVal <- s.Stop

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}()

	return <-retVal
}
