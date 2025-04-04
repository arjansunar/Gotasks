package main

import (
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
	Id          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func NewTask(id int, description string) Task {
	createdAt := time.Now().UTC()
	return Task{
		id, description, NOT_DONE, createdAt, createdAt,
	}
}

func (t Task) String() string {
	return fmt.Sprintf("Task (%d)| %s | created -> %s", t.Id, t.Description, t.CreatedAt.Format(time.RFC3339))
}

func main() {
	example := NewTask(1, "New desc")
	fmt.Println(example)
}
