package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// Define the Todo struct
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// Define the Todos as splice of Todo
type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	// Append the latest todo we created
	// using the title string to the splice Todos
	*todos = append(*todos, todo)
}

// Vaidate functions that gets the current
// *Todos in the current context and check its
// length against the index pass to the function below
func (todos *Todos) validateIndex(index int) error {
	// check if index is less than 0
	// or index is greater than or equal
	// to the length of todos splice using
	// the len() build in function of go
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	// This line get all the content of the slice and create
	// a new slice, here is how its down:
	// `t[:index]` gets all the items from the start
	// of the slice up to the specified index
	// `t[index+1:]` gets all the items from the index
	// upto the end of the slice, and then using
	// `...` to unpack the 2nd part of the slice
	// this is because of how the `append()` function of GO:
	// --------------
	// First Slice as Base: In the expression append(t[:index], t[index+1:]...),
	// the first argument t[:index] is the base slice to which you want to append elements.
	// This slice contains all elements before the specified index.
	// Unpacking the Second Slice: The second argument t[index+1:] is a slice
	// that contains all elements after the specified index.
	// By using the ... operator, you are unpacking this slice so that each element
	// in t[index+1:] is appended individually to the first slice.
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	// validate if the index is valid or existing
	if err := t.validateIndex(index); err != nil {
		return err
	}
	// get the current completed status of
	// the task here
	isCompleted := t[index].Completed

	// if its not yet marked as completed
	// get the current datetime now and assign it into completionTime
	// then reference the compl
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	// Set the table headers
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	// loop the todos
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			// adding this additional check for completedAt
			// so we can make sure we are not going to throw
			// and error when its not and try to convert it
			// and vice versa
			if t.CompletedAt != nil {
				// convert t.CompletedAt as string date
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		// inserting the roaws into the table with this order
		// index,title,completed status,created at,completed at
		// here are the values we set and what we did
		//
		// column 0 = we converted the index into string using strconv.Itoa
		// column 1 = we set the value from the current t.Title,
		// column 2 = we set the completed value we evaluated above
		// column 3 = we set the created at value and convert it using Format(time.RFC1123)
		// column 4 = we set the last column with the evaluated value of completedAt
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	// show the table
	table.Render()
}
