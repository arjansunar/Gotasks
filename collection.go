package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Db struct {
	tasks  []Task
	length int
}

func (db *Db) Add(desc string) Task {
	task := NewTask(db.length+1, desc)
	db.tasks = append(db.tasks, task)
	return task
}

func (db *Db) Delete(id int) []Task {
	newTasks := []Task{}
	for _, v := range db.tasks {
		if v.Id != id {
			newTasks = append(newTasks, v)
		}
	}
	return newTasks
}

func (db *Db) Find(id int) (Task, error) {
	for _, task := range db.tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("no task found with id %d", id)
}

type Filter struct {
	status Status
}

const (
	ALL_FILTER Status = "all"
)

func (db *Db) List(filter *Filter) []Task {
	if filter != nil {
		filtered := []Task{}
		for _, t := range db.tasks {
			if t.Status == filter.status {
				filtered = append(filtered, t)
			}
		}
		return filtered
	}
	return db.tasks
}

func (db *Db) Render(filter *Filter) string {
	res := ""
	tasks := db.List(filter)
	for _, task := range tasks {
		res = fmt.Sprintf("%s\n%s", res, task.Render())
	}
	return res
}

func (db *Db) Save() {
	file, _ := os.Create(getPath())
	defer file.Close()
	data, err := json.Marshal(db.tasks)
	if err != nil {
		fmt.Println("Unable to save", err)
		os.Exit(1)
	}
	_, werr := file.Write(data)
	if werr != nil {
		fmt.Println("Unable to write", werr)
		os.Exit(1)
	}
}

func getPath() string {
	return "db.json"
}

func readFromJson(r io.Reader) Db {
	decoder := json.NewDecoder(r)
	decoder.Token()

	data := []Task{}
	var task Task
	for decoder.More() {
		decoder.Decode(&task)
		data = append(data, task)
	}

	return Db{data, len(data)}
}

func prepareDump(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		defer file.Close()
		file, err = os.Create(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		return filename
	}
	return filename
}
