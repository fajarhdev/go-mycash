package models

import (
	"time"

	// "gorm.io/gorm"
)

type User struct {
	CreatedAt   time.Time	`json:"createdat"`
  	UpdatedAt   time.Time	`json:"updatedat"`
	Fullname 	string 		`json:"fullname"`
	Age 		int16 		`json:"age"`		
	PhoneNumber string 		`json:"phonenumber"`
	Email 		string 		`json:"email"`	
	Password 	string		`json:"password"`	
}