package main

import (
	"context"
	"log"
	"net"

	"github.com/edgarrps/grpc-go/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	println("Running gRPC server")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
