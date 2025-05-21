package main

import (
	"log"
	"wallet-api/config"
	"wallet-api/internal/controller"
	"wallet-api/internal/models"
	"wallet-api/internal/repository"
	"wallet-api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate tables
	db.AutoMigrate(&models.User{}, &models.Transaction{})

	// Repositories and services
	userRepo := repository.NewUserRepository(db)
	txnRepo := repository.NewTransactionRepository(db)

	userService := service.NewUserService(userRepo)
	txnService := service.NewTransactionService(txnRepo)

	// Setup Gin
	r := gin.Default()

	// Initialize controller
	ctrl := &controller.Controller{
		UserService:        userService,
		TransactionService: txnService,
	}

	// Route definitions
	r.GET("/dashboard/health", ctrl.DashboardHealth)
	r.POST("/register", ctrl.Register)
	r.POST("/login", ctrl.Login)
	r.POST("/topup", ctrl.TopUp)
	r.POST("/pay", ctrl.Pay)
	r.POST("/transfer", ctrl.Transfer)
	r.GET("/transactions", ctrl.GetTransactions)
	r.PUT("/profile", ctrl.UpdateProfile)

	r.Run(":8080")
}
