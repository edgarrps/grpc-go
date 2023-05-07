package main

import "github.com/edgarrps/grpc-go/pb"

type Server struct {
	pb.UnimplementedHelloServer
}
