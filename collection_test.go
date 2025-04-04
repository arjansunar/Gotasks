package main

import (
	"os"
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

func TestCreation(t *testing.T) {
	db := readFromJson(strings.NewReader(""))
	AssertEqual(t, 0, len(db.List(nil)))

	db = readFromJson(strings.NewReader("[]"))
	AssertEqual(t, 0, len(db.List(nil)))

	db = readFromJson(strings.NewReader(`[{"id":1,"description":"testing","status":"in-progress","createdAt":"2025-04-04T06:53:24.375688299Z","updatedAt":"2025-04-04T06:53:24.375688299Z"}]`))
	AssertEqual(t, 1, len(db.List(nil)))
	AssertEqual(t, 0, len(db.List(&Filter{TODO})))
	AssertEqual(t, 1, len(db.List(&Filter{IN_PROGRESS})))
	AssertEqual(t, 0, len(db.List(&Filter{DONE})))
}

func TestSave(t *testing.T) {
	db := getEmptyDb()
	db.Add("testing")
	db.Add("testing")
	db.Add("testing")
	db.Save()

	path, err := prepareDump(getPath())
	AssertEqual(t, true, err == nil)
	file, err := os.Open(path)
	AssertEqual(t, true, err == nil)
	db = readFromJson(file)
	AssertEqual(t, 3, len(db.List(nil)))
	AssertEqual(t, 3, len(db.List(&Filter{TODO})))
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

func TestRender(t *testing.T) {
	db := getEmptyDb()
	AssertEqual(t, "", db.Render(nil))
	db.Add("testing")
	AssertEqual(t, "- [ ] testing", strings.TrimSpace(db.Render(nil)))
}

func TestPrepareDump(t *testing.T) {
	t.Run("file exists", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "testfile-")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(tempFile.Name()) // Clean up after the test

		result, err := prepareDump(tempFile.Name())
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if result != tempFile.Name() {
			t.Errorf("Expected %s, but got %s", tempFile.Name(), result)
		}
	})

	t.Run("file does not exist and is created", func(t *testing.T) {
		tempFileName := "testfile-does-not-exist.txt"
		defer os.Remove(tempFileName) // Clean up after the test

		result, err := prepareDump(tempFileName)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if result != tempFileName {
			t.Errorf("Expected %s, but got %s", tempFileName, result)
		}

		if _, err := os.Stat(tempFileName); os.IsNotExist(err) {
			t.Errorf("Expected file %s to be created, but it does not exist", tempFileName)
		}
	})

	t.Run("error creating file", func(t *testing.T) {
		tempFileName := "/root/testfile-no-permission.txt" // Modify based on your system

		_, err := prepareDump(tempFileName)

		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}
