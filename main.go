package main

type Status string

const (
	IN_PROGRESS Status = "in-progress"
	DONE        Status = "done"
	NOT_DONE    Status = "not-done"
)

type Task struct {
	status    Status
	name      string
	completed bool
}

func main() {
	print("Hello, World!")
}
