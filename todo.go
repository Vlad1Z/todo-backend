package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

// Todo представляет задачу в списке дел.
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo

// Функции обработчики для HTTP запросов
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(rand.Intn(1000000)) // Не используйте в продакшене
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			var todo Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = params["id"]
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func patchTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			var updatedFields struct {
				Title *string `json:"title,omitempty"`
				Done  *bool   `json:"done,omitempty"`
			}
			_ = json.NewDecoder(r.Body).Decode(&updatedFields)
			if updatedFields.Title != nil {
				todos[index].Title = *updatedFields.Title
			}
			if updatedFields.Done != nil {
				todos[index].Done = *updatedFields.Done
			}
			json.NewEncoder(w).Encode(todos[index])
			return
		}
	}
}
