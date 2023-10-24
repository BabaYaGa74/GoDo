package handlers

import (
	"database/sql"
	"fmt"
	"go_todoApp/config"
	"html/template"
	"net/http"
)

var db *sql.DB
var templates *template.Template

type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()

	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{}

	result, err := db.Query("SELECT id, task, done FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for result.Next() {
		var task Task
		err := result.Scan(&task.ID, &task.Task, &task.Done)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
		fmt.Print(tasks)
	}
	w.Header().Set("Content-Type", "text/html")
	templates.ExecuteTemplate(w, "tasks.html", tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	taskName := r.FormValue("task")
	if taskName == "" {
		http.Error(w, "Task is required", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO tasks (task) VALUES (?)", taskName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
