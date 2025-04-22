package service

import (
	"github.com/clpowe/oop-todo/internal/model"
	"github.com/clpowe/oop-todo/internal/repository"
)

type TodoService interface {
	Create(title string) (*model.Todo, error)
	List() ([]*model.Todo, error)
	Complete(id string) error
	Incomplete(id string) error
	Delete(id string) error
}

type DefaultTodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *DefaultTodoService {
	return &DefaultTodoService{repo: repo}
}
