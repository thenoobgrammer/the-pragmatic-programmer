package todo_test

import (
	"os"
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New task should not be completed")
	}
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	l.Delete(1)

	if len(l) > 0 {
		t.Errorf("Expected list to be empty, got lenght of %d.", len(l))
	}
}

func TestSave(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	filename := "Todo"
	l.Save(filename)
	bytes, err := os.ReadFile(filename)
	if err != nil {
		if err == os.ErrNotExist {
			t.Errorf("Expected file %q to exist.", filename)
		}
	}
	if len(bytes) == 0 {
		t.Errorf("Expected file %q to have data inside.", filename)
	}

	clearfile(t, filename)
}

func clearfile(t *testing.T, filename string) {
	err := os.Remove(filename)
	if err != nil {
		if err == os.ErrNotExist {
			t.Errorf("Expected file %q to exist.", filename)
		}
	}
}
