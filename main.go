package main

import (
	"fmt"
	"go_todoApp/pkg"
	"os"
)

func main() {
	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. List Task")
		fmt.Println("3. Exit")
		fmt.Println()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTaskUI()
		case 2:
			listTasksUI()
			fmt.Println()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again!")
		}
	}
}

func addTaskUI() {
	var title, description string
	fmt.Println("Enter the title : ")
	fmt.Scanln(&title)
	fmt.Println("Enter the description: ")
	fmt.Scanln(&description)

	task := pkg.Task{
		ID:          len(pkg.TaskList) + 1,
		Title:       title,
		Description: description,
	}
	pkg.AddTask(task)
	fmt.Println("Task added successfully!")
	fmt.Println()

}

func listTasksUI() {
	tasks := pkg.ListTasks()
	fmt.Println("Task lists: ")
	for _, task := range tasks {
		fmt.Printf(" ID: %d\n Title: %s\n Description: %s", task.ID, task.Title, task.Description)
		fmt.Println()
	}
}
