package service

import (
	"simple-todo-list/internal/api/presenter"
	"simple-todo-list/internal/entities"
	"simple-todo-list/internal/repositries"
)

// Service is an interface from which our api module can access our repository of all our models
type BookService interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	FetchBooks() (*[]presenter.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	RemoveBook(ID string) error
}

type bookService struct {
	repository repositries.BookRepository
}

// NewService is used to create a single instance of the service
func NewBookService(r repositries.BookRepository) BookService {
	return &bookService{
		repository: r,
	}
}

// InsertBook is a service layer that helps insert book in BookShop
func (s *bookService) InsertBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.CreateBook(book)
}

// FetchBooks is a service layer that helps fetch all books in BookShop
func (s *bookService) FetchBooks() (*[]presenter.Book, error) {
	return s.repository.ReadBook()
}

// UpdateBook is a service layer that helps update books in BookShop
func (s *bookService) UpdateBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.UpdateBook(book)
}

// RemoveBook is a service layer that helps remove books from BookShop
func (s *bookService) RemoveBook(ID string) error {
	return s.repository.DeleteBook(ID)
}
