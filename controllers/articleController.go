package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Get articles",
	})
}