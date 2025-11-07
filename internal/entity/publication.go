package entity

import (
	"time"

	"github.com/google/uuid"
)

type Publication struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserID;onDelete:CASCADE"`
}
