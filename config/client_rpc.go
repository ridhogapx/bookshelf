package config

import (
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("gRPC address", "localhost:50051", "Address for dialing gRPC Server")
)

var errorRPC error
var RpcClient *grpc.ClientConn

func ConnectRPC() {
	flag.Parse()

	RpcClient, errorRPC = grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if errorRPC != nil {
		panic("Failed to connect gRPC Server!")
	}

}
