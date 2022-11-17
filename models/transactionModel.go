package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID 		int64
	TotalAmount int64
	Status 		string
	User 		User
}