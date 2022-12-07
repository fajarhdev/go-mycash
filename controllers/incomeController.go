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

	query.AddIncome(&incomeBody)

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"message": "Successfuly created income",
		"data": incomeBody,
	})
}

// get income
func GetIncome (c *gin.Context){
	var income []models.Income

	id := c.Params.ByName("id")

	// var user models.User

	// result := initializers.DB.Model(&user).Where("userid = ?", 4).Association("User").Find(&income)

	query.GetAllIncome(&income, id)

	c.JSON(http.StatusOK, gin.H{
		"status":http.StatusOK,
		"message":"Found the data",
		"data":income,
	})
}

// update income
func UpdateIncome (c *gin.Context){
	var income models.Income
	var incomeBody models.Income

	if err := c.BindJSON(&incomeBody); err != nil {
		return
	}

	id := c.Params.ByName("id")

	query.UpdateIncome(&income, incomeBody, id)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message":"Success updating",
		"data":incomeBody,
	})
}

// delete income
func DeleteIncome(c *gin.Context)  {
	var income models.Income
	
	id := c.Params.ByName("id")

	query.DeleteIncome(&income, id)

	c.JSON(http.StatusOK, gin.H{
		"status":http.StatusOK,
		"message":"Income has successfuly deleted",
	})
}

