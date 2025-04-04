package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	path := prepareDump(getPath())
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db := readFromJson(file)
	if len(os.Args) < 2 {
		fmt.Println("Expected subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCommand(&db)
	case "list":
		list(&db)
	case "delete":
		deleteCommand(&db)
	case "update":
		updateCommand(&db)
	case "mark-in-progress":
		markCommand(&db, IN_PROGRESS)
	case "mark-done":
		markCommand(&db, DONE)
	case "help":
		helpCommand()
	default:
		fmt.Println("Expected a subcommand")
		os.Exit(1)
	}
}

func list(db *Db) {
	var filter *Filter
	if len(os.Args) == 3 {
		switch os.Args[2] {
		case "done":
			filter = &Filter{DONE}
		case "in-progress":
			filter = &Filter{IN_PROGRESS}
		case "todo":
			filter = &Filter{TODO}
		}
	}
	fmt.Println(db.Render(filter))
}

func addCommand(db *Db) {
	defer db.Save()
	if len(os.Args) < 3 {
		fmt.Println("Expected a task description")
		os.Exit(1)
	}

	desc := os.Args[2]
	t := db.Add(desc)
	fmt.Printf("Task added successfully (ID: %d)", t.Id)
}

func deleteCommand(db *Db) {
	defer db.Save()
	if len(os.Args) < 3 {
		fmt.Println("Expected a task id")
		os.Exit(1)
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Ids should be numbers: %d", id)
		os.Exit(1)
	}
	db.Delete(id)
	fmt.Printf("Task %d deleted", id)
}

func markCommand(db *Db, status Status) {
	defer db.Save()
	if len(os.Args) < 3 {
		fmt.Println("Expected a task id")
		os.Exit(1)
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Ids should be numbers: %d", id)
		os.Exit(1)
	}
	db.Mark(id, status)
}

func updateCommand(db *Db) {
	defer db.Save()
	if len(os.Args) < 4 {
		fmt.Println("Expected task id and description")
		os.Exit(1)
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Ids should be numbers: %d", id)
		os.Exit(1)
	}
	desc := os.Args[3]
	db.Update(id, desc)
}

func helpCommand() {
	fmt.Println(
		`
Usage: task-cli <command> [options]

Commands:
  add <task_description>              Add a new task
  update <task_id> <new_description>  Update an existing task
  delete <task_id>                    Delete a task
  mark-in-progress <task_id>          Mark a task as in progress
  mark-done <task_id>                 Mark a task as done
  list                                List all tasks
  list <status>                       List tasks by status (done, todo, in-progress)
  help                                Show this message and exit

Examples:
  task-cli add "Buy groceries"
  task-cli update 1 "Buy groceries and cook dinner"
  task-cli delete 1
  task-cli mark-in-progress 1
  task-cli mark-done 1
  task-cli list
  task-cli list done
  task-cli list todo
  task-cli list in-progress`)
}
