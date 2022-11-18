package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup
func UserCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":"Create user",
	})
}


// login
func GetUser(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Get user",
	})
}