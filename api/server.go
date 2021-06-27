package main

import (
	"context"
	"log"
	"net"

	pb "grpc-web-sample/helloworld/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9090"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", request.GetName())
	return &pb.HelloReply{Message: "Hello: " + request.GetName()}, nil
}

func (s server) SayRepeatHello(request *pb.RepeatHelloRequest, server pb.Greeter_SayRepeatHelloServer) error {
	panic("implement me")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	log.Printf("start gRPC server")
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
