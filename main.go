package main

import (
	"PetProject/controllers"
	"PetProject/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	db := repositories.NewDatabase()

	r := gin.Default()

	var repository = repositories.BookRepository{DB: db}

	var controller = controllers.BookController{
		Repository: &repository,
	}

	r.GET("/books", controller.GetBooks)
	r.GET("/books/:id", controller.GetBook)
	r.POST("/books", controller.PostBooks)
	r.PATCH("/books/:id", controller.PatchBooks)
	r.DELETE("/books/:id", controller.DeleteBook)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
