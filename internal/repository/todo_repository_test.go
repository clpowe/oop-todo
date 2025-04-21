package repository

import (
	"testing"

	"www.github.com/clpowe/oop-todo/internal/model"
)

func TestInMemoryTodoRepository_Add_GetAll(t *testing.T) {
	repo := NewInMemoryRepo()

	t1 := model.NewTodo("1", "Buy milk")
	err := repo.Add(t1)
	if err != nil {
		t.Fatalf("Unexpected error adding todo: %v", err)
	}

	t2 := model.NewTodo("2", "Buy eggs")
	err = repo.Add(t2)
	if err != nil {
		t.Fatalf("Unexpected error adding todo: %v", err)
	}

	all, err3 := repo.GetAll()
	if err3 != nil {
		t.Fatalf("Unexpected error getting all todos: %v", err3)
	}

	if len(all) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(all))
	}
}

func TestInMemoryTodoRepository_Add_Duplicate(t *testing.T) {
	repo := NewInMemoryRepo()
	t1 := model.NewTodo("dup", "Duplacate Task")

	// First Add
	if err := repo.Add(t1); err != nil {
		t.Fatalf("Unexpected error on first add: %v", err)
	}

	// Duplicate Add
	err := repo.Add(t1)
	if err == nil {
		t.Error("Expected error when adding duplicate ID, got none")
	}
}

func TestInMemory_Update(t *testing.T) {
	repo := NewInMemoryRepo()
	todo := model.NewTodo("u1", "Old Title")
	repo.Add(todo)

	// Modify and update
	todo.Title = "New Title"
	err := repo.Update(todo)
	if err != nil {
		t.Fatalf("Unexpected error updating todo: %v", err)
	}

	all, _ := repo.GetAll()
	if len(all) != 1 || all[0].Title != "New Title" {
		t.Errorf("Expected updated title \"New Title\", got %q", all[0].Title)
	}
}

func TestInMemoryTodoRepository_Update_NotFoind(t *testing.T) {
	// Attempt update without adding
	repo := NewInMemoryRepo()
	todo := model.NewTodo("missing", "Nope")
	err := repo.Update(todo)
	if err == nil {
		t.Error("Expected error updated non-existent todo, got none")
	}
}

func TestInMemoryTodoRepository_Delete(t *testing.T) {
	repo := NewInMemoryRepo()
	todo := model.NewTodo("d1", "To Delete")
	repo.Add(todo)

	err := repo.Delete("d1")
	if err != nil {
		t.Fatalf("Unexpected error deleting todo: %v", err)
	}

	all, _ := repo.GetAll()
	if len(all) != 0 {
		t.Errorf("Expected 0 todos after delete, got %d", len(all))
	}
}

func TestInMemoryTodoRepository_Delete_NotFound(t *testing.T) {
	repo := NewInMemoryRepo()
	err := repo.Delete("missing")
	if err == nil {
		t.Error("Expected error deleting non-existent todo, got none")
	}
}
