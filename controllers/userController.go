package controllers

import (
	"fmt"
	"net/http"

	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/fajarhdev/go-mycash/models"
	"github.com/gin-gonic/gin"
)

// signup
func UserCreate(c *gin.Context) {

	var body struct {
		Fullname 	string
		Age 		int16
		Phonenumber string
		Email 		string
		Password 	string
	}

	c.Bind(&body)

	post := models.User{Fullname: body.Fullname, Age: body.Age, PhoneNumber: body.Phonenumber, Email: body.Email, Password: body.Password}
	result := initializers.DB.Create(&post)

	if result != nil {
		fmt.Println(result.Error)
	}else {
		fmt.Println(result.RowsAffected)
	}

	c.JSON(http.StatusCreated, gin.H{
		"status" : http.StatusCreated,
		"user": post,
	})
}


// login
func GetUser(c *gin.Context){
	var userBody struct {
		Email 		string
		Password 	string
	}

	c.Bind(&userBody)

	var user []models.User
	initializers.DB.Where("email = ?", userBody.Email).Find(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusFound,
		"user":user,
	})
	
}