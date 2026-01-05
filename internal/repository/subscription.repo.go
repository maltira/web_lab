package repository

import (
	"web-lab/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	Create(userID, targetID uuid.UUID) error
	Delete(userID, targetID uuid.UUID) error
	GetUserSubscriptions(ID uuid.UUID) ([]entity.Subscription, error)
	GetUserSubscribers(ID uuid.UUID) ([]entity.Subscription, error)
	CheckIsSubscribed(userID, targetID uuid.UUID) bool
}

type subscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

func (r *subscriptionRepository) Create(userID, targetID uuid.UUID) error {
	return r.db.Create(&entity.Subscription{
		UserID:   userID,
		TargetID: targetID,
	}).Error
}

func (r *subscriptionRepository) Delete(userID, targetID uuid.UUID) error {
	return r.db.Where("user_id = ? AND target_id = ?", userID, targetID).Delete(&entity.Subscription{}).Error
}

func (r *subscriptionRepository) GetUserSubscriptions(ID uuid.UUID) ([]entity.Subscription, error) {
	var subscriptions []entity.Subscription
	err := r.db.Where("user_id = ?", ID).Preload("TargetUser").Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}
func (r *subscriptionRepository) GetUserSubscribers(ID uuid.UUID) ([]entity.Subscription, error) {
	var subscriptions []entity.Subscription
	err := r.db.Where("target_id = ?", ID).Preload("SubscriberUser").Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (r *subscriptionRepository) CheckIsSubscribed(userID, targetID uuid.UUID) bool {
	err := r.db.Where("user_id = ? AND target_id = ?", userID, targetID).First(&entity.Subscription{}).Error
	if err != nil {
		return false
	}
	return true
}
