package service

import (
	"web-lab/internal/entity"
	"web-lab/internal/repository"

	"github.com/google/uuid"
)

type SubscriptionService interface {
	Create(userID, targetID uuid.UUID) error
	Delete(userID, targetID uuid.UUID) error
	GetUserSubscriptions(ID uuid.UUID) ([]entity.Subscription, error)
	GetUserSubscribers(ID uuid.UUID) ([]entity.Subscription, error)
	CheckIsSubscribed(userID, targetID uuid.UUID) bool
}

type subscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{repo: repo}
}

func (s *subscriptionService) Create(userID, targetID uuid.UUID) error {
	return s.repo.Create(userID, targetID)
}

func (s *subscriptionService) Delete(userID, targetID uuid.UUID) error {
	return s.repo.Delete(userID, targetID)
}

func (s *subscriptionService) GetUserSubscriptions(ID uuid.UUID) ([]entity.Subscription, error) {
	return s.repo.GetUserSubscriptions(ID)
}

func (s *subscriptionService) GetUserSubscribers(ID uuid.UUID) ([]entity.Subscription, error) {
	return s.repo.GetUserSubscribers(ID)
}

func (s *subscriptionService) CheckIsSubscribed(userID, targetID uuid.UUID) bool {
	return s.repo.CheckIsSubscribed(userID, targetID)
}
