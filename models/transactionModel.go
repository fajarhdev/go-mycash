package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TotalAmount string	`json:"totalamount"`
	Status 		string	`json:"status"`
	UserID 		int64	`json:"userid"`
	User 		User	`json:"user" gorm:"foreignKey:UserID"`
}