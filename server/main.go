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
	fmt.Println("Running gRPC...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		fmt.Println("Error:", err)
		panic("Failed to listen gRPC server!")
	}

	s := grpc.NewServer()

	pb.RegisterBookServiceServer(s, &service.BookSrv{})

	fmt.Printf("gRPC Server running on port %v", *port)

	if errS := s.Serve(lis); errS != nil {
		panic("Failed to start gRPC Server!")
	}

}
