package service

import (
	"web-lab/internal/entity"
	"web-lab/internal/repository"

	"github.com/google/uuid"
)

type GroupService interface {
	GetByID(ID uuid.UUID) (*entity.Group, error)
	GetAll() ([]entity.Group, error)
}

type groupService struct {
	repo repository.GroupRepository
}

func NewGroupService(repo repository.GroupRepository) GroupService {
	return &groupService{repo: repo}
}

func (g *groupService) GetByID(ID uuid.UUID) (*entity.Group, error) {
	return g.repo.Get(ID)
}
func (g *groupService) GetAll() ([]entity.Group, error) {
	return g.repo.GetAll()
}
