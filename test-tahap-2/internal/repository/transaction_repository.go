package repository

import (
	"wallet-api/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *models.Transaction) error
	GetByUserID(userID string) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(txn *models.Transaction) error {
	return r.db.Create(txn).Error
}

func (r *transactionRepository) GetByUserID(userID string) ([]models.Transaction, error) {
	var txns []models.Transaction
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&txns).Error
	return txns, err
}
