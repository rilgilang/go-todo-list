package routes

import (
	"github.com/gofiber/fiber/v2"
	"simple-todo-list/internal/api/handlers"
	"simple-todo-list/internal/service"
)

// BookRouter is the Router for GoFiber App
func LoginRouter(app fiber.Router, service service.AuthService) {
	app.Post("/login", handlers.Login(service))
	app.Post("/register", handlers.Register(service))
}
