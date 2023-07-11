package main

import (
	"bookshelf/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// Intialize Gin
	r := gin.Default()

	// Grouping Route
	v1 := r.Group("/api/v1")
	{
		v1.POST("/book", controller.AddBook)
		v1.GET("/book/:owner", controller.OwnerBook)
		v1.GET("/book", controller.AllBook)
		v1.PUT("/book/:id", controller.EditBook)
	}

	r.Run(":3000")

}
