package repository

import (
	"errors"
	"sync"

	"www.github.com/clpowe/oop-todo/internal/model"
)

type TodoRepository interface {
	Add(todo *model.Todo) error
	GetAll() ([]*model.Todo, error)
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

func (r *InMemoryTodoRepository) GetAll() ([]*model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*model.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		list = append(list, todo)
	}
	return list, nil
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
