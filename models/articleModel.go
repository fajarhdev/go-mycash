package models

import "gorm.io/gorm"

// schema fot article table
type Article struct {
	gorm.Model
	Title 	string
	Author 	string
	Text 	string
	Img 	string
}