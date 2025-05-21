package controller

import (
	"net/http"
	"time"
	"wallet-api/internal/models"
	"wallet-api/internal/service"
	"wallet-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	UserService        service.UserService
	TransactionService service.TransactionService
}

func (ctrl *Controller) DashboardHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Dashboard running"})
}

func (ctrl *Controller) Register(c *gin.Context) {
	var req struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		PIN         string `json:"pin"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	hashedPIN, _ := utils.HashPIN(req.PIN)
	user := &models.User{
		UserID:      uuid.NewString(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		PIN:         hashedPIN,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := ctrl.UserService.CreateUser(user)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Phone Number already registered")
		return
	}
	utils.RespondSuccess(c, user)
}

func (ctrl *Controller) Login(c *gin.Context) {
	var req struct {
		PhoneNumber string `json:"phone_number"`
		PIN         string `json:"pin"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request")
		return
	}
	user, err := ctrl.UserService.GetByPhone(req.PhoneNumber)
	if err != nil || utils.CheckPIN(user.PIN, req.PIN) != nil {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	token, _ := utils.GenerateJWT(user.UserID)
	utils.RespondSuccess(c, gin.H{"access_token": token})
}

func (ctrl *Controller) TopUp(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id"`
		Amount int64  `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
		utils.RespondError(c, http.StatusBadRequest, "Invalid topup request")
		return
	}
	user, err := ctrl.UserService.GetByID(req.UserID)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "User not found")
		return
	}
	before := user.Balance
	user.Balance += req.Amount
	ctrl.UserService.UpdateBalance(user)
	tx := &models.Transaction{
		ID:              uuid.NewString(),
		UserID:          user.UserID,
		TransactionType: "CREDIT",
		Amount:          req.Amount,
		BalanceBefore:   before,
		BalanceAfter:    user.Balance,
		CreatedAt:       time.Now(),
	}
	ctrl.TransactionService.CreateTransaction(tx)
	utils.RespondSuccess(c, tx)
}

func (ctrl *Controller) Pay(c *gin.Context) {
	var req struct {
		UserID  string `json:"user_id"`
		Amount  int64  `json:"amount"`
		Remarks string `json:"remarks"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
		utils.RespondError(c, http.StatusBadRequest, "Invalid pay request")
		return
	}
	user, err := ctrl.UserService.GetByID(req.UserID)
	if err != nil || user.Balance < req.Amount {
		utils.RespondError(c, http.StatusBadRequest, "Insufficient balance")
		return
	}
	before := user.Balance
	user.Balance -= req.Amount
	ctrl.UserService.UpdateBalance(user)
	tx := &models.Transaction{
		ID:              uuid.NewString(),
		UserID:          user.UserID,
		TransactionType: "DEBIT",
		Amount:          req.Amount,
		Remarks:         req.Remarks,
		BalanceBefore:   before,
		BalanceAfter:    user.Balance,
		CreatedAt:       time.Now(),
	}
	ctrl.TransactionService.CreateTransaction(tx)
	utils.RespondSuccess(c, tx)
}

func (ctrl *Controller) Transfer(c *gin.Context) {
	var req struct {
		FromUserID string `json:"user_id"`
		ToUserID   string `json:"target_user"`
		Amount     int64  `json:"amount"`
		Remarks    string `json:"remarks"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
		utils.RespondError(c, http.StatusBadRequest, "Invalid transfer request")
		return
	}
	fromUser, err := ctrl.UserService.GetByID(req.FromUserID)
	if err != nil || fromUser.Balance < req.Amount {
		utils.RespondError(c, http.StatusBadRequest, "Sender has insufficient balance")
		return
	}
	toUser, err := ctrl.UserService.GetByID(req.ToUserID)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Receiver not found")
		return
	}
	fromBefore := fromUser.Balance
	toBefore := toUser.Balance

	fromUser.Balance -= req.Amount
	toUser.Balance += req.Amount
	ctrl.UserService.UpdateBalance(fromUser)
	ctrl.UserService.UpdateBalance(toUser)

	tx1 := &models.Transaction{
		ID:              uuid.NewString(),
		UserID:          fromUser.UserID,
		TransactionType: "DEBIT",
		Amount:          req.Amount,
		Remarks:         "Transfer to " + toUser.UserID + ": " + req.Remarks,
		BalanceBefore:   fromBefore,
		BalanceAfter:    fromUser.Balance,
		CreatedAt:       time.Now(),
	}
	tx2 := &models.Transaction{
		ID:              uuid.NewString(),
		UserID:          toUser.UserID,
		TransactionType: "CREDIT",
		Amount:          req.Amount,
		Remarks:         "Transfer from " + fromUser.UserID + ": " + req.Remarks,
		BalanceBefore:   toBefore,
		BalanceAfter:    toUser.Balance,
		CreatedAt:       time.Now(),
	}
	ctrl.TransactionService.CreateTransaction(tx1)
	ctrl.TransactionService.CreateTransaction(tx2)

	utils.RespondSuccess(c, gin.H{"message": "Transfer successful"})
}

func (ctrl *Controller) GetTransactions(c *gin.Context) {
	userID := c.Query("user_id")
	txns, err := ctrl.TransactionService.GetUserTransactions(userID)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to fetch transactions")
		return
	}
	utils.RespondSuccess(c, txns)
}

func (ctrl *Controller) UpdateProfile(c *gin.Context) {
	var req struct {
		UserID    string `json:"user_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address   string `json:"address"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid profile update request")
		return
	}
	user, err := ctrl.UserService.GetByID(req.UserID)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "User not found")
		return
	}
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Address = req.Address
	user.UpdatedAt = time.Now()
	ctrl.UserService.UpdateUser(user)
	utils.RespondSuccess(c, user)
}
