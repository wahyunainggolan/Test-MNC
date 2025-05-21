package controller_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"wallet-api/internal/controller"
	"wallet-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockUserService struct{}

func (m *mockUserService) CreateUser(u *models.User) error { return nil }
func (m *mockUserService) GetByPhone(p string) (*models.User, error) {
	return &models.User{UserID: "123", PIN: "$2a$14$yxbhL"}, nil
}
func (m *mockUserService) GetByID(id string) (*models.User, error) {
	return &models.User{UserID: id, Balance: 100000}, nil
}
func (m *mockUserService) UpdateBalance(u *models.User) error { return nil }
func (m *mockUserService) UpdateUser(u *models.User) error    { return nil }

type mockTxnService struct{}

func (m *mockTxnService) CreateTransaction(tx *models.Transaction) error { return nil }
func (m *mockTxnService) GetUserTransactions(userID string) ([]models.Transaction, error) {
	return []models.Transaction{
		{ID: "tx1", UserID: userID, Amount: 10000, TransactionType: "CREDIT", CreatedAt: time.Now()},
	}, nil
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	c := &controller.Controller{UserService: &mockUserService{}, TransactionService: &mockTxnService{}}

	r.POST("/register", c.Register)
	r.POST("/login", c.Login)
	r.POST("/topup", c.TopUp)
	r.GET("/transactions", c.GetTransactions)
	return r
}

func TestRegisterHandler(t *testing.T) {
	r := setupRouter()
	body := `{"first_name":"Jane","last_name":"Doe","phone_number":"0811","address":"Jl. A","pin":"1234"}`
	req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestLoginHandler(t *testing.T) {
	r := setupRouter()
	body := `{"phone_number":"0811","pin":"1234"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestTopUpHandler(t *testing.T) {
	r := setupRouter()
	body := `{"user_id":"123", "amount":5000}`
	req, _ := http.NewRequest("POST", "/topup", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetTransactionsHandler(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/transactions?user_id=123", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
