package data

import (
	"database/sql"
	"errors"
	"time"
)

type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" validate:"required"`
	IsCompleted bool      `json:"is_completed" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"-"`
	Version     int32     `json:"version"`
}

type TodoModel struct {
	DB *sql.DB
}

func (t TodoModel) Insert(todo *Todo) error {
	qry := `INSERT INTO todos (title, is_completed) VALUES ($1, $2) RETURNING id, created_at, version`

	args := []interface{}{todo.Title, todo.IsCompleted}

	return t.DB.QueryRow(qry, args...).Scan(&todo.ID, &todo.CreatedAt, &todo.Version)
}

func (t TodoModel) Get(id int64) (*Todo, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	qry := `SELECT id, title, is_completed, created_at, updated_at, version FROM todos WHERE id = $1`

	var todo Todo

	err := t.DB.QueryRow(qry, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.IsCompleted,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &todo, nil
}

func (t TodoModel) GetAll(title string, isCompleted bool) ([]*Todo, error) {
	qry := `SELECT id, title, is_completed, created_at, updated_at, version FROM todos ORDER BY id`

	rows, err := t.DB.Query(qry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*Todo{}

	for rows.Next() {
		var todo Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.IsCompleted,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.Version,
		)
		if err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (t TodoModel) Update(todo *Todo) error {
	qry := `UPDATE todos SET title = $1, is_completed = $2, updated_at = $3, version = version + 1 WHERE id = $4 AND version = $5 RETURNING version`

	args := []interface{}{
		todo.Title,
		todo.IsCompleted,
		todo.UpdatedAt,
		todo.ID,
		todo.Version,
	}

	err := t.DB.QueryRow(qry, args...).Scan(&todo.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (t *TodoModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	qry := `DELETE FROM todos WHERE id = $1`

	result, err := t.DB.Exec(qry, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
