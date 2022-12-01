package controllers

import (
	"net/http"

	"github.com/fajarhdev/go-mycash/models"
	"github.com/fajarhdev/go-mycash/query"
	"github.com/gin-gonic/gin"
)

// insert income
func PostIncome(c *gin.Context) {

	var incomeBody models.Income
	if err := c.BindJSON(&incomeBody); err != nil {
		return
	}

	query.AddIncome(incomeBody)

	c.JSON(http.StatusCreated, incomeBody)
}

// get income
func GetIncome (c *gin.Context){
	var income []models.Income

	query.GetAllIncome(&income)

	c.JSON(http.StatusOK, income)
}

