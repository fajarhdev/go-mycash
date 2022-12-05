package controllers

import (
	"net/http"

	"github.com/fajarhdev/go-mycash/models"
	"github.com/fajarhdev/go-mycash/query"
	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
	var expense models.Expense

	if err := c.BindJSON(&expense); err != nil {
		return
	}

	query.AddExpense(expense)
	
	c.JSON(http.StatusCreated, gin.H{
		"message":"Created",
		"status":http.StatusCreated,
	})
}

func GetAllExpense(c *gin.Context){
	var expense []models.Expense

	query.GetAllExpenses(&expense)

	c.JSON(http.StatusFound, expense)
}