package repository

import (
	"web-lab/internal/entity"

	"gorm.io/gorm"
)

type TutorialRepository interface {
	GetAll() ([]entity.Tutorial, error)
}

type tutorialRepository struct {
	db *gorm.DB
}

func NewTutorialRepository(db *gorm.DB) TutorialRepository {
	return &tutorialRepository{db: db}
}

func (r *tutorialRepository) GetAll() ([]entity.Tutorial, error) {
	var tutorials []entity.Tutorial
	err := r.db.Find(&tutorials).Error

	if err != nil {
		return nil, err
	}
	return tutorials, nil
}
