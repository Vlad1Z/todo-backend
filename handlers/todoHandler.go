// handlers/todoHandler.go
package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"myapi/models"
	"myapi/store" // Обратите внимание на измененный путь импорта
	"net/http"
	"strconv"
	"strings"
)

// GetTodos обрабатывает запросы на получение всех задач с возможностью фильтрации и поиска.
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	filterDone := r.URL.Query().Get("done")
	searchTitle := r.URL.Query().Get("title")

	var filteredTodos []models.Todo

	for _, todo := range store.Todos {
		// Проверяем выполнена ли задача, если параметр done присутствует
		if filterDone != "" {
			done, _ := strconv.ParseBool(filterDone)
			if todo.Done != done {
				continue
			}
		}

		// Проверяем входит ли подстрока searchTitle в название задачи, если параметр title присутствует
		if searchTitle != "" && !strings.Contains(strings.ToLower(todo.Title), strings.ToLower(searchTitle)) {
			continue
		}

		filteredTodos = append(filteredTodos, todo)
	}

	json.NewEncoder(w).Encode(filteredTodos)
}

// getTodo обрабатывает запросы на получение одной задачи по ID.
func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range store.Todos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createTodo обрабатывает запросы на создание новой задачи.
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo models.Todo // Используйте models1.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(rand.Intn(1000000)) // Не используйте в продакшене
	store.Todos = append(store.Todos, todo)    // Используйте store.Todos
	json.NewEncoder(w).Encode(todo)
}

// updateTodo обрабатывает запросы на обновление задачи.
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range store.Todos {
		if item.ID == params["id"] {
			store.Todos = append(store.Todos[:index], store.Todos[index+1:]...)
			var todo models.Todo // Используйте models1.Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = params["id"]
			store.Todos = append(store.Todos, todo) // Используйте store.Todos
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
}

// deleteTodo обрабатывает запросы на удаление задачи.
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range store.Todos {
		if item.ID == params["id"] {
			store.Todos = append(store.Todos[:index], store.Todos[index+1:]...) // Используйте store.Todos
			break
		}
	}
	json.NewEncoder(w).Encode(store.Todos) // Используйте store.Todos
}

// patchTodo обрабатывает PATCH запросы для частичного обновления задачи.
func PatchTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range store.Todos {
		if item.ID == params["id"] {
			var updatedFields struct {
				Title *string `json:"title,omitempty"`
				Done  *bool   `json:"done,omitempty"`
			}
			_ = json.NewDecoder(r.Body).Decode(&updatedFields)
			if updatedFields.Title != nil {
				store.Todos[index].Title = *updatedFields.Title // Используйте store.Todos[index]
			}
			if updatedFields.Done != nil {
				store.Todos[index].Done = *updatedFields.Done // Используйте store.Todos[index]
			}
			json.NewEncoder(w).Encode(store.Todos[index]) // Используйте store.Todos[index]
			return
		}
	}
}
