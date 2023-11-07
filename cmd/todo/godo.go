package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	todo "go_todoApp"
	"io"
	"os"
	"strings"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "Add a new todo")
	done := flag.Int("done", 0, "Marks the todo as complete")
	del := flag.Int("del", 0, "Deletes the todo")
	delAll := flag.Bool("delAll", false, "Deletes all todos")
	list := flag.Bool("list", false, "Lists all todos")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.AddTodo(task)
		err = todos.StoreTodo(todoFile)
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

	case *delAll:
		todos.DeleteAll()
		err := todos.StoreTodo(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println("Deleted all todos successfully!")
		os.Exit(0)

	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "Invalid command!")
		os.Exit(0)
	}

}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("todo cannot be empty")
	}

	return text, nil
}
