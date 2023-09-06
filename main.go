package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"simple-todo-list/bootstrap"
	routes2 "simple-todo-list/internal/api/routes"
	"simple-todo-list/internal/middlewares/jwt"
	"simple-todo-list/internal/repositries"
	"simple-todo-list/internal/service"
	"simple-todo-list/migrations"
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

	//bookRepo := repositries.NewBookRepo(db)
	userRepo := repositries.NewUserRepo(db)
	todoRepo := repositries.NewTodoRepo(db)

	middleware := jwt.NewAuthMiddleware(userRepo)

	//bookService := service.NewBookService(bookRepo)
	todoService := service.NewTodoService(todoRepo)
	userService := service.NewAuthService(middleware, userRepo)

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
			fiber.MethodPut,
			fiber.MethodDelete,
		}, ","),
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture todo list!"))
	})

	api := app.Group("/api")
	//routes2.BookRouter(api, middleware, bookService)
	routes2.TodoRouter(api, middleware, todoService)
	routes2.LoginRouter(api, userService)

	log.Fatal(app.Listen(":8080"))
}
