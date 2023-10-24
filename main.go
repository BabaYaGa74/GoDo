package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:Biplove@123@localhost/go_tododb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/", homeFunc)

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}

// func createTask(w http.)

func getTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a home page"))
}
