package main

import (
	"context"
	"log"
	"net"

	"test/plugins/scenarios/grpc/api"

	_ "github.com/apache/skywalking-go"
	"google.golang.org/grpc"
)

type Echo struct {
	api.UnimplementedEchoServer
}

func (e *Echo) UnaryEcho(ctx context.Context, req *api.EchoRequest) (*api.EchoResponse, error) {
	log.Printf("Recved: %v", req.GetMessage())
	resp := &api.EchoResponse{Message: req.GetMessage()}
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
		return
	}

	s := grpc.NewServer()

	api.RegisterEchoServer(s, &Echo{})

	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}
