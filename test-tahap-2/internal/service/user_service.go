package service

import (
	"wallet-api/internal/models"
	"wallet-api/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetByPhone(phone string) (*models.User, error)
	GetByID(id string) (*models.User, error)
	UpdateBalance(user *models.User) error
	UpdateUser(user *models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetByPhone(phone string) (*models.User, error) {
	return s.repo.GetByPhone(phone)
}

func (s *userService) GetByID(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) UpdateBalance(user *models.User) error {
	return s.repo.UpdateBalance(user)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}
