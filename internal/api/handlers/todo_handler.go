package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"net/http"
	"simple-todo-list/internal/api/presenter"
	"simple-todo-list/internal/api/request_model"
	"simple-todo-list/internal/consts"
	"simple-todo-list/internal/entities"
	"simple-todo-list/internal/helper"
	"simple-todo-list/internal/service"
)

func AddTodo(service service.TodoService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request_model.CreateTodo
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		if requestBody.Title == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TodoErrorResponse(errors.New(
				"Please specify title and author")))
		}

		todo := entities.Todo{
			UserID: helper.InterfaceToInt(c.Locals(consts.UserId)),
			Title:  requestBody.Title,
			Status: false,
		}

		result, err := service.InsertTodo(&todo)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		return c.JSON(presenter.TodoSuccessResponse(result))
	}
}

func UpdateTodo(service service.TodoService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request_model.UpdateTodo
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TodoErrorResponse(err))
		}

		todo := entities.Todo{
			ID:     requestBody.Id,
			UserID: helper.InterfaceToInt(c.Locals(consts.UserId)),
			Status: requestBody.Status,
		}

		result, err := service.UpdateTodo(&todo)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}

		//to do not found
		if result == nil {
			c.Status(http.StatusNotFound)
			return c.JSON(presenter.TodoErrorResponse(errors.New("todo not found")))
		}

		return c.JSON(presenter.TodoSuccessUpdateResponse(result))
	}
}

func RemoveTodo(service service.TodoService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request_model.DeleteTodo
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		userId := helper.InterfaceToInt(c.Locals(consts.UserId))
		todoId := requestBody.Id
		err = service.RemoveTodo(userId, todoId)
		if err != nil {

			if err.Error() == "todo not found" {
				c.Status(http.StatusNotFound)
				return c.JSON(presenter.TodoErrorResponse(err))
			}

			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated successfully",
			"err":    nil,
		})
	}
}

func GetTodos(service service.TodoService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := helper.InterfaceToInt(c.Locals(consts.UserId))
		fetched, err := service.FetchTodo(userId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TodoErrorResponse(err))
		}
		return c.JSON(presenter.TodosSuccessResponse(fetched))
	}
}
