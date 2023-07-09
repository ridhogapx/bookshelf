package service

import (
	"bookshelf/db"
	"bookshelf/model"
	pb "bookshelf/proto"
	"context"
)

type BookSrv struct {
	pb.UnimplementedBookServiceServer
}

func (*BookSrv) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := req.GetBook()

	data := model.Book{
		ID:     book.GetId(),
		Title:  book.GetTitle(),
		Author: book.GetAuthor(),
		Owner:  book.GetOwner(),
	}

	db.DB.Create(&data)

	return &pb.CreateBookResponse{
		Book: &pb.Book{
			Id:     book.GetId(),
			Title:  book.GetTitle(),
			Author: book.GetAuthor(),
			Owner:  book.GetOwner(),
		},
	}, nil

}
