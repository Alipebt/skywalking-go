package main

import (
	"fmt"
	"net"

	pb "test/plugins/scenarios/grpc/api"

	_ "github.com/apache/skywalking-go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSendMsgServer
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Printf("listen port error: %v\n", err)
		return
	}

	// 新建grpc的服务
	grpcServer := grpc.NewServer()

	// 注册服务
	pb.RegisterSendMsgServer(grpcServer, &server{})

	// 启动
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("start grpcServer error: %v\n", err)
		return
	}
}
