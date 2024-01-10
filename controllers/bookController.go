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

func (bc *BookController) GetBooks(c *gin.Context) {
	books, err := bc.Repository.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, books)
	}
}

func (bc *BookController) GetBook(c *gin.Context) {
	id := c.Param("id")

	book, err := bc.Repository.GetBookById(id)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func (bc *BookController) PostBook(c *gin.Context) {
	var book repositories.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		errs := errors.Wrap(err, "invalid json body")
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}

	bookID, err := bc.Repository.Make(book)
	if err != nil {
		errs := errors.Wrap(err, "failed to create book")
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created", "bookID": bookID})
}

func (bc *BookController) PatchBook(c *gin.Context) {
	var updateBook repositories.Book
	if err := c.BindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad_request"})
		return
	}

	id := c.Param("id")

	err := bc.Repository.Update(updateBook, id)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated", "book": updateBook})
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := bc.Repository.DeleteBookById(id)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
