package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Enabled     bool      `json:"enabled"`
}

type TodoRequest struct {
	Name        string `json:"name" validate:"required,max=100,min=3"`
	Description string `json:"description" validate:"max=1000"`
}
