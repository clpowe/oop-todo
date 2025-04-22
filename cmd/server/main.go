package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/clpowe/oop-todo/internal/handler"
	"github.com/clpowe/oop-todo/internal/repository"
	"github.com/clpowe/oop-todo/internal/service"
)

func main() {
	engine := html.New("internal/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./web/static")

	repo := repository.NewInMemoryRepo()
	svc := service.NewTodoService(repo)
	h := handler.NewTodoHandler(svc)

	h.RegisterRoutes(app)

	log.Println("Starting server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
