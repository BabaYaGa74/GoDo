package pkg

var TaskList []Task

func AddTask(task Task) {
	TaskList = append(TaskList, task)
}

func ListTasks() []Task {
	return TaskList
}
