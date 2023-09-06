package service

import (
	"github.com/pkg/errors"
	"simple-todo-list/internal/api/presenter"
	"simple-todo-list/internal/entities"
	"simple-todo-list/internal/repositries"
)

type TodoService interface {
	InsertTodo(book *entities.Todo) (*entities.Todo, error)
	FetchTodo(userId int) (*[]presenter.Todo, error)
	UpdateTodo(todo *entities.Todo) (*entities.Todo, error)
	RemoveTodo(userId int, ID int) error
}

type todoService struct {
	repository repositries.TodoRepository
}

func NewTodoService(r repositries.TodoRepository) TodoService {
	return &todoService{
		repository: r,
	}
}

func (s *todoService) InsertTodo(book *entities.Todo) (*entities.Todo, error) {
	return s.repository.CreateTodo(book)
}

func (s *todoService) FetchTodo(userId int) (*[]presenter.Todo, error) {
	return s.repository.ReadTodo(userId)
}

func (s *todoService) UpdateTodo(todo *entities.Todo) (*entities.Todo, error) {
	todoExist, err := s.repository.CheckTodo(todo.UserID, todo.ID)
	if err != nil {
		return nil, err
	}

	if todoExist == nil {
		return nil, nil
	}

	return s.repository.UpdateTodo(todo)
}

func (s *todoService) RemoveTodo(userId int, todoId int) error {
	todoExist, err := s.repository.CheckTodo(userId, todoId)
	if err != nil {
		return err
	}

	if todoExist == nil {
		return errors.New("todo not found")
	}

	return s.repository.DeleteTodo(userId, todoId)
}
