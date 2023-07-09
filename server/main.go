package main

import (
	pb "bookshelf/proto"
	"bookshelf/service"
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic("Failed to listen gRPC server!")
	}

	s := grpc.NewServer()

	pb.RegisterBookServiceServer(s, &service.BookSrv{})

	if err := s.Serve(lis); err != nil {
		panic("Failed to start gRPC Server!")
	}
}
