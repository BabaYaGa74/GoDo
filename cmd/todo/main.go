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
	done := flag.Int("done", 0, "Marks the todo as complete")
	del := flag.Int("del", 0, "Deletes the todo")
	list := flag.Bool("list", false, "Lists all todos")
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

	case *done > 0:
		err := todos.MarkComplete(*done)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.StoreTodo(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *del > 0:
		err := todos.DeleteTodo(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.StoreTodo(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *list:

	default:
		fmt.Fprintln(os.Stdout, "Invalid command!")
		os.Exit(0)
	}

}
