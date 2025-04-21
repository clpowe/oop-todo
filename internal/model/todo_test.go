package model

import (
	"testing"
	"time"
)

func TestNewTodo(t *testing.T) {
	id := "123"
	title := "Buy milk"
	todo := NewTodo(id, title)

	if todo.ID != id {
		t.Errorf("Expected ID %s, got %s", id, todo.ID)
	}

	if todo.Title != title {
		t.Errorf("Expected Title %s, got %s", title, todo.Title)
	}

	if todo.Done {
		t.Errorf("Expected Done to be false, got true")
	}

	if todo.CreatedAt.IsZero() {
		t.Errorf("Expected UreatedAt to be set, got zero")
	}

	if todo.UpdatedAt.IsZero() {
		t.Errorf("Expected UpdatedAt to be set, got zero")
	}
}

func TestMarkComplete(t *testing.T) {
	todo := NewTodo("123", "Buy milk")
	prevUpdated := todo.UpdatedAt
	time.Sleep(time.Millisecond)
	todo.MarkComplete()

	if !todo.Done {
		t.Errorf("Expected Done to be true after MarkComplete, got false")
	}

	if !todo.UpdatedAt.After(prevUpdated) {
		t.Errorf("Expected UpdatedAt to be updated after MarkComplete")
	}
}

func TestMarkIncomplete(t *testing.T) {
	todo := NewTodo("123", "Buy milk")
	todo.MarkComplete()
	prevUpdated := todo.UpdatedAt
	time.Sleep(time.Millisecond)
	todo.MarkIncomplete()

	if todo.Done {
		t.Errorf("Expected Done false after INcomplete")
	}

	if !todo.UpdatedAt.After(prevUpdated) {
		t.Error("Expected UpdateAt to be updated after TestMarkIncomplete")
	}
}

func TestIsCompleted(t *testing.T) {
	todo := NewTodo("123", "Buy milk")
	if todo.IsCompleted() {
		t.Error("Expected IsCompleted to return false for a new Todo")
	}
	todo.MarkComplete()
	if !todo.IsCompleted() {
		t.Error("Expected IsCompleted true after MarkComplete")
	}
}
