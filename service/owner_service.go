package service

import (
	"bookshelf/config"
	"bookshelf/model"
	pb "bookshelf/proto"
	"context"
	"errors"
)

type OwnerSrv struct {
	pb.UnimplementedOwnerServiceServer
}

// Should add new record
func (*OwnerSrv) CreateOwner(ctx context.Context, req *pb.CreateOwnerRequest) (*pb.CreateOwnerResponse, error) {
	arg := req.GetOwner()

	data := model.Owner{
		ID:   arg.GetId(),
		Name: arg.GetName(),
	}

	result := config.DB.Create(&data)

	if result.Error != nil {
		return nil, errors.New("failed to add new record")
	}

	return &pb.CreateOwnerResponse{
		Owner: &pb.Owner{
			Id:   data.ID,
			Name: data.Name,
		},
	}, nil

}
