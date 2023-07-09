package service

import (
	pb "bookshelf/proto"
)

type BookSrv struct {
	pb.UnimplementedBookServiceServer
}
