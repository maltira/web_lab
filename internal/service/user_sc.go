package service

import (
	"web-lab/internal/dto"
	"web-lab/internal/entity"
	"web-lab/internal/repository"
	utils "web-lab/pkg/utils"

	"github.com/google/uuid"
)

type UserService interface {
	Create(usr *dto.CreateUserRequest) error
	Update(user *dto.UpdateUserRequest) error
	Delete(ID uuid.UUID) error
	GetByID(ID uuid.UUID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetAll() ([]entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) Create(usr *dto.CreateUserRequest) error {
	passwordHash, err := utils.HashPassword(usr.Password)
	if err != nil {
		return err
	}
	user := &entity.User{
		Name:     usr.Name,
		Email:    usr.Email,
		Password: passwordHash,
		GroupID:  usr.GroupID,
	}
	return u.repo.Create(user)
}
func (u *userService) Update(user *dto.UpdateUserRequest) error {
	oldUser, err := u.repo.GetByID(user.ID)
	if err != nil {
		return err
	}
	if user.Name != "" {
		oldUser.Name = user.Name
	}
	if user.Email != "" {
		oldUser.Email = user.Email
	}
	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return err
		}
		oldUser.Password = hashedPassword
	}
	if user.GroupID != uuid.Nil {
		oldUser.GroupID = user.GroupID
	}
	return u.repo.Update(oldUser)
}
func (u *userService) Delete(ID uuid.UUID) error {
	return u.repo.Delete(ID)
}
func (u *userService) GetByID(ID uuid.UUID) (*entity.User, error) {
	return u.repo.GetByID(ID)
}
func (u *userService) GetByEmail(email string) (*entity.User, error) {
	return u.repo.GetByEmail(email)
}
func (u *userService) GetAll() ([]entity.User, error) {
	return u.repo.GetAll()
}
