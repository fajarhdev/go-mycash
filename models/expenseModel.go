package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Amount 	string	`json:"amount"`
	UserID 	int
	User 	*User
}