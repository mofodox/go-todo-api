package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mofodox/go-todo/internal/data"
	// "github.com/mofodox/go-todo/internal/validator"
)

var v *validator.Validate

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title" validate:"required"`
		IsCompleted bool   `json:"is_completed" validate:"required"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	todo := &data.Todo{
		Title:       input.Title,
		IsCompleted: input.IsCompleted,
	}

	v = validator.New()
	err = v.Struct(todo)
	if err != nil {
		if e, ok := err.(*validator.InvalidValidationError); ok {
			app.errorResponse(w, r, http.StatusUnprocessableEntity, e.Error())
			return
		}
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}
}
