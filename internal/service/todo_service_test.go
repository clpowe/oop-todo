package service

import (
	"testing"

	"github.com/clpowe/oop-todo/internal/repository"
)

func TestDefaultTodoService_Create_List(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	svc := NewTodoService(repo)

	// Test empty title error
	if _, err := svc.Create(""); err == nil {
		t.Error("expected error when creating todo with empty title, got none")
	}

	// Test successful create
	todo, err := svc.Create("My Task")
	if err != nil {
		t.Fatalf("unexpected error creating todo: %v", err)
	}
	if todo.Title != "My Task" {
		t.Errorf("expected Title 'My Task', got '%s'", todo.Title)
	}
	if todo.Done {
		t.Error("expected new todo Done=false")
	}

	// Test List returns the created todo
	all, err := svc.List()
	if err != nil {
		t.Fatalf("unexpected error listing todos: %v", err)
	}
	if len(all) != 1 {
		t.Errorf("expected 1 todo, got %d", len(all))
	}
	found, ok := all[todo.ID]
	if !ok {
		t.Errorf("expected todo with ID %s in list", todo.ID)
	}
	if found.ID != todo.ID {
		t.Errorf("expected ID '%s', got '%s'", todo.ID, found.ID)
	}
}

func TestDefaultTodoService_Complete_Incomplete(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	svc := NewTodoService(repo)

	todo, _ := svc.Create("Task for Status")

	// Complete existing todo
	if err := svc.Complete(todo.ID); err != nil {
		t.Fatalf("unexpected error completing todo: %v", err)
	}
	listed, _ := svc.List()
	if !listed[todo.ID].IsCompleted() {
		t.Error("expected todo to be completed after Complete()")
	}

	// Incomplete existing todo
	if err := svc.Incomplete(todo.ID); err != nil {
		t.Fatalf("unexpected error incompleting todo: %v", err)
	}
	listed, _ = svc.List()
	if listed[todo.ID].IsCompleted() {
		t.Error("expected todo to be incomplete after Incomplete()")
	}

	// Complete non-existent
	if err := svc.Complete("nope"); err == nil {
		t.Error("expected error completing non-existent todo, got none")
	}
	// Incomplete non-existent
	if err := svc.Incomplete("nope"); err == nil {
		t.Error("expected error incompleting non-existent todo, got none")
	}
}

func TestDefaultTodoService_Delete(t *testing.T) {
	repo := repository.NewInMemoryRepo()
	svc := NewTodoService(repo)

	todo, _ := svc.Create("Task to Delete")

	// Delete existing
	if err := svc.Delete(todo.ID); err != nil {
		t.Fatalf("unexpected error deleting todo: %v", err)
	}
	all, _ := svc.List()
	if len(all) != 0 {
		t.Errorf("expected 0 todos after delete, got %d", len(all))
	}

	// Delete non-existent
	if err := svc.Delete("missing"); err == nil {
		t.Error("expected error deleting non-existent todo, got none")
	}
}
