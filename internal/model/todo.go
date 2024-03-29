package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	CreatedAt  time.Time `json:"created-at"`
	FinishedAt time.Time `json:"finished-at,omitempty"`
	Name       string    `json:"name"`
	Id         uuid.UUID `json:"id"`
	Finished   bool      `json:"finished"`
	Active     bool      `json:"active"`
}

func NewTodo(name string) *Todo {
	return &Todo{
		Id:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		Finished:  false,
		Active:    true,
	}
}

func (t *Todo) Finish() {
	t.FinishedAt = time.Now()
	t.Finished = true
}
