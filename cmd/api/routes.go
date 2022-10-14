package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialise httprouter router instance
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/todos", app.createTodoHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/todos/:id", app.showTodoHandler)

	return router
}
