package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("gRPC address", "localhost:50051", "Address for dialing gRPC Server")
)

func main() {
	flag.Parse()

	// Connection
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("Failed to dialing gRPC server")
	}

	defer conn.Close()

	// Intialize Gin
	r := gin.Default()

}
