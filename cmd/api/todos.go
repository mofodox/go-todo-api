package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mofodox/go-todo/internal/data"
)

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new todo")
}

func (app *application) showTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	todo := data.Todo{
		ID:          id,
		Title:       "Test 1",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Version:     1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"todo": todo}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}