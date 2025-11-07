package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"not null;uniqueIndex"`
	Password string    `json:"password" gorm:"not null"`
	GroupID  uuid.UUID `json:"group_id" gorm:"type:uuid;not null"`
	IsBlock  bool      `json:"is_block" gorm:"not null;default:false"`

	LastVisitAt time.Time `json:"last_visit_at" gorm:"not null;default:now()"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Group Group `gorm:"foreignKey:GroupID;onDelete:CASCADE"`
}
