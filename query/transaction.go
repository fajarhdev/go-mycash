package query

import (

	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
)

func Income(income *models.Income, id int) {
	// results := initializers.DB.Select("amount").Where("user_id = ?", id).Find(&income)
	// fmt.Println(results)
	initializers.DB.Raw("SELECT SUM(amount) AS amount FROM incomes WHERE user_id = ?", id).Scan(&income)
}

func Expense(expense *models.Expense, id int) {
	initializers.DB.Raw("SELECT SUM(amount) AS amount FROM expenses WHERE user_id = ?", id).Scan(&expense)
}

// post to db transaction
func PostTransaction(transaction models.Transaction) {
	initializers.DB.Omit("User").Create(&transaction)
}