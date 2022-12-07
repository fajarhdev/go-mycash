package query

import (
	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
)

func AddExpense(expense *models.Expense) {
	initializers.DB.Omit("User").Create(&expense)

}

func GetAllExpense(expense *[]models.Expense, id string) {
	initializers.DB.Where("user_id = ?", id).Joins("User").Find(&expense)
}

func UpdateExpense(expense *models.Expense, updateExpense models.Expense, id string) {
	initializers.DB.Model(&expense).Where("id = ?", id).Updates(&updateExpense)
}

// delete specific expense
func DeleteExpense(expense *models.Expense, id string) (err error){
	initializers.DB.Delete(&expense, id)
	return
}


