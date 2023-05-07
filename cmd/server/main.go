package main

import (
	"context"

	"github.com/edgarrps/grpc-go/pb"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + in.GetName()}, nil
}

func main() {

}
