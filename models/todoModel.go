// models/todoModel.go

package models

// Todo представляет задачу в списке дел.
type Todo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Priority string `json:"priority,omitempty"` // Новое поле для приоритета
	Archived bool   `json:"archived,omitempty"` // Новое поле для архивирования
}
