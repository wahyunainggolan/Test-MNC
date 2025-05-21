package service_test

import (
	"testing"
	"time"

	"wallet-api/internal/models"
	"wallet-api/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByPhone(phone string) (*models.User, error) {
	args := m.Called(phone)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id string) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateBalance(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userSvc := service.NewUserService(mockRepo)

	newUser := &models.User{
		UserID:      "test-uuid",
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "08123456789",
		Address:     "Jl. Example",
		PIN:         "hashed-pin",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("Create", newUser).Return(nil)

	err := userSvc.CreateUser(newUser)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
