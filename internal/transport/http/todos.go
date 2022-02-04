package http

import (
	"encoding/json"
	"net/http"

	"github.com/anfelo/gotodo/internal/todos"
	"github.com/anfelo/gotodo/internal/transport/errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetTodo - retrieve a todo by ID
func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid todo id")
		RespondJson(w, restErr.Status, restErr)
		return
	}

	todo, err := h.Service.GetTodo(id)
	if err != nil {
		restErr := errors.NewNotFoundError("todo not found")
		RespondJson(w, restErr.Status, restErr)
		return
	}

	RespondJson(w, http.StatusOK, todo)
}

// GetAllTodos - retrieve all todos from the todos service
func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Service.GetAllTodos()
	if err != nil {
		restErr := errors.NewInternatServerError("internal server error")
		RespondJson(w, restErr.Status, restErr)
		return
	}
	RespondJson(w, http.StatusOK, todos)
}

// CreateTodo - creates a new todo
func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo todos.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		RespondJson(w, restErr.Status, restErr)
		return
	}
	todo, err := h.Service.CreateTodo(todo)
	if err != nil {
		restErr := errors.NewInternatServerError("internal server error")
		RespondJson(w, restErr.Status, restErr)
		return
	}
	RespondJson(w, http.StatusCreated, todo)
}

// UpdateTodo - updates a todo by ID
func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo todos.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		RespondJson(w, restErr.Status, restErr)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid todo id")
		RespondJson(w, restErr.Status, restErr)
		return
	}

	todo, err = h.Service.UpdateTodo(id, todo)
	if err != nil {
		restErr := errors.NewInternatServerError("internal server error")
		RespondJson(w, restErr.Status, restErr)
		return
	}
	RespondJson(w, http.StatusOK, todo)
}

// DeleteTodo - deletes a todo by ID
func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid todo id")
		RespondJson(w, restErr.Status, restErr)
		return
	}

	err = h.Service.DeleteTodo(id)
	if err != nil {
		restErr := errors.NewInternatServerError("internal server error")
		RespondJson(w, restErr.Status, restErr)
		return
	}
	RespondJson(w, http.StatusOK, map[string]string{"success": "true"})
}
