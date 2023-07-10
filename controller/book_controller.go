package controller

import (
	"bookshelf/entity"

	pb "bookshelf/proto"

	"github.com/gin-gonic/gin"
)

var client = pb.NewBookServiceClient()

func AddBook(ctx gin.Context) {
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

}
