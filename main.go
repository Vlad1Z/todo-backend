package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Маршруты
	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todo", createTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	r.HandleFunc("/todo/{id}", patchTodo).Methods("PATCH")

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8000", r))
}
