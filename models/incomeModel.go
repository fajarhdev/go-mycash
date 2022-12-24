package models

import "gorm.io/gorm"

// the schema for income table
type Income struct {
	gorm.Model
	Amount 	int 	`json:"amount"`
	UserID 	int
	User	*User
}