package dto

import (
	"github.com/google/uuid"
)

type PublicationRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id"`
	Categories  string    `json:"categories"`
}

type PublicationUpdateRequest struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Categories  string    `json:"categories"`
}
