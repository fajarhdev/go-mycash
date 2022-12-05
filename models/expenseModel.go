package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Amount 	string	`json:"amount"`
	UserID 	int64	`json:"userid"`
	User 	User	`json:"user" gorm:"foreignKey:UserID"`
}