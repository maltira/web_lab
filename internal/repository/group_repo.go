package repository

import (
	"web-lab/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupRepository interface {
	Get(ID uuid.UUID) (*entity.Group, error)
	GetAll() ([]entity.Group, error)
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (g *groupRepository) Get(ID uuid.UUID) (*entity.Group, error) {
	var group entity.Group
	err := g.db.Where("id = ?", ID).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (g *groupRepository) GetAll() ([]entity.Group, error) {
	var groups []entity.Group
	err := g.db.Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}
