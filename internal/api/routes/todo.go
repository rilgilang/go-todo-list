package routes

import (
	"github.com/gofiber/fiber/v2"
	"simple-todo-list/internal/api/handlers"
	"simple-todo-list/internal/middlewares/jwt"
	"simple-todo-list/internal/service"
)

// BookRouter is the Router for GoFiber App
func TodoRouter(app fiber.Router, middleware jwt.AuthMiddleware, service service.TodoService) {
	app.Get("/todos", middleware.ValidateToken(), handlers.GetTodos(service))
	app.Post("/todo", middleware.ValidateToken(), handlers.AddTodo(service))
	app.Put("/todo", middleware.ValidateToken(), handlers.UpdateTodo(service))
	app.Delete("/todo", middleware.ValidateToken(), handlers.RemoveTodo(service))
}
