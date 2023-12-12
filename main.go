package main

import (
	"PetProject/controllers"
	"PetProject/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var repository = repositories.BookRepository{}

	var controller = controllers.BookController{
		Repository: &repository,
	}

	r.GET("/books", controller.GetBooks)
	r.POST("/books", controller.PostBooks)

	fmt.Println("попробую откатить изменения")

	r.Run(":8080")

}
