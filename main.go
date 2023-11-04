package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Task!!")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Task")
		fmt.Println("3. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTaskUI()
		case 2:
			listTasksUI()
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

	task := Task{
		ID:          len(TaskList) + 1,
		Title:       title,
		Description: description,
	}
	AddTask(task)
	fmt.Println("Task added successfully!")
}

func listTasksUI() {
	tasks := ListTasks()
	fmt.Println("Task lists: ")
	for _, task := range tasks {
		fmt.Printf("ID: %d\n Title: %s\n Description: %s", task.ID, task.Title, task.Description)
	}
}
