package query

import (
	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
)

func Transaction(income *[]models.Income, id string) {
	initializers.DB.Select("amount").Where("user_id = ?", id).Find(&income)
}