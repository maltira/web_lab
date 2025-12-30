package dto

import (
	"github.com/google/uuid"
)

type PublicationRequest struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	UserID          uuid.UUID `json:"user_id"`
	BackgroundColor string    `json:"background_color"`
	IsDraft         bool      `json:"is_draft"`

	Categories []PublicationCategoryRequest
}

type PublicationUpdateRequest struct {
	ID              uuid.UUID `json:"id"`
	Title           *string   `json:"title"`
	Description     *string   `json:"description"`
	BackgroundColor *string   `json:"background_color"`
	IsDraft         *bool     `json:"is_draft"`
	Categories      *[]PublicationCategoryUpdateRequest
}
