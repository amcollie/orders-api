package application

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (app *App) Start(ctx context.Context) error {
	loggedRouter := handlers.LoggingHandler(os.Stdout, app.router)
	server := &http.Server{
		Addr:    ":3000",
		Handler: loggedRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to listen to server: %w", err)
	}

	return nil
}
