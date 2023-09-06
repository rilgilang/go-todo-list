package presenter

import (
	"github.com/gofiber/fiber/v2"
	"simple-todo-list/internal/entities"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func TodoSuccessResponse(data *entities.Todo) *fiber.Map {
	todo := Todo{
		ID:        data.ID,
		Title:     data.Title,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
	}
	return &fiber.Map{
		"status": true,
		"data":   todo,
		"error":  nil,
	}
}

func TodoSuccessUpdateResponse(data *entities.Todo) *fiber.Map {
	todo := Todo{
		ID:     data.ID,
		Status: data.Status,
	}
	return &fiber.Map{
		"status": true,
		"data":   todo,
		"error":  nil,
	}
}

func TodosSuccessResponse(data *[]Todo) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func TodoErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
