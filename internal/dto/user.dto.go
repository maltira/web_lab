package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	ID               uuid.UUID `json:"id" binding:"required"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Description      *string   `json:"description"`
	Avatar           *string   `json:"avatar"`
	IsGreetingClosed *bool     `json:"is_greeting_closed"`
	IsBlock          *bool     `json:"is_block"`
	GroupID          uuid.UUID `json:"group_id"`
	LastVisitTime    time.Time `json:"last_visit_time"`
}
type UpdateUserGreetingRequest struct {
	ID               uuid.UUID `json:"id" binding:"required"`
	IsGreetingClosed *bool     `json:"is_greeting_closed"`
}

type AuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
