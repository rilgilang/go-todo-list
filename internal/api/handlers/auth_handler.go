package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"net/http"
	"simple-todo-list/internal/api/presenter"
	"simple-todo-list/internal/consts"
	entities2 "simple-todo-list/internal/entities"
	"simple-todo-list/internal/service"
)

func Login(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities2.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		if requestBody.Username == "" || requestBody.Password == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(errors.New(
				"Please specify title and author")))
		}
		user, token, err := service.Login(&requestBody)
		//only internal server error
		if err != nil && err.Error() == consts.InternalServerError {
			c.Status(500)
			return c.JSON(presenter.AuthErrorResponse(err))
		}

		//can be unauthorized or something else
		if err != nil {
			c.Status(401)
			return c.JSON(presenter.AuthErrorResponse(err))
		}

		c.Status(200)
		return c.JSON(presenter.AuthSuccesResponse(user, *token))
	}
}

func Register(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities2.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		if requestBody.Username == "" || requestBody.Password == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(errors.New(
				"Please specify title and author")))
		}

		user, token, err := service.Register(&requestBody)
		//only internal server error
		if err != nil && err.Error() == consts.InternalServerError {
			c.Status(500)
			return c.JSON(presenter.AuthErrorResponse(err))
		}

		//can be unauthorized or something else
		if err != nil {
			c.Status(401)
			return c.JSON(presenter.AuthErrorResponse(err))
		}

		c.Status(200)
		return c.JSON(presenter.AuthSuccesResponse(user, *token))
	}
}
