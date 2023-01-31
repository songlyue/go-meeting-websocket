package main

import (
	"context"
	"fmt"
	proto "go-meeting-websocket/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (reply *proto.HelloReply, e error) {
	return &proto.HelloReply{Message: "hello" + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50054")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("启动成")
		log.Fatalf("failed to serve: %v", err)
	}

}
