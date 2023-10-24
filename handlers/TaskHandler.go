package handlers

import (
	"database/sql"
	"fmt"
	"go_todoApp/config"
	"net/http"
)

var db *sql.DB

type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	// tasks:= []Task{}
	fmt.Println("Hello world")

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
