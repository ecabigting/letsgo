package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCommandFlags() *CommandFlags {
	cf := CommandFlags{}
	// Using the flag package native in go, which is used for command-line flag parsing.
	flag.StringVar(&cf.Add, "add", "", "Add a new Task by title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an existing task by specifying id # and new task title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a task by id #")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle task status by id #")
	flag.BoolVar(&cf.List, "list", false, "List all task")

	flag.Parse()
	return &cf
}

func (cf *CommandFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Print("Invalid edit arguments, Use id:title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid task id # to edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Invalid command!")

	}
}
