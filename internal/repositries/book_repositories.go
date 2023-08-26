package repositries

import (
	"gorm.io/gorm"
	"simple-todo-list/internal/api/presenter"
	"simple-todo-list/internal/entities"
	"time"
)

// Repository interface allows us to access the CRUD Operations in sql here.
type BookRepository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	ReadBook() (*[]presenter.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(ID string) error
}
type bookRepository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewBookRepo(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

// CreateBook is a gorm repository that helps to create books
func (r *bookRepository) CreateBook(book *entities.Book) (*entities.Book, error) {
	result := r.db.Create(book)

	return book, result.Error
}

// ReadBook is a gorm repository that helps to fetch books
func (r *bookRepository) ReadBook() (*[]presenter.Book, error) {
	var books []presenter.Book
	result := r.db.Where("deleted_at IS NULL").Find(&books)
	return &books, result.Error
}

// UpdateBook is a gorm repository that helps to update books
func (r *bookRepository) UpdateBook(book *entities.Book) (*entities.Book, error) {
	book.UpdatedAt = time.Now()
	result := r.db.Save(book)

	return book, result.Error
}

// DeleteBook is a mongo repository that helps to delete books
func (r *bookRepository) DeleteBook(ID string) error {
	err := r.db.Where("id = ?", ID).Delete(&entities.Book{}).Error

	return err
}
