package repository

import (
	"web-lab/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(ID uuid.UUID) error
	GetByID(ID uuid.UUID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetAll() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(user *entity.User) error {
	return u.db.Create(&user).Error
}

func (u *userRepository) Update(user *entity.User) error {
	return u.db.Save(&user).Error
}

func (u *userRepository) Delete(ID uuid.UUID) error {
	return u.db.Delete(&entity.User{}, "id = ?", ID).Error
}

func (u *userRepository) GetByID(ID uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("id = ?", ID).Preload("Group").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.db.Where("email = ?", email).Preload("Group").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	err := u.db.Preload("Group").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
