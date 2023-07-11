package controller

import (
	"bookshelf/config"
	"bookshelf/entity"

	pb "bookshelf/proto"

	"github.com/gin-gonic/gin"
)

var client = pb.NewBookServiceClient(config.RpcClient)

func AddBook(ctx *gin.Context) {
	var book entity.BookEntity

	err := ctx.ShouldBind(&book)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	data := &pb.Book{
		Id:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Owner:  book.Owner,
	}

	res, err := client.CreateBook(ctx, &pb.CreateBookRequest{
		Book: data,
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Successfull adding data",
		"book":    res.Book,
	})

}

func OwnerBook(ctx *gin.Context) {

	res, err := client.GetBook(ctx, &pb.ReadBookRequest{
		Owner: ctx.Param("owner"),
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(200, entity.BookEntity{
		ID:     res.Book.Id,
		Title:  res.Book.Title,
		Author: res.Book.Author,
		Owner:  res.Book.Owner,
	})
}

func AllBook(ctx *gin.Context) {
	res, err := client.GetBooks(ctx, &pb.ReadBooksRequest{})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(200, res)

}

func EditBook(ctx *gin.Context) {
	var book entity.BookEntity

	bookID := ctx.Param("id")

	err := ctx.ShouldBind(&book)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	data := &pb.Book{
		Id:     bookID,
		Title:  book.Title,
		Author: book.Author,
		Owner:  book.Owner,
	}

	res, err := client.UpdateBook(ctx, &pb.UpdateBookRequest{
		Book: data,
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(201, res)

}
