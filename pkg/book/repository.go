package book

import (
	"gorm.io/gorm"
	"simple-todo-list/api/presenter"
	"simple-todo-list/pkg/entities"
	"time"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	ReadBook() (*[]presenter.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(ID string) error
}
type repository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewRepo(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// CreateBook is a gorm repository that helps to create books
func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {
	result := r.db.Create(book)

	return book, result.Error
}

// ReadBook is a gorm repository that helps to fetch books
func (r *repository) ReadBook() (*[]presenter.Book, error) {
	var books []presenter.Book
	result := r.db.Where("deleted_at IS NULL").Find(&books)
	return &books, result.Error
}

// UpdateBook is a gorm repository that helps to update books
func (r *repository) UpdateBook(book *entities.Book) (*entities.Book, error) {
	book.UpdatedAt = time.Now()
	result := r.db.Save(book)

	return book, result.Error
}

// DeleteBook is a mongo repository that helps to delete books
func (r *repository) DeleteBook(ID string) error {
	err := r.db.Where("id = ?", ID).Delete(&entities.Book{}).Error

	return err
}
