package models

import "gorm.io/gorm"

type Income struct {
	gorm.Model
	UserID 	int64
	Amount 	int64
	User	User
}