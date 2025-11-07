package repository

import (
	"time"
	"web-lab/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicationRepository interface {
	Create(publication *entity.Publication) error
	Delete(publicationID uuid.UUID) error
	Update(publication *entity.Publication) error

	FindByID(publicationID uuid.UUID) (*entity.Publication, error)
	FindAll() ([]entity.Publication, error)
}

type publicationRepository struct {
	db *gorm.DB
}

func NewPublicationRepository(db *gorm.DB) PublicationRepository {
	return &publicationRepository{db: db}
}

func (p *publicationRepository) Create(publication *entity.Publication) error {
	return p.db.Create(&publication).Error
}

func (p *publicationRepository) Delete(publicationID uuid.UUID) error {
	return p.db.Delete(&entity.Publication{}, publicationID).Error
}

func (p *publicationRepository) Update(publication *entity.Publication) error {
	response := p.db.Model(&entity.Publication{}).Where("id = ?", publication.ID).Updates(map[string]interface{}{
		"title":       publication.Title,
		"description": publication.Description,
		"categories":  publication.Categories,
		"updated_at":  time.Now(),
	})
	return response.Error
}

func (p *publicationRepository) FindByID(publicationID uuid.UUID) (*entity.Publication, error) {
	var publication entity.Publication
	err := p.db.Preload("User").First(&publication, publicationID).Error
	if err != nil {
		return nil, err
	}
	return &publication, nil
}

func (p *publicationRepository) FindAll() ([]entity.Publication, error) {
	var publications []entity.Publication
	if err := p.db.Preload("User").Find(&publications).Error; err != nil {
		return nil, err
	}
	return publications, nil
}
