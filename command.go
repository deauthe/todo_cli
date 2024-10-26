package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdflags() *cmdFlags {
	cf := cmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index, format: index:new title")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()
	return &cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	switch {
		case cf.Add != "":
			todos.add(cf.Add)
		case cf.Del != -1:
			todos.delete(cf.Del)
		case cf.Edit != "":
			parts := strings.SplitN(cf.Edit, ":", 2)
			if len(parts) != 2 {
				fmt.Print("invalid edit format")
				os.Exit(1)
			}
			index, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Print("invalid index")
				os.Exit(1)
			}
			newTitle := parts[1]
			todos.edit(index, newTitle)
		case cf.Toggle != -1:
			todos.toggle(cf.Toggle)
		case cf.List :
			todos.print()
		default:
			fmt.Print("invalid command")
	}
}