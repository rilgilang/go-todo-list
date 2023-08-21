package routes

import (
	"github.com/gofiber/fiber/v2"
	"simple-todo-list/api/handlers"
	"simple-todo-list/pkg/book"
)

// BookRouter is the Router for GoFiber App
func LoginRouter(app fiber.Router, service book.Service) {
	app.Get("/login", handlers.GetBooks(service))
	app.Post("/register", handlers.AddBook(service))
}
