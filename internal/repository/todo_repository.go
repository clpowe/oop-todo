package repository

import (
	"errors"
	"sync"

	"github.com/clpowe/oop-todo/internal/model"
)

type TodoRepository interface {
	Add(todo *model.Todo) error
	GetAll() (map[string]*model.Todo, error)
	GetByID(id string) (*model.Todo, error)
	Update(todo *model.Todo) error
	Delete(id string) error
}

type InMemoryTodoRepository struct {
	mu    sync.RWMutex
	todos map[string]*model.Todo
}

func NewInMemoryRepo() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[string]*model.Todo),
	}
}

func (r *InMemoryTodoRepository) Add(todo *model.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.todos[todo.ID]; exists {
		return errors.New("todo with given ID already exists")
	}
	r.todos[todo.ID] = todo
	return nil
}

func (r *InMemoryTodoRepository) GetAll() (map[string]*model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := r.todos
	if len(list) == 0 {
		return nil, errors.New("no todos found")
	}
	return list, nil
}

func (r *InMemoryTodoRepository) GetByID(id string) (*model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if _, exists := r.todos[id]; !exists {
		return nil, errors.New("todo not found")
	}
	return r.todos[id], nil
}

func (r *InMemoryTodoRepository) Update(todo *model.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.todos[todo.ID]; !exists {
		return errors.New("todo not found")
	}
	r.todos[todo.ID] = todo
	return nil
}

func (r *InMemoryTodoRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.todos[id]; !exists {
		return errors.New("todo not found")
	}
	delete(r.todos, id)
	return nil
}
