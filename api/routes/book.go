package routes

import (
	"github.com/gofiber/fiber/v2"
	"simple-todo-list/api/handlers"
	"simple-todo-list/pkg/book"
)

// BookRouter is the Router for GoFiber App
func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/books", handlers.AddBook(service))
	app.Put("/books", handlers.UpdateBook(service))
	app.Delete("/books", handlers.RemoveBook(service))
}
