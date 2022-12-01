package query

import (
	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
)

func AddIncome(income models.Income) {
	initializers.DB.Create(&income)
}

func GetAllIncome(income *[]models.Income) {
	initializers.DB.Find(&income)
}

func UpdateIncome(income *models.Income, updateIncome *models.Income) {
	// initializers.DB.Model(&income).Update(&updateIncome)
}

func DeleteIncome(){
	
}