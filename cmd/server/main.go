package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// App - contain application information
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (a *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    a.Name,
			"AppVersion": a.Version,
		}).Info("Setting up application")

	http.HandleFunc("/", Home)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are on the home page")
}

func main() {
	log.Info("Go Todo App")
	app := App{
		Name:    "Todo App",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error("Error starting up our Web App")
		log.Fatal(err)
	}
}
