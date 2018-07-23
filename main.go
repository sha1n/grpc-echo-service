package main

import (
	"context"
	"github.com/sha1n/grpc-test/server"
	"google.golang.org/grpc"
	"log"
	"time"
	"github.com/sha1n/grpc-test/echo"
)

const address = "localhost:50051"

func main() {
	// Start a grpc echo echoServer
	stop := server.Start(address)
	defer stop()

	// connect
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// create a grpc client
	c := echo.NewEchoServiceClient(conn)

	message := "Hey, I'm just playing with grpc..."
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call the bastard
	r, err := c.Echo(ctx, &echo.Request{Message: message})

	if err != nil {
		log.Fatalf("echoServer request has failed :_( : %v", err)
	} else {
		log.Printf("the echoServer responded with: %s", r.Message)
	}
}
