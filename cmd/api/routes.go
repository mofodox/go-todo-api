package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Initialise httprouter router instance
	router := httprouter.New()

	// Using our custom not found and method not allowed functions
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/todos", app.createTodoHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/todos/:id", app.showTodoHandler)
	router.HandlerFunc(http.MethodPatch, "/api/v1/todos/:id", app.updateTodoHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/todos/:id", app.deleteTodoHandler)

	return router
}
