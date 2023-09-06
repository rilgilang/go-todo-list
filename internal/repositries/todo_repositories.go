package repositries

import (
	"gorm.io/gorm"
	"simple-todo-list/internal/api/presenter"
	"simple-todo-list/internal/consts"
	"simple-todo-list/internal/entities"
	"time"
)

type TodoRepository interface {
	CheckTodo(userId int, Id int) (*entities.Todo, error)
	CreateTodo(book *entities.Todo) (*entities.Todo, error)
	ReadTodo(userId int) (*[]presenter.Todo, error)
	UpdateTodo(book *entities.Todo) (*entities.Todo, error)
	DeleteTodo(userId int, ID int) error
}
type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) CheckTodo(userId int, Id int) (*entities.Todo, error) {
	var todo entities.Todo
	err := r.db.Where("id = ?", Id).Where("user_id = ?", userId).First(&todo).Error
	if err != nil && err.Error() == consts.SqlNoRow {
		return nil, nil
	}
	return &todo, err
}

func (r *todoRepository) CreateTodo(todo *entities.Todo) (*entities.Todo, error) {
	result := r.db.Create(todo)

	return todo, result.Error
}

func (r *todoRepository) ReadTodo(userId int) (*[]presenter.Todo, error) {
	var todos []presenter.Todo
	result := r.db.Where("user_id = ?", userId).Where("deleted_at IS NULL").Find(&todos)
	return &todos, result.Error
}

func (r *todoRepository) UpdateTodo(todo *entities.Todo) (*entities.Todo, error) {
	todo.UpdatedAt = time.Now()
	result := r.db.Model(&todo).Update("status", todo.Status).Update("updated_at", todo.UpdatedAt).Where("id", todo.ID)

	return todo, result.Error
}

func (r *todoRepository) DeleteTodo(userId int, ID int) error {
	err := r.db.Where("id = ?", ID).Where("user_id = ?", userId).Delete(&entities.Todo{}).Error

	return err
}
