package service

import (
	"web-lab/internal/dto"
	"web-lab/internal/entity"
	"web-lab/internal/repository"

	"github.com/google/uuid"
)

type PublicationService interface {
	Create(publication *dto.PublicationRequest) error
	Delete(publicationID uuid.UUID) error
	FindByID(publicationID uuid.UUID) (*entity.Publication, error)
	FindAll() ([]entity.Publication, error)
}

type publicationService struct {
	repo repository.PublicationRepository
}

func NewPublicationService(repo repository.PublicationRepository) PublicationService {
	return &publicationService{repo: repo}
}

func (s *publicationService) Create(publication *dto.PublicationRequest) error {
	p := &entity.Publication{
		Title:       publication.Title,
		Description: publication.Description,
		UserID:      publication.UserID,
	}
	return s.repo.Create(p)
}

func (s *publicationService) Delete(publicationID uuid.UUID) error {
	return s.repo.Delete(publicationID)
}

func (s *publicationService) FindByID(publicationID uuid.UUID) (*entity.Publication, error) {
	return s.repo.FindByID(publicationID)
}

func (s *publicationService) FindAll() ([]entity.Publication, error) {
	return s.repo.FindAll()
}
