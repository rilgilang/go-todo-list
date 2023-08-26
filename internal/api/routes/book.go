package routes

import (
	"github.com/gofiber/fiber/v2"
	"simple-todo-list/internal/api/handlers"
	"simple-todo-list/internal/middlewares/jwt"
	"simple-todo-list/internal/service"
)

// BookRouter is the Router for GoFiber App
func BookRouter(app fiber.Router, middleware jwt.AuthMiddleware, service service.BookService) {
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/books", handlers.AddBook(service))
	app.Put("/books", handlers.UpdateBook(service))
	app.Delete("/books", handlers.RemoveBook(service))
}
