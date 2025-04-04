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

	db.Delete(1)
	err := db.Update(1, "should not happen")
	AssertEqual(t, "no task found with id 1", err.Error())
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
	db.Delete(1)
	err := db.Mark(1, TODO)
	AssertEqual(t, "no task found with id 1", err.Error())
}

func TestDeleteTask(t *testing.T) {
	db := getEmptyDb()
	db.Add("testing")
	db.Delete(1)
	_, err := db.Find(1)
	AssertEqual(t, "no task found with id 1", err.Error())
}

func TestListTask(t *testing.T) {
	db := getEmptyDb()
	db.Add("testing")
	db.Add("testing")
	db.Add("testing")

	all := db.List(nil)
	AssertEqual(t, 3, len(all))
	todo := db.List(&Filter{TODO})
	AssertEqual(t, 3, len(todo))

	done := db.List(&Filter{DONE})
	AssertEqual(t, 0, len(done))
	inprog := db.List(&Filter{IN_PROGRESS})
	AssertEqual(t, 0, len(inprog))

	db.Mark(2, IN_PROGRESS)
	db.Mark(3, DONE)
	done = db.List(&Filter{DONE})
	AssertEqual(t, 1, len(done))
	inprog = db.List(&Filter{IN_PROGRESS})
	AssertEqual(t, 1, len(inprog))
}
