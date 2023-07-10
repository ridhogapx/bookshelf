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
	}

	r.Run(":3000")

}
