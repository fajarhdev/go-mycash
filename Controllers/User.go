package Controllers

import (
	config "go-api-sql/Config"
	models "go-api-sql/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	result := config.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": result.Value})
}

// This struct is to validate the input parameters
type newInput struct {
	Uname string `json:"uname" binding:"required"`
	Pass  string `json:"pass" binding:"required"`
}

func CreateUser(c *gin.Context) {

	var new_input newInput

	validate := c.ShouldBindJSON(&new_input)
	if validate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": validate.Error()})
		return
	}

	new_user := models.User{Uname: new_input.Uname, Pass: new_input.Pass}
	config.DB.Create(&new_user)
	c.JSON(http.StatusOK, gin.H{"New Data": new_user})
}

func GetUserByID(c *gin.Context) {
	var userSelected models.User

	result := config.DB.Where("id = ?", c.Param("id")).First(&userSelected)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User " + c.Param("id") + " is not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userSelected})

	// [To select just some column]
	// Set the return value to []string
	// var some_column []string
	// result := config.DB.Model(&User{}).Pluck("Uname", &orang)
}

type updateInput struct {
	Uname string `json:"uname"`
	Pass  string `json:"pass"`
}

func UpdateUser(c *gin.Context) {

	// Check the user want to update is exist or not
	var userSelected models.User
	check := config.DB.Where("id = ?", c.Param("id")).First(&userSelected)
	if check.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User " + c.Param("id") + " is not found!"})
		return
	}

	// Validate the user input
	var update_input updateInput
	validate := c.ShouldBindJSON(&update_input)
	if validate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": validate.Error()})
		return
	}

	config.DB.Model(&userSelected).Updates(&update_input)
	c.JSON(http.StatusOK, gin.H{"data": userSelected})
}

func DeleteUser(c *gin.Context) {

	var userSelected models.User

	check := config.DB.Where("id = ?", c.Param("id")).First(&userSelected)
	if check.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User " + c.Param("id") + " is not found!"})
		return
	}

	config.DB.Delete(&userSelected)
	c.JSON(http.StatusOK, gin.H{"message": "User " + c.Param("id") + " has been deleted!"})
}
