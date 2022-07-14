package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// item struct represents a To-Do item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of To-Do items
type List []item

// Add appends a new task to the List of tasks
//
// Note: we use the pointer to the type *List as the receiver type.
// Otherwise, the method would change a copy of the list instead.
// We define the receiver as a pointer to the type (*SomeType)
// when the method needs to modify the content of the receiver.
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	// Note: we dereference the pointer to the List type with '*l' in the append call
	// to access the underlying slice
	*l = append(*l, t)
}

// Complete method marks a To-Do item as completed
// by setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	// Note: Strictly speaking, we do not modify the list, only an item in the list,
	// so we don't need a pointer receiver for this method.
	// But it's a good practice to keep the entire method set of a single type, i.e. List,
	// with the same receiver type. Hence, we decided to declare the Complete method
	// with a pointer receiver as well.

	return nil
}

// Delete deletes a To-Do item from the List
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		fmt.Errorf("item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Save method encodes the List as JSON
// and saved is using the provided filename
func (l *List) Save(filename string) error {
	if jsn, err := json.Marshal(l); err == nil {
		return ioutil.WriteFile(filename, jsn, 0644)
	} else {
		return err
	}
}

// Get method decodes the JSON data from the provided filename
// and deserializes it into a List.
//
// This will return an error for empty and non-existent files.
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
