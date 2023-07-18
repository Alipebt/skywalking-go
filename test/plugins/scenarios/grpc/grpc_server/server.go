package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "test/plugins/scenarios/grpc/api"

	_ "github.com/apache/skywalking-go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SimpleRPC(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("client call simpleRPC...")
	log.Println(in)
	return &pb.HelloResponse{Reply: "Hello " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
		return
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
	log.Println("gRPC server starts running...")

	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}
