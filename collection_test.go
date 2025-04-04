package main

import (
	"strings"
	"testing"
)

func AssertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func getEmptyDb() Db {
	empty := ""
	return readFromJson(strings.NewReader(empty))
}

func TestAddTask(t *testing.T) {
	db := getEmptyDb()
	db.Add("testing")
	AssertEqual(t, 1, len(db.List(nil)))
	tsk, _ := db.Find(1)
	AssertEqual(t, "testing", tsk.Description)
}

func TestAddRemoveTask(t *testing.T) {
	db := getEmptyDb()
	db.Add("testing")
	db.Delete(1)
	AssertEqual(t, 0, len(db.List(nil)))
}

func TestUpdateTask(t *testing.T) {
	db := getEmptyDb()
	db.Add("testing")
	db.Update(1, "new value")

	tsk, _ := db.Find(1)
	AssertEqual(t, "new value", tsk.Description)
}

func TestMarkTask(t *testing.T) {
	db := getEmptyDb()
	db.Add("testing")

	db.Mark(1, IN_PROGRESS)
	tsk, _ := db.Find(1)
	AssertEqual(t, IN_PROGRESS, tsk.Status)

	db.Mark(1, DONE)
	tsk, _ = db.Find(1)
	AssertEqual(t, DONE, tsk.Status)

	db.Mark(1, TODO)
	tsk, _ = db.Find(1)
	AssertEqual(t, TODO, tsk.Status)
}
