package models

import "time"

type User struct {
	UserID      string    `gorm:"primaryKey" json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `gorm:"unique" json:"phone_number"`
	Address     string    `json:"address"`
	PIN         string    `json:"-"` // encrypted
	Balance     int64     `json:"balance"`
	CreatedAt   time.Time `json:"created_date"`
	UpdatedAt   time.Time `json:"updated_date"`
}

type Transaction struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	UserID          string    `json:"user_id"`
	TransactionType string    `json:"transaction_type"` // DEBIT or CREDIT
	Amount          int64     `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   int64     `json:"balance_before"`
	BalanceAfter    int64     `json:"balance_after"`
	CreatedAt       time.Time `json:"created_at"`
}
