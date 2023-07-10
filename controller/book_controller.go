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
