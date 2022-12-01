package routes

import (
	"github.com/fajarhdev/go-mycash/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.CreateUser)
		user.GET("/:id", controllers.GetUserByID)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
	
	income := r.Group("/income")
	{
		income.POST("/", controllers.PostIncome)
		income.GET("/", controllers.GetIncome)
	}
	return r
}