package controllers

import (
	"PetProject/errors"
	"PetProject/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
	Repository repositories.BookRepositoryInterface
}

func (res *BookController) GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, res.Repository.Get())
}

func (res *BookController) PostBooks(c *gin.Context) {
	var newBooks repositories.Book
	if err := c.BindJSON(&newBooks); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.Err{"bad_request"})
		return
	}

	res.Repository.Create(newBooks)
	c.IndentedJSON(http.StatusCreated, newBooks)
}
