package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anfelo/gotodo/internal/todos"
	"github.com/gorilla/mux"
)

// GetTodo - retrieve a todo by ID
func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Unable to parse UINT from ID", Error: err.Error()})
		return
	}

	todo, err := h.Service.GetTodo(uint(i))
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Error Retrieving Todo By", Error: err.Error()})
		return
	}

	RespondJson(w, http.StatusOK, todo)
}

// GetAllTodos - retrieve all todos from the comment service
func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Service.GetAllTodos()
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to retrieve all todos", Error: err.Error()})
		return
	}
	RespondJson(w, http.StatusOK, todos)
}

// CreateTodo - creates a new todo
func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo todos.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		fmt.Fprintf(w, "Failed to decode JSON Body")
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to decode JSON Body", Error: err.Error()})
		return
	}
	todo, err := h.Service.CreateTodo(todo)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to post new todo", Error: err.Error()})
		return
	}
	RespondJson(w, http.StatusCreated, todo)
}

// UpdateTodo - updates a todo by ID
func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo todos.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to decode JSON Body", Error: err.Error()})
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	todoID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to parse uint from ID", Error: err.Error()})
		return
	}

	todo, err = h.Service.UpdateTodo(uint(todoID), todo)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to update todo", Error: err.Error()})
		return
	}
	RespondJson(w, http.StatusOK, todo)
}

// DeleteTodo - deletes a todo by ID
func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todoID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to parse uint from ID", Error: err.Error()})
		return
	}

	err = h.Service.DeleteTodo(uint(todoID))
	if err != nil {
		RespondJson(w, http.StatusInternalServerError,
			Response{Message: "Failed to delete todo by todo ID", Error: err.Error()})
		return
	}
	RespondJson(w, http.StatusOK, Response{Message: "Successfully deleted todo"})
}
