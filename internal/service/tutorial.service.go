package service

import (
	"web-lab/internal/entity"
	"web-lab/internal/repository"
)

type TutorialService interface {
	GetAll() ([]entity.Tutorial, error)
}

type tutorialService struct {
	repo repository.TutorialRepository
}

func NewTutorialService(repo repository.TutorialRepository) TutorialService {
	return &tutorialService{repo: repo}
}

func (t *tutorialService) GetAll() ([]entity.Tutorial, error) {
	return t.repo.GetAll()
}
