package controllers

import (
	"fmt"
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

	// var body struct {
	// 	Id int64 `json:"id"`
	// }
	
	var user models.User

	c.BindJSON(&user)

	query.GetAllIncome(&income, &user)

	c.JSON(http.StatusOK, income)
	fmt.Println(user)
}

