package controllers

import (
	"fmt"
	"net/http"

	"github.com/fajarhdev/go-mycash/models"
	"github.com/fajarhdev/go-mycash/query"
	"github.com/gin-gonic/gin"
)

func Transaction(c *gin.Context) {
	id := c.Params.ByName("id")
	var income []models.Income

	query.Transaction(&income, id)

	fmt.Println(income)
	c.JSON(http.StatusOK, income)
}

// get all of the income transaction of specific user
// get all of the expense transaction od specific user
// fetch only the amount of each transaction
// calculate each transaction
// calculate the income - expense
// if the result is minus the deficit or not health
// if the result is zero than it balance
// if the result is plus than its surplus or healthy