package service

import (
	"errors"

	"github.com/clpowe/oop-todo/internal/model"
	"github.com/clpowe/oop-todo/internal/repository"

	"github.com/google/uuid"
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

func (s *DefaultTodoService) Create(title string) (*model.Todo, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	id := uuid.NewString()
	todo := model.NewTodo(id, title)
	if err := s.repo.Add(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *DefaultTodoService) List() (map[string]*model.Todo, error) {
	return s.repo.GetAll()
}

func (s *DefaultTodoService) Complete(id string) error {
	todos, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	todo, ok := todos[id]
	if !ok {
		return errors.New("todo not found")
	}

	todo.MarkComplete()
	return s.repo.Update(todo)
}

func (s *DefaultTodoService) Incomplete(id string) error {
	todos, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	todo, ok := todos[id]
	if !ok {
		return errors.New("todo not found")
	}

	todo.MarkIncomplete()
	return s.repo.Update(todo)
}

func (s *DefaultTodoService) Delete(id string) error {
	return s.repo.Delete(id)
}
