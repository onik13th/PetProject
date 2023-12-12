package repositories

type BookRepositoryInterface interface {
	Create(Book)
	Get() []Book
}

type BookRepository struct {
	BookRepositoryInterface
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []Book{
	{ID: "1", Title: "Капитал", Author: "Карл Маркс", Price: 69.69},
	{ID: "2", Title: "Государство и революция", Author: "Владимир Ленин", Price: 49.99},
	{ID: "3", Title: "Что делать?", Author: "Николай Чернышевский", Price: 35.99},
	{ID: "4", Title: "Мартин Иден", Author: "Джек Лондон", Price: 40.00},
}

func (res *BookRepository) Create(newElement Book) {
	books = append(books, newElement)
}

func (res *BookRepository) Get() []Book {
	return books
}
