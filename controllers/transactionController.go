package controllers

import (
	"net/http"
	"strconv"
	
	"github.com/fajarhdev/go-mycash/models"
	"github.com/fajarhdev/go-mycash/query"
	"github.com/gin-gonic/gin"
)

// this function is the logic of total transaction per user
func TransactionCore(c *gin.Context) (UserStatus string, TotalAmount int) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var income models.Income
	var expense models.Expense
	var transactions models.Transaction

	// get sum income
	query.Income(&income, id)
	sumIncome := income.Amount
	
	// get sum expense
	query.Expense(&expense, id)
	sumExpense := expense.Amount

	// user status logic
	if TotalAmount = sumIncome - sumExpense; TotalAmount > 0 {
		UserStatus = "Health"
	}else if TotalAmount < 0 {
		UserStatus = "Minus"
	} else {
		UserStatus = "Balance"
	}

	transaction := models.Transaction{
		TotalAmount: TotalAmount,
		Status: UserStatus,
		UserID: id,
	}

	// post it to the db
	if result := query.GetAllTransactionByUser(&transactions, id); result > 0 {
		query.UpdateTransaction(&transactions, &transaction, id)
	}else {
		query.PostTransaction(transaction)
	}

	return 
}

// this will add and update the transaction total per user
func Transaction(c *gin.Context) {
	var userStatus, totalAmount = TransactionCore(c)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user_status": userStatus,
		"total_amount": totalAmount,
	})

}