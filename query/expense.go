package query

import (
	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
)

func AddExpense(expense models.Expense){
	initializers.DB.Create(&expense)
}

func GetExpensesByUser(expense *[]models.Expense, id int) {
	initializers.DB.Where("id <> ?", id).Find((&expense))
}

func GetAllExpenses(expense *[]models.Expense){
	initializers.DB.Find(&expense)
}


