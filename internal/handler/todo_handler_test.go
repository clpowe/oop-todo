package handler

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/clpowe/oop-todo/internal/model"
	"github.com/gofiber/fiber/v2"
)

// stubService provides a fake implementation of service.TodoService.
type stubService struct{}

func (s *stubService) Create(title string) (*model.Todo, error) {
	return nil, errors.New("create error")
}

func (s *stubService) List() (map[string]*model.Todo, error) {
	return nil, errors.New("list error")
}

func (s *stubService) Get(id string) (*model.Todo, error) {
	return nil, errors.New("get error")
}

func (s *stubService) Complete(id string) error {
	return errors.New("complete error")
}

func (s *stubService) Incomplete(id string) error {
	return errors.New("incomplete error")
}

func (s *stubService) Delete(id string) error {
	return errors.New("delete error")
}

func setupApp() *fiber.App {
	app := fiber.New()
	h := NewTodoHandler(&stubService{})
	h.RegisterRoutes(app)
	return app
}

func TestGetTodos_Error(t *testing.T) {
	app := setupApp()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, _ := app.Test(req)
	if res.StatusCode != fiber.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", res.StatusCode)
	}
	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "list error") {
		t.Errorf("expected error 'list error', got %s", string(body))
	}
}

func TestGetTodo_NotFound(t *testing.T) {
	app := setupApp()
	req := httptest.NewRequest(http.MethodGet, "/todos/123", nil)
	res, _ := app.Test(req)
	if res.StatusCode != fiber.StatusNotFound {
		t.Errorf("expected status 404, got %d", res.StatusCode)
	}
	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "get error") {
		t.Errorf("expected error 'get error', got %s", string(body))
	}
}

func TestCreateTodo_BadRequest(t *testing.T) {
	app := setupApp()
	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader("title=task"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, _ := app.Test(req)
	if res.StatusCode != fiber.StatusBadRequest {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}
	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "create error") {
		t.Errorf("expected error 'create error', got %s", string(body))
	}
}

func TestCompleteTodo_NotFound(t *testing.T) {
	app := setupApp()
	req := httptest.NewRequest(http.MethodPost, "/todos/123/complete", nil)
	res, _ := app.Test(req)
	if res.StatusCode != fiber.StatusNotFound {
		t.Errorf("expected status 404, got %d", res.StatusCode)
	}
	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "complete error") {
		t.Errorf("expected error 'complete error', got %s", string(body))
	}
}

func TestDeleteTodo_NotFound(t *testing.T) {
	app := setupApp()
	req := httptest.NewRequest(http.MethodDelete, "/todos/123", nil)
	res, _ := app.Test(req)
	if res.StatusCode != fiber.StatusNotFound {
		t.Errorf("expected status 404, got %d", res.StatusCode)
	}
	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "delete error") {
		t.Errorf("expected error 'delete error', got %s", string(body))
	}
}
