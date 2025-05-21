package service_test

import (
	"testing"
	"time"
	"wallet-api/internal/models"
	"wallet-api/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTxnRepo struct {
	mock.Mock
}

func (m *MockTxnRepo) Create(tx *models.Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockTxnRepo) GetByUserID(userID string) ([]models.Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func TestCreateTransaction(t *testing.T) {
	mockRepo := new(MockTxnRepo)
	svc := service.NewTransactionService(mockRepo)

	tx := &models.Transaction{
		ID: "tx1", UserID: "user1", Amount: 50000, TransactionType: "CREDIT", CreatedAt: time.Now(),
	}

	mockRepo.On("Create", tx).Return(nil)

	err := svc.CreateTransaction(tx)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUserTransactions(t *testing.T) {
	mockRepo := new(MockTxnRepo)
	svc := service.NewTransactionService(mockRepo)

	mockData := []models.Transaction{
		{ID: "tx1", UserID: "user1", Amount: 50000, TransactionType: "CREDIT"},
	}
	mockRepo.On("GetByUserID", "user1").Return(mockData, nil)

	txns, err := svc.GetUserTransactions("user1")
	assert.NoError(t, err)
	assert.Len(t, txns, 1)
	assert.Equal(t, int64(50000), txns[0].Amount)
}