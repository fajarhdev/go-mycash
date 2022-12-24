package models

import "gorm.io/gorm"

// the schema for expense table
type Expense struct {
	gorm.Model
	Amount 	int	`json:"amount"`
	UserID 	int
	User 	*User
}