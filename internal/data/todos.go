package data

import "time"

type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"-"`
	Version     int32     `json:"version"`
}
