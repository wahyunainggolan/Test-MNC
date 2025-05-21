package repository

import (
	"wallet-api/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Create(user *models.User) error
	GetByPhone(phone string) (*models.User, error)
	GetByID(id string) (*models.User, error)
	UpdateBalance(user *models.User) error
	UpdateUser(user *models.User) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("phone_number = ?", phone).First(&user).Error
	return &user, err
}

func (r *userRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Where("user_id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepository) UpdateBalance(user *models.User) error {
	return r.db.Model(user).Update("balance", user.Balance).Error
}

func (r *userRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}
