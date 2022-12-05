package controllers

import (
	"github.com/fajarhdev/go-mycash/models"
	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
	var expense models.Expense

	if err := c.BindJSON(&expense); err != nil {
		return
	}

	
}