package query

import (
	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
)

// the function for query adding expense transaction to database
func AddExpense(expense *models.Expense) {
	initializers.DB.Omit("User").Create(&expense)
}

// the function for query fetch expense transaction with specific id user
func GetAllExpense(expense *[]models.Expense, id int) {
	initializers.DB.Where("user_id = ?", id).Joins("User").Find(&expense)
}

// the function for query update expense transaction with specific id user
func UpdateExpense(expense *models.Expense, updateExpense models.Expense, id int) {
	initializers.DB.Model(&expense).Where("id = ?", id).Updates(&updateExpense)
}

// the function for query delete expense transaction with specific id user
func DeleteExpense(expense *models.Expense, id int) (err error){
	initializers.DB.Delete(&expense, id)
	return
}


