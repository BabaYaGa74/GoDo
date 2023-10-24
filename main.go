package main

import (
	"go_todoApp/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", serverIndex)

	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func serverIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
