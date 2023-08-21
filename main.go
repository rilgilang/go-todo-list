package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"simple-todo-list/api/routes"
	"simple-todo-list/bootstrap"
	"simple-todo-list/migrations"
	"simple-todo-list/pkg/book"
	"strings"
)

func main() {
	db, err := bootstrap.DatabaseConnection()
	if err != nil {
		log.Fatal(fmt.Sprintf(`db connection error got : %v`, err))
	}

	fmt.Println("Database connection success!")

	migrations.AutoMigration(db)

	if err != nil {
		log.Fatal(fmt.Sprintf(`error auto migrate got : %v`, err))
	}

	fmt.Println("Migration success!")

	bookRepo := book.NewRepo(db)
	bookService := book.NewService(bookRepo)

	app := fiber.New()
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	routes.BookRouter(api, bookService)

	log.Fatal(app.Listen(":8080"))
}
