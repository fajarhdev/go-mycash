package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	UserID 	int64
	Amount 	int64
	User 	User
}