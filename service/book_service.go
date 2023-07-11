package service

import (
	"bookshelf/config"
	"bookshelf/model"
	pb "bookshelf/proto"
	"context"
	"errors"
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

func (*BookSrv) GetBook(ctx context.Context, req *pb.ReadBookRequest) (*pb.ReadBookResponse, error) {
	var book model.Book

	reqBook := req.GetOwner()

	result := config.DB.First(&book, "owner = ?", reqBook)
	if result.Error != nil {
		return nil, errors.New("data is not found")
	}

	return &pb.ReadBookResponse{
		Book: &pb.Book{
			Id:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Owner:  book.Owner,
		},
	}, nil
}

func (*BookSrv) GetBooks(ctx context.Context, req *pb.ReadBooksRequest) (*pb.ReadBooksResponse, error) {
	var books = []*pb.Book{}

	result := config.DB.Find(&books)

	if result.Error != nil {
		return nil, errors.New("record not found")
	}

	return &pb.ReadBooksResponse{
		Books: books,
	}, nil
}

func (*BookSrv) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	arg := req.GetBook()
	var book model.Book

	q := config.DB.Model(&book).Where("id = ?", arg.GetId()).Updates(model.Book{
		ID:     arg.GetId(),
		Title:  arg.GetTitle(),
		Author: arg.GetAuthor(),
		Owner:  arg.GetOwner(),
	})

	if q.Error != nil {
		return nil, errors.New("failed to update data")
	}

	return &pb.UpdateBookResponse{
		Book: &pb.Book{
			Id:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Owner:  book.Owner,
		},
	}, nil
}

func (*BookSrv) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	arg := req.GetId()
	var book model.Book

	result := config.DB.Where("id = ?", arg).Delete(&book)

	if result.Error != nil {
		return &pb.DeleteBookResponse{
			Success: false,
		}, errors.New("failed to delete record")
	}

	return &pb.DeleteBookResponse{
		Success: true,
	}, nil
}
