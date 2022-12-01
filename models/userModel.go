package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname 	string 	`json:"fullname"`
	Age 		int16 	`json:"age"`		
	PhoneNumber string 	`json:"phonenumber"`
	Email 		string 	`json:"email"`	
	Password 	string	`json:"password"`	
}