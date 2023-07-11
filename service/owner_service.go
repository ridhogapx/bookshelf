package service

import (
	pb "bookshelf/proto"
	"context"
)

type OwnerSrv struct {
	pb.UnimplementedOwnerServiceServer
}

// Should add new record
func (*OwnerSrv) CreateOwner(ctx context.Context, req *pb.CreateOwnerRequest) (*pb.CreateOwnerResponse, error) {

}
