package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"simple-todo-list/bootstrap"
	"simple-todo-list/config/yaml"
	routes2 "simple-todo-list/internal/api/routes"
	"simple-todo-list/internal/middlewares/jwt"
	"simple-todo-list/internal/repositries"
	"simple-todo-list/internal/service"
	"simple-todo-list/migrations"
)

func main() {
	cfg, err := yaml.NewConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf(`read cfg yaml got error : %v`, err))
	}

	db, err := bootstrap.DatabaseConnection(cfg)
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

	middleware := jwt.NewAuthMiddleware(userRepo, cfg)

	//bookService := service.NewBookService(bookRepo)
	todoService := service.NewTodoService(todoRepo)
	userService := service.NewAuthService(middleware, userRepo)

	app := fiber.New()
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture todo list!"))
	})

	api := app.Group("/api")
	//routes2.BookRouter(api, middleware, bookService)
	routes2.TodoRouter(api, middleware, todoService)
	routes2.LoginRouter(api, userService)

	log.Fatal(app.Listen(fmt.Sprintf(`:%s`, cfg.App.Port)))
}
