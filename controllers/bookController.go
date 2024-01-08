package controllers

import (
	"PetProject/errors"
	"PetProject/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		c.IndentedJSON(http.StatusBadRequest, errors.Err{Error: "bad_request"})
		return
	}

	err := res.Repository.Make(newBooks)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, newBooks)
}

func (res *BookController) PatchBooks(c *gin.Context) {
	var updateBook repositories.Book
	if err := c.BindJSON(&updateBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.Err{Error: "bad_request"})
		return
	}

	id := c.Param("id")

	err := res.Repository.Update(updateBook, id)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, updateBook)
}

func (res *BookController) GetBook(c *gin.Context) {
	id := c.Param("id")

	book, err := res.Repository.GetBookById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// если книга по указанному id не найдена
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		} else {
			// любая другая ошибка
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func (res *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := res.Repository.DeleteBookById(id)
	if err != nil {
		// если книга по указанному id не найдена
		if err == gorm.ErrRecordNotFound {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		} else {
			// любая другая ошибка
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	//Если удаление прошло успешно, возвращается статус-код 204, указывающий на отсутствие содержимого
	c.Status(http.StatusNoContent)
}
