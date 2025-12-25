package service

import (
	"time"
	"web-lab/internal/dto"
	"web-lab/internal/entity"
	"web-lab/internal/repository"
	"web-lab/pkg/utils"

	"github.com/google/uuid"
)

type UserService interface {
	Create(usr *dto.CreateUserRequest) (*entity.User, error)
	Update(user *dto.UpdateUserRequest) error
	UpdateLastVisitTime(userID uuid.UUID, time time.Time) error
	UpdatePassword(userID uuid.UUID, password string) error
	UpdateGreetingClosed(userID uuid.UUID, value bool) error
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

func (u *userService) Create(usr *dto.CreateUserRequest) (*entity.User, error) {
	passwordHash, err := utils.HashPassword(usr.Password)
	if err != nil {
		return nil, err
	}
	user := &entity.User{
		Name:     usr.Name,
		Email:    usr.Email,
		Password: passwordHash,
		GroupID:  uuid.MustParse("700c704d-f5c9-4a95-ad9e-c040b4429050"),
	}
	return u.repo.Create(user)
}

func (u *userService) Update(req *dto.UpdateUserRequest) error {
	oldUser, err := u.repo.GetByID(req.ID)
	if err != nil {
		return err
	}
	if req.Name != "" && req.Name != oldUser.Name {
		oldUser.Name = req.Name
	}
	if req.Email != "" && req.Email != oldUser.Email {
		oldUser.Email = req.Email
	}
	if req.IsBlock != nil {
		oldUser.IsBlock = *req.IsBlock
	}
	if req.GroupID != uuid.Nil && oldUser.GroupID != req.GroupID {
		oldUser.GroupID = req.GroupID
	}
	if !req.LastVisitTime.IsZero() {
		oldUser.LastVisitAt = req.LastVisitTime
	}
	if req.Description != oldUser.Description {
		oldUser.Description = req.Description
	}
	if oldUser.Avatar != req.Avatar {
		oldUser.Avatar = req.Avatar
	}
	if req.IsGreetingClosed != nil && oldUser.IsGreetingClosed != req.IsGreetingClosed {
		oldUser.IsGreetingClosed = req.IsGreetingClosed
	}
	return u.repo.Update(oldUser)
}
func (u *userService) UpdateLastVisitTime(userID uuid.UUID, time time.Time) error {
	oldUser, err := u.repo.GetByID(userID)
	if err != nil {
		return err
	}
	oldUser.LastVisitAt = time
	return u.repo.Update(oldUser)
}
func (u *userService) UpdatePassword(userID uuid.UUID, password string) error {
	oldUser, err := u.repo.GetByID(userID)
	if err != nil {
		return err
	}
	newPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	oldUser.Password = newPassword
	return u.repo.Update(oldUser)
}
func (u *userService) UpdateGreetingClosed(userID uuid.UUID, value bool) error {
	oldUser, err := u.repo.GetByID(userID)
	if err != nil {
		return err
	}
	*oldUser.IsGreetingClosed = value
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
