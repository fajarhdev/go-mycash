package models

import "gorm.io/gorm"

type Income struct {
	gorm.Model
	Amount 	int64 	`json:"amount"`
	UserID 	int64 	`json:"userid"`
	User	User	`json:"user" gorm:"foreignKey:UserID"`
}