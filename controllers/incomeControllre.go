package controllers

import (
	"fmt"
	"net/http"

	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
	"github.com/gin-gonic/gin"
)

// insert income
func PostIncome(c *gin.Context) {

	var incomeBody models.Income

	c.BindJSON(&incomeBody)
	post := models.Income{Amount: incomeBody.Amount}
	result := initializers.DB.Create(&post)

	fmt.Println(incomeBody)
	if result != nil {
		fmt.Println(result.Error)
	}else {
		fmt.Println(result.RowsAffected)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"income": post,
	})
}

// get income
func getIncome (c *gin.Context){
	
}

