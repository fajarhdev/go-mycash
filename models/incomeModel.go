package models

import "gorm.io/gorm"

type Income struct {
	gorm.Model
	Amount 	string 	`json:"amount"`
	UserID 	int
	User	*User
}