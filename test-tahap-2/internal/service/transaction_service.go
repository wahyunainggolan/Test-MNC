package service

import (
	"time"
	"wallet-api/internal/models"
	"wallet-api/internal/repository"
)

type TransactionService interface {
	CreateTransaction(tx *models.Transaction) error
	GetUserTransactions(userID string) ([]models.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) CreateTransaction(tx *models.Transaction) error {
	tx.CreatedAt = time.Now()
	return s.repo.Create(tx)
}

func (s *transactionService) GetUserTransactions(userID string) ([]models.Transaction, error) {
	return s.repo.GetByUserID(userID)
}
