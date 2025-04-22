package handler

import (
	"github.com/clpowe/oop-todo/internal/service"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	svc service.TodoService
}

func NewTodoHandler(svc service.TodoService) *TodoHandler {
	return &TodoHandler{svc}
}

func (h *TodoHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/", h.GetTodos)
	app.Get("/todos/:id", h.GetTodo)
	app.Post("/todos", h.CreateTodo)
	app.Post("/todos/:id/complete", h.CompleteTodo)
	app.Post("/todos/:id/incomplete", h.IncompleteTodo)
	app.Delete("/todos/:id", h.DeleteTodo)
}

func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	todos, err := h.svc.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "We're sorry something went wrong", "error": err.Error()})
	}
	return c.Render("todos.html", fiber.Map{"Todos": todos})
}

func (h *TodoHandler) GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := h.svc.Get(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "We're sorry something went wrong", "error": err.Error()})
	}
	return c.Render("todo.html", fiber.Map{"Todo": todo})
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	title := c.FormValue("title")
	if _, err := h.svc.Create(title); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "We're sorry something went wrong", "error": err.Error()})
	}
	return h.GetTodos(c)
}

func (h *TodoHandler) CompleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.Complete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "We're sorry something went wrong", "error": err.Error()})
	}
	return h.GetTodo(c)
}

func (h *TodoHandler) IncompleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.Incomplete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "We're sorry something went wrong", "error": err.Error()})
	}
	return h.GetTodo(c)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.Delete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "We're sorry something went wrong", "error": err.Error()})
	}
	return h.GetTodos(c)
}
