package model

import "time"

type Todo struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTodo(id, title string) *Todo {
	now := time.Now()
	return &Todo{
		ID:        id,
		Title:     title,
		Done:      false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t *Todo) MarkComplete() {
	t.Done = true
	t.UpdatedAt = time.Now()
}

func (t *Todo) MarkIncomplete() {
	t.Done = false
	t.UpdatedAt = time.Now()
}

func (t *Todo) IsCompleted() bool {
	return t.Done
}
