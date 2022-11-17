package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname 	string 	
	Age 		int16 			
	PhoneNumber string 	
	Email 		string 		
	Password 	string		
}