package http

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router *mux.Router
}

// Response - an object to store responses from the api
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
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

	h.Router.HandleFunc("/", h.Home).Methods("GET")

	h.Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		RespondJson(w, http.StatusOK, Response{Message: "I am Alive"})
	})
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	type data struct{}
	tmpl := template.Must(template.ParseFiles(
		"internal/templates/layout.html",
		"internal/templates/home.html",
	))
	tmpl.Execute(w, data{})
}
