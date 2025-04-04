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
	NOT_DONE    Status = "not-done"
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
		id, description, NOT_DONE, createdAt, createdAt,
	}
}

func Decode(content string) error {
	var payload Task
	return json.Unmarshal([]byte(content), &payload)
}

func (t *Task) Encode() ([]byte, error) {
	return json.Marshal(t)
}

func (t Task) String() string {
	return fmt.Sprintf("Task (%d)| %s | created -> %s", t.Id, t.Description, t.CreatedAt.Format(time.RFC3339))
}
