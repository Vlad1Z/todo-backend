// main.go
package main

import (
	"github.com/gorilla/mux"
	"log"
	"myapi/handlers" // Обратите внимание на измененный путь импорта
	"net/http"
)

func main() {
	// Настройка роутера
	router := mux.NewRouter()

	// Регистрация обработчиков
	router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", handlers.GetTodo).Methods("GET")
	router.HandleFunc("/todo", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", handlers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todo/{id}", handlers.PatchTodo).Methods("PATCH")

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8000", router))
}
