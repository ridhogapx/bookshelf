package service

import (
	"bookshelf/config"
	"bookshelf/model"
	pb "bookshelf/proto"
	"context"
)

/*
* Todo:
* Implement method for retrieve all book data.
* Implement method for retrieve single book by its ID.
* Implement method for update book
* Implement method for delete book
 */

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

	config.DB.Create(&data)

	return &pb.CreateBookResponse{
		Book: &pb.Book{
			Id:     book.GetId(),
			Title:  book.GetTitle(),
			Author: book.GetAuthor(),
			Owner:  book.GetOwner(),
		},
	}, nil

}
