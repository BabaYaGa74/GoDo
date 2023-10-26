package main

import (
	"fmt"
	"go_todoApp/handlers"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/static/{rest:.*}", serveStatic)
	router.HandleFunc("/", serverIndex)

	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/addTasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/delete", handlers.DeleteTask).Methods("POST")
	router.HandleFunc("/update", handlers.UpdateTask).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func serverIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	path := "./" + r.URL.Path
	fmt.Println("Trying to serve:", path)
	if strings.HasSuffix(path, ".css") {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	}
	http.ServeFile(w, r, path)
}
