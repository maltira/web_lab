package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
	GroupID  uuid.UUID `json:"group_id" binding:"required"`
}

type UpdateUserRequest struct {
	ID            uuid.UUID `json:"id" binding:"required"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	GroupID       uuid.UUID `json:"group_id"`
	LastVisitTime time.Time `json:"last_visit_time"`
}

type AuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
