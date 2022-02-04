package http

import (
	"html/template"
	"net/http"

	"github.com/anfelo/gotodo/internal/todos"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *todos.Service
}

// Response - an object to store responses from the api
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *todos.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// LoggingMiddleware - a handy middleware function that logs out incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
			}).Info("handled request")
		log.Info("Endpoint hit!")
		next.ServeHTTP(w, r)
	})
}

// BasicAuth - a handy middleware function that logs out incoming requests
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(w, r)
		} else {
			RespondJson(w, http.StatusUnauthorized,
				Response{Message: "User is not authorized", Error: "not authorized"})
		}
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	log.Info("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	// WebApp routes
	h.Router.HandleFunc("/", h.Home).Methods("GET")

	// Todos REST API
	h.Router.HandleFunc("/api/todos/{id}", h.GetTodo).Methods("GET")
	h.Router.HandleFunc("/api/todos", h.GetAllTodos).Methods("GET")
	h.Router.HandleFunc("/api/todos", h.CreateTodo).Methods("POST")
	h.Router.HandleFunc("/api/todos/{id}", h.UpdateTodo).Methods("PUT")
	h.Router.HandleFunc("/api/todos/{id}", h.DeleteTodo).Methods("DELETE")

	// TodoLists REST API
	h.Router.HandleFunc("/api/lists/{id}", h.GetTodoList).Methods("GET")
	h.Router.HandleFunc("/api/lists", h.GetAllTodoLists).Methods("GET")
	h.Router.HandleFunc("/api/lists", h.CreateTodoList).Methods("POST")
	// h.Router.HandleFunc("/api/todos/{id}", h.UpdateTodo).Methods("PUT")
	// h.Router.HandleFunc("/api/todos/{id}", h.DeleteTodo).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		RespondJson(w, http.StatusOK, Response{Message: "I am Alive"})
	})
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Lists []todos.TodoList
	}
	tmpl := template.Must(template.ParseFiles(
		"internal/templates/layout.html",
		"internal/templates/home.html",
	))
	todoLists, err := h.Service.GetAllTodoLists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data{Lists: todoLists})
}
