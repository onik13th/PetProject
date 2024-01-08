package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type BookRepositoryInterface interface {
	Get() []Book
	Make(Book) error
	Update(Book, string) error
	GetBookById(string) (*Book, error)
	DeleteBookById(string) error
}

type BookRepository struct {
	*gorm.DB
}

type Book struct {
	gorm.Model
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

func (db *BookRepository) Get() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (db *BookRepository) Make(book Book) error {
	result := db.Create(&book)
	return result.Error
}

func (db *BookRepository) Update(book Book, id string) error {
	// Сначала находим книгу по идентификатору.
	var existingBook Book
	if err := db.DB.First(&existingBook, "id = ?", id).Error; err != nil {
		return err
	}

	// Теперь обновляем книгу новыми данными
	return db.DB.Model(&existingBook).Updates(book).Error
}

func (db *BookRepository) GetBookById(id string) (*Book, error) {
	var book Book
	if err := db.DB.First(&book, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book with id '%s' not found", id)
		}
		// Другие ошибки
		return nil, err
	}
	return &book, nil
}

func (db *BookRepository) DeleteBookById(id string) error {
	result := db.DB.Delete(&Book{}, "id = ?", id)
	if result.Error != nil {
		return result.Error // ошибка связанная с бд
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no book found with id '%s'", id) // нет записи для удаления
	}
	return fmt.Errorf("deletion was successful", nil)
}
