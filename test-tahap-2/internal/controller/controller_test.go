package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"wallet-api/internal/controller"
	"wallet-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct{}

func (m *MockUserService) CreateUser(user *models.User) error { return nil }
func (m *MockUserService) GetByPhone(phone string) (*models.User, error) {
	return &models.User{UserID: "test-uuid", PIN: "$2a$14$Q9xF2kOjA1dSTDbW4YIkme2GOrw4Akw9iYOz8n3eE6m0RlT3qI4X2"}, nil
}
func (m *MockUserService) GetByID(id string) (*models.User, error) {
	return &models.User{UserID: id, Balance: 100000}, nil
}
func (m *MockUserService) UpdateBalance(user *models.User) error { return nil }
func (m *MockUserService) UpdateUser(user *models.User) error    { return nil }

type MockTransactionService struct{}

func (m *MockTransactionService) CreateTransaction(tx *models.Transaction) error { return nil }
func (m *MockTransactionService) GetUserTransactions(userID string) ([]models.Transaction, error) {
	return []models.Transaction{
		{ID: "tx1", UserID: userID, Amount: 50000, TransactionType: "CREDIT", CreatedAt: time.Now()},
	}, nil
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	userSvc := &MockUserService{}
	txnSvc := &MockTransactionService{}
	controller.Controller(r, userSvc, txnSvc)

	return r
}

func TestRegisterEndpoint(t *testing.T) {
	r := setupRouter()
	reqBody := `{
        "first_name": "John",
        "last_name": "Doe",
        "phone_number": "0811111111",
        "address": "Jl. Test",
        "pin": "123456"
    }`

	req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginEndpoint(t *testing.T) {
	r := setupRouter()
	reqBody := `{
        "phone_number": "0811111111",
        "pin": "123456"
    }`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTopUpEndpoint(t *testing.T) {
	r := setupRouter()
	data := map[string]interface{}{
		"user_id": "test-user-id",
		"amount":  50000,
	}
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/topup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPayEndpoint(t *testing.T) {
	r := setupRouter()
	data := map[string]interface{}{
		"user_id": "test-user-id",
		"amount":  5000,
		"remarks": "test payment",
	}
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/pay", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTransferEndpoint(t *testing.T) {
	r := setupRouter()
	data := map[string]interface{}{
		"user_id":     "test-user-id",
		"target_user": "target-user-id",
		"amount":      10000,
		"remarks":     "test transfer",
	}
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/transfer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateProfileEndpoint(t *testing.T) {
	r := setupRouter()
	data := map[string]interface{}{
		"user_id":    "test-user-id",
		"first_name": "Jane",
		"last_name":  "Smith",
		"address":    "New Address",
	}
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("PUT", "/profile", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTransactionsEndpoint(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/transactions?user_id=test-user-id", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
