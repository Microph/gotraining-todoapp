package todo_test

import (
	"os"
	"testing"
	"todoapp/todo"
)

func TestNew(t *testing.T) {
	err := todo.New("2019-02-26.td")
	if err != nil {
		t.Error("cannot create new todo list file:", err)
	}
	os.Remove("2019-02-26.td")
}

func TestPyramid(t *testing.T) {
	err := todo.Pyramid("7")
	if err != nil {
		t.Error("fail to print pyramid:", err)
	}
}
