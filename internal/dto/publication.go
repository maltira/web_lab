package dto

import "github.com/google/uuid"

type PublicationRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id"`
}
