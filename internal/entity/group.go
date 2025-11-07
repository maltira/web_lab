package entity

import "github.com/google/uuid"

type Group struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey; not null"`
	Name            string    `json:"name" gorm:"not null;uniqueIndex"`
	IsAvailable     bool      `json:"is_available" gorm:"not null; default:true"`
	CanPublishPosts bool      `json:"can_publish_posts" gorm:"not null; default:true"`
}
