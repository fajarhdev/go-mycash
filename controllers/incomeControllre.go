package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostIncome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":"Get income",
	})
}
