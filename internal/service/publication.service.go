package service

import (
	"errors"
	"fmt"
	"web-lab/internal/dto"
	"web-lab/internal/entity"
	"web-lab/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicationService interface {
	Create(publication *dto.PublicationRequest) error
	Delete(publicationID uuid.UUID) error
	Update(publication *dto.PublicationUpdateRequest) error

	FindByID(publicationID uuid.UUID) (*entity.Publication, error)
	FindByUserID(userID uuid.UUID, isDraft bool) ([]entity.Publication, error)
	FindAll(isDraft bool) ([]entity.Publication, error)
}

type publicationService struct {
	repo repository.PublicationRepository
	db   *gorm.DB
}

func NewPublicationService(repo repository.PublicationRepository, db *gorm.DB) PublicationService {
	return &publicationService{repo: repo, db: db}
}

func (s *publicationService) Create(publication *dto.PublicationRequest) error {
	tx := s.db.Begin()

	p := &entity.Publication{
		Title:           publication.Title,
		Description:     publication.Description,
		UserID:          publication.UserID,
		BackgroundColor: publication.BackgroundColor,
		IsDraft:         publication.IsDraft,
	}

	// * 1. Создание публикации
	if err := tx.Create(p).Error; err != nil {
		tx.Rollback()
		return err
	}

	// * 2. Обрабатываем категории
	if len(publication.Categories) <= 4 {
		for i, cat := range publication.Categories {
			category := &entity.Category{Name: cat.Category.Name}

			// * 1. Создаем или находим категорию
			result := tx.Where("name = ?", cat.Category.Name).FirstOrCreate(category)
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			}

			// * 2. Создаем связь с публикацией и стилями
			pubCategory := &entity.PublicationCategories{
				PublicationID:   p.ID,
				CategoryID:      category.ID,
				BackgroundColor: cat.BackgroundColor,
				TextColor:       cat.TextColor,
				DisplayOrder:    i,
			}

			if err := tx.Create(pubCategory).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	} else {
		tx.Rollback()
		return errors.New("указано более 4-х категорий")
	}

	tx.Commit()
	return nil
}

func (s *publicationService) Update(publication *dto.PublicationUpdateRequest) error {
	tx := s.db.Begin()

	// * Находим публикацию
	p, err := s.FindByID(publication.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// * Обновляем базовые поля
	if publication.Title != nil && *publication.Title != "" && p.Title != *publication.Title {
		p.Title = *publication.Title
	}
	if publication.Description != nil && *publication.Description != "" && p.Description != *publication.Description {
		p.Description = *publication.Description
	}
	if publication.BackgroundColor != nil && *publication.BackgroundColor != "" && p.BackgroundColor != *publication.BackgroundColor {
		p.BackgroundColor = *publication.BackgroundColor
	}
	if publication.IsDraft != nil && *publication.IsDraft != p.IsDraft {
		p.IsDraft = *publication.IsDraft
	}
	if err = tx.Save(p).Error; err != nil {
		tx.Rollback()
		return err
	}

	// * Обновляем категории
	if publication.Categories != nil && len(*publication.Categories) > 0 {
		if len(*publication.Categories) > 4 {
			tx.Rollback()
			return errors.New("указано более 4 категорий")
		}

		// * Удаляем существующие связи этой публикации с категориями
		if err = tx.Where("publication_id = ?", p.ID).
			Delete(&entity.PublicationCategories{}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("ошибка удаления старых категорий: %w", err)
		}

		// * Создаем новые связи
		for i, catReq := range *publication.Categories {

			// * Находим или создаем категорию
			var category entity.Category
			if err = tx.
				Where("name = ?", catReq.Category.Name).
				FirstOrCreate(&category, entity.Category{Name: catReq.Category.Name}).
				Error; err != nil {
				tx.Rollback()
				return err
			}

			// * Создаем новую связь
			pubCategory := &entity.PublicationCategories{
				PublicationID:   p.ID,
				CategoryID:      category.ID,
				BackgroundColor: catReq.BackgroundColor,
				TextColor:       catReq.TextColor,
				DisplayOrder:    i,
			}

			if err = tx.Create(pubCategory).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	} else {
		tx.Rollback()
		return errors.New("категории не указаны")
	}

	tx.Commit()

	return nil
}

func (s *publicationService) Delete(publicationID uuid.UUID) error {
	return s.repo.Delete(publicationID)
}

func (s *publicationService) FindByID(publicationID uuid.UUID) (*entity.Publication, error) {
	return s.repo.FindByID(publicationID)
}

func (s *publicationService) FindByUserID(userID uuid.UUID, isDraft bool) ([]entity.Publication, error) {
	return s.repo.FindByUserID(userID, isDraft)
}

func (s *publicationService) FindAll(isDraft bool) ([]entity.Publication, error) {
	return s.repo.FindAll(isDraft)
}
