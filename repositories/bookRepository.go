package repositories

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepositoryInterface interface {
	Get() ([]Book, error)
	Make(Book) (string, error)
	Update(Book, string) error
	GetBookById(string) (*Book, error)
	DeleteBookById(string) error
}

type BookRepository struct {
	*gorm.DB
}

type Book struct {
	//gorm.Model
	ID     string  `gorm:"primarykey"`
	Title  string  `json:"title" validate:"required,max=256"`
	Author string  `json:"author" validate:"required,max=256"`
	Price  float64 `json:"price" validate:"required,min=0"`
}

func (db *BookRepository) Get() ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (db *BookRepository) Make(book Book) (string, error) {
	validate := validator.New()
	// Валидация полей структуры Book
	if err := validate.Struct(book); err != nil {
		return "", err
	}

	book.ID = uuid.New().String() // устанавливаем id книги используя UUID
	// Создаем новую запись книги в бд
	if err := db.Create(&book).Error; err != nil {
		return "", err
	}
	// Если ошибки нет, возвращаем UUID новой книги и nil в качестве значения ошибки
	return book.ID, nil
}

func (db *BookRepository) Update(book Book, id string) error {
	validate := validator.New()
	// Валидация полей структуры Book
	if err := validate.Struct(book); err != nil {
		return err
	}

	if err := db.Where("id = ?", id).First(&Book{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book with id '%s' not found", id)
		}
		return err
	}
	return db.Model(&Book{}).Where("id = ?", id).Updates(book).Error
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
	result := db.Where("id = ?", id).Delete(&Book{})
	if result.Error != nil {
		return result.Error // Ошибка связанная с бд
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no book found with id '%s'", id) // Нет записи для удаления
	}
	return nil
}
