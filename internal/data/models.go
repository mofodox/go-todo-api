package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Todos TodoModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Todos: TodoModel{DB: db},
	}
}
