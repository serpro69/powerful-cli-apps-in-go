package todo_test

import (
	"io.github.serpro69/todo"
	"io/ioutil"
	"os"
	"testing"
)

func TestList_Add(t *testing.T) {
	l := todo.List{}
	newTask := "New Task"
	l.Add(newTask)
	if task := l[0].Task; task != newTask {
		t.Errorf("Expected %q, got %q instead.", newTask, task)
	}
}

func TestList_Complete(t *testing.T) {
	l := createTasks()
	if l[0].Done {
		t.Errorf("Task %q should not be completed", l[0].Task)
	}
	_ = l.Complete(1)
	if !l[0].Done {
		t.Errorf("Task %q should be completed", l[0].Task)
	}
}

func TestList_Delete(t *testing.T) {
	l := createTasks()
	l.Delete(1)
	if l[0].Task != "Task Two" {
		t.Errorf("Expected %q, got %q instead", "Task Two", l[0].Task)
	}

	if len(l) != 2 {
		t.Errorf("Expected list lenght %d, got %d instead", 2, len(l))
	}
}

func TestList_Save(t *testing.T) {

}

func TestList_GetSave(t *testing.T) {
	l1 := createTasks()
	l2 := todo.List{}

	tf, _ := ioutil.TempFile("", "")

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task", l1[0].Task, l2[0].Task)
	}
	if len(l1) != len(l2) {
		t.Errorf("Expected list lenght %d, got %d instead", len(l1), len(l2))
	}
}

func createTasks() todo.List {
	l := todo.List{}
	for _, t := range []string{"Task One", "Task Two", "Task Three"} {
		l.Add(t)
	}
	return l
}
