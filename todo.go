package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)


type Todo struct {
Title string
Completed bool
CreatedAt time.Time
CompletedAt *time.Time //this is a pointer because it can be null
}

type Todos []Todo

func(todos *Todos) add(title string){
	todo := Todo {
		Title : title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos){
		err:= errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error{
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil

}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed;

	if !isCompleted{
		now := time.Now()
		t[index].CompletedAt = &now
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit (index int, title string) error {
	t:= *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) print() {
	Table := table.New(os.Stdout)
	Table.SetRowLines(false)
	Table.SetHeaders([]string{"#", "Title", "Completed", "Created At", "Completed At"}...)
	for i, todo := range *todos {
		completed := " X "
		completedAt := ""
		
		if todo.Completed {
			completed = "✓"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)

			}
		}
		createdAt := todo.CreatedAt.Format(time.RFC1123)
		Table.AddRow(strconv.Itoa(i + 1), todo.Title, completed, createdAt, completedAt)
	}
	Table.Render()
}