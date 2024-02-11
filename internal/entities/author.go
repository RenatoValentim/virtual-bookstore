package entities

import "github.com/google/uuid"

type Author struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at"`
}
