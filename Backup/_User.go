package Controllers

import (
	models "go-api-sql/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	result := models.GetAllUsers()

	if result != nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// func GetUsers(c *gin.Context) {
// 	var user []models.User
// 	err := models.GetAllUsers(&user)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		fmt.Print(user)
// 		c.JSON(http.StatusOK, user)
// 	}
// }

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	result := models.CreateUser(&user)

	if result != nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.AbortWithStatus(http.StatusNotFound)

	}
}

// func CreateUser(c *gin.Context) {
// 	var user models.User
// 	c.BindJSON(&user)
// 	err := models.CreateUser(&user)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

func GetUserByID(c *gin.Context) {

	id := c.Params.ByName("id")
	var user models.User
	c.BindJSON(&user)

	result := models.GetUserByID(&user, id)

	if result != nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"User " + id: "is not found!"})
	}
}

// func GetUserByID(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user models.User
// 	c.BindJSON(&user)
// 	err := models.GetUserByID(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func UpdateUser(c *gin.Context) {

// 	id := c.Params.ByName("id")
// 	var user models.User
// 	// c.BindJSON(&user)

// 	search_user_id := models.GetUserByID(&user, id)
// 	if search_user_id != nil {
// 		c.JSON(http.StatusNotFound, user)
// 	}

// 	c.BindJSON(&user)
// 	update_user := models.UpdateUser(&user, id)
// 	if update_user != nil {
// 		c.JSON(http.StatusOK, user)
// 	} else {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// }

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	// c.BindJSON(&user)

	check_user_id := models.GetUserByID(&user, id)
	if check_user_id != nil {
		c.BindJSON(&user)
		check_user_id = models.UpdateUser(&user)

		// If check_user_id not nill, mean the data is exist
		if check_user_id != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"User " + id: "is not found!"})
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	c.BindJSON(&user)
	err := models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted!"})
	}
}
