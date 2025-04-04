package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	db := readFromJson(getPath())
	if len(os.Args) < 2 {
		fmt.Println("Expected subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCommand(&db)
		db.Save()
	case "list":
		list(&db)
	case "delete":
		deleteCommand(&db)
		db.Save()
	case "find":
		task, err := db.Find(1)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(task)
	default:
		fmt.Println("Expected a subcommand")
		os.Exit(1)
	}
}

func list(db *Db) {
	fmt.Println(db.Render())
}

func addCommand(db *Db) {
	if len(os.Args) < 3 {
		fmt.Println("Expected a task description")
		os.Exit(1)
	}

	desc := os.Args[2]
	t := db.Add(desc)
	fmt.Printf("Task added successfully (ID: %d)", t.Id)
}

func deleteCommand(db *Db) {
	if len(os.Args) < 3 {
		fmt.Println("Expected a task id")
		os.Exit(1)
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Ids should be numbers: %d", id)
		os.Exit(1)
	}
	db.Remove(id)
	fmt.Printf("Task %d deleted", id)
}
