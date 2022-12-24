package models

import (
	"gorm.io/gorm"
)

// the schema for user table
type User struct {
	gorm.Model
	Fullname 	string 		`json:"fullname"`
	Age 		string 		`json:"age"`		
	PhoneNumber string 		`json:"phonenumber"`
	Email 		string 		`json:"email"`	
	Password 	string		`json:"password"`
	Income		[]Income
	Expense		[]Expense
}