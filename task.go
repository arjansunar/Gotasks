package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Status string

const (
	IN_PROGRESS Status = "in-progress"
	DONE        Status = "done"
	TODO        Status = "todo"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updatedAt"`
}

func NewTask(id int, description string) Task {
	createdAt := time.Now().UTC()
	return Task{
		id, description, TODO, createdAt, createdAt,
	}
}

func Decode(content string) error {
	var payload Task
	return json.Unmarshal([]byte(content), &payload)
}

func (t *Task) Encode() ([]byte, error) {
	return json.Marshal(t)
}

func (t Task) Render() string {
	completedRender := " "
	if t.Status == DONE {
		completedRender = "X"
	}
	if t.Status == IN_PROGRESS {
		completedRender = "~"
	}
	return fmt.Sprintf("- [%s] %s", completedRender, t.Description)
}

func (t Task) String() string {
	return fmt.Sprintf("Task (%d)| %s | created -> %s", t.Id, t.Description, t.CreatedAt.Format(time.RFC3339))
}
