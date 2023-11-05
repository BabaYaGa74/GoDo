package main

import (
	"flag"
	"fmt"
	todo "go_todoApp"
	"os"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "Add a new todo")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.AddTodo("A new task!")
		err := todos.StoreTodo(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stdout, "Invalid command!")
		os.Exit(0)
	}

}
