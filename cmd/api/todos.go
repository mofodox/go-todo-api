package main

import (
	"errors"
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

	err = app.models.Todos.Insert(todo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/api/v1/todos/%d", todo.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"todo": todo}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	todo, err := app.models.Todos.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"todo": todo}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string
		IsCompleted bool
	}

	todos, err := app.models.Todos.GetAll(input.Title, input.IsCompleted)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"todos": todos}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       *string `json:"title"`
		IsCompleted *bool   `json:"is_completed"`
	}

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	todo, err := app.models.Todos.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		todo.Title = *input.Title
	}

	if input.IsCompleted != nil {
		todo.IsCompleted = *input.IsCompleted
	}

	todo.UpdatedAt = time.Now()

	v = validator.New()
	err = v.Struct(todo)
	if err != nil {
		if e, ok := err.(*validator.InvalidValidationError); ok {
			app.errorResponse(w, r, http.StatusUnprocessableEntity, e.Error())
			return
		}
	}

	err = app.models.Todos.Update(todo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"todo": todo}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Todos.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": fmt.Sprintf("todo with id %d successfully deleted", id)}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
